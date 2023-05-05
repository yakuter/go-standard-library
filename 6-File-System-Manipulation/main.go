package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Remove a directory recursively
	err := os.RemoveAll("mydir")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new directory
	err = os.Mkdir("mydir", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Write some data to a file
	data := []byte("hello world\n")
	err = os.WriteFile("mydir/myfile.txt", data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Rename a file
	err = os.Rename("mydir/myfile.txt", "mydir/newfile.txt")
	if err != nil {
		log.Fatal(err)
	}

	// List the contents of a directory
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}

	// Recursively walk a directory
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// Remove a file
	err = os.Remove("mydir/newfile.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Remove a directory recursively
	err = os.RemoveAll("mydir")
	if err != nil {
		log.Fatal(err)
	}
}
