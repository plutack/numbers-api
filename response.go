package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// successResponse contains fields for a 200 OK response
type successResponse struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	FunFact    string   `json:"fun_fact"`
}

// successResponse contains fields for a 400 BadRequest response
type errorResponse struct {
	Number string `json:"number"`
	Error  bool   `json:"error"`
}

// sendResponse is an helper function to write certain headers into a response.
func sendResponse(w http.ResponseWriter, statusCode int, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		log.Printf("failed to encode json response: %s", err)
	}
}

// newSuccessResponse creates an instance of successResponse.
//
// It makes use of go routines to run all functions asyncronously
// to improve the speed of generating a response.
//
// Returns successResponse
func newSuccessResponse(n int) successResponse {
	var wg sync.WaitGroup
	var mu sync.Mutex
	response := successResponse{Number: n}

	// Running computations concurrently
	wg.Add(4)

	go func() {
		defer wg.Done()
		isPrime := isPrime(n)
		mu.Lock()
		response.IsPrime = isPrime
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		isPerfect := isPerfect(n)
		mu.Lock()
		response.IsPerfect = isPerfect
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		properties := computeProperties(isArmstrong(n), isOdd(n))
		mu.Lock()
		response.Properties = properties
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		funFact := getFunFact(n)
		mu.Lock()
		response.FunFact = funFact
		mu.Unlock()
	}()

	wg.Wait()
	return response
}
