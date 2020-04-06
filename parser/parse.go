package parser

import (
	"fmt"

	"go/ast"
	"go/parser"
	"go/token"
)

func Parse(fset *token.FileSet, filename string) (*ast.File, error) {
	f, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func Generate(f *ast.File, fset *token.FileSet) {
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
