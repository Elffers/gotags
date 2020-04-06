package main

import (

	"fmt"

	"gotags/parser"

	"go/token"
)

func main() {
	fset := token.NewFileSet()
	path := "./parser"

	m, err := parser.Parse(fset, path)

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, pkg := range m {
		for _, f := range pkg.Files {
		parser.Generate(f, fset)
		}
	}


}
