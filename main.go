package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"
	var result string

	ss := strings.Split(input, " ")
	max := ss[0]
	min := ss[0]
	for i := 0; i < len(ss); i++ {

		if min < ss[i] {
			min = ss[i]
		}
		if max > ss[i] {
			max = ss[i]
		}
	}
	result = min + " " + max
	fmt.Println(result)

}
