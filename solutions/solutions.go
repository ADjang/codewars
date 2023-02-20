package solutions

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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

func PrinterError(s string) string {
	count := 0
	for _, v := range s {
		if v > 'm' {
			count++
		}
	}
	return strconv.Itoa(count) + "/" + strconv.Itoa(len(s))
}

func Collatz(n int) int {
	i := 0
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
		i++
	}
	i++
	return i
}

func Gps(s int, x []float64) int {
	if len(x) <= 1 {
		return 0
	}
	v := 0.0
	arr := []float64{}
	vel := []float64{}
	for i := 0; i < len(x); i++ {
		if i != len(x)-1 {
			arr = append(arr, x[i+1]-x[i])
		}
	}
	for i := 0; i < len(arr); i++ {
		v = (3600 * arr[i]) / float64(s)
		vel = append(vel, v)
	}
	max := vel[0]
	for _, speed := range vel {
		if speed > max {
			max = speed
		}
	}
	return int(max)
}

func Vaporcode(s string) string {
	str := strings.ToUpper(s)
	str = strings.Trim(str, "\t \f \v 10")
	r := []rune(str)
	vapor := ""
	for i := 0; i < len(r); i++ {
		if i != len(r)-1 && r[i] != ' ' {
			vapor += string(r[i]) + "  "
		} else if i == len(r)-1 && r[i] != ' ' {
			vapor += string(r[i])
		}
	}
	return vapor
}

func FindNextSquare(sq int64) int64 {
	if check(math.Sqrt(float64(sq))) {
		return -1
	} else {
		for i := sq; ; i++ {
			//   0, 1, 4, 5, 6, 9
			if check(math.Sqrt(float64(i))) {
				return int64(i)
			}
		}
	}
}

func check(n float64) bool {
	num := strconv.FormatFloat(float64(n), 10, -1, 64)
	fmt.Println(num)
	if IsDigitsOnly(num) {
		return num[len(num)-1] == '0' || num[len(num)-1] == '1' || num[len(num)-1] == '4' || num[len(num)-1] == '5' || num[len(num)-1] == '6' || num[len(num)-1] == '9'
	} else {
		return false
	}
}

func IsDigitsOnly(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func MergeArrays(arr1, arr2 []int) []int {
	arr := arr1
	union := []int{}
	for j := 0; j < len(arr2); j++ {
		arr = append(arr, arr2[j])
	}

	// sorting
	k := len(arr)
	for k != 0 {
		for j := 0; j < len(arr); j++ {
			if j != len(arr)-1 && arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
		k--
	}

	for i := 0; i < len(arr); i++ {
		if i != len(arr)-1 && arr[i] != arr[i+1] {
			union = append(union, arr[i])
		} else if i == len(arr)-1 {
			union = append(union, arr[i])
		}
	}
	return union
}

func ContainsInt(a int, b []int) bool {
	for j := 0; j < len(b); j++ {
		if a == b[j] {
			return true
		}
	}

	return false
}

func wave(words string) []string {
	s := ""
	arr := []string{}
	index := []int{}
	k := 0

	for i := 0; i < len(words); i++ {
		if words[i] >= 'a' && words[i] <= 'z' {
			k++
			index = append(index, i)
		}
	}

	for i := 0; i < k; i++ {

		for j := 0; j < len(words); j++ {
			if words[j] >= 'a' && words[j] <= 'z' {
				if j == index[i] {
					s += string(rune(words[j] - 32))
				} else {
					s += string(rune(words[j]))
				}
			} else {
				s += string(rune(words[j]))
			}
		}
		arr = append(arr, s)
		s = ""
	}
	return arr
}

type Tuple struct {
	Char  string
	Count int
}

func OrderedCount(text string) []Tuple {
	num := 0
	s := []rune(text)
	check := ""
	var pair Tuple
	var tu []Tuple
	if len(text) == 0 {
		return []Tuple{}
	}
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if !strings.Contains(check, string(s[j])) {
				check += string(s[j])
			}
		}
	}
	for i := 0; i < len(check); i++ {
		for j := 0; j < len(s); j++ {
			if rune(check[i]) == s[j] {
				num++
			}
		}
		pair.Char = string(check[i])
		pair.Count = num
		tu = append(tu, pair)
		num = 0
	}
	return tu
}
