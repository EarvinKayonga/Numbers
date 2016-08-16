package main

import (
	"strings"
	"strconv"
	"log"
)

func process(data string) {
	var numbers []int64
	var tabString = strings.Split(data, " ")
	if  len(tabString)> 0 {
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
}
