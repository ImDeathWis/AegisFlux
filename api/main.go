package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Event struct {
    Type string json:"type"
    Data string json:"data"
}

var events []Event

func main() {
    http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            var e Event
            json.NewDecoder(r.Body).Decode(&e)
            events = append(events, e)
            w.WriteHeader(201)
            json.NewEncoder(w).Encode(e)
            return
        }

        json.NewEncoder(w).Encode(events)
    })

    log.Println("API running on :9090")
    http.ListenAndServe(":9090", nil)
}
