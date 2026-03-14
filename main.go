package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var directory string
var defaultDirectory string = "."

var port int
var defaultPort int = 8000

func main() {
	port = defaultPort
	if s := os.Getenv("PORT"); s != "" {
		if p, err := strconv.Atoi(s); err == nil {
			port = p
		}
	}

	directory = defaultDirectory
	if s := os.Getenv("DIR"); s != "" {
		directory = s
	}

	log.Printf(`serving directory "%s" on port "%d"`, directory, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(directory))))
}
