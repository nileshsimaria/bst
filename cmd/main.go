package main

import (
	"fmt"
	"log"

	"github.com/nileshsimaria/bst"
)

func main() {
	t := bst.NewBST()
	if err := t.AddBulk(1, 2, 10, 20, 7, 5, 21, 3); err != nil {
		log.Fatal(err)
	}

	c := make(chan int)
	func() {
		go t.Walk(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
