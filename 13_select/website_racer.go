package websiteracer

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

// func WebsiteRacer(url1, url2 string) string {
// 	url1ResTime := measureTime(url1)
// 	url2ResTime := measureTime(url2)

// 	if url1ResTime > url2ResTime {
// 		return url2
// 	} else {
// 		return url1
// 	}

// }

// func WebsiteRacer(url1, url2 string) string {
// 	var url1ResTime time.Duration
// 	var url2ResTime time.Duration
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go func() {
// 		defer wg.Done()
// 		url1ResTime = measureTime(url1)
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		url2ResTime = measureTime(url2)
// 	}()

// 	wg.Wait()
// 	if url1ResTime > url2ResTime {
// 		return url2
// 	} else {
// 		return url1
// 	}

// }
func WebsiteRacer(url1, url2 string) (string, error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(10 * time.Second):
		return "", errors.New("ERROR - 10 sec passed")
	}
}
func measureTime(url string) time.Duration {
	starTime := time.Now()
	res, err := http.Get(url)
	if err != nil {
	}
	defer res.Body.Close()
	fmt.Println("Error ", err)
	return time.Since(starTime)

}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		res, err := http.Get(url)
		if err != nil {
			return
		}
		defer res.Body.Close()
		close(ch)

	}()
	return ch
}
