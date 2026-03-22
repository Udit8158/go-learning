package concurrency

import (
	"fmt"
	"time"
)

type WebsiteChecker func(string) bool
type resultUrl struct {
	url  string
	isUp bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	result := make(map[string]bool)
	s := time.Now()
	// concurrently run wc for each url
	// var wg sync.WaitGroup
	resultChannel := make(chan resultUrl)

	for _, url := range urls {
		go func(url string) {
			// we will write to channel instead of writting to map
			resultChannel <- resultUrl{
				url:  url,
				isUp: wc(url),
			}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		result[r.url] = r.isUp
	}

	// for _, url := range urls {
	// 	result[url] = wc(url)
	// }
	e := time.Since(s)
	fmt.Println("Took ", e)

	return result
}
