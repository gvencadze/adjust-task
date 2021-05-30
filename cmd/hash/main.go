package main

import (
	"log"

	"github.com/gvencadze/adjust-task/cli/flags"
	"github.com/gvencadze/adjust-task/cli/requests"
)

func main() {
	conn, links := getArguments()

	URLsAndHashes, err := requests.SendHTTPRequests(conn, links)

	if err != nil {
		log.Fatalf("Error in sending http requests: %s", err.Error())
	}

	for key, value := range URLsAndHashes {
		log.Printf("%s %s\n", key, value)
	}
}

// get number of connections and links from flags
func getArguments() (uint, []string) {
	opts := flags.InitFlags()

	err := opts.GetFlags()
	if err != nil {
		return 0, nil
	}

	opts.CheckAndFixURLs()

	return opts.NumberOfConnections, opts.Links
}
