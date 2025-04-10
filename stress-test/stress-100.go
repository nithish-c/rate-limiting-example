package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func fetchMessage(url string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		results <- fmt.Sprintf("Error fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		results <- fmt.Sprintf("Error reading body from %s: %v", url, err)
		return
	}

	results <- string(body)
}

func main() {
	url := "http://localhost:8080/hello" // Replace with your API endpoint
	numRequests := 100                   // Number of requests to make

	results := make(chan string, numRequests)
	var wg sync.WaitGroup

	startTime := time.Now()

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go fetchMessage(url, results, &wg)
	}

	wg.Wait()
	close(results)

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	fmt.Println("Stress test completed.")
	fmt.Println("Duration:", duration)
	fmt.Println("Results:")

	for result := range results {
		fmt.Println(result)
	}
}
