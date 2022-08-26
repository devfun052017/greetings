package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	})
	for i := 0; i < 10000; i++ {
		time.Sleep(5)
		test()
	}
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func test() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Print(resp)
}
