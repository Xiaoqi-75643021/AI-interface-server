package handlers

import (
    "chat-api-server/config"
	"chat-api-server/ernie-bot-wrapper/baidu"
	"chat-api-server/openai-wrapper/openai"
	"chat-api-server/utils"
    
	"net/http"
	"encoding/json"
	"bytes"
	"io"
    
	"fmt"
)

func ChatHandler(cfg *config.Configuration) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
            return
        }

        // This struct is used to determine which service to route to
        type ServiceRequest struct {
            Service         string `json:"service"`
            InterfaceName   string `json:"interface_name"`
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

        // Find the service config for the requested service
        serviceConfig, err := findServiceConfig(cfg, serviceReq.Service, serviceReq.InterfaceName)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        switch serviceReq.Service {
        case "OpenAI":
            OpenAIChatHandler(serviceConfig, w, handlerBody)
        case "Baidu":
            BaiduChatHandler(serviceConfig, w, handlerBody)
        default:
            http.Error(w, "Service not supported", http.StatusBadRequest)
        }
    }
}

// findServiceConfig finds the service config for a given service and interface name
func findServiceConfig(cfg *config.Configuration, serviceName, interfaceName string) (*config.Service, error) {
    for _, api := range cfg.APIs {
        if api.Name == serviceName {
            for _, service := range api.Categories["Chat"] {
                if service.InterfaceName == interfaceName {
                    return &service, nil
                }
            }
        }
    }
    return nil, fmt.Errorf("service or interface not found")
}

// OpenAIChatHandler handles chat requests for the OpenAI service
func OpenAIChatHandler(serviceConfig *config.Service, w http.ResponseWriter, body io.ReadCloser) {
    var chatRequest openai.ChatRequest
    if err := json.NewDecoder(body).Decode(&chatRequest); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Check if the requested model is in the list of available models
    if _, exists := serviceConfig.ModelsList[chatRequest.Model]; !exists {
        http.Error(w, "Model not supported", http.StatusBadRequest)
        return
    }

    client := openai.NewClient(serviceConfig.APIKey, serviceConfig.APIURL)
    chatResponse, err := client.Chat(&chatRequest)
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
func BaiduChatHandler(serviceConfig *config.Service, w http.ResponseWriter, body io.ReadCloser) {
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

    accessToken, err := utils.GetAccessToken(serviceConfig.APIURL, serviceConfig.APIKey, serviceConfig.SecretKey)
    if err != nil {
        http.Error(w, "Error getting access token", http.StatusInternalServerError)
        return
    }

    client := baidu.NewClient(serviceConfig.APIURL, accessToken)
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
