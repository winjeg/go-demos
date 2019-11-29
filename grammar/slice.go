package grammar

import (
	"fmt"
	"reflect"
)

func sliceTest() {

	// b[:]  是吧特定数组类型转换成 slice  b[:6:8] 把新slice的cap 设置为8
	// 只有 slice 可以加三个点变成变参

	var a = [3]int{1, 2, 3}
	// b是slice
	var b = a[:]
	ad(a)
	ad(b)
	// 展开变参
	ab(b...)
	// 需要[]interface类型的展开
	// ac(b...)

	// []int 整体可以作为单个interface 参数
	ac(b)
}

func ab(o ...int) {
	fmt.Println(reflect.TypeOf(o))
}

func ac(o ...interface{}) {
	fmt.Println(reflect.TypeOf(o))
}

func ad(o interface{}) {
	fmt.Println(reflect.TypeOf(o))
}
