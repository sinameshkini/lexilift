package ollama

import (
	"context"
	"errors"
	"fmt"
	"resty.dev/v3"
	"strings"
)

type Client struct {
	rc    *resty.Client
	model string
}

// Request defines the payload for the generate endpoint
type Request struct {
	Model   string `json:"model"`
	Prompt  string `json:"prompt"`
	Stream  bool   `json:"stream"`
	Options struct {
		Temperature float64 `json:"temperature"`
	} `json:"options"`
}

// Response defines the structure of the response from Ollama
type Response struct {
	Response string `json:"response"`
}

// New url example "http://localhost:11434/api"
func New(url, model string, debug bool) *Client {
	rc := resty.New().SetBaseURL(url).SetDebug(debug)

	return &Client{
		rc:    rc,
		model: model,
	}
}

func (c *Client) GetRandomWord(ctx context.Context, existed []string) (string, error) {
	var (
		resp     Response
		excluded = strings.Join(existed, ", ")
		prompt   = fmt.Sprintf(
			"Provide one random, interesting English word. "+
				"It MUST NOT be any of these: [%s]. "+
				"Respond with ONLY the word.", excluded,
		)
		//prompt  = "Provide a new random, interesting English word. Respond with ONLY the word, no punctuation or sentences."
		reqBody = Request{
			Model:  c.model,
			Prompt: prompt,
			Stream: false,
			Options: struct {
				Temperature float64 `json:"temperature"`
			}{Temperature: 1.0},
		}
	)

	r, err := c.rc.R().WithContext(ctx).
		SetBody(reqBody).
		SetResult(&resp).
		Post("/generate")
	if err != nil {
		return "", err
	}

	if r.IsError() {
		return "", errors.New(r.String())
	}

	return strings.TrimSpace(resp.Response), nil
}
