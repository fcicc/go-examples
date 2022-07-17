package main

import "sync"

func BubbleSort(list []int, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	// Do not use this in production!
	for i := 1; i < len(list); i++ {
		for j, k := 0, 1; j < len(list)-i; j, k = j+1, k+1 {
			if list[j] > list[k] {
				list[j], list[k] = list[k], list[j]
			}
		}
	}
}
