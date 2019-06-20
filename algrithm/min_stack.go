/// @Author winjeg,  winjeg@qq.com
/// All rights reserved to winjeg

package algrithm

import "math"

type MinStack struct {
	Min    int
	Values []int
	Sorted []int
}

/** initialize your data structure here. */
func Constructor2() MinStack {
	return MinStack{
		Min:    math.MaxInt32,
		Values: make([]int, 0, 5),
		Sorted: make([]int, 0, 5),
	}
}

func (this *MinStack) Push(x int) {
	if x <= this.Min {
		this.Min = x
	}
	if len(this.Sorted) == 0 {
		this.Sorted = append(this.Sorted, x)
	} else {
		sl := len(this.Sorted)
		if this.Sorted[0] >= x {
			w :=  this.Sorted
			this.Sorted = append([]int{}, x)
			this.Sorted = append(this.Sorted, w...)
		} else if this.Sorted[sl-1] <= x {
			this.Sorted = append(this.Sorted, x)
		} else {
			for i := range this.Sorted {
				if this.Sorted[i] >= x {
					ori := this.Sorted
					this.Sorted = append([]int{}, ori[0:i]...)
					this.Sorted = append(this.Sorted, x)
					this.Sorted = append(this.Sorted, ori[i:]...)
					break
				}
			}
		}

	}
	this.Values = append(this.Values, x)
}

func (this *MinStack) Pop() {
	l := len(this.Values)
	if l < 1 {
		return
	}
	if l == 1 {
		this.Values = []int{}
		this.Sorted = []int{}
		this.Min = math.MaxInt32
		return
	}
	popped := this.Values[l-1]
	this.Values = this.Values[0:l -1]
	if popped == this.Min {
		this.Sorted = this.Sorted[1:]
		this.Min = this.Sorted[0]
		return
	}
	if popped == this.Sorted[len(this.Sorted) -1] {
		this.Sorted = this.Sorted[0:len(this.Sorted) -1]
		this.Min = this.Sorted[0]
		return
	}
	ori := make([]int, len(this.Sorted), len(this.Sorted))
	copy(ori, this.Sorted)
	for i := range ori {
		if ori[i] == popped {
			this.Sorted = append([]int{}, ori[0:i]...)
			this.Sorted = append(this.Sorted, ori[i+1:]...)
		}
	}
	this.Min = this.Sorted[0]
}

func (this *MinStack) Top() int {
	return this.Values[len(this.Values)-1]
}

func (this *MinStack) GetMin() int {
	return this.Min
}



type MinStack2 struct {
	stack []int
	minStack []int
}

// 主要还是要想到一点，  更小的元素一定会出现在 minStack 的最后面， 如果先pop， 肯定会先pop到最小的
func (this *MinStack2) Push(x int) {
	this.stack = append(this.stack, x)

	if len(this.minStack) == 0 || this.minStack[len(this.minStack)-1] >= x {
		this.minStack = append(this.minStack, x)
	}
}

func (this *MinStack2) Pop() {
	popValue := this.stack[len(this.stack)-1]

	if popValue == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[0 : len(this.minStack)-1]
	}
	this.stack = this.stack[0 : len(this.stack)-1]

}

