package main

import (
	"fmt"
	"os"
)

func main() {
	content := make([]byte, 1000)
	file, error := os.Open(os.Args[1])

	if error != nil {
		fmt.Println("Error: ", error)
	}

	file.Read(content)

	fmt.Println(string(content))
}
