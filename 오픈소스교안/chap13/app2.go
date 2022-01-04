package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func responseSize(url string, channel chan WebPage) {
	//fmt.Println("URL 주소 ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s", body)
	//fmt.Println(len(body))
	channel <- WebPage{URL: url, Size: len(body)}
}

type WebPage struct {
	URL  string
	Size int
}

func main() {
	pages := make(chan WebPage)
	urls := []string{"https://www.inhatc.ac.kr", "https://www.naver.com", "https://www.daum.net", "https://www.nate.com"}

	for _, url := range urls {
		go responseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}
	// go responseSize("https://www.inhatc.ac.kr")
	// go responseSize("https://www.naver.com")
	// go responseSize("https://www.daum.net")
	// go responseSize("https://www.nate.com")
	//response, err := http.Get("https://www.inhatc.ac.kr")
	//time.Sleep(time.Second * 5) //이거 안주면 너무 빨리끝나서 아무것도 안뜸

}

// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"time"
// )

// func responseSize(url string) {
// 	fmt.Println("URL 주소 ", url)
// 	response, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer response.Body.Close()
// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//fmt.Printf("%s", body)
// 	fmt.Println(len(body))
// }

// func main() {
// 	go responseSize("https://www.inhatc.ac.kr")
// 	go responseSize("https://www.naver.com")
// 	go responseSize("https://www.daum.net")
// 	go responseSize("https://www.nate.com")
// 	//response, err := http.Get("https://www.inhatc.ac.kr")
// 	time.Sleep(time.Second * 5) //이거 안주면 너무 빨리끝나서 아무것도 안뜸

// }
