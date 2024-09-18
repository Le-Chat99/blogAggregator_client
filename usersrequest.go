package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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
