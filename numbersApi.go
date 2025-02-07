package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// getFunFact is a function that gets a random math funfact using http://numbersapi.com/{n}/math
// where n is the number.
//
// Returns a string.
func getFunFact(n int) string {
	fmt.Printf("http://numbersapi.com/%d/math", n)
	resp, err := http.Get(fmt.Sprintf("http://numbersapi.com/%d/math", n))
	if err != nil {
		log.Printf("request error: %s", err)
		return "No fun fact available."
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read body error: %s", err)
		return "Error retrieving fun fact."
	}
	return string(body)
}
