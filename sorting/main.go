package main

import (
	"fmt"
	"sort"
)

func main() {
	people := []Person{
		{"Fernando", 30},
		{"Paul", 33},
		{"John", 23},
		{"Alice", 33},
		{"Mark", 19},
		{"Bob", 35},
	}

	sort.Sort(ByAgeThenName(people))

	fmt.Println(people)
}
