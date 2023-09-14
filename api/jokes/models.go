package jokes

import "fmt"

type rawJoke struct {
	Joke       string   `json:"joke"`
	Categories []string `json:"categories"`
}

func (j rawJoke) ToDTO(id int, firstName, lastName string) *FactResponse {
	return &FactResponse{
		ID:         id,
		Joke:       fmt.Sprintf(j.Joke, firstName, lastName),
		Categories: j.Categories,
	}
}

type CountResponse struct {
	Count int `json:"count"`
}

type FactResponse struct {
	ID         int      `json:"id,omitempty"`
	Joke       string   `json:"joke"`
	Categories []string `json:"categories"`
}
