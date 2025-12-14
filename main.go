package main

import (
	"fmt"
	"os"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting current working director:", err)
		return
	}
	fmt.Println("current working directory: ", cwd)
}
