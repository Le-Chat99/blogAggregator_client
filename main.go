package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Res struct {
	Feed Feed         `json:"feed"`
	FF   FeedFollowed `json:"feed_follow"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollowed struct {
	ID        uuid.UUID `json:"id"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func main() {
	us1, err := postUserReq("lane")
	if err != nil {
		fmt.Printf("fail add user: %v", err)
	}
	us2, err := postUserReq("Khiem")
	if err != nil {
		fmt.Printf("fail add user: %v", err)
	}
	fmt.Println(us1)
	fmt.Println(us2)
	us1get, err := getUserReq(us1.ApiKey)
	if err != nil {
		fmt.Printf("fail get user: %v", err)
	}
	fmt.Println(us1get)
	feed1, err := postFeedsReq("ff", "https://wagslane.dev/index.xml", us1.ApiKey)
	if err != nil {
		fmt.Printf("fail post feed: %v", err)
	}
	feed2, err := postFeedsReq("ff", "https://blog.boot.dev/index.xml", us2.ApiKey)
	if err != nil {
		fmt.Printf("fail post feed: %v", err)
	}
	fmt.Println(feed1)
	fmt.Println(feed2)
}

func postUserReq(name string) (User, error) {
	userData := map[string]string{
		"name": name,
	}

	// Convert the user data to JSON format
	jsonData, err := json.Marshal(userData)
	if err != nil {
		return User{}, fmt.Errorf("Error marshalling JSON: %s\n", err)
	}

	// Define the endpoint
	url := "http://localhost:8080/v1/users"

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return User{}, fmt.Errorf("Error creating request: %s\n", err)
	}

	// Set the content type to application/json
	req.Header.Set("Content-Type", "application/json")

	// Execute the request using an HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return User{}, fmt.Errorf("Error sending request: %s\n", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode == http.StatusCreated {
		fmt.Println("User created successfully!")
	} else {
		fmt.Printf("Failed to create user: %d\n", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return User{}, fmt.Errorf("Error read body: %s\n", err)
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return User{}, fmt.Errorf("Error unmarshal: %s\n", err)
	}
	return user, nil
}

func getUserReq(apiKey string) (User, error) {
	url := "http://localhost:8080/v1/users"

	// Create the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return User{}, fmt.Errorf("Error geting request: %s\n", err)
	}

	// Set the content type to application/json
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "ApiKey "+apiKey)

	// Execute the request using an HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return User{}, fmt.Errorf("Error sending request: %s\n", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode == http.StatusOK {
		fmt.Println("User got successfully!")
	} else {
		fmt.Printf("Failed to get user: %d\n", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return User{}, fmt.Errorf("Error read body: %s\n", err)
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return User{}, fmt.Errorf("Error unmarshal: %s\n", err)
	}
	return user, nil
}

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
