package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	filename := "./main.go"
	f, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Output: %+v", f)
}
