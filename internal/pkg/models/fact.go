package models

import (
	"encoding/json"
	"fmt"
	respModels "github.com/baez90/go-icndb/models"
	packr "github.com/gobuffalo/packr/v2"
	"html"
	"math/rand"
)

type Fact struct {
	// joke
	// Required: true
	Joke string `json:"joke"`

	// categories
	// Required: true
	Categories []string `json:"categories"`
}

func (f *Fact) ToFactResponse(id int64, firstName *string, lastName *string) *respModels.FactResponse {
	formattedJoke := fmt.Sprintf(f.Joke, html.EscapeString(*firstName), html.EscapeString(*lastName))
	return &respModels.FactResponse{
		ID:         id,
		Joke:       &formattedJoke,
		Categories: f.Categories,
	}
}

type Facts struct {
	Facts []Fact
}

func (facts Facts) Length() int64 {
	return int64(len(facts.Facts))
}

func (facts Facts) GetRandomFact() (idx int64, fact Fact) {
	idx = rand.Int63n(int64(facts.Length()))
	fact = facts.Facts[idx]
	return
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
