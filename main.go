package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	// Define the number of goroutines
	const numGoroutines = 30000

	// URL to make GET requests to
	const url = "http://34.64.117.93:9090/work/"

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Function to perform the GET request
	request := func(id int) {
		defer wg.Done()

		// Make the GET request
		resp, err := http.Get(url + time.Now().String())
		if err != nil {
			fmt.Printf("Goroutine %d: error making GET request: %v\n", id, err)
			return
		}

		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Goroutine %d: error reading response body: %v\n", id, err)
			return
		}

		fmt.Printf("[%s] Goroutine %d: received response with status %s - and body: %s\n", time.Now().String(), id, resp.Status, string(body))
	}

	// Start the goroutines
	for i := 0; i < numGoroutines; i++ {
		go request(i)
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Println("FIN")
	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All requests completed.")
}
