// baidu/client.go
package baidu

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	BaseURL     string
	AccessToken string
}

func NewClient(baseURL, accessToken string) *Client {
	return &Client{
		BaseURL:     baseURL,
		AccessToken: accessToken,
	}
}

func (c *Client) Chat(messages []Message) (*ChatResponse, error) {
	// url := c.BaseURL + "/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions_pro?access_token=" + c.AccessToken
	url := c.BaseURL + "/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions_pro?access_token=" + c.AccessToken
	requestBody, err := json.Marshal(ChatRequest{Messages: messages})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
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
