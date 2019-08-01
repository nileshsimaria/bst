package bst_test

import (
	"fmt"

	bst "github.com/nileshsimaria/bst"
)

func ExampleBST_Add() {
	bstTree := bst.NewBST()
	if err := bstTree.Add(5); err != nil {
		fmt.Println(err)
	}

	if err := bstTree.Add(1); err != nil {
		fmt.Println(err)
	}

	if err := bstTree.Add(5); err != nil {
		fmt.Println(err)
	}
	if err := bstTree.Add(1); err != nil {
		fmt.Println(err)
	}

	//Output:
	// Found duplicate: data=5
	// Found duplicate: data=1
}
func ExampleBST_AddBulk() {
	bstTree := bst.NewBST()
	c := make(chan int)

	if err := bstTree.AddBulk(10, 5, 6, 1, 11, 2, 9); err == nil {
		func() {
			go bstTree.Walk(c)
		}()

		r := make([]int, 0)
		for d := range c {
			r = append(r, d)
		}
		fmt.Println(r)
		//Output: [1 2 5 6 9 10 11]
	}
}

func ExampleBST_Search() {
	bstTree := bst.NewBST()
	bstTree.Add(10)
	bstTree.Add(8)
	bstTree.Add(17)
	if n := bstTree.Search(8); n != nil {
		fmt.Println("found")
	} else {
		fmt.Println("not found")
	}
	//Output: found
}
