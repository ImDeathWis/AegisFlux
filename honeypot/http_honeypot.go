package main

import (
    "log"
    "net/http"
    "os"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    ip := r.RemoteAddr
    ua := r.UserAgent()

    log.Printf("[HTTP Honeypot] Visit from %s UA=%s PATH=%s\n", ip, ua, r.URL.Path)

    f, _ := os.OpenFile("honeypot_http.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    f.WriteString(time.Now().Format(time.RFC3339) + " " + ip + " " + ua + " " + r.URL.Path + "\n")
    f.Close()

    w.Write([]byte("Hello from AegisFlux Honeypot"))
}

func main() {
    http.HandleFunc("/", handler)
    log.Println("HTTP Honeypot running on :8088")
    http.ListenAndServe(":8088", nil)
}
