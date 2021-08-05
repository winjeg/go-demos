/// @Author winjeg,  winjeg@qq.com
/// All rights reserved to winjeg

package leetcode

import (
	"fmt"
	"testing"
)

func TestMedian(t *testing.T) {
	result := findMedianSortedArrays([]int{1}, []int{1})
	fmt.Print(result)
}

func TestSolution_Shuffle(t *testing.T) {

	nums := []int{2, 3, 7, 6}

	obj := Constructor(nums)
	param_1 := obj.Reset()
	param_2 := obj.Shuffle()
	fmt.Println(param_1)
	fmt.Println(param_2)

	obj2 := Constructor2()
	obj2.Push(-2)
	obj2.Push(0)
	obj2.Push(-1)
	obj2.GetMin()
	obj2.Pop()
	obj2.Top()
	obj2.GetMin()

	obj2.Pop()
	obj2.Pop()

	obj2.GetMin()
	obj2.Push(2)
	obj2.Push(2)
	obj2.Push(-1)
	obj2.Push(-1)
	obj2.Pop()
	obj2.Pop()
	obj2.GetMin()
	param_3 := obj2.Top()
	param_4 := obj2.GetMin()

	a := []int{}
	fmt.Println(a[0:0])
	fmt.Println(param_3, param_4)
}

func TestCountPrimes(t *testing.T) {
	fmt.Println(countPrimes2(11999983))
	fmt.Println(countPrimes(11999983))
}
