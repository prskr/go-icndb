package main

import (
	"encoding/json"
	"github.com/gobuffalo/packr"
	"math/rand"
)

type CountResponse struct {
	Count int
}

type FactResponse struct {
	*Fact
	Id int `json:"id"`
}

type Fact struct {
	Joke       string   `json:"joke"`
	Categories []string `json:"categories"`
}

func (f *Fact) ToFactResponse(id int) *FactResponse {
	return &FactResponse{
		Id: id,
		Fact: f,
	}
}

type Facts struct {
	Facts []Fact
}

func (facts Facts) Length() int {
	return len(facts.Facts)
}

func (facts Facts) GetRandomFact() (int, Fact) {
	idx := rand.Intn(facts.Length())
	return idx, facts.Facts[idx]
}

func LoadFacts(box *packr.Box, key string) (*Facts, error) {
	var facts []Fact

	bytes, err := box.Find(key)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &facts)

	if err != nil {
		return nil, err
	}

	return &Facts{
		Facts: facts,
	}, nil
}
