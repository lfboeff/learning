package main

import "fmt"

func main() {

	arr := []int32{1, 1, 3, 2, 1}
	fmt.Println(arr)

	// result := make([]int32, 100)
	result := make([]int32, 4)
	fmt.Println(result)

	for _, value := range arr {
		result[value] += 1
	}

	fmt.Println(result)

	var arrSorted []int32

	for index, value := range result {
		for i := 0; i < int(value); i++ {
			arrSorted = append(arrSorted, int32(index))
		}
	}

	fmt.Println(arrSorted)
}
