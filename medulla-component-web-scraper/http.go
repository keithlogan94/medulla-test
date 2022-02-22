package main

import (
	"io"
	"net/http"
)

func GetHtmlAt(webAddress string) string {
	resp, err := http.Get(webAddress)
	CheckError(err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		CheckError(err)
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	CheckError(err)
	return string(body)
}
