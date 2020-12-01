package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
)

const target = 2020

func main() {
	// Determine filepath
	cwd, err := os.Getwd()
	check(err)
	filepath := path.Join(cwd, "1-1", "input", "data")
	// Read input
	lines, err := fileToLines(filepath)
	check(err)

	// Parse values into ints
	values := make([]int, len(lines))
	for idx, line := range lines {
		value, err := strconv.Atoi(line)
		check(err)
		values[idx] = value
	}

	// Seed map with input
	entries := seedMap(values)

	firstValue, secondValue := getOperandsForTarget(entries, target)
	fmt.Printf("%d + %d = %d\n", firstValue, secondValue, firstValue+secondValue)
	fmt.Printf("%d * %d = %d\n", firstValue, secondValue, firstValue*secondValue)
	println("Solution:")
	println(firstValue * secondValue)

}

func getOperandsForTarget(entries map[int]struct{}, target int) (minuend int, subtrahend int) {
	// For each key in the map, subtract the first value from the target to get the wanted value.
	// Then search the map for that wanted value to see if it exists. Loop through keys until the second value is found.
	for first, _ := range entries {
		wanted := target - first
		// Check the entries map to see if second value exists
		if _, ok := entries[wanted]; ok {
			return first, wanted
		}
	}
	return 0,0
}

func seedMap(values []int) map[int]struct{} {
	entries := make(map[int]struct{}, len(values))
	for _, value := range values {
		entries[value] = struct{}{}
	}
	return entries
}

func fileToLines(filePath string) (lines []string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
