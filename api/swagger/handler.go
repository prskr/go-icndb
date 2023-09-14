package swagger

import (
	"embed"
	_ "embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"path"
	txttemplate "text/template"

	"github.com/julienschmidt/httprouter"

	"github.com/baez90/go-icndb/assets"
)

var (
	//go:embed templates
	templates     embed.FS
	textTemplates *txttemplate.Template
)

func init() {
	var err error
	textTemplates, err = txttemplate.ParseFS(templates, "templates/*.*")
	if err != nil {
		panic(err)
	}
}

func SetupRouter(r *httprouter.Router) error {
	handler := Handler{}

	sub, err := fs.Sub(assets.FS, "swagger-ui")
	if err != nil {
		return err
	}

	handler.staticSwaggerUiHandler = http.StripPrefix("/swagger/ui/", http.FileServer(http.FS(sub)))

	r.GET("/swagger/ui/*filepath", handler.SwaggerUI)
	r.GET("/swagger/swagger.json", handler.SwaggerSpec)

	return nil
}

type Handler struct {
	staticSwaggerUiHandler http.Handler
}

func (h *Handler) SwaggerUI(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	switch params.ByName("filepath") {
	case "/swagger-initializer.js":
		h.swaggerInitializer(writer, req)
		return
	default:
		h.staticSwaggerUiHandler.ServeHTTP(writer, req)
	}
}

func (h *Handler) SwaggerSpec(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)

	specFile, err := assets.FS.Open(path.Join("api", "swagger.json"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	defer func() {
		_ = specFile.Close()
	}()

	if _, err := io.Copy(writer, specFile); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) swaggerInitializer(writer http.ResponseWriter, req *http.Request) {
	type templateData struct {
		BaseUrl string
	}

	scheme := "http"
	if req.TLS != nil {
		scheme = "https"
	}

	err := textTemplates.ExecuteTemplate(writer, "swagger-initializer.tmpl.js", templateData{
		fmt.Sprintf("%s://%s", scheme, req.Host),
	})

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
}
