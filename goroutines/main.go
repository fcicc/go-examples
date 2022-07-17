package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	const n = 50000
	runSync(n)
	runAsync(n)
}

func runSync(n int) {
	rand.Seed(time.Now().Unix())
	list := rand.Perm(n)

	start := time.Now()
	BubbleSort(list, nil)
	elapsed := time.Since(start)

	fmt.Println("--- sync version ---")
	fmt.Printf("is ordered: %t", IsOrdered(list))
	fmt.Println()
	fmt.Printf("we took %s to run the function", elapsed)
	fmt.Println()
	fmt.Println()
}

func runAsync(n int) {
	rand.Seed(time.Now().Unix())
	list := rand.Perm(n)

	var wg sync.WaitGroup
	start := time.Now()
	wg.Add(1)
	go BubbleSort(list, &wg)
	wg.Wait()
	elapsed := time.Since(start)

	fmt.Println("--- async version ---")
	fmt.Printf("is ordered: %t", IsOrdered(list))
	fmt.Println()
	fmt.Printf("we took %s to run the function", elapsed)
	fmt.Println()
	fmt.Println()
}
