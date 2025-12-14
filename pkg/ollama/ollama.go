package ollama

import (
	"context"
	"encoding/json"
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

type WritingEvaluation struct {
	GrammarCorrect bool     `json:"grammar_correct"`
	KeywordExists  bool     `json:"keyword_exists"`
	Score          int      `json:"score"` // 0-100
	Feedback       string   `json:"feedback"`
	Suggestions    []string `json:"suggestions"`
	BetterSentence string   `json:"better_sentence"`
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

func (c *Client) EvaluateWriting(ctx context.Context, keyword, userSentence string) (*WritingEvaluation, error) {
	var resp Response

	prompt := fmt.Sprintf(`
       Evaluate this English writing challenge:
       Target Keyword: "%s"
       User Sentence: "%s"

       Instructions:
       1. Check if the sentence is grammatically correct.
       2. Check if the target keyword is used correctly and naturally.
       3. Provide a 'better_sentence' which is a more natural, idiomatic, or sophisticated version of the user's sentence while keeping the same meaning and using the target keyword.
       4. Calculate a score from 0 to 100 based on accuracy and naturalness.
       5. Provide brief feedback and specific suggestions for improvement.

       Respond ONLY in JSON format with these exact keys: 
       "grammar_correct" (bool), 
       "keyword_exists" (bool), 
       "score" (int), 
       "better_sentence" (string), 
       "feedback" (string), 
       "suggestions" (array of strings).
    `, keyword, userSentence)

	reqBody := map[string]interface{}{
		"model":  c.model,
		"prompt": prompt,
		"stream": false,
		"format": "json",
		"options": map[string]interface{}{
			"temperature": 0.3, // Slightly increased for more creative "better sentences"
		},
	}

	r, err := c.rc.R().WithContext(ctx).
		SetBody(reqBody).
		SetResult(&resp).
		Post("/generate")
	if err != nil {
		return nil, err
	}

	if r.IsError() {
		return nil, errors.New(r.String())
	}

	// Parse the JSON string inside the Ollama response into our struct
	var evaluation WritingEvaluation
	err = json.Unmarshal([]byte(resp.Response), &evaluation)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AI response: %v", err)
	}

	return &evaluation, nil
}
