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
	BP          int    `json:"bp",omitempty`
	SpO2        int    `json:"spo2",omitempty`
}

func main() {
	for i := 1; i <= 50; i++ {
		// fmt.Printf("%.2f \n", temperature())
		fmt.Println(heartRate())
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

// respRate returns a random respiratory rate (breaths per minute) for an adult
func respRate() float64 {
	// respiratory rates for an adult person at rest range from 12
	// to 18 breaths per minute
	mean := 15.0
	stdDev := 3.0

	sample := rand.NormFloat64()*stdDev + mean
	return math.Floor(sample)
}

// heartRate returns a random heart rate (beats per minute)
// TODO: would be good to favour larger numbers (more common)
func heartRate() int {
	// normal range is 50-90 beats per minute, for athletes this might
	// be between 33 and 50 bpm
	max := 95
	min := 28

	sample := rand.Intn(max-min) + min
	return sample
}
