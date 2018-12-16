package models

import (
	"encoding/json"
	"fmt"
	"github.com/gobuffalo/packr"
	"html"
	"math/rand"
	respModels "github.com/baez90/go-icndb/models"
)

type Fact struct {

	// joke
	// Required: true
	Joke       string   `json:"joke"`

	// categories
	// Required: true
	Categories []string `json:"categories"`
}

func (f *Fact) ToFactResponse(id int64, firstName *string, lastName *string) *respModels.FactResponse {
	return &respModels.FactResponse{
		ID: id,
		Joke: fmt.Sprintf(f.Joke, html.EscapeString(*firstName), html.EscapeString(*lastName)),
		Categories: f.Categories,
	}
}

type Facts struct {
	Facts []Fact
}

func (facts Facts) Length() int64 {
	return int64(len(facts.Facts))
}

func (facts Facts) GetRandomFact() (int64, Fact) {
	idx := rand.Int63n(int64(facts.Length()))
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