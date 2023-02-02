package leetcode

func maxArea1(height []int) int {
	i, j := 0, len(height)-1
	maxArea := 0
	for i < j {
		var curArea int
		width := j - i
		if height[i] < height[j] {
			curArea = width * height[i]
			i++
		} else {
			curArea = width * height[j]
			j--
		}
		if curArea > maxArea {
			maxArea = curArea
		}
	}
	return maxArea
}
