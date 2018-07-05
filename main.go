package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Response struct {
	ID          string `json:"id",omitempty`
	Description string `json:"desc",omitempty`
	Timestamp   string `json:"timestamp",omitempty`
	Temp        int    `json:"temp",omitempty`
	HeartRate   int    `json:"heartrate",omitempty`
	RespRate    int    `json:"resprate",omitempty`
	BPM         int    `json:"bpm",omitempty`
	SO2         int    `json:"so2",omitempty`
}

func main() {
	for i := 1; i <= 50; i++ {
		// fmt.Printf("%.2f \n", temperature())
		fmt.Println(respRate())
	}
}

// temperature generates a random body temperature in degrees celcius
func temperature() float64 {
	// average body temperatures should fall between 36.5 and 37,
	// with a standard deviation of about 0.4 degrees
	mean := 36.8
	stdDev := 0.4

	// use rand.Norm for a normal distribution curve
	sample := rand.NormFloat64()*stdDev + mean
	return sample
}

// respRate returns a random respiration rate (breaths per minute) for an adult
func respRate() float64 {
	// respiration rates for an adult person at rest range from 12
	// to 18 breaths per minute
	mean := 15.0
	stdDev := 3.0

	sample := rand.NormFloat64()*stdDev + mean
	return math.Floor(sample)
}
