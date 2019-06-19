/// @Author winjeg,  winjeg@qq.com
/// All rights reserved to winjeg

package storage

import (
	"fmt"
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	_ "github.com/pingcap/tidb/types/parser_driver"
)

func ParseSql(sql string) {
	digest := parser.DigestHash(sql)
	fmt.Println(digest)
	p := parser.New()
	stmt, warn, err := p.Parse(sql, "utf8", "utf8_general_ci")
	fmt.Println(stmt[0].Text())
	fmt.Printf("stmt:%v\t, warn:%v\t,error:%v", stmt, warn, err)
}

type MyVisitor struct {
}

func (m MyVisitor) Enter(n ast.Node) (node ast.Node, skipChildren bool) {
	if n != nil {
		fmt.Println("enter:" + n.Text())
		x := n.(*ast.SelectStmt)
		y := x.From.TableRefs.Left.(*ast.TableSource)
		tn := y.Source.(*ast.TableName)
		fmt.Println(tn.Name)
		return nil, true
	}
	return  nil, true
}

func (m MyVisitor) Leave(n ast.Node) (node ast.Node, ok bool) {
	if n != nil {
		fmt.Println("leave:" + n.Text())
	}
	return nil, true
}
