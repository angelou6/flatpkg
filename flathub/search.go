package flathub

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const flathubUrl = "https://flathub.org/api/v2/search?locale=en"

type body struct {
	Query string `json:"query"`
}

func Search(name string) (Results, error) {
	searchBody, _ := json.Marshal(body{Query: name})
	res, err := http.Post(flathubUrl, "application/json", bytes.NewBuffer(searchBody))
	if err != nil {
		return Results{}, err
	}

	var results Results
	err = json.NewDecoder(res.Body).Decode(&results)
	if err != nil {
		return Results{}, err
	}

	return results, nil
}
