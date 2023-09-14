package jokes

import (
	_ "embed"
	"encoding/json"
	"math/rand"
	"time"
)

var (
	//go:embed jokes.json
	jokesRaw []byte
	jokes    = jokesRepo{
		random: rand.New(rand.NewSource(time.Now().Unix())),
	}
)

func init() {
	if err := json.Unmarshal(jokesRaw, &jokes); err != nil {
		panic(err)
	}
}

var _ json.Unmarshaler = (*jokesRepo)(nil)

type jokesRepo struct {
	jokes  []rawJoke
	random *rand.Rand
}

func (j *jokesRepo) ById(id int, firstName, lastName string) *FactResponse {
	if id < 0 || id >= len(j.jokes) {
		return nil
	}

	return j.jokes[id].ToDTO(id, firstName, lastName)
}

func (j *jokesRepo) Random(firstName, lastName string) *FactResponse {
	idx := int(j.random.Int31n(int32(len(j.jokes) - 1)))

	return j.jokes[idx].ToDTO(idx, firstName, lastName)
}

func (j *jokesRepo) Count() int {
	return len(j.jokes)
}

func (j *jokesRepo) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &j.jokes)
}
