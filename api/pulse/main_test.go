package main

import (
	"fmt"
	"testing"
)

func TestGetHnNewestUrls(t *testing.T) {
	last := 10
	links := getHnNewestUrls(last)
	fmt.Println(links)

	if len(links) == 0 {
		t.Errorf("Expected at least one link, got 0")
	}
}
