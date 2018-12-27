package main

import (
	"fmt"
	"os"
)

func main() {

	jsonFile, err := os.Open("user.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully Opened user.json")

	defer jsonFile.Close()
}
