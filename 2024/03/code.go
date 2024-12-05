package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	if part2 {
		return p2(input)
	}

	return p1(input)
}

var OPRegex = regexp.MustCompile(`mul\(\d+,\d+\)`)
var OPRegex2 = regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don\'t\(\)`)

func p1(input string) any {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	var sum int64 = 0

	for _, line := range lines {
		ops := OPRegex.FindAllString(line, -1)

		for _, op := range ops {
			parts := strings.Split(op[4:len(op)-1], ",")

			a, err := strconv.Atoi(parts[0])

			if err != nil {
				panic(err)
			}

			b, err := strconv.Atoi(parts[1])

			if err != nil {
				panic(err)
			}

			sum += int64(a * b)
		}
	}

	return sum
}

func p2(input string) any {
	line := strings.ReplaceAll(strings.ReplaceAll(input, "\r\n", "\n"), "\n", "")

	enabled := true
	sum := int64(0)

	ops := OPRegex2.FindAllString(line, -1)

	for _, op := range ops {
		if op == "do()" {
			enabled = true
			continue
		}

		if op == "don't()" {
			enabled = false
			continue
		}

		if !enabled {
			continue
		}

		parts := strings.Split(op[4:len(op)-1], ",")

		a, err := strconv.Atoi(parts[0])

		if err != nil {
			panic(err)
		}

		b, err := strconv.Atoi(parts[1])

		if err != nil {
			panic(err)
		}

		sum += int64(a * b)
	}

	return sum
}
