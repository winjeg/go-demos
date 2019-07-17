/// @Author winjeg,  winjeg@qq.com
/// All rights reserved to winjeg

package leetcode

import (
	"math/rand"
)

type Solution struct {
	Arr []int
}

func Constructor(nums []int) Solution {
	return Solution{
		Arr: nums,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.Arr
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	arrLen := len(this.Arr)
	na := make([]int, arrLen, arrLen)
	copy(na, this.Arr)
	rand.Shuffle(arrLen, func(i, j int) {
		na[i], na[j] = na[j], na[i]
	})
	return na
}


// XX 算法
func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	m := make(map[int]bool, n)
	for i := 2; i < n; i++ {
		m[i] = true
	}
	count := 0
	for i := 2; i < n; {
		if !m[i] {
			i++
			continue
		}
		count++
		delete(m, i)
		if i*i > n {
			count += len(m)
			break
		}
		for k := range m {
			if k%i == 0 {
				delete(m, k)
			}
		}
		i++
	}
	return count
}

//
func countPrimes2(n int) int {
	res := 0
	var p = make([]bool, n)
	for i:=2;i<n;i++ {
		if !p[i] {
			res ++
			for j:=2;i*j<n;j++ {
				p[i*j] = true
			}
		}
	}
	return res
}
