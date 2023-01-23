package leetcode

import "strconv"

func isPalindrome11(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 {
		return false
	}
	if x%10 == 0 {
		return false
	}

	arr := make([]int, 0, 32)
	for x > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}
	for i, j := 0, len(arr)-1; i <= j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}

func isPalindrome12(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	s := strconv.Itoa(x)
	length := len(s)
	for i := 0; i <= length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}
