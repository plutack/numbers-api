package main

import (
	"math/rand/v2"
	"net/http"
	"strconv"
)

// classifyNumberHandler is a handle function for the api route to to write an http response
func classifyNumberHandler(w http.ResponseWriter, r *http.Request) {
	numberString := r.URL.Query().Get("number")

	if numberString == "" {
		n := rand.IntN(1000) // if no number query is added, this randomly generates for a number between 1 and 1000
		sendResponse(w, http.StatusOK, newSuccessResponse(n))
		return
	}

	parsedInt, err := strconv.Atoi(numberString)
	if err != nil {
		sendResponse(w, http.StatusBadRequest, errorResponse{Number: numberString, Error: true})
		return
	}

	sendResponse(w, http.StatusOK, newSuccessResponse(parsedInt))
}
