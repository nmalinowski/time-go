package main

import (
	"net/http"
	"time-go/internal/timeapi" // Assuming the package is in a directory named pkg
)

func main() {
	http.HandleFunc("/time", timeapi.GetTime) // Use the GetTime function from the timeapi package
	http.ListenAndServe(":8080", nil)
}
