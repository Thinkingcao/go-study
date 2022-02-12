package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Name  string  `json:"name"`
	Page  int     `json:page`
	Price float32 `json:price`
}

func main() {
	book := Book{"go in actin", 345, 99.99}
	jsonStr, err := json.Marshal(book)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)
	myBook := Book{}
	json.Unmarshal(jsonStr, myBook)
	if err != nil {
		fmt.Println("json unmarshal error ", err)
		return
	}
	fmt.Printf("%v\n", myBook)
}
