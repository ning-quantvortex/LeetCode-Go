package leetcode

var lookup = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt1(s string) int {
	n := len(s)
	total := 0
	for i := range s {
		if i < n-1 && lookup[s[i]] < lookup[s[i+1]] {
			total -= lookup[s[i]]
		} else {
			total += lookup[s[i]]
		}
	}
	return total
}
