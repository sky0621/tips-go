package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	// src は AST を調べたい入力です。
	src := `
package p

// methodABC stringを受け取ってBigStrを返す。
func methodABC(s string) (BigStr, error) {
  // s がブランクならエラー
  if s == "" {
    return BigStr{}, errors.New("err")
  }
  return BigStr{}, nil
}

func IsXyz() bool {
  return true
}
`

	// src を解析して AST を作成します。
	fset := token.NewFileSet() // 位置は fset に対する相対位置です
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	// AST を調べて，すべての ID とリテラルを表示します。
	ast.Inspect(f, func(n ast.Node) bool {
		var s string
		switch x := n.(type) {
		//case *ast.BasicLit:
		//	s = x.Value
		//case *ast.Ident:
		//	s = x.Name
		case *ast.FuncDecl:
			//for _, b := range x.Body.List {
			//	fmt.Println(b)
			//}
			fmt.Println(x.Doc.Text())
			fmt.Println(x.Name.String())
		}
		if s != "" {
			fmt.Printf("%s:\t%s\n", fset.Position(n.Pos()), s)
		}
		return true
	})

}
