//go:generate go run -mod=mod github.com/magefile/mage downloadSwaggerUi

package main

import (
	"context"
	"errors"
	"flag"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"

	"github.com/baez90/go-icndb/api/health"
	"github.com/baez90/go-icndb/api/jokes"
	"github.com/baez90/go-icndb/api/swagger"
	"github.com/baez90/go-icndb/internal/logging"
	"github.com/baez90/go-icndb/internal/metrics"
)

var (
	appCfg struct {
		HTTP struct {
			Address           string
			ReadHeaderTimeout time.Duration
		}
		Logging struct {
			Level logging.LevelVar
		}
		Jokes struct {
			DefaultFirstName string
			DefaultLastName  string
		}
	}
)

func main() {
	appCfg.Logging.Level = logging.LevelVar{
		Value: new(slog.LevelVar),
	}

	setupLogging()

	flagSet := flag.NewFlagSet("icndb", flag.ExitOnError)

	prepareFlags(flagSet)

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		slog.Error("failed to parse flags", slog.String("error", err.Error()))
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)

	err := run(ctx)
	cancel()

	if err != nil {
		slog.Error("Error occurred", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	logger := slog.Default()
	router := httprouter.New()

	promExporter, err := prometheus.New()
	if err != nil {
		return err
	}

	provider := metric.NewMeterProvider(metric.WithReader(promExporter))
	meter := provider.Meter("github.com/baez90/go-icndb")

	router.Handler(http.MethodGet, "/metrics", promhttp.Handler())

	if err := swagger.SetupRouter(router); err != nil {
		return err
	}

	health.SetupRouter(router)
	router.Handler(http.MethodGet, "/", http.RedirectHandler("/swagger/ui", http.StatusPermanentRedirect))
	jokes.SetupRouter(router, appCfg.Jokes.DefaultFirstName, appCfg.Jokes.DefaultLastName)

	measuredHandler, err := metrics.RequestMetrics(router, meter)
	if err != nil {
		return err
	}

	srv := &http.Server{
		Addr:              appCfg.HTTP.Address,
		Handler:           logging.RequestLogger(measuredHandler, logger),
		ReadHeaderTimeout: 100 * time.Millisecond,
		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
	}

	logger.Info("Starting server", slog.String("address", appCfg.HTTP.Address))

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Error occurred", slog.String("error", err.Error()))
		}
	}()

	<-ctx.Done()

	return srv.Shutdown(context.Background())
}

func setupLogging() {
	opts := slog.HandlerOptions{
		Level: appCfg.Logging.Level.Value,
	}

	handler := slog.NewJSONHandler(os.Stderr, &opts)

	slog.SetDefault(slog.New(handler))
}

func prepareFlags(flagSet *flag.FlagSet) {
	flagSet.StringVar(
		&appCfg.HTTP.Address,
		"http.address",
		envOr("ICNDB_HTTP_ADDRESS", ":3000", noOpParse),
		"Listener address, the HTTP server will open - ICNDB_HTTP_ADDRESS",
	)

	flagSet.StringVar(
		&appCfg.Jokes.DefaultFirstName,
		"jokes.default-first-name",
		envOr("ICNDB_JOKES_DEFAULT_FIRST_NAME", "Chuck", noOpParse),
		"Name to replace Chuck with in Joke - ICNDB_JOKES_DEFAULT_FIRST_NAME",
	)

	flagSet.StringVar(
		&appCfg.Jokes.DefaultLastName,
		"jokes.default-last-name",
		envOr("ICNDB_JOKES_DEFAULT_LAST_NAME", "Norris", noOpParse),
		"Name to replace Norris with in Joke - ICNDB_JOKES_DEFAULT_LAST_NAME",
	)

	flagSet.Var(
		&appCfg.Logging.Level,
		"log.level",
		"Logging level, one of: debug, info, warn, error",
	)
}

func envOr[T any](key string, defaultValue T, parse func(v string) (T, error)) T {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	envVal = os.ExpandEnv(envVal)

	if parsed, err := parse(envVal); err != nil {
		return defaultValue
	} else {
		return parsed
	}
}

func noOpParse(v string) (string, error) {
	return v, nil
}
