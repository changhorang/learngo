package main // src에서 main.go로 실행 가능

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type result struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request is failed!")

func main() {
	chl := make(chan result)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitURL(url, chl)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-chl)
	}
}

func hitURL(url string, chl chan result) {
	status := "OK"
	fmt.Println("Checking: ", url)

	resp, err := http.Get(url)
	fmt.Println(err, resp.StatusCode)

	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	chl <- result{url: url, status: status}
}

func goCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "GO COUNT", i)
		time.Sleep(time.Second)
	}

}

func goCnt(person string, chl chan string) {
	time.Sleep(time.Second)
	fmt.Println(person)
	chl <- person + " go count"
}
