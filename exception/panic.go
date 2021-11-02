package main

import "fmt"

func executePanic() {
	panic("This is Panic Situation")
	fmt.Println("The function executes Completely")
}

func main() {
	executePanic()
	fmt.Println("Main block is executed completely...")
}
