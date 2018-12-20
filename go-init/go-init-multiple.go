package main

import "fmt"

var initCounter int

func init() {
	fmt.Println("Called First in Order of Declaration")
	initCounter++
}

func init() {
	fmt.Println("Called Second in Order of Declaration")
	initCounter++
}

func main() {
	fmt.Println("Does nothing of any significance")
	fmt.Printf("Init Counter: %d\n", initCounter)
}
