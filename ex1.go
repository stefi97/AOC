package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Ex1 face ceva
func Ex1() {
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

		s = s + calculFuel(m)
	}

	fmt.Printf("Suma este %d", s)
}

func readLines(path string) ([]string, error) {
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

func calculFuel(masa int) int {

	m := masa/3 - 2
	return m
}
