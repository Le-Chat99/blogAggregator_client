package main

import (
	"fmt"
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
	fmt.Println("---------")
	fmt.Println(us1)
	fmt.Println("---------")
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
	fmt.Println("---------")
	fmt.Println(feed1)
	fmt.Println("---------")
	fmt.Println(feed2)
}
