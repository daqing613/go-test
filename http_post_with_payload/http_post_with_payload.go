package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {

	response, err := http.PostForm(
		"http://example.com/form",
		url.Values{
			"username": {"MyUserName"},
			"password": {"123"},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	log.Println(response.Header) // Print the response headers
	// To upload a file, use Post instead of PostForm, provide a content type like application/json
	// or application/octet-stream, and then provide the an io.Reader with the data

	// http.Post("https://example.com/upload", "image/jpeg", &buff)
}
