package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// err := solve1("sample.txt")
	err := solve1("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to solve part 1: %v", err)
		os.Exit(1)
	}

	// err = solve2("sample1.txt")
	// err = solve2("sample2.txt")
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

	var numbers []int

	for scanner.Scan() {
		text := scanner.Text()

		// Process the line
		var first, last int
		for _, char := range text {
			if unicode.IsDigit(char) {
				last, _ = strconv.Atoi(string(char))
				if first == 0 {
					first = last
				}
			}
		}
		// We have the first and second digit of the number
		num := first*10 + last
		numbers = append(numbers, num)
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	sum := 0
	for _, n := range numbers {
		sum += n
	}

	fmt.Println(sum)

	return nil
}

// NOTE(alvaro): The names of the numbers can be overlapping, so when parsing a number
// that is spelled out, you cannot skip the full word, and instead must check the rest
// of the string rune by rune
func solve2(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var numbers []int

	for scanner.Scan() {
		text := scanner.Text()

		// Process the line
		rest := []rune(text)
		var first, last int
		for {
			// Try to parse a number
			num := parseNumber(rest)
			if num != 0 {
				if first == 0 {
					first = num
				}
				last = num
			}

			rest = rest[1:]

			// Check if we are done processing
			if len(rest) == 0 {
				break
			}
		}

		if first == 0 || last == 0 {
			panic("Did not find a number on this line")
		}
		num := first*10 + last
		// fmt.Printf("%s: %d %d -> %d\n", text, first, last, num)
		numbers = append(numbers, num)
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	sum := 0
	for _, n := range numbers {
		sum += n
	}

	fmt.Println(sum)

	return nil
}

func parseNumber(input []rune) int {
	// Check if the first part of the input is a number
	for _, r := range input {
		if unicode.IsDigit(r) {
			num, _ := strconv.Atoi(string(r))
			// fmt.Printf("Found digit: %v\n", num)
			return num
		}
		break
	}

	words := []string{
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	// Check if the current string is pointing to a number spelled as a word
	rest := string(input)
	for i, numWord := range words {
		if strings.HasPrefix(rest, numWord) {
			// fmt.Printf("Found spelled digit: %v (%v)\n", numWord, i+1)
			return i + 1
		}
	}

	// Nothing matched, so we don't return a thing
	// NOTE(alvaro): The zero is not a possible parsed value, so we use it to mark
	// failure to parse
	return 0
}
