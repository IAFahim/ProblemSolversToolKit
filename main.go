package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Sites struct {
	Site []struct {
		Name    string   `json:"name"`
		Domains []string `json:"domains"`
		Dom     struct {
			Title  string `json:"title"`
			Input  string `json:"input"`
			Output string `json:"output"`
		} `json:"dom"`
	} `json:"site"`
}

func main() {
	content, _ := ioutil.ReadFile("DataBase.json")

	var cod Sites

	err := json.Unmarshal(content, &cod)

	if err != nil {
		fmt.Println(err)
	}
	print(cod.Site[0].Domains[0])
}
