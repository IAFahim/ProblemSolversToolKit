package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type domSelector struct {
	Title  string
	Input  []string
	Output []string
}

type site struct {
	Name    string `json:"name"`
	domains []string
	dom     domSelector
}

func getSites() {
	content, err := ioutil.ReadFile("DataBase.json")
	if err != nil {
		panic(err)
	}
	data := site{}

	err = json.Unmarshal(content, &data)
	if err != nil {
		print(err)
	}
	fmt.Println(data)
}
