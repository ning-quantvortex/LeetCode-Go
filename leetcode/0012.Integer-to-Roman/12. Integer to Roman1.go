package leetcode

func intToRoman1(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result, i := "", 0
	for num > 0 {
		for num < values[i] {
			i++
		}
		num -= values[i]
		result += symbols[i]
	}
	return result
}
