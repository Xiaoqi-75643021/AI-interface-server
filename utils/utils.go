package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

func GetAccessToken(baseUrl, apiKey, secretKey string) (string, error) {
    data := url.Values{}
    data.Set("grant_type", "client_credentials")
    data.Set("client_id", apiKey)
    data.Set("client_secret", secretKey)

    fullUrl := baseUrl + "/oauth/2.0/token?" + data.Encode()

    req, err := http.NewRequest("POST", fullUrl, nil)
    if err != nil {
        return "", err
    }
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var result map[string]interface{}
    err = json.Unmarshal(body, &result)
    if err != nil {
        return "", err
    }

    token, ok := result["access_token"].(string)
    if !ok {
        return "", errors.New("access_token not found or is not a string")
    }

    return token, nil
}
