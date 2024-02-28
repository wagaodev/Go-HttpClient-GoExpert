package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	client := http.Client{Timeout: time.Second}
	result, err := client.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
