package Yandex_Taxi_Project

import "fmt"

const (
	minPrice = 99.0
	maxPrice = 20000.0
)

func ApplyPriceLimits(price float64) float64 {
	if price < minPrice {
		price = minPrice
	} else if price > maxPrice {
		price = maxPrice
	}
	return price
}

func SuckDick() {
	fmt.Println("Suck suck suck...")
}
