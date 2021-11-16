package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Infof("User: %s", getUsername())
	fmt.Println("sum 2 and 3 is", sum(2, 3))

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {})

	mux := http.NewServeMux()
	mux.HandleFunc("/status", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Service: 39_go_CI\nStatus: OK\n")
	})
	// Starting server to listen <port>
	s := http.Server{Addr: "0.0.0.0:8080", Handler: mux}
	log.Fatal(s.ListenAndServe())
}
