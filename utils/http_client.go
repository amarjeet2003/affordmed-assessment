package utils

import (
	"net/http"
	"time"
)

const BaseURL = "http://localhost:8080/test/companies/"

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func Get(url string) (*http.Response, error) {
	return httpClient.Get(url)
}
