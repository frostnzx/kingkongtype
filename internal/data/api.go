package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"kingkongtype/internal/domain"
	"log"
	"net/http"
	"time"
)

type QuoteJson struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

func fetchQuote(difficulty string) (*domain.Quote, error) {
	url := "https://thequoteshub.com/api/random-quote" // public quote api
	client := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if res.Body != nil {
		defer res.Body.Close()
		var quote QuoteJson
		body, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
			return nil, readErr
		}
		jsonErr := json.Unmarshal(body, &quote)
		if jsonErr != nil {
			log.Fatal(jsonErr)
			return nil, jsonErr
		}
		fmt.Println("text :", quote.Text, " author : ", quote.Author)
		return &domain.Quote{
			Text:   quote.Text,
			Author: quote.Author,
		}, nil
	}
	return nil, errors.New("Error : No response body")
}
