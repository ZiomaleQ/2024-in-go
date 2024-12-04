package main

import (
	"math"
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
		return p2(input)
	}
	// solve part 1 here
	return p1(input)
}

func p1(input string) any {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	safeCount := 0

	for _, line := range lines {
		nums := make([]int, 0)
		for _, val := range strings.Split(line, " ") {
			val, err := strconv.Atoi(val)

			if err != nil {
				return err
			}

			nums = append(nums, val)
		}

		if checkSafety(nums) {
			safeCount++
		}
	}

	return safeCount
}

func p2(input string) any {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	safeCount := 0

	for _, line := range lines {
		nums := make([]int, 0)

		for _, val := range strings.Split(line, " ") {
			val, err := strconv.Atoi(val)

			if err != nil {
				return err
			}

			nums = append(nums, val)
		}

		isSafe := false

		for i := 0; i < len(nums); i++ {
			copyNums := make([]int, len(nums)-1)
			copyIdx := 0

			for j := 0; j < len(nums); j++ {
				if j == i {
					continue
				}

				copyNums[copyIdx] = nums[j]
				copyIdx++
			}

			if checkSafety(copyNums) {
				isSafe = true
				break
			}
		}

		if isSafe {
			safeCount++
		}
	}

	return safeCount
}

func checkSafety(elements []int) bool {
	diff := make([]int, len(elements)-1)

	for i := 0; i < len(elements)-1; i++ {
		diff[i] = elements[i+1] - elements[i]
	}

	firstDiff := diff[0]

	inc := firstDiff > 0

	for idx, d := range diff {
		if math.Abs(float64(d)) > 3 || d == 0 {
			return false
		}

		if idx == 0 {
			continue
		}

		if inc && d > 0 || !inc && d < 0 {
			continue
		} else {
			return false
		}
	}

	return true
}
