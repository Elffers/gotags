package main

import (

	"fmt"

	"gotags/parser"

	"go/ast"
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
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			token := x.Name
			filename := fset.File(n.Pos()).Name()
			regex := fmt.Sprintf("/^func %s(/", token)
			fmt.Printf("%s\t%s\t%s", token, filename, regex)
		}
		return true
	})
}
