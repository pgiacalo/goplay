package main

import (
	"fmt"
	"os"
)

func createFile(name string) error {
	fmt.Println("Defer.createFile() called")
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func main() {
	fmt.Println("Defer.main() called")
	err := createFile("testfile.txt")
	if err != nil {
		fmt.Println("Defer.main() error: ", err)
	}
	fmt.Println("Defer.main() exiting")
}
