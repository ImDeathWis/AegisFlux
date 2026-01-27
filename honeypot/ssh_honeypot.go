package main

import (
    "log"
    "net"
    "os"
    "time"
)

func main() {
    ln, err := net.Listen("tcp", ":2222")
    if err != nil {
        log.Fatal(err)
    }
    log.Println("SSH Honeypot on port 2222")

    for {
        conn, _ := ln.Accept()
        go func(c net.Conn) {
            defer c.Close()
            ip := c.RemoteAddr().String()
            log.Println("[SSH Honeypot] Connection from:", ip)
            f, _ := os.OpenFile("honeypot_ssh.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
            f.WriteString(time.Now().Format(time.RFC3339) + " " + ip + "\n")
            f.Close()
            c.Write([]byte("SSH-2.0-OpenSSH_8.0 FakeServer\r\n"))
        }(conn)
    }
}
