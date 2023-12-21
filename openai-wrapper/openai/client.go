package openai

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
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
func (c *Client) Chat(requestPayload *ChatRequest, logger *log.Logger) (*ChatResponse, error) {
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
        logger.Printf("Error making request to OpenAI: %v", err)
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
        logger.Printf("Error reading response body: %v", err)
        return nil, err
    }

    // logger.Printf("OpenAI response body: %s", body)

    // Log the response to the log file using the passed logger
    // logResponse(&chatResponse, logger)

    return &chatResponse, nil
}

func logResponse(chatResponse *ChatResponse, logger *log.Logger) {
    // Log each choice's content
    for _, choice := range chatResponse.Choices {
        logger.Printf("Model: %s | Content: %s | Time: %s\n", chatResponse.Model, choice.Message.Content, time.Now().Format(time.RFC3339))
    }
}
