package main

import (
	"fmt"
	"go/ast"
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
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			token := x.Name
			filename := fset.File(n.Pos()).Name()
			// token  filename  regex
			regex := fmt.Sprintf("/^func %s(/", token)
			fmt.Printf("%s\t%s\t%s", token, filename, regex)
		}
		return true
	})
}
