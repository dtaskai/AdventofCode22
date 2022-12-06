package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	ROCK     int = 1
	PAPER        = 2
	SCISSORS     = 3
	WIN          = 6
	DRAW         = 3
	LOSS         = 0
)

func main() {
	var scorePartOne int
	var scorePartTwo int

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal("Error opening input file.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hands := strings.Split(scanner.Text(), " ")
		scorePartOne += calculateResultPartOne(hands[0], hands[1])
		scorePartTwo += calculateResultPartTwo(hands[0], hands[1])
	}

	fmt.Printf("Part One: %d\n", scorePartOne)
	fmt.Printf("Part Two: %d\n", scorePartTwo)
}

func calculateResultPartOne(elf string, player string) int {
	switch elf {
	case "A":
		switch player {
		case "X":
			return ROCK + DRAW
		case "Y":
			return PAPER + WIN
		case "Z":
			return SCISSORS + LOSS
		}
	case "B":
		switch player {
		case "X":
			return ROCK + LOSS
		case "Y":
			return PAPER + DRAW
		case "Z":
			return SCISSORS + WIN
		}
	case "C":
		switch player {
		case "X":
			return ROCK + WIN
		case "Y":
			return PAPER + LOSS
		case "Z":
			return SCISSORS + DRAW
		}
	}
	return 0
}

func calculateResultPartTwo(elf string, result string) int {
	switch elf {
	case "A":
		switch result {
		case "X":
			return SCISSORS + LOSS
		case "Y":
			return ROCK + DRAW
		case "Z":
			return PAPER + WIN
		}
	case "B":
		switch result {
		case "X":
			return ROCK + LOSS
		case "Y":
			return PAPER + DRAW
		case "Z":
			return SCISSORS + WIN
		}
	case "C":
		switch result {
		case "X":
			return PAPER + LOSS
		case "Y":
			return SCISSORS + DRAW
		case "Z":
			return ROCK + WIN
		}
	}
	return 0
}
