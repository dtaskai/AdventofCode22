package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var overlappingPartOne int
	var overlappingPartTwo int
	var e1s, e1e, e2s, e2e int

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal("Error opening input file.")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Sscanf(line, "%d-%d,%d-%d", &e1s, &e1e, &e2s, &e2e)

		if (e1s <= e2s && e1e >= e2e) || (e1s >= e2s && e1e <= e2e) {
			overlappingPartOne++
		}

		if e1s <= e2e && e1e >= e2s {
			overlappingPartTwo++
		}

	}

	fmt.Println(overlappingPartOne)
	fmt.Println(overlappingPartTwo)
}
