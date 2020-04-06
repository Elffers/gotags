package parser

import (
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
