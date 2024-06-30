package main

import (
	"fmt"
	"strings"
)

func main() {
	var num int
	var binArr []string

	fmt.Scan(&num)

	binArr = strings.Split(fmt.Sprintf("%b", num), "")
	fmt.Println(binArr)

	aux := 0
	arrTotal := []int{}

	for i, v := range binArr {
		if v == "1" {
			aux++
		} else {
			if aux != 0 {
				arrTotal = append(arrTotal, aux)
				aux = 0
			}
		}
		if i == len(binArr)-1 {
			arrTotal = append(arrTotal, aux)
		}
	}
	// fmt.Println(arrTotal)

	aux = arrTotal[0]
	for _, v := range arrTotal[1:] {
		if v > aux {
			aux = v
		}
	}
	fmt.Println(aux)

}
