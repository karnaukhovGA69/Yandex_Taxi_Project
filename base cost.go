package Yandex_Taxi_Project

const pricePerKm = 10.0
const pricePerMinute = 2.0

type TripParameters struct {
	Distance float64
	Duration float64
}

func CalculateBasePrice(trip TripParameters) float64 {
	return trip.Distance*pricePerKm + trip.Duration*pricePerMinute
}
