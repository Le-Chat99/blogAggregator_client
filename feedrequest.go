package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func postFeedsReq(name, url, apiKey string) (Res, error) {
	feedData := map[string]string{
		"name": name,
		"url":  url,
	}

	// Convert the feed data to JSON format
	jsonData, err := json.Marshal(feedData)
	if err != nil {
		return Res{}, fmt.Errorf("Error marshalling JSON: %s\n", err)
	}

	// Define the endpoint
	ReqUrl := "http://localhost:8080/v1/feeds"

	// Create the HTTP request
	req, err := http.NewRequest("POST", ReqUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return Res{}, fmt.Errorf("Error creating request: %s\n", err)
	}

	// Set the content type to application/json
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "ApiKey "+apiKey)

	// Execute the request using an HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Res{}, fmt.Errorf("Error sending request: %s\n", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode == http.StatusCreated {
		fmt.Println("Feed created successfully!")
	} else {
		fmt.Printf("Failed to create feed: %d\n", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Res{}, fmt.Errorf("Error read body: %s\n", err)
	}

	var feed Res
	err = json.Unmarshal(body, &feed)
	if err != nil {
		return Res{}, fmt.Errorf("Error unmarshal: %s\n", err)
	}
	return feed, nil
}
