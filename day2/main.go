package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputFile = flag.String("input", "", "Input file path")

type set struct {
	red   int
	green int
	blue  int
}

func (s set) power() int {
	return s.red * s.green * s.blue
}

type game struct {
	name    int
	reveals []set
}

var bag = set{red: 12, green: 13, blue: 14}

func main() {
	flag.Parse()

	input, err := os.Open(*inputFile)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	var (
		invalidGames []game
		minBags      []set
	)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		game := parseGame(scanner.Text())
		if game.isValid(bag) {
			invalidGames = append(invalidGames, game)
		}
		minBags = append(minBags, game.minBag())
	}

	var sumInvalid int
	for _, game := range invalidGames {
		sumInvalid += game.name
	}
	fmt.Println("sum of invalid games:", sumInvalid)

	var sumPowers int
	for _, minBag := range minBags {
		sumPowers += minBag.power()
	}
	fmt.Println("sum of powers:", sumPowers)
}

func parseGame(line string) game {
	lineParts := strings.Split(line, ": ")

	var game game
	game.name, _ = strconv.Atoi(strings.Split(lineParts[0], " ")[1])

	for _, revealDescr := range strings.Split(lineParts[1], "; ") {
		colorsDescr := strings.Split(revealDescr, ", ")

		var s set
		for _, colorDescr := range colorsDescr {
			colorParts := strings.Split(colorDescr, " ")

			v, _ := strconv.Atoi(colorParts[0])

			switch colorParts[1] {
			case "red":
				s.red = v
			case "green":
				s.green = v
			case "blue":
				s.blue = v
			}
		}

		game.reveals = append(game.reveals, s)
	}

	return game
}

func (g game) isValid(bag set) bool {
	for _, reveal := range g.reveals {
		if reveal.red > bag.red || reveal.green > bag.green || reveal.blue > bag.blue {
			return false
		}
	}
	return true
}

func (g game) minBag() set {
	var res set
	for _, reveal := range g.reveals {
		if res.red < reveal.red {
			res.red = reveal.red
		}
		if res.green < reveal.green {
			res.green = reveal.green
		}
		if res.blue < reveal.blue {
			res.blue = reveal.blue
		}
	}
	return res
}
