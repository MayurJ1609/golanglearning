package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to creating JSON")
	encodeJson()
}

func encodeJson() {
	lcoCourses := []course{
		{"React JS Bootcamp", 299, "youtube.com", "abc123", []string{"web", "dev", "js"}},
		{"MERN Bootcamp", 199, "youtube.com", "abcd123", []string{"web", "dev", "js"}},
		{"NESTJS Bootcamp", 399, "youtube.com", "abcde123", nil},
	}

	// Package data as JSON
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
