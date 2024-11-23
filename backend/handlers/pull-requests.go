package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gofr.dev/pkg/gofr"
)

func FetchPullRequests(ctx *gofr.Context)(interface{}, error) {
	const (
		owner = "gofr-dev"
		repo = "gofr"
		apiVersion = "2022-11-28"
	)

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls?state=all", owner, repo)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println("Error creating request: " + err.Error())
	}

	token := os.Getenv("GITHUB_TOKEN")

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-GitHub-Api-Version", apiVersion)

	client := &http.Client{}

	response, err :=client.Do(req)

	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", response.StatusCode)
	}

	// Read and display the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	log.Println(string(body))

	return string(body), nil
}
