package main

import (
	"fmt"
	"testing"
)

func TestGetHnNewestUrls(t *testing.T) {
	links, err := getHnNewestPostsIds()
	if err != nil {
		t.Fatalf("failed to get HN newest posts IDs: %v", err)
	}
	fmt.Println(links)
}

func TestHnPostProcess(t *testing.T) {
	processHnPost("38701099")
}
