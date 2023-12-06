package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"unicode"
)

var inputFile = flag.String("input", "", "Input file path")

type symbol struct {
	ch       rune
	row, col int
}

type part struct {
	number  int
	symbols map[symbol]bool
}

func main() {
	flag.Parse()

	input, err := os.Open(*inputFile)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	var lines []string

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var missingParts []part
	for i, line := range lines {
		var p part

		for j, ch := range line {
			if !unicode.IsNumber(ch) {
				if len(p.symbols) > 0 {
					missingParts = append(missingParts, p)
				}
				p = part{}
				continue
			}

			p.number = p.number*10 + int(ch-'0')

			for k := i - 1; k <= i+1; k++ {
				for m := j - 1; m <= j+1; m++ {
					if k >= 0 && k < len(lines) && m >= 0 && m < len(line) {
						sym := symbol{
							ch:  rune(lines[k][m]),
							row: k,
							col: m,
						}
						if !unicode.IsNumber(sym.ch) && sym.ch != '.' {
							if p.symbols == nil {
								p.symbols = make(map[symbol]bool)
							}
							p.symbols[sym] = true
						}
					}
				}
			}
		}

		if len(p.symbols) > 0 {
			missingParts = append(missingParts, p)
		}
	}

	var sum int
	for _, p := range missingParts {
		sum += p.number
	}
	fmt.Println("parts sum:", sum)

	stars := make(map[symbol][]part)
	for _, p := range missingParts {
		for sym := range p.symbols {
			if sym.ch == '*' {
				stars[sym] = append(stars[sym], p)
			}
		}
	}

	var gearsSum int
	for _, pns := range stars {
		if len(pns) == 2 {
			gearsSum += pns[0].number * pns[1].number
		}
	}
	fmt.Println("gears sum:", gearsSum)
}
