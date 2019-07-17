/// @Author winjeg,  winjeg@qq.com
/// All rights reserved to winjeg

package leetcode
func plusOne(digits []int) []int {
	l := len(digits)
	if l < 1 {
		return digits
	}
	i := l - 1
	digits[i]++
	r := digits[i]
	for r > 9 && i >= 0 {
		digits[i] = 0
		if i == 0 {
			digits = append([]int{1}, digits...)
			return digits
		}
		i--
		digits[i]++
		r = digits[i]
	}
	return digits
}
