package solutions

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
