package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
)

type Hotel struct {
	HotelID             int         `json:"hotel_id"`
	ChainID             int         `json:"chain_id"`
	ChainName           string      `json:"chain_name"`
	BrandID             int         `json:"brand_id"`
	BrandName           string      `json:"brand_name"`
	HotelName           string      `json:"hotel_name"`
	HotelFormerlyName   string      `json:"hotel_formerly_name"`
	HotelTranslatedName string      `json:"hotel_translated_name"`
	Addressline1        string      `json:"addressline1"`
	Addressline2        string      `json:"addressline2"`
	Zipcode             string      `json:"zipcode"`
	City                string      `json:"city"`
	State               string      `json:"state"`
	Country             string      `json:"country"`
	Countryisocode      string      `json:"countryisocode"`
	StarRating          float32     `json:"star_rating"`
	Longitude           float64     `json:"longitude"`
	Latitude            float64     `json:"latitude"`
	URL                 string      `json:"url"`
	Checkin             string      `json:"checkin"`
	Checkout            string      `json:"checkout"`
	Numberrooms         int         `json:"numberrooms"`
	Numberfloors        interface{} `json:"numberfloors"`
	Yearopened          int         `json:"yearopened"`
	Yearrenovated       int         `json:"yearrenovated"`
	Photo1              string      `json:"photo1"`
	Photo2              string      `json:"photo2"`
	Photo3              string      `json:"photo3"`
	Photo4              string      `json:"photo4"`
	Photo5              string      `json:"photo5"`
	Overview            string      `json:"overview"`
	RatesFrom           int         `json:"rates_from"`
	ContinentID         int         `json:"continent_id"`
	ContinentName       string      `json:"continent_name"`
	CityID              int         `json:"city_id"`
	CountryID           int         `json:"country_id"`
	NumberOfReviews     int         `json:"number_of_reviews"`
	RatingAverage       float64     `json:"rating_average"`
	RatesCurrency       string      `json:"rates_currency"`
}

func main() {
	// strBytes, err := os.ReadFile("texto.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(strBytes))

	strBytes, err := os.ReadFile("json.json")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(strBytes))

	var hotels []Hotel

	if err = json.Unmarshal(strBytes, &hotels); err != nil {
		log.Fatal(err)
	}

	var continent []string

	for _, item := range hotels {
		if !slices.Contains(continent, item.ContinentName) {
			continent = append(continent, item.ContinentName)
		}
	}
	sort.Slice(continent, func(i, j int) bool { return continent[i] < continent[j] })
	for _, value := range continent {
		fmt.Println(value)
	}

	fmt.Println()

	arr := []int{3, 4, 5, 6, 7}
	fmt.Println(arr)
	for index, value := range arr {
		if index >= len(arr)/2 {
			break
		}
		aux := value
		arr[index] = arr[len(arr)-index-1]
		arr[len(arr)-index-1] = aux
	}
	fmt.Println(arr)
}
