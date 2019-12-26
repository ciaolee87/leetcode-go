package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	fmt.Println(reverse(1534236469))

}

func reverse(x int) int {
	var rev string
	sign := x < 0

	if sign {
		x = -x
	}

	stringX := strconv.Itoa(x)

	for i := len(stringX); 0 < i; i-- {
		rev = rev + stringX[i-1:i]
	}

	num, _ := strconv.Atoi(rev)

	if sign {
		num = -num
	}

	if (num < math.MinInt32) || (math.MaxInt32 < num) {
		return 0
	} else {
		return num
	}

}
