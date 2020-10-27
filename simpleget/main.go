package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values{ //쿼리 문자열 선언
		"query": {"Hello world"},
	}
	//url 끝에 문자열로 만든 결과를 추가한다.
	resp, err := http.Get("http://localhost:18888" + "?" + values.Encode()) //문자열로 만든다.
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	log.Println("Status:", resp.Status)
	log.Println("StatusCode:", resp.StatusCode)
	log.Println("Headers:", resp.Header)
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
}
