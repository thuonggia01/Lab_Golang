package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	b, err := ioutil.ReadFile("./todos.json")
	if err != nil {
		log.Print(err)
	}
	fmt.Println(string(b))

}
