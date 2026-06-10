package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"kingkongtype/internal/domain"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type QuoteJson struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

func FetchPrompt(settings domain.Settings) (*domain.Quote, error) {
	if settings.Mode == 0 {
		return generateWordPrompt(settings.Difficulty), nil
	}

	return FetchQuote(settings.Difficulty)
}

func generateWordPrompt(difficulty int) *domain.Quote {
	words := wordsForDifficulty(difficulty)
	count := wordCountForDifficulty(difficulty)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	prompt := make([]string, count)

	for i := range prompt {
		prompt[i] = words[random.Intn(len(words))]
	}

	return &domain.Quote{
		Text:   strings.Join(prompt, " "),
		Author: "Word mode",
	}
}

func FetchQuote(difficulty int) (*domain.Quote, error) {
	var lastQuote *domain.Quote

	for i := 0; i < 5; i++ {
		quote, err := fetchRandomQuote()
		if err != nil {
			return nil, err
		}
		lastQuote = quote

		if IsDifficultyMatch(quote.Text, difficulty) {
			return quote, nil
		}
	}

	return lastQuote, nil
}

func IsDifficultyMatch(text string, difficulty int) bool {
	length := len(text)

	switch difficulty {
	case 0:
		return length <= 80
	case 1:
		return length > 80 && length <= 160
	case 2:
		return length > 160
	default:
		return true
	}
}

func wordsForDifficulty(difficulty int) []string {
	switch difficulty {
	case 0:
		return []string{
			"air", "and", "book", "cat", "day", "go", "home", "kind", "light", "make",
			"name", "open", "play", "read", "run", "see", "sun", "time", "tree", "walk",
		}
	case 2:
		return []string{
			"abstract", "balance", "capacity", "decision", "evidence", "fragment", "generate", "horizon", "incident", "journey",
			"language", "momentum", "navigate", "original", "practice", "question", "resource", "strategy", "terminal", "velocity",
		}
	default:
		return []string{
			"answer", "better", "change", "direct", "effect", "follow", "ground", "handle", "inside", "letter",
			"memory", "notice", "option", "people", "reason", "simple", "system", "typing", "window", "yellow",
		}
	}
}

func wordCountForDifficulty(difficulty int) int {
	switch difficulty {
	case 0:
		return 25
	case 2:
		return 75
	default:
		return 50
	}
}

func fetchRandomQuote() (*domain.Quote, error) {
	url := "https://thequoteshub.com/api/random-quote" // hard coded public quote api
	client := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.Body != nil {
		defer res.Body.Close()
		if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
			return nil, fmt.Errorf("quote api returned status %d", res.StatusCode)
		}

		var quote QuoteJson
		body, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			return nil, readErr
		}
		jsonErr := json.Unmarshal(body, &quote)
		if jsonErr != nil {
			return nil, jsonErr
		}
		return &domain.Quote{
			Text:   quote.Text,
			Author: quote.Author,
		}, nil
	}
	return nil, errors.New("Error : No response body")
}
