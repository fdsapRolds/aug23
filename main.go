package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	newlog := log.New(os.Stdout, "account-api", log.LstdFlags)
	newHandler := handlers.NewHandler(newlog)

	myServeMux := http.NewServeMux()
	myServeMux.Handle("/", newHandler)

	http.ListenAndServe(":9090", nil)
}
