package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
)

// proxy запускает TCP-листенер на localAddr и проксирует трафик к remoteAddr.
func proxy(localAddr, remoteAddr string) {
	ln, err := net.Listen("tcp", localAddr)
	if err != nil {
		log.Fatalf("не удалось слушать %s: %v", localAddr, err)
	}
	log.Printf("TCP proxy %s → %s", localAddr, remoteAddr)

	for {
		clientConn, err := ln.Accept()
		if err != nil {
			log.Printf("Accept error on %s: %v", localAddr, err)
			continue
		}
		go func() {
			defer clientConn.Close()

			serverConn, err := net.Dial("tcp", remoteAddr)
			if err != nil {
				log.Printf("не удалось подключиться к %s: %v", remoteAddr, err)
				return
			}
			defer serverConn.Close()

			go io.Copy(serverConn, clientConn)
			io.Copy(clientConn, serverConn)
		}()
	}
}

func main() {
	raw := os.Getenv("PROXIES")
	if raw == "" {
		log.Fatal("не задана переменная окружения PROXIES")
	}
	for _, part := range strings.Split(raw, ",") {
		parts := strings.Split(part, "=")
		if len(parts) != 2 {
			log.Fatalf("неверный формат прокси: %s", part)
		}
		local := parts[0]
		remote := parts[1]
		go proxy(local, remote)
	}

	select {}
}
