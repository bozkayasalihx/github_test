package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
)

type Response struct {
    Message string `json:"message"`
    Status  string `json:"status"`
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        response := Response{
            Message: "Hello, World!",
            Status:  "success",
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    })

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        response := Response{
            Message: "OK",
            Status:  "healthy",
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(response)
    })

    log.Printf("Server starting on port %s...", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}
