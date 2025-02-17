package main

import (
	"fmt"
	"git_test/prototype/product"
)

func main() {
	// create file
	file1 := product.NewFile("file1.txt")
	file2 := product.NewFile("file2.txt")
	file3 := product.NewFile("file3.txt")

	// create folder
	folder1 := &product.Folder{
		Name:     "Folder1",
		Children: []product.INode{file1},
	}

	folder2 := &product.Folder{
		Children: []product.INode{folder1, file2, file3},
		Name:     "Folder2",
	}

	fmt.Println("\n Printing hierarcy for Folder2")
	folder2.Print("   ")

	// clone folder 2
	folder2Clone := folder2.Clone()
	fmt.Println("\n Printing hierarcy for Folder2")
	folder2Clone.Print("   ")
}
