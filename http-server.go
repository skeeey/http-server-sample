package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const ServerPortKey = "SERVER_PORT"

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Bonjour\n")
}

func main() {
	http.HandleFunc("/hello", hello)

	serverPort, existing := os.LookupEnv(ServerPortKey)
	if !existing {
		serverPort = "8000"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil))
}
