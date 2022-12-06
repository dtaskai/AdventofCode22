package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var scorePartOne int
	var scorePartTwo int
	var badgeGroup = [3]string{}
	var lineCount = 0

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal("Error opening input file.")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		scorePartOne += MapCharToScore(FindCommonItems(strings.Split(line[:len(line)/2], ""), strings.Split(line[len(line)/2:], ""))[0])

		badgeGroup[lineCount] = line
		if lineCount == 2 {
			common := FindCommonItems(strings.Split(badgeGroup[0], ""), strings.Split(badgeGroup[1], ""))
			scorePartTwo += MapCharToScore(FindCommonItems(common, strings.Split(badgeGroup[2], ""))[0])

			badgeGroup = [3]string{}
			lineCount = 0
		} else {
			lineCount++
		}
	}

	fmt.Println(scorePartOne)
	fmt.Println(scorePartTwo)
}

func FindCommonItems(group1 []string, group2 []string) []string {
	var result []string

	for _, x := range group1 {
	out:
		for _, y := range group2 {
			if x == y {
				result = append(result, x)
				break out
			}
		}
	}

	return result
}

func MapCharToScore(char string) int {
	if char == strings.ToUpper(char) {
		score := 27
		for i := 'A'; i <= 'Z'; i++ {
			if string(i) == char {
				return score
			}
			score++
		}
	} else {
		score := 1
		for i := 'a'; i <= 'z'; i++ {
			if string(i) == char {
				return score
			}
			score++
		}
	}
	return 0
}
