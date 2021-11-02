package main

import "fmt"

func hello() {
	fmt.Println("Hello123")
}

func who() {
	fmt.Println("Go")
}

func main() {
	defer who()
	hello()
	defer fmt.Println("Hello")
	fmt.Println("World")
}
