package main

import (
	"fmt"
	"net/http"
	"reflect"
	"time"
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

func process(any) {

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

	type Vector struct {
		x, y, z int
	}
	// Case 1 — stays on stack (compiler sees it doesn't escape)
	vs := Vector{1, 2, 3}
	// process(v) // passed by value, v never leaves this scope

	// Case 2 — escapes to heap (pointer leaves the scope)
	// v := Vector{1, 2, 3}
	p := &vs
	process(p) // now v's address is passed out, compiler moves v to heap

	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	fmt.Println(typeOfT, s, s.CanSet())
	// for i := 0; i < s.NumField(); i++ {
	// 	f := s.Field(i)
	// 	fmt.Printf("%d: %s %s = %v\n", i,
	// 		typeOfT.Field(i).Name, f.Type(), f.Interface())
	// }

}

// tcp connection
// func main() {
// 	// Start listening on port 8080
// 	listener, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		fmt.Println("Error starting server:", err)
// 		return
// 	}
// 	defer listener.Close()

// 	fmt.Println("Server listening on port 8080...")

// 	for {
// 		// Block here until a client connects
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			fmt.Println("Error accepting connection:", err)
// 			continue
// 		}

// 		fmt.Println("Client connected:", conn.RemoteAddr())

// 		// Handle the connection
// 		handleConnection(conn)
// 	}
// }

// func handleConnection(conn net.Conn) {
// 	defer conn.Close()

// 	// Read what the client sent
// 	buf := make([]byte, 1024)
// 	n, err := conn.Read(buf)
// 	if err != nil {
// 		fmt.Println("Error reading:", err)
// 		return
// 	}

// 	message := string(buf[:n])
// 	fmt.Println("Received:", message)

// 	// Write a response back
// 	conn.Write([]byte("Hello from server!\n"))
// }
