package main

import (
	"fmt"
	"go/token"
	"gotags/parser"
)

func main() {
	fset := token.NewFileSet()
	path := "."
	files, err := parser.ParseAll(path, fset)

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, f := range files {
		parser.Generate(f, fset)
	}
}
