package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mt sync.Mutex

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	successfullEndpoints := []string{}
	if err != nil {
		fmt.Printf("Can not get to %s", endpoint)
	} else {

		mt.Lock()
		successfullEndpoints = append(successfullEndpoints, endpoint)
		mt.Unlock()

		fmt.Printf("Response received from %s is %d\n", endpoint, res.StatusCode)
	}
}

func main() {
	endpoints := []string{
		"https://www.google.com",
	}

	for _, endpoint := range endpoints {
		go getStatusCode(endpoint)
		wg.Add(1)
	}

	wg.Wait()
}
