package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err) //calls panic
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func main() {
	f := "tmp"
	result, _ := ReadFile(f)
	fmt.Println("Result: %v", result)
}
