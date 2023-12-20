package main

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestGetHnNewestUrls(t *testing.T) {
	links, err := getHnNewestPostsIds()
	if err != nil {
		t.Fatalf("failed to get HN newest posts IDs: %v", err)
	}
	fmt.Println(links)
}

func TestHnPostProcess(t *testing.T) {
	expected := []Link{
		{
			title: "Ask HN: What Should I Do with Broken Solar Panels?",
			url:   "item?id=38708816",
		},
	}
	data := processHnPost("38708816")
	assert.Equal(t, expected, data)
}
