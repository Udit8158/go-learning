package main

import (
	"fmt"
	"net/http"
	"time"

	concurrency "github.com/Udit8158/go-learning/12_concurrency"
)

func Change(arr []int) {
	arr[0] = 0
	arr = append(arr, 6)
	fmt.Println("Array changed, ", arr)
}

type F struct {
	name  string
	email string
}

func (f *F) nameChange(newName string) {
	fmt.Printf("Memory address of f Type F in nameChange method %p\n", f)
	fmt.Println((*f))
	f.name = newName
}

type RealSleeper struct{}

func (rs *RealSleeper) Sleep() {
	time.Sleep(time.Duration(1) * time.Second)
}
func checkUrl(url string) bool {
	res, err := http.Get(url)

	// http error occured
	if err != nil {
		fmt.Printf("HTTP ERROR - %v\n", err)
		return false
	}

	// getting the 200 ok response - so url exists
	// if res.StatusCode == http.StatusOK {
	// 	return true
	// }

	// fmt.Println("Not 200ok - status code,", res.StatusCode, url)
	return res.StatusCode < 500
}
func main() {

	// s := []int{1, 2, 3, 4, 5} // this a way to define a slice
	// s2 := make([]int, 5, 6)   // len -> 5, cap -> 8 slice (but this is better ig, more control)

	// // interesting
	// // s2 = s2[:6]
	// // s2[5] = 12

	// fmt.Println(custom_append.CustomAppend(s2, 12))
	// fmt.Printf("After custom append %p\n", &s2[0])

	// fmt.Println("Current cap", cap(s2))
	// s2 = append(s2, 100)
	// s2 = append(s2, 100)
	// s2 = append(s2, 100)
	// s2 = append(s2, 100)
	// s2 = append(s2, 100)
	// fmt.Println(s2, len(s2), cap(s2))
	// // Change(s)

	// fmt.Println(cap(s)) // capacity -> 5
	// s = append(s, 6)
	// fmt.Println(cap(s)) // capacity -> 10
	// s = append(s, 7)
	// fmt.Println(cap(s)) // capacity -> 10
	// fmt.Println(s)

	// var friend F
	// fmt.Printf("Memory address of friend %p\n", &friend)

	// fmt.Printf("%v\n", friend)
	// friend.name = "i"
	// friend.email = "k"
	// fmt.Printf("%#v\n", friend)

	// friend.nameChange("udit")
	// fmt.Printf("%#v\n", friend)
	//
	// dependencyinjetion.Greet(os.Stdout, "Udit")
	// os.Stdout.WriteString("hi there\n")

	// var sleeper RealSleeper
	// countdown.Countdown(os.Stdout, &sleeper)

	// websites := []string {
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"hello",
	// 	"https://google.com",
	// 	"https://google.com",
	// 	"https://google.com",
	// 	"https://google.com",
	// 	"https://google.com",
	// 	"https://google.com",
	// 	"https://google.com",
	// 	"https://google.com",
	// 	"https://google.com",
	// 	"https://google.com",
	// 	"https://google.com",
	// 	"https://google.com",
	// 	"https://google.com",
	// 	"https://google.com",
	// }
	urlData := concurrency.CheckWebsites(checkUrl, []string{"hello", "https://google.com", "https://notion.so", "https://google.com",
		"https://github.com",
		"https://golang.org",
		"https://go.dev",
		"https://stackoverflow.com",
		"https://openai.com",
		"https://wikipedia.org",
		"https://amazon.com",
		"https://microsoft.com",
		"https://apple.com"})
	fmt.Printf("URL DATA %#v\n", urlData)
}
