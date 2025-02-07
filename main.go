package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/api/classify-number", classifyNumberHandler)

	//activate CORS
	mWithCORS := cors.Default().Handler(m)
	log.Println("Server running on port 3000...")

	http.ListenAndServe(":3000", mWithCORS)
}
