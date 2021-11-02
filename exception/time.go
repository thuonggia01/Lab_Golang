package main

import (
	"fmt"
	"time"
)

func main() {
	tNow := time.Now()
	Stime := tNow.Format(time.RFC3339)
	fmt.Println("String to Time")
	timeCo, _ := time.Parse(time.RFC3339, Stime)
	fmt.Printf("Type: %T \n", Stime)
	fmt.Printf("Time : %v - > Type : %T\n", timeCo, timeCo)
	fmt.Printf("===========================================")

	fmt.Println("Add")
	stime1 := "2021-10-18T22:08:41+00:00"

	stime1C, _ := time.Parse(time.RFC3339, stime1)
	stime2 := stime1C.AddDate(0, 2, 0)
	fmt.Println("time 1:", stime1C.Format("02/01/2006"))
	fmt.Println("time 2:", stime2.Format("02/01/2006"))

	fmt.Println("Minus")
	fmt.Println("time 1:", stime1C.Format("02/01/2006"))
	stime2c := stime2.AddDate(0, -2, 0)
	fmt.Println("time 3:", stime2c.Format("02/01/2006"))
}
