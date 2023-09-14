package jokes

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

var (
	_ httprouter.Handle = Handler{}.GetJokesCount
	_ httprouter.Handle = Handler{}.GetRandomJoke
	_ httprouter.Handle = Handler{}.GetJokeByID
)

func SetupRouter(r *httprouter.Router, firstNameFallback, lastNameFallback string) {
	handler := Handler{
		firstNameFallback: firstNameFallback,
		lastNameFallback:  lastNameFallback,
		random:            rand.New(rand.NewSource(time.Now().Unix())),
	}

	r.GET("/api/joke/:id", handler.GetJokeByID)
	r.GET("/api/jokes/count", handler.GetJokesCount)
	r.GET("/api/jokes/random", handler.GetRandomJoke)
}

type Handler struct {
	random            *rand.Rand
	firstNameFallback string
	lastNameFallback  string
}

func (h Handler) GetJokesCount(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	encoder := json.NewEncoder(writer)

	if err := encoder.Encode(CountResponse{Count: jokes.Count()}); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (h Handler) GetRandomJoke(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	encoder := json.NewEncoder(writer)

	if err := encoder.Encode(jokes.Random(h.getNameValues(request))); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (h Handler) GetJokeByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := getParam(params, "id", strconv.Atoi)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	firstName, lastName := h.getNameValues(request)

	joke := jokes.ById(id, firstName, lastName)
	if joke == nil {
		http.Error(writer, "joke not found", http.StatusNotFound)
		return
	}

	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(joke); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (h Handler) getNameValues(request *http.Request) (firstName string, lastName string) {
	firstName = h.firstNameFallback
	lastName = h.lastNameFallback

	query := request.URL.Query()
	if query.Has("firstName") {
		firstName = query.Get("firstName")
	}

	if query.Has("lastName") {
		lastName = query.Get("lastName")
	}
	return firstName, lastName
}

func getParam[T any](params httprouter.Params, key string, parse func(v string) (T, error)) (T, error) {
	var parsed T
	val := params.ByName(key)
	if val == "" {
		return parsed, fmt.Errorf("required parameter %s is missing", key)
	}

	return parse(val)
}
