package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	inputFile = flag.String("input", "", "Input file path")

	digitMap = map[string]int{
		"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	}
)

func main() {
	flag.Parse()

	input, err := os.Open(*inputFile)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	var codes []int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var (
			leftDigit  int
			leftIndex  = -1
			rightDigit int
			rightIndex = -1
		)

		for text, digit := range digitMap {
			if index := strings.Index(scanner.Text(), text); index != -1 {
				if leftIndex == -1 || index < leftIndex {
					leftDigit = digit
					leftIndex = index
				}
			}

			if index := strings.LastIndex(scanner.Text(), text); index != -1 {
				if rightIndex == -1 || index > rightIndex {
					rightDigit = digit
					rightIndex = index
				}
			}
		}

		code := 10*leftDigit + rightDigit
		codes = append(codes, code)
	}

	var sum int
	for _, code := range codes {
		sum += code
	}
	fmt.Println(sum)
}
