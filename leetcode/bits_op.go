/// @Author winjeg,  winjeg@qq.com
/// All rights reserved to winjeg

package leetcode

import "fmt"

func reverseBits(num uint32) uint32 {
	t := num
	result := uint32(0)
	for i := uint32(0); i < 32; i++ {
		fmt.Printf("%32b\n", num)
		x := num << uint32(31)
		x >>= i
		fmt.Printf("%32b\n", x)
		result &= x
		num = t >> (i + 1)
	}
	return result
}

func hammingDistance(x int, y int) int {
	r := x ^ y
	count := 0
	for r != 0 {
		r &= r - 1
		count++
	}
	return count
}
func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		num &= num - 1
		count++
	}
	return count
}
