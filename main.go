package main

import (

	"fmt"

	"gotags/parser"

	"go/token"
)

func main() {
	fset := token.NewFileSet()
	filename := "./main.go"

	f, err := parser.Parse(fset, filename)

	if err != nil {
		fmt.Println(err)
		return
	}
	parser.Generate(f, fset)
}
