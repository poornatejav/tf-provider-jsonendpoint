package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "sync"
)

var (
    mu       sync.Mutex
    lastData []byte
)

func submitHandler(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "read error", 500)
        return
    }
    mu.Lock()
    lastData = body
    mu.Unlock()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "created"})
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "read error", 500)
        return
    }
    mu.Lock()
    lastData = body
    mu.Unlock()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "updated"})
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    lastData = nil
    mu.Unlock()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
}

func fetchHandler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    data := lastData
    mu.Unlock()
    w.Header().Set("Content-Type", "application/json")
    if data == nil {
        json.NewEncoder(w).Encode(map[string]string{"error": "no data"})
    } else {
        w.Write(data)
    }
}

func main() {
    http.HandleFunc("/api/item/123", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodPost:
            submitHandler(w, r)
        case http.MethodPut:
            updateHandler(w, r)
        case http.MethodDelete:
            deleteHandler(w, r)
        case http.MethodGet:
            fetchHandler(w, r)
        default:
            http.Error(w, "method not allowed", 405)
        }
    })
    log.Println("Mock server listening on http://localhost:9000")
    log.Fatal(http.ListenAndServe(":9000", nil))
}
