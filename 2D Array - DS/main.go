package main

import "fmt"

func main()  {
	arr := generate2DArray()
	fmt.Println(hourglassSum(arr))
}

func generate2DArray() [][]int32 {
	arr := [][]int32{
		{1,1,1,0,0,0},
		{0,1,0,0,0,0},
		{1,1,1,0,0,0},
		{0,0,2,4,4,0},
		{0,0,0,2,0,0},
		{0,0,1,2,4,0},
	}
	return arr
}

func hourglassSum(arr [][]int32) int32 {
	var max int32
	max = -1000
	for x := 0; x < len(arr) - 2; x++ {
		for y := 0; y < len(arr[x]) - 2; y++ {
			top := arr[x][y] + arr[x][y+1] + arr[x][y+2]
			middle := arr[x+1][y+1]
			bottom := arr[x+2][y] + arr[x+2][y+1] + arr[x+2][y+2]
			sum := top + middle + bottom
			if  sum > max {
				max = sum
			}
		}
	}
	return max
}
