package main

import "fmt"

func main() {
	n1 := 0
	n2 := 1
	nth := 0
	sum := 0
	for nth <= 4000000 {
		nth = n1 + n2
		n1 = n2
		n2 = nth
		if nth%2 == 0 {
			sum += nth
		}
	}
	fmt.Println(sum)
}
