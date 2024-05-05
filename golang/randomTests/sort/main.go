package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {

	// sort array of integers:
	fmt.Println("Sort array of integers:")

	int_array := []int{5, 7, 3, 7, 8, 3, 4, 9, 0, 3, 4}
	fmt.Println(int_array)

	sort.Ints(int_array)
	fmt.Println(int_array)

	// sort array of floats:
	fmt.Println("\nSort array of floats:")

	flt64_array := []float64{5.3242, 7.3434, 3.4343, 7.4343, 8.24325, 3.4343, 4, 97867.766, 0.546, -3.4343, 4.654}
	fmt.Println(flt64_array)

	sort.Float64s(flt64_array)
	fmt.Println(flt64_array)

	// sort array of strings (ALERT: case-sensitive):
	fmt.Println("\nSort array of strings (ALERT: case-sensitive):")

	str_array := []string{"fdf", "sf", "fdsf", "jyjypkp", "trerwq", "Dsdsa", "vfjbifn", "ndsidsai", "dsdsa", "asdasd"}
	fmt.Println(str_array)

	sort.Strings(str_array)
	fmt.Println(str_array)

	// sort array of string (case-insensitive):
	fmt.Println("\nSort array of string (case-insensitive):")

	str_array = []string{"fdf", "sf", "fdsf", "jyjypkp", "trerwq", "Dsdsa", "vfjbifn", "ndsidsai", "dsdsa", "asdasd"}
	fmt.Println(str_array)

	sort.Slice(str_array, func(i, j int) bool {
		return strings.ToLower((str_array[i])) < strings.ToLower((str_array[j]))
	})
	fmt.Println(str_array)

	// sort in reversed order:
	fmt.Println("\nSort in reversed order:")

	sort.Slice(int_array, func(i, j int) bool {
		return int_array[j] < int_array[i]
	})
	fmt.Println(int_array)

	sort.Slice(flt64_array, func(i, j int) bool {
		return flt64_array[j] < flt64_array[i]
	})
	fmt.Println(flt64_array)

	sort.Slice(str_array, func(i, j int) bool {
		return strings.ToLower((str_array[j])) < strings.ToLower((str_array[i]))
	})
	fmt.Println(str_array)

	// sort struct by number-field (age) in ascending order:
	fmt.Println("\nSort struct by number-field in ascending order:")

	type Book struct {
		Title         string    `json:"title"`
		Artist        string    `json:"artist"`
		Age           int       `json:"age,omitempty"`
		TotalPages    int       `json:"total_pages"`
		FavoritePages []int     `json:"favorite_pages,omitempty"`
		Color         string    `json:"color"`
		IsActive      bool      `json:"is_active"`
		Lastmodified  time.Time `json:"lastmodified"`
	}

	fileBS, err := os.ReadFile("file.json")
	if err != nil {
		log.Fatal(err)
	}

	var books []Book

	if err = json.Unmarshal(fileBS, &books); err != nil {
		log.Fatal(err)
	}

	for _, v := range books {
		fmt.Printf("%+v\n", v)
	}

	sort.Slice(books, func(i, j int) bool {
		return books[i].Age < books[j].Age
	})

	fmt.Println()
	for _, v := range books {
		fmt.Printf("%+v\n", v)
	}

	// sort struct by date-field (Lastmodified) in ascending order:
	fmt.Println("\nSort struct by date-field (Lastmodified) in ascending order:")

	sort.Slice(books, func(i, j int) bool {
		return books[i].Lastmodified.Before(books[j].Lastmodified)
	})

	for _, v := range books {
		fmt.Printf("%+v\n", v)
	}

	// sort struct by date-field (Lastmodified) in descending order:
	fmt.Println("\nSort struct by date-field (Lastmodified) in descending order:")

	sort.Slice(books, func(i, j int) bool {
		return books[j].Lastmodified.Before(books[i].Lastmodified)
	})

	for _, v := range books {
		fmt.Printf("%+v\n", v)
	}

	// sort struct by multiple fields (IsActive > Lastmodified > Title) in descending > descending > ascending order:
	fmt.Println("\nSort struct by multiple fields (IsActive > Lastmodified > Title) in descending > descending > ascending order:")

	sort.Slice(books, func(i, j int) bool {
		if books[i].IsActive != books[j].IsActive {
			return books[i].IsActive
		}
		if books[i].Lastmodified != books[j].Lastmodified {
			return books[j].Lastmodified.Before(books[i].Lastmodified)
		}
		return strings.ToLower(books[i].Title) < strings.ToLower(books[j].Title)
	})

	for _, v := range books {
		fmt.Printf("%+v\n", v)
	}

	fmt.Println("\nTestes:")
	for _, v := range books {
		fmt.Println(v.FavoritePages)
	}

	// save json data to a file:
	fileBS, err = json.Marshal(books)
	// fileBS, err = json.MarshalIndent(books, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%v\n", string(fileBS))

	if err = os.WriteFile("new_file.json", fileBS, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n\n\n...")

	arr := []int{5, 2, 3, 5, 4}
	fmt.Println(arr)

	for i := range arr {
		for j := range arr[i:] {
			if arr[j+i] < arr[i] {
				arr[i], arr[j+i] = arr[j+i], arr[i]
			}
		}
	}
	fmt.Println(arr)
}
