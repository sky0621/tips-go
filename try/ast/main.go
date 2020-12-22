package main

import (
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
	//ast.Inspect()
	if err := ast.Print(fSet, f); err != nil {
		log.Fatal(err)
	}
	return nil
}
