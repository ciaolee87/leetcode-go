package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	println(myAtoi("20000000000000000000"))
}

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	regex, _ := regexp.Compile("^[-+]?[0-9]+")

	ext := regex.FindString(str)
	num, _ := strconv.Atoi(ext)

	if num < math.MinInt32 {
		return math.MinInt32
	} else if math.MaxInt32 < num {
		return math.MaxInt32
	} else {
		return num
	}
}
