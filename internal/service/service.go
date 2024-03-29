package word

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	mword "github.com/justarandomlearner/SignificAPI/internal/model"
	rword "github.com/justarandomlearner/SignificAPI/internal/repository"
)

type Service struct {
	rword.Repository
}

func NewService() Service {
	return Service{
		Repository: rword.Repository{
			// Conn: conn,
		},
	}
}

func (s Service) FindWord(ctx context.Context, word string) (*mword.Payload, error) {
	r, err := http.Get("https://api.dicionario-aberto.net/word/" + word)
	if err != nil {
		return nil, err
	}

	var results []mword.Result
	dec := json.NewDecoder(r.Body)

	err = dec.Decode(&results)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("word not found")
	}

	payload := &mword.Payload{Word: word}

	payload.Meanings = make([]mword.Meaning, 0)

	for _, result := range results {
		var e mword.Entry

		if err := xml.Unmarshal([]byte(result.XML), &e); err != nil {
			return nil, err
		}

		for _, sense := range e.Senses {
			var m mword.Meaning

			m.SenseToMeaning(sense)

			payload.Meanings = append(payload.Meanings, m)
		}
	}

	return payload, nil
}
