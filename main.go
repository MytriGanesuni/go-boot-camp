package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {

	urls := []string{
		"https://gdg.community.dev/gdg-cochin/",
		"https://golang.org",
		"https://httpstat.us/500",
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.twitter.com/",
		"https://www.instagram.com/",
		"https://site-not-present.io",
		"https://www.youtube.com/",
		"https://www.linkedin.com/",
		"https://www.github.com/",
		"https://www.stackoverflow.com/",
		"https://www.reddit.com/",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var wg sync.WaitGroup

	ch := make(chan string, len(urls))
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			statusCode, _ := checkURL(ctx, u)
			printStatus(u, statusCode)
			wg.Done()
			ch <- u
		}(url)
	}
	wg.Wait()

	select {
	case <-ctx.Done():
		fmt.Println("Context cancelled")
	case <-ch:
		if len(ch) == len(urls)-1 {
			fmt.Println("✓ All URLs checked")
		}
	}
}

func printStatus(url string, statusCode int) {
	if statusCode == 200 {
		fmt.Println("URL: ", url, "is up and running with status code: ", statusCode)
	} else {
		fmt.Println("URL: ", url, "is down with status code: ", statusCode)
	}
}

func checkURL(ctx context.Context, url string) (int, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
