package main

type ByAgeThenName []Person

func (arr ByAgeThenName) Len() int {
	return len(arr)
}

func (arr ByAgeThenName) Less(i, j int) bool {
	if arr[i].Age == arr[j].Age {
		return arr[i].Name < arr[j].Name
	}
	return arr[i].Age < arr[j].Age
}

func (arr ByAgeThenName) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
