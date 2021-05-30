package requests

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// SendHTTPRequests and get responses hashes
// pass max number of connections
// pass array of links from flags
func SendHTTPRequests(connections uint, links []string) (map[string]string, error) {
	var linkAndHashes = map[string]string{}

	for i := uint(0); i != connections; i++ {
		wg := sync.WaitGroup{}

		for _, link := range links {
			wg.Add(1)

			go func(link string) {
				defer wg.Done()

				response, err := http.Get(link)
				if err != nil {
					log.Fatalf("Error in http request: %s", err)
				}

				body, err := ioutil.ReadAll(response.Body)
				if err != nil {
					log.Fatalf("Error in ReadAll: %s", err)
				}

				hasher := md5.New()
				hasher.Write(body)

				linkAndHashes[link] = hex.EncodeToString(hasher.Sum(nil))
			}(link)
		}
		wg.Wait()
	}

	return linkAndHashes, nil
}
