package flags

import (
	"os"
	"strings"
	"testing"
)

type MockFlags struct {
	Links               []string
	NumberOfConnections uint
}

// Test if CLI is getting arguments properly (for test it's vk.com/ac and 5 parallel connections)
func Test_GetFlags(t *testing.T) {
	os.Args = []string{"./main", "vk.com/ac"}

	flags := InitFlags()
	err := flags.GetFlags()

	if err != nil {
		t.Fatal("failed to set up test")
	}

	if flags.Links[0] != "vk.com/ac" {
		t.Fatal("link argument was captured incorrectly")
	}

	// default number of connections is 10 - otherwise throw error
	if flags.NumberOfConnections != uint(10) {
		t.Fatal("default -conn flag argument was captured incorrectly")
	}
}

// Test if links without scheme is adding scheme properly
func Test_CheckAndFixURLs(t *testing.T) {
	var links []string

	mock := MockFlags{
		NumberOfConnections: 10,
	}

	links = append(links, "https://vk.com/ac", "ozon.ru", "avito.ru", "lamoda.ru")
	mock.Links = links

	opts := InitFlags()
	opts.NumberOfConnections = mock.NumberOfConnections
	opts.Links = mock.Links

	opts.CheckAndFixURLs()

	for _, link := range opts.Links {
		if !strings.Contains(link, "https://") {
			t.Fatal("schemes weren't added to links") // if scheme wasn't added to link -> throw error
		}
	}
}
