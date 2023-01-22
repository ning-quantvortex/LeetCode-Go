package leetcode

import (
	"fmt"
	"math"
	"strconv"
)

func myAtoi1(s string) int {
	if len(s) == 0 {
		return 0
	}
	temp := make([]byte, 0)
	positive, spaceAllowed, signAllowed := true, true, true
	for _, ch := range s {
		if ch == ' ' && spaceAllowed {
			continue
		}
		if signAllowed {
			if ch == '-' {
				positive = false
				signAllowed = false
				spaceAllowed = false
				continue
			}
			if ch == '+' {
				signAllowed = false
				spaceAllowed = false
				continue
			}
		}
		if ch < '0' || ch > '9' {
			break
		}
		signAllowed, spaceAllowed = false, false
		temp = append(temp, byte(ch))
	}
	if len(temp) == 0 {
		return 0
	}
	result, _ := strconv.ParseInt(string(temp), 10, 64)
	if !positive {
		result = result * (-1)
	}
	if result > math.MaxInt32 {
		result = math.MaxInt32
	}
	if result < math.MinInt32 {
		result = math.MinInt32
	}
	fmt.Println(math.MinInt)
	return int(result)
}
