package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ex2() {

	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var s int
	for _, line := range lines {

		m, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		s = s + calculFuelRequired(m)
	}

	fmt.Printf("Suma este %d", s)

}
func readallLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func calculFuelRequired(x int) int {
	var sum int
	for x > 0 {
		x = x/3 - 2
		if x < 0 {
			break
		}
		sum += x
	}

	return sum
}
