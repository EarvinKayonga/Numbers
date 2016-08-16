package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func toNumSlice(data string) []int64 {
	var numbers []int64
	var tabString = strings.Split(data, " ")
	if len(tabString) > 0 {
		for _, item := range tabString {
			if item != "" {
				number, err := strconv.ParseInt(item, 10, 64)
				if err != nil {
					log.Fatal(err)
				}
				numbers = append(numbers, number)
			}
		}
	}

	return (numbers)
}

func decode(numbers []int64) {
	for _, number := range numbers {
		letter := number % 26
		fmt.Printf(Letter(letter))
	}
}

/*
	Takes an INT between 0 and 25, return letter from the alphabet
*/
func Letter(number int64) string {
	return alphabet[number]
}
