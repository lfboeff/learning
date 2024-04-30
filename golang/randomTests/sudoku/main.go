package main

import (
	"fmt"
	"slices"
	"sudoku_mod/examples"
)

func main() {

	fmt.Println("Sudoku!")

	listExamples := examples.CreateExamples()

	board := listExamples[0].Board

	showSudoku(board)

	rowIdx, colIdx := findNextPosition(board)
	// log.Fatal("a:", a, "b:", b)

	// rowIdx, colIdx = 2, 1

	row := getRowArray(board, rowIdx)
	fmt.Println(row)
	fmt.Println(findPossibleValuesArray(row))

	col := getColArray(board, colIdx)
	fmt.Println(col)
	fmt.Println(findPossibleValuesArray(col))

	block := getBlock(board, rowIdx, colIdx)
	fmt.Println(block)
	fmt.Println(findPossibleValuesArray(block))

	fmt.Printf(">> %v\n", resultPossibleValues(
		findPossibleValuesArray(row),
		findPossibleValuesArray(col),
		findPossibleValuesArray(block)),
	)

	if possibleResults := resultPossibleValues(
		findPossibleValuesArray(row),
		findPossibleValuesArray(col),
		findPossibleValuesArray(block)); len(possibleResults) == 1 {
		fmt.Println("Yes!")
	} else {
		fmt.Println("Not yet!")
	}

}

func findNextPosition(board [][]int) (int, int) {

	for row := range board {
		for col := range board[row] {
			if board[row][col] == 0 {
				return row, col
			}
		}
	}

	return 0, 0
}

func showSudoku(matrix [][]int) {

	fmt.Println("|-----------------------|")

	for outerIndex, row := range matrix {

		fmt.Printf("|")

		for innerIndex, digit := range row {

			if (innerIndex)%3 == 0 && innerIndex != 0 {
				fmt.Printf(" |")
			}

			if digit == 0 {
				fmt.Printf(" _")
			} else {
				fmt.Printf(" %v", digit)
			}
		}

		fmt.Println(" |")

		if (outerIndex+1)%3 == 0 && outerIndex != 0 {
			fmt.Println("|-----------------------|")
		}
	}
}

// A
func findPossibleValuesArray(array []int) []int {

	var possibleValues []int

	for _, v := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		if !slices.Contains(array, v) {
			possibleValues = append(possibleValues, v)
		}
	}

	return possibleValues
}

// A
func getRowArray(matrix [][]int, row int) []int {

	return matrix[row]
}

// A
func getColArray(matrix [][]int, col int) []int {

	var array []int

	for index := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		array = append(array, matrix[index][col])
	}

	return array
}

// A
func getBlock(matrix [][]int, rowIdx, colIdx int) []int {

	var array []int

	if rowIdx <= 2 && colIdx <= 2 {
		for _, v := range []int{0, 1, 2} {
			array = append(array, matrix[v][:3]...)
		}
	} else if rowIdx <= 5 && colIdx <= 2 {
		for _, v := range []int{3, 4, 5} {
			array = append(array, matrix[v][:3]...)
		}
	} else if rowIdx <= 8 && colIdx <= 2 {
		for _, v := range []int{6, 7, 8} {
			array = append(array, matrix[v][:3]...)
		}
	} else if rowIdx <= 2 && colIdx <= 5 {
		for _, v := range []int{0, 1, 2} {
			array = append(array, matrix[v][3:6]...)
		}
	} else if rowIdx <= 5 && colIdx <= 5 {
		for _, v := range []int{3, 4, 5} {
			array = append(array, matrix[v][3:6]...)
		}
	} else if rowIdx <= 8 && colIdx <= 5 {
		for _, v := range []int{6, 7, 8} {
			array = append(array, matrix[v][3:6]...)
		}
	} else if rowIdx <= 2 && colIdx <= 8 {
		for _, v := range []int{0, 1, 2} {
			array = append(array, matrix[v][6:]...)
		}
	} else if rowIdx <= 5 && colIdx <= 8 {
		for _, v := range []int{3, 4, 5} {
			array = append(array, matrix[v][6:]...)
		}
	} else if rowIdx <= 8 && colIdx <= 8 {
		for _, v := range []int{6, 7, 8} {
			array = append(array, matrix[v][6:]...)
		}
	}

	return array
}

// A
func resultPossibleValues(row, col, block []int) []int {

	var result []int

	for _, v := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		if slices.Contains(row, v) && slices.Contains(col, v) && slices.Contains(block, v) {
			result = append(result, v)
		}
	}

	return result
}
