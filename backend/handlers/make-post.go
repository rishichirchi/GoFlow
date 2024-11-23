package handlers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gofr.dev/pkg/gofr"
)

// MakeTweet creates a new tweet using Twitter API v2 with OAuth 1.0a User Context
func MakeTweet(ctx *gofr.Context) (interface{}, error) {

	// Retrieve the content of the tweet from the request
	tweetContent := ctx.Request.Param("content")
	if tweetContent == "" {
		return nil, errors.New("content parameter is required")
	}

	backendUrl := "http://localhost:3000/tweets" + "?content=" + tweetContent

	// Make a POST request to the backend
	resp, err := http.Get(backendUrl)
	if err != nil {
		log.Fatalf("Error making the GET request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Print the response status and body
	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Response Body: %s\n", body)

	return body, nil
}
