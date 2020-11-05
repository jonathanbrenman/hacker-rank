package main

import "fmt"

func main() {
	arr := generateArr()
	counter := minimumSwaps(arr)
	fmt.Println(counter)
}

func generateArr() []int32 {
	return []int32{4,3,1,2}
}

func minimumSwaps(arr []int32) int32 {
	var swaps int32 = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != int32(i) + 1 {
			index := retrieveIndex(arr, i)
			tmp := arr[i]
			arr[i] = arr[index]
			arr[index] = tmp
			swaps++
		}
	}
	return swaps
}

func retrieveIndex(arr []int32, index int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == int32(index) + 1 {
			return i
			break
		}
	}
	return -1
}