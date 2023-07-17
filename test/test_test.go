package test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"testing"
)

func TestName(t *testing.T) {
	// 使用go语法表示的bool表达式，in_array为函数调用
	expr := `a == "3" && b == "0" && in_array(c, []string{"900","1100"}) && d == "0"`

	// 使用go parser解析上述表达式，返回结果为一颗ast
	parseResult, err := parser.ParseExpr(expr)
	if err != nil {
		fmt.Println(err)

		return
	}

	// 打印该ast
	ast.Print(nil, parseResult)
}
