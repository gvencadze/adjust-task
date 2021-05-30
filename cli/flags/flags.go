package flags

import (
	"errors"
	"flag"
	"log"
	"net/url"
)

// Flags struct with program arguments
type Flags struct {
	Links               []string
	NumberOfConnections uint
}

// InitFlags with default number of connections
func InitFlags() *Flags {
	return &Flags{
		NumberOfConnections: 10,
	}
}

// GetFlags from program arguments
func (r *Flags) GetFlags() error {
	var connections uint

	flag.UintVar(&connections, "p", 10, "Number of parallel connections")
	flag.Parse()

	hosts := flag.Args()

	if len(hosts) == 0 {
		return errors.New("an argument is required")
	}

	r.Links = hosts
	r.NumberOfConnections = connections

	return nil
}

// CheckAndFixURLs | add scheme if not presented
func (r *Flags) CheckAndFixURLs() {
	for id, link := range r.Links {
		var err error

		u, err := url.ParseRequestURI(link)

		if err != nil || u.Host == "" {
			u, err = url.ParseRequestURI("https://" + link)
			if err != nil {
				log.Println(err)
			}

			err = nil
		}

		link = u.Scheme + "://" + u.Host + u.Path

		r.Links[id] = link
	}
}
