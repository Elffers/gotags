package parser

import (
	"fmt"

	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"os"
)

func ParseAll(path string, fset *token.FileSet) ([]*ast.File, error) {
	out := []*ast.File{}
	dirsToSkip := ".git"

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == dirsToSkip {
			return filepath.SkipDir
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".go" {
			f, err := Parse(fset, path)
			if err != nil {
				return err
			}

			out = append(out, f)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return out, nil
}

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
			fmt.Printf("%s\t%s\t%s\n", token, filename, regex)
		}
		return true
	})
}
