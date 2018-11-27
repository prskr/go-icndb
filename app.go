package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type App struct {
	Router *mux.Router
	Facts  *Facts
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/health", a.getHealth).Methods("GET")
	a.Router.HandleFunc("/jokes/count", a.getJokeCount).Methods("GET")
	a.Router.HandleFunc("/jokes/random", a.getRandomJoke).Methods("GET")
	a.Router.HandleFunc("/joke/{id:[0-9]+}", a.getJoke).Methods("GET")
}

func (a *App) getHealth(w http.ResponseWriter, r *http.Request) {
	log.Println("/health got called")
	w.WriteHeader(200)
}

func (a *App) getRandomJoke(w http.ResponseWriter, r *http.Request) {
	id, fact := a.Facts.GetRandomFact()
	returnJson(w, 200, fact.ToFactResponse(id))
}

func (a *App) getJokeCount(w http.ResponseWriter, r *http.Request) {
	returnJson(w, 200, CountResponse{
		Count: a.Facts.Length(),
	})
}

func (a *App) getJoke(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		returnError(w, http.StatusBadRequest, "Invalid joke id")
	}
	if id < 0 || id >= a.Facts.Length() {
		returnError(w, http.StatusNotFound, "No joke with matching id present")
	}

	returnJson(w, 200, a.Facts.Facts[id].ToFactResponse(id))
}

func returnError(w http.ResponseWriter, code int, message string) {
	returnJson(w, code, map[string]string{"error": message})
}

func returnJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
