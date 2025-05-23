package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	endpoints := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
		"https://www.github.com",
	}

	for _, endpoint := range endpoints {
		go getStatusCode(endpoint)
		wg.Add(1)
	}

	wg.Wait()
}

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)
	defer wg.Done()
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Printf("%d is the status code for website %s\n", res.StatusCode, endpoint)
		defer res.Body.Close()
	}
}
