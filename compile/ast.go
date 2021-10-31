/**
抽象语法树等编译器相关demo，学习golang标准库如何进行词法、语法分析
*/
package compile

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/scanner"
	"go/token"
	"os"
)

// Scan 扫描源码文件的字符，生成token
func Scan(sourceCode []byte) {
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(sourceCode))
	s.Init(file, sourceCode, nil, 0)

	for {
		pos, tok, lit := s.Scan()
		fmt.Printf("%-6s%-8s%q\n", fset.Position(pos), tok, lit)

		if tok == token.EOF {
			break
		}
	}
}

// Parse...
func Parse(sourceCode []byte) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", sourceCode, 0)
	if err != nil {
		panic(err)
	}

	// 打印AST
	_ = ast.Print(fset, f)
}

// Inspect 解析特定文件，形成抽象语法树，并处理特定的节点，比如return
func Inspect(filePath string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		// 发现return语句
		ret, ok := n.(*ast.ReturnStmt)
		if ok {
			fmt.Printf("return statement found on line %v:\n", fset.Position(ret.Pos()))
			printer.Fprint(os.Stdout, fset, ret)
			fmt.Printf("\n")
			return true
		}
		return true
	})
}
