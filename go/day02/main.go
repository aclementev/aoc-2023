package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_RED   = 12
	MAX_GREEN = 13
	MAX_BLUE  = 14
)

func main() {
	// err := solve1("sample.txt")
	err := solve1("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to solve part 1: %v", err)
		os.Exit(1)
	}

	// err = solve2("sample.txt")
	err = solve2("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to solve part 2: %v", err)
		os.Exit(1)
	}
}

func solve1(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// var numbers []int

	var sum int

	for scanner.Scan() {
		text := scanner.Text()
		gameid, valid, err := checkLine(text, part1Verifier)
		if err != nil {
			return err
		}
		if valid {
			sum += gameid
		}
	}

	fmt.Println(sum)

	return nil
}

func checkLine(line string, verifierFunc func(string, int) bool) (int, bool, error) {
	// Example line: Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	gameParts := strings.Split(line, ":")
	game_part, rest := gameParts[0], gameParts[1]

	parts := strings.Split(game_part, " ")
	gameID, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, false, nil
	}

	// Check the colors
	for _, hand := range strings.Split(rest, ";") {
		for _, cube := range strings.Split(hand, ",") {
			cubeParts := strings.Split(strings.TrimSpace(cube), " ")
			num, err := strconv.Atoi(cubeParts[0])
			if err != nil {
				return 0, false, nil
			}
			color := cubeParts[1]
			if !verifierFunc(color, num) {
				return gameID, false, nil
			}
		}
	}

	return gameID, true, nil
}

func part1Verifier(color string, value int) bool {

	max := map[string]int{
		"red":   MAX_RED,
		"green": MAX_GREEN,
		"blue":  MAX_BLUE,
	}[color]
	return value > 0 && value <= max
}

func solve2(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// var numbers []int

	var sum int

	for scanner.Scan() {
		text := scanner.Text()
		power, err := gameMinPower(text)
		if err != nil {
			return err
		}
		sum += power
	}

	fmt.Println(sum)

	return nil
}

func gameMinPower(line string) (int, error) {
	// Example line: Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	gameParts := strings.Split(line, ":")
	_, rest := gameParts[0], gameParts[1]

	// Check the colors
	maximums := map[string]int64{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, hand := range strings.Split(rest, ";") {
		for _, cube := range strings.Split(hand, ",") {
			cubeParts := strings.Split(strings.TrimSpace(cube), " ")
			numInt, err := strconv.Atoi(cubeParts[0])
			if err != nil {
				return 0, nil
			}
			num := int64(numInt)
			color := cubeParts[1]
			if oldMax := maximums[color]; num > oldMax {
				maximums[color] = num
			}
		}
	}

	// Compute the power
	var power int64 = 1

	for _, value := range maximums {
		power *= value
	}

	return int(power), nil
}
