package main

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHealthEndpoint(t *testing.T) {
    req, err := http.NewRequest("GET", "/health", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        response := Response{
            Message: "OK",
            Status:  "healthy",
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(response)
    })

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    var response Response
    if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
        t.Fatal(err)
    }

    if response.Status != "healthy" {
        t.Errorf("handler returned unexpected status: got %v want %v",
            response.Status, "healthy")
    }
}
