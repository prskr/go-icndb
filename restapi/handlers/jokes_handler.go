package handlers

import (
	"fmt"
	"github.com/baez90/go-icndb/internal/pkg/metrics"
	"github.com/baez90/go-icndb/internal/pkg/models"
	respModels "github.com/baez90/go-icndb/models"
	"github.com/baez90/go-icndb/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"os"
)

const (
	defaultFirstName = "Chuck"
	defaultLastName  = "Norris"
)

type getJokesByIDHandler struct {
	facts *models.Facts
}

type getRandomJokeHandler struct {
	facts   *models.Facts
	metrics *metrics.AppMetricsCollectors
}

type getJokesCountHandler struct {
	facts *models.Facts
}

func NewJokesByIDHandler(facts *models.Facts) *getJokesByIDHandler {
	return &getJokesByIDHandler{
		facts: facts,
	}
}

func NewRandomJokeHandler(facts *models.Facts, metrics *metrics.AppMetricsCollectors) *getRandomJokeHandler {
	return &getRandomJokeHandler{
		facts:   facts,
		metrics: metrics,
	}
}

func NewGetJokesCountHandler(facts *models.Facts) *getJokesCountHandler {
	return &getJokesCountHandler{
		facts: facts,
	}
}

// Handle the GetJokeById operation
// /api/jokes/{id}
func (h *getJokesByIDHandler) Handle(params operations.GetJokeByIDParams) middleware.Responder {
	fact := h.facts.Facts[params.ID]

	if envVar := os.Getenv("DEPLOYMENT_ENV"); envVar != "" {
		fact.Joke = fmt.Sprintf("[%s] %s", envVar, fact.Joke)
	}

	return operations.NewGetJokeByIDOK().WithPayload(fact.ToFactResponse(params.ID, getOrElse(params.FirstName, defaultFirstName), getOrElse(params.LastName, defaultLastName)))
}

// Handle the GetRandomJoke operation
// /api/jokes/random
func (h *getRandomJokeHandler) Handle(params operations.GetRandomJokeParams) middleware.Responder {
	id, fact := h.facts.GetRandomFact()

	if envVar := os.Getenv("DEPLOYMENT_ENV"); envVar != "" {
		fact.Joke = fmt.Sprintf("[%s] %s", envVar, fact.Joke)
	}

	h.metrics.RandomJokesCounter.WithLabelValues().Inc()

	return operations.NewGetRandomJokeOK().WithPayload(fact.ToFactResponse(id, getOrElse(params.FirstName, defaultFirstName), getOrElse(params.LastName, defaultLastName)))
}

// Handle the GetJokesCount operation
// /api/jokes/count
func (h *getJokesCountHandler) Handle(params operations.GetJokesCountParams) middleware.Responder {
	factsLength := h.facts.Length()
	return operations.NewGetJokesCountOK().WithPayload(&respModels.CountResponse{
		Count: &factsLength,
	})
}

func getOrElse(value *string, fallback string) *string {
	if value == nil {
		return &fallback
	}
	return value
}
