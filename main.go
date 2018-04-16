package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cloudfoundry/go-socks5"
)

func main() {
	server := createServer()

	go func() {
		addr := ":1080"

		if err := server.ListenAndServe("tcp", addr); err != nil {
			panic(err)
		}
	}()

	waitForSignal()
}

func createServer() *socks5.Server {
	// Set credentials
	user := os.Getenv("PROXY_USER")
	password := os.Getenv("PROXY_PASSWORD")

	if user == "" {
		panic("PROXY_USER env variable is not set")
	}

	if password == "" {
		panic("PROXY_PASSWORD env variable is not set")
	}

	credentials := socks5.StaticCredentials{user: password}

	// Create server
	conf := &socks5.Config{Credentials: credentials}

	server, err := socks5.New(conf)

	if err != nil {
		panic(err)
	}

	return server
}

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	log.Printf("Got signal: %v, exiting.", s)
}
