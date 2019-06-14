package qrcode

import (
	"testing"
)

func TestGenerateQrCode(t *testing.T) {
	GenerateQrCode("中间件测试工作你会做到什么时候？")
	// t.Fail()
}

type Query4Audit struct {
	Query string
}

type Rule struct {
	//blabla
	Position int `json:"Position"`
	//增加一个变量类型int
	Func     func(*Query4Audit, int) Rule
}
func (q *Query4Audit) RuleSameAlias() Rule {
	return Rule{}
}

func (q *Query4Audit) test(a int) Rule {
	return Rule{}
}

//源码比较复杂，简化为说明该场景的模式
func testx(){
	q := &Query4Audit{Query: ""}

	f1 := (*Query4Audit).test
	f1(q, 3)

	f2 := q.test
	f2(3)
}