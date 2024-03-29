package main

import (
	"log"
	"net/http"
	"os"
	"time-go/internal/timeapi"
)

func main() {
	http.HandleFunc("/time", timeapi.GetTime) // Use the GetTime function from the timeapi package

	certFilePath := os.Getenv("CERT_FILE_PATH")
	keyFilePath := os.Getenv("KEY_FILE_PATH")

	var err error
	if certFilePath == "" || keyFilePath == "" {
		err = http.ListenAndServe(":8080", nil)
	} else {
		err = http.ListenAndServeTLS(":8080", certFilePath, keyFilePath, nil)
	}
	log.Fatal(err)
}
