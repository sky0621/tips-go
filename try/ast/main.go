package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
)

func main() {
	searchRoot := os.Getenv("SEARCH_ROOT")
	if searchRoot == "" {
		log.Fatal("need SEARCH_ROOT")
	}
	if err := filepath.Walk(searchRoot, analyze); err != nil {
		log.Fatal(err)
	}
}

func analyze(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	// FIXME:
	if info.Name() != "resolver_advertiser.go" {
		return nil
	}
	fSet := token.NewFileSet()
	f, err := parser.ParseFile(fSet, path, nil, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	ast.Inspect(f, func(n ast.Node) bool {

		if fn, ok := n.(*ast.FuncDecl); ok {
			fmt.Println("=============================")
			fmt.Println(fn.Name)
			fmt.Println("--- Recv --------------------")
			if fn.Recv != nil {
				for _, recv := range fn.Recv.List {
					fmt.Println(recv.Type)
				}
			}
			fmt.Println("--- Params ------------------")
			if fn.Type != nil {
				if fn.Type.Params != nil {
					for _, param := range fn.Type.Params.List {
						for _, name := range param.Names {
							fmt.Println(name.Name)
						}
						fmt.Println(param.Type)
					}
				}
			}
			fmt.Println("--- Body --------------------")
			if fn.Body != nil {
				for _, body := range fn.Body.List {
					fmt.Println(body)
					//if ident, ok := body.(ast.CallExpr); ok {
					//	fmt.Println(ident.Name)
					//}
				}
			}
			fmt.Println("=============================")
		}
		return true
	})
	return nil
}
