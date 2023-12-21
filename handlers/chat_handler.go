package handlers

import (
	"bytes"
	"chat-api-server/config"
	"chat-api-server/ernie-bot-wrapper/baidu"
	"chat-api-server/openai-wrapper/openai"
	"chat-api-server/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func ChatHandler(cfg *config.Configuration) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
            return
        }

        // This struct is used to determine which service to route to
        type ServiceRequest struct {
            Service string `json:"service"`
        }

        // Read the entire request body
        bodyBytes, err := io.ReadAll(r.Body)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        // Close the original body
        defer r.Body.Close()

        // Decode the first part of the request to determine the service
        var serviceReq ServiceRequest
        if err := json.Unmarshal(bodyBytes, &serviceReq); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        fmt.Println("Service requested: ", serviceReq.Service)

        // Create a new buffer with the same content for the specific handler
        handlerBody := io.NopCloser(bytes.NewBuffer(bodyBytes))

        switch serviceReq.Service {
        case "openai":
            // Call the OpenAI chat handler code with the new buffer
            OpenAIChatHandler(cfg, w, handlerBody)
        case "baidu":
            // Call the Baidu chat handler code with the new buffer
            BaiduChatHandler(cfg, w, handlerBody)
        default:
            http.Error(w, "Service not supported", http.StatusBadRequest)
        }
    }
}

// OpenAIChatHandler handles chat requests for the OpenAI service
func OpenAIChatHandler(cfg *config.Configuration, w http.ResponseWriter, body io.ReadCloser) {
    var chatRequest openai.ChatRequest
    if err := json.NewDecoder(body).Decode(&chatRequest); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Check if the requested model is in the list of available models
    if _, exists := cfg.OpenAI.ModelsList[chatRequest.Model]; !exists {
        http.Error(w, "Model not supported", http.StatusBadRequest)
        return
    }

    client := openai.NewClient(cfg.OpenAI.APIKey, cfg.OpenAI.APIURL)
    chatResponse, err := client.Chat(&chatRequest, log.New(os.Stderr, "LOG: ", log.LstdFlags))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    if err := json.NewEncoder(w).Encode(chatResponse); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// BaiduChatHandler handles chat requests for the Baidu service
func BaiduChatHandler(cfg *config.Configuration, w http.ResponseWriter, body io.ReadCloser) {
    // Define a new struct that matches the expected request body
    type BaiduServiceRequest struct {
        Messages []baidu.Message `json:"messages"`
    }

    var serviceRequest BaiduServiceRequest
    if err := json.NewDecoder(body).Decode(&serviceRequest); err != nil {
        http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
        return
    }

    messages := serviceRequest.Messages

    accessToken, err := utils.GetAccessToken(cfg.Baidu.BaseURL, cfg.Baidu.APIKey, cfg.Baidu.SecretKey)
    if err != nil {
        http.Error(w, "Error getting access token", http.StatusInternalServerError)
        return
    }

    client := baidu.NewClient(cfg.Baidu.BaseURL, accessToken)
    chatResponse, err := client.Chat(messages)
    if err != nil {
        http.Error(w, "Error during chat", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(chatResponse); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
