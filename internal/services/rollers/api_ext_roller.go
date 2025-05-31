package rollers

import (
	"encoding/json"
	"io"
	"net/http"
)

// Внешний генератор (API)
type APIRoller struct {
	url string
}

func NewAPIRoller(apiURL string) *APIRoller {
	return &APIRoller{url: apiURL}
}

func (a *APIRoller) Roll() int {
	resp, err := http.Get(a.url)
	if err != nil {
		return -1
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var result struct{ Value int }
	json.Unmarshal(body, &result)
	return result.Value
}

func (m *APIRoller) Name() string {
	return "APIRoller"
}
