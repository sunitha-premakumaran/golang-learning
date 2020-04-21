package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	cyoa "../cyoa"
)

func main() {
	f, err := os.Open("gopher.json")
	if err != nil {
		panic(err)
	}
	story, error := cyoa.JSONStory(f)
	if error != nil {
		panic(err)
	}
	handler := cyoa.NewHandler(story)
	fmt.Println("Starting server on :3000")
	log.Fatal(http.ListenAndServe(":7070", handler))
}
