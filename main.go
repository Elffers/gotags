package main

import (
	"fmt"
	"go/token"
	"gotags/parser"
	"sort"
)

func main() {
	fset := token.NewFileSet()
	path := "."
	files, err := parser.ParseAll(path, fset)

	if err != nil {
		fmt.Println(err)
		return
	}
	tags := []string{}
	for _, f := range files {
		parser.Generate(f, fset, &tags)
	}

	// Print special fields
	fmt.Println("!_TAG_FILE_FORMAT\t2\t/extended format/")
	fmt.Println("!_TAG_FILE_SORTED\t1\t/sorted/")
	fmt.Println("!_TAG_PROGRAM_AUTHOR\tHsing-Hui Hsu\thhhsu1@gmail.com")
	fmt.Println("!_TAG_PROGRAM_NAME\tgotags\t")
	fmt.Println("!_TAG_PROGRAM_URL\thttps://github.com/SoManyHs/gotags")
	sort.Strings(tags)
	for _, tag := range tags {
		fmt.Printf(tag)
	}

}
