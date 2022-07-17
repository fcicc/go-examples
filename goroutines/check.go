package main

func IsOrdered(list []int) bool {
	for i, j := 0, 1; j < len(list); i, j = i+1, j+1 {
		if list[i] > list[j] {
			return false
		}
	}
	return true
}
