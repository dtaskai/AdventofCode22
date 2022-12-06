package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var max = 0
	var sum = 0
	var topThree = []int{0, 0, 0}
	var topThreeSum = 0

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal("Error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() != "" {
			num, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatalf("Error converting string to int [value={%s}]", scanner.Text())
			}
			sum += num
		} else {
			i := Min(topThree)
			if sum > topThree[i] {
				topThree[i] = sum
			}
			if sum > max {
				max = sum
			}
			sum = 0
		}
	}

	for _, v := range topThree {
		topThreeSum += v
	}

	fmt.Printf("Part One: %d\n", max)
	fmt.Printf("Part Two: %d\n", topThreeSum)

}

func Min(array []int) int {
	minimum := array[0]
	index := 0
	for i, v := range array[1:] {
		if v < minimum {
			minimum = v
			index = i + 1
		}
	}
	return index
}
