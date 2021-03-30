package main

import (
	"fmt"
	"sample/algo/sort"
	"sample/algo/trie"
)

func fiboIter() func() int {
	fib1, fib2 := 0, 1
	return func() int {
		retFib := fib1
		fib1, fib2 = fib2, fib1+fib2
		return retFib
	}
}

func main() {
	t := trie.New()
	t.AddWords("Apple", "Ape", "Water", "What", "Walter", "when")
	fmt.Println(t.Exists("Apple"))
	fmt.Println(t.Search("W"))

	arr := []int{11, 10, 9, 0, -1}
	sort.InsertionSort(arr)
	fmt.Println(arr)
}
