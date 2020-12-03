package main

import (
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/JackRamey/GoAdvent2020/pkg/iohelpers"
)

const target = 2020

func main() {
	// Determine filepath
	cwd, err := os.Getwd()
	check(err)
	filepath := path.Join(cwd, "1-2", "input", "data")
	// Read input
	lines, err := iohelpers.FileToLines(filepath)
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

	// For each key in the map, subtract the first value from the target to get the secondary target. This secondary target
	// Then search the map for that second value to see if it exists. Loop through keys until the second value is found.
	for firstValue, _ := range entries {
		secondaryTarget := target - firstValue
		entriesMinusKey := seedMapWithExclusion(values, firstValue)
		secondValue, thirdValue, found := getOperandsForTarget(entriesMinusKey, secondaryTarget)
		if found {
			fmt.Printf("%d + %d + %d = %d\n", firstValue, secondValue, thirdValue, firstValue+secondValue+thirdValue)
			fmt.Printf("%d * %d * %d = %d\n", firstValue, secondValue, thirdValue, firstValue*secondValue*thirdValue)
			println("Solution:")
			println(firstValue * secondValue * thirdValue)
			return
		}
	}
}

func getOperandsForTarget(entries map[int]struct{}, target int) (minuend int, subtrahend int, found bool) {
	// For each key in the map, subtract the first value from the target to get the wanted value.
	// Then search the map for that wanted value to see if it exists. Loop through keys until the second value is found.
	for first, _ := range entries {
		wanted := target - first
		// Check the entries map to see if second value exists
		if _, ok := entries[wanted]; ok {
			return first, wanted, true
		}
	}
	return 0, 0, false
}

func seedMap(values []int) map[int]struct{} {
	entries := make(map[int]struct{}, len(values))
	for _, value := range values {
		entries[value] = struct{}{}
	}
	return entries
}

func seedMapWithExclusion(values []int, exclusion int) map[int]struct{} {
	entries := make(map[int]struct{}, len(values))
	for _, value := range values {
		if value != exclusion {
			entries[value] = struct{}{}
		}
	}
	return entries
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
