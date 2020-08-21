package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	go func() {
		for {
			m.Store("beijin", 100)
			fmt.Println(m.Load("beijin"))
		}
	}()

	go func() {
		for {
			m.Store("beijin", 1000)
			fmt.Println(m.Load("beijin"))
		}
	}()
	var i int
	fmt.Scan(&i)
}
