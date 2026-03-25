# Concurrent Web Scrapper

## Description
This is a backend development challenge to practice my knowledge about Golang concurrency programming. This project will covers anything that I have learnt before in Golang concurrency programming about : 
- Goroutines
- Channels
- Wait groups
- Select statement
- Context
- Mutexes
I hope, by doing this challenge, I could improve my understanding about Golang concurrency programming.

## Scenario
You are building a backend service for a search engine. You have a list of 20 website URLs that need to be scraped for their text. However, network requests are slow, so doing it one by one is unacceptable. You need to build a concurrent crawler.

### The Requirements
- The Target Data: You have a list of 20 simulated URLs (e.g., "http://website-1.com", "http://website-2.com", etc.).
- The Shared State: You need to track the totalWords scraped across all websites. This is a single integer that all workers will share.
- The Worker Pool: You must spawn exactly 4 worker goroutines.
- The Work Simulation: * When a worker receives a URL, it should simulate a network request by sleeping for 600 milliseconds.
- Every website successfully scraped yields exactly 500 words.
- The worker must safely add those 500 words to the totalWords counter.
- The Hard Deadline (Timeout): Your search engine guarantees a response to the user in 2 seconds max. You must use a Context to enforce this. If the 2 seconds run out, all workers must immediately drop what they are doing, print a shutdown message, and return.
- Clean Shutdown: Your main function must wait for all 4 workers to finish (either because they scraped everything, or because they were killed by the timeout) before it prints the final totalWords count.

### The Checklist of Tools
To complete this successfully, your code must contain at least one of each of the following:

- go (to spawn the workers)
- chan (a buffered channel to hold the 20 URLs)
- select (to listen to the channel OR the timeout)
- context.WithTimeout (for the 2-second deadline)
- sync.Mutex (to protect the totalWords integer)
- sync.WaitGroup (to keep the main function from exiting too early)

## Expected Behavior
Since 4 workers each take 600ms per website, they can scrape about 4 websites each in 2.4 seconds. But your timeout is 2 seconds! This means your workers will not be able to finish all 20 URLs. They should process about 12 websites total, add 6,000 words to the counter, catch the timeout signal, print that they are shutting down, and exit safely.