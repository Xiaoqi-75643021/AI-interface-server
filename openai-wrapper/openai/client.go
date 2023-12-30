package openai

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// Client is a struct that will hold any necessary information for making requests to OpenAI
type Client struct {
	APIKey string
	APIURL string
}

// NewClient creates a new instance of Client
func NewClient(apiKey string, apiURL string) *Client {
	return &Client{APIKey: apiKey, APIURL: apiURL}
}

// Chat sends a request to the OpenAI API and returns the response
func (c *Client) Chat(requestPayload *ChatRequest) (*ChatResponse, error) {
	requestBody, err := json.Marshal(requestPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.APIURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", "Bearer "+c.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var chatResponse ChatResponse
	err = json.Unmarshal(body, &chatResponse)
	if err != nil {
		return nil, err
	}

	return &chatResponse, nil
}