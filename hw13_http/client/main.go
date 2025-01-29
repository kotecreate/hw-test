package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	urlFlag     = flag.String("url", "", "url server")
	pathFlag    = flag.String("path", "/", "path resource")
	methodFlag  = flag.String("method", http.MethodGet, "HTTP method")
	contentType = "application/json"
	requestBody = []byte(`{"message":"hello"}`)
)

func sendGetRequest(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func sendPostRequest(url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	postBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return postBody, nil
}

func main() {
	flag.Parse()

	if *urlFlag == "" || (*methodFlag != http.MethodGet && *methodFlag != http.MethodPost) {
		flag.Usage()
		os.Exit(1)
	}

	url := fmt.Sprintf("%s%s", *urlFlag, *pathFlag)

	var responseBody []byte
	var err error

	switch *methodFlag {
	case http.MethodGet:
		responseBody, err = sendGetRequest(url)
	case http.MethodPost:
		responseBody, err = sendPostRequest(url, requestBody)
	}

	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Println(string(responseBody))
}
