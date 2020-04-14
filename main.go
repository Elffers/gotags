package main

import (

	"fmt"

	"gotags/parser"

	"go/token"
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
