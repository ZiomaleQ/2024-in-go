package main

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return part2solution(input)
	}
	// solve part 1 here
	return part1solution(input)
}

func part2solution(input string) any {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	left := make([]int, 0)
	right := make(map[int]int, 0)

	for _, line := range lines {
		raws := strings.Split(line, "   ")

		{
			leftParsed, err := strconv.Atoi(raws[0])

			if err != nil {
				return err
			}

			left = append(left, leftParsed)
		}

		{
			rightParsed, err := strconv.Atoi(raws[1])

			if err != nil {
				return err
			}

			right[rightParsed]++
		}
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	var sum int64 = 0

	for _, leftValue := range left {
		rightValue, ok := right[leftValue]

		if !ok {
			continue
		}

		sum += int64(rightValue) * int64(leftValue)
	}

	return sum
}

func part1solution(input string) any {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	left := make([]int, 0)
	right := make([]int, 0)

	for _, line := range lines {
		raws := strings.Split(line, "   ")

		{
			leftParsed, err := strconv.Atoi(raws[0])

			if err != nil {
				return err
			}

			left = append(left, leftParsed)
		}

		{
			rightParsed, err := strconv.Atoi(raws[1])

			if err != nil {
				return err
			}

			right = append(right, rightParsed)
		}
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	sum := 0.0

	for i := 0; i < len(left); i++ {
		sum += math.Abs(float64(left[i]) - float64(right[i]))
	}

	return int64(sum)
}
