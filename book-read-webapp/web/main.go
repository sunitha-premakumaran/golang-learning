package main

import (
	"encoding/json"
	"fmt"
	"os"

	cyoa "../cyoa"
)

func main() {
	f, err := os.Open("gopher.json")

	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(f)
	var story cyoa.Story
	if err := decoder.Decode(&story); err != nil {
		panic(err)
	}
	fmt.Printf("%v", story)
}
