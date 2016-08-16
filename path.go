package main

import (
	"log"
	"os"
)

func getFilePath() string {
	var filePath string = "numbers.txt"
	if len(os.Args) == 2 {
		filePath = os.Args[1]
	}
	return (filePath)
}

func read(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Read " + string(count) + " bytes")
	return string(data)
}
