package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"
    "strings"
    "time"
)

func readInt64(path string) int64 {
    b, err := os.ReadFile(path)
    if err != nil {
        return 0
    }
    s := strings.TrimSpace(string(b))
    v, _ := strconv.ParseInt(s, 10, 64)
    return v
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
    rx := readInt64("/sys/class/net/eth0/statistics/rx_bytes")
    tx := readInt64("/sys/class/net/eth0/statistics/tx_bytes")
    fmt.Fprintf(w, "aegisflux_network_rx_bytes %d\n", rx)
    fmt.Fprintf(w, "aegisflux_network_tx_bytes %d\n", tx)
}

func main() {
    http.HandleFunc("/metrics", metricsHandler)
    log.Println("AegisFlux Agent running on :9100")
    log.Fatal(http.ListenAndServe(":9100", nil))
}
