package main

import (
	"context"
	"fmt"
	"net/url"
	"sync"
	"time"
)

type Website struct {
	URL string
}

func generateRandomWebsites(websites chan Website) {
	for i := 1; i <= 20; i++ {
		websites <- Website{URL: fmt.Sprintf("https://website.example-%d.com", i)}
	}
}

func worker(workerId int, websites chan Website, ctx context.Context, wg *sync.WaitGroup, mu *sync.Mutex, totalWords *int) {
	defer wg.Done()

	// Listen for incoming websites queue
	for {
		select {
		case website, ok := <-websites:

			// Validate website queue
			if !ok {
				fmt.Printf("Worker %d sees no more website to scrapped. Closing worker!", workerId)
				return
			}

			// Validate request url
			u, err := url.ParseRequestURI(website.URL)
			if err != nil {
				fmt.Printf("Worker %d received an invalid url. Skipping process!", workerId)
			}

			time.Sleep(time.Millisecond * 600)

			mu.Lock()
			*totalWords += 500
			fmt.Printf("Worker %d successfully scrapping website: %s \n", workerId, u.Host)
			mu.Unlock()
		case <-ctx.Done():
			fmt.Printf("Worker %d shutting down! \n", workerId)
			return
		}
	}
}

func main() {

	// Initiate necessary variables
	var mu sync.Mutex
	var wg sync.WaitGroup
	var totalWords int

	// Initiate channels for website queue
	websites := make(chan Website, 20)

	// Initiate context
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)

	defer cancel()

	// Assign random websites to queue
	generateRandomWebsites(websites)

	// Assign scrapper workers
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go worker(i, websites, ctx, &wg, &mu, &totalWords)
	}

	wg.Wait()
	fmt.Println("------------------------------")
	fmt.Printf("Suceed scrapping websites, got %d words total \n", totalWords)
}
