// main.go
package main

import (
    "log"
    "net/http"
    "chat-api-server/config"
    "chat-api-server/handlers"
)

func main() {
    cfg, err := config.LoadConfig("config/config.json")
    if err != nil {
        log.Fatal(err)
    }

    http.Handle("/chat", handlers.ChatHandler(cfg))

    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
