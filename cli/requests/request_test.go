package requests

import (
	"testing"
)

// Test if http request is working properly
func Test_SendHTTPRequests(t *testing.T) {
	connections := uint(5)
	links := []string{"https://vk.com/ac"}

	hashes, err := SendHTTPRequests(connections, links)
	if err != nil {
		t.Fatal(err)
	}

	if hashes == nil {
		t.Fatal("No hash from response")
	}
}

func Test_SendInvalidHTTPRequests(t *testing.T) {
	connections := uint(1)
	links := []string{"https://zz.yy"}

	hashes, err := SendHTTPRequests(connections, links)
	if err != nil {
		t.Fatal(err)
	}

	if hashes == nil {
		t.Fatal("No hash from response")
	}
}
