package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func cat(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	fmt.Println(string(content))
	return nil
}


func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage: go run cat.go [file1] [file2] ....")
		os.Exit(1)
	}

	for _, filename := range args {
		err := cat(filename)
		if err != nil {
			fmt.Println("Error reading %s: %s\n", filename, err)
		}
	}
}
