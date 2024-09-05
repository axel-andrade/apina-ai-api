package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Match struct {
	Message      string `json:"message"`
	ShortMsg     string `json:"shortMessage"`
	Offset       int    `json:"offset"`
	Length       int    `json:"length"`
	Replacements []struct {
		Value string `json:"value"`
	} `json:"replacements"`
}

type LanguageToolResponse struct {
	Matches []Match `json:"matches"`
}

type LanguageToolHandler struct{}

func BuildLanguageToolHandler() *LanguageToolHandler {
	return &LanguageToolHandler{}
}

func (e *LanguageToolHandler) CorrectText(text string) (string, error) {
	apiURL := "https://api.languagetool.org/v2/check"

	data := url.Values{}
	data.Set("text", text)
	data.Set("language", "en-US")

	resp, err := http.Post(apiURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		return "", fmt.Errorf("unexpected content type: %s, body: %s", resp.Header.Get("Content-Type"), string(body))
	}

	var response LanguageToolResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("error parsing JSON response: %v, body: %s", err, string(body))
	}

	correctedText := []rune(text)
	offsetCorrection := 0
	for _, match := range response.Matches {
		if len(match.Replacements) > 0 {
			replacement := []rune(match.Replacements[0].Value)
			correctedText = append(correctedText[:match.Offset+offsetCorrection], append(replacement, correctedText[match.Offset+offsetCorrection+match.Length:]...)...)
			offsetCorrection += len(replacement) - match.Length
		}
	}

	return string(correctedText), nil
}
