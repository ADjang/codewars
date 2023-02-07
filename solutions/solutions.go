package solutions

import (
	"math"
	"strconv"
)

func Cats(start, finish int) int {
	jumps := 0
	if start == finish {
		return 0
	} else {
		i := start
		for i < finish {
			if finish-i >= 1 && finish-i < 3 {
				i += 1
				jumps++
			} else {
				i += 3
				jumps++
			}
		}
		return jumps
	}
}

func SevenWondersScience(compasses int, gears int, tablets int) int {
	total, x, y := 0, 0, 0

	arr := []int{compasses, gears, tablets}

	min := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	x = min * 7

	y += int(math.Pow(float64(compasses), 2))
	y += int(math.Pow(float64(gears), 2))
	y += int(math.Pow(float64(tablets), 2))

	total += (x + y)
	return total
}

func IsNegativeZero(n float64) bool {
	a := 1.0 / math.Inf(-1)
	if a == -0 {
		return true
	} else {
		return false
	}
}

func PrinterError(s string) string {
	count := 0
	for _, v := range s {
		if v > 'm' {
			count++
		}
	}
	return strconv.Itoa(count) + "/" + strconv.Itoa(len(s))
}
