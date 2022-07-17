package main

import "sync"

func BubbleSort(list []int, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	// Do not use this in production!
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j]
				list[j] = list[j+1]
				list[j+1] = temp
			}
		}
	}
}
