package main

import (
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

func p1(input string) any {
	return 42
}

func p2(input string) any {
	return "not implemented"
}
