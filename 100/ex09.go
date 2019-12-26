package main

import "strconv"

func main() {
	println(isPalindrome(100011))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	stringX := strconv.Itoa(x)

	strSize := len(stringX)
	half := strSize / 2
	for i := 0; i < half; i++ {
		if stringX[i] != stringX[strSize-1-i] {
			return false
		}
	}

	return true
}
