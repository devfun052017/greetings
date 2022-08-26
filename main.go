package main

import (
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("https://www.google.com")
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(resp)
}
