package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
)

var (
	age    int
	weight int
	height int
	gender string
)

func init() {
	flag.IntVar(&age, "age", 0, "Patient's age")
	flag.IntVar(&weight, "weight", 0, "Patient's weight in kg")
	flag.IntVar(&height, "height", 0, "Patient's height in cm")
	flag.StringVar(&gender, "gender", "", "Patient's gender (f/m)")
}

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
	flag.Parse()
	fmt.Println("Patient's age", age)
	fmt.Println("Patient's gender", gender)
	fmt.Println("Patient's weight", weight)
	fmt.Println("Patient's height", height)

	for i := 1; i <= 50; i++ {
		// fmt.Printf("%.2f \n", temperature())
		fmt.Println(spo2())
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

// SpO2 is the measurement of the amount of oxygen in the blood
// normal SpO2 values vary between 95 and 100%
// TODO: would be good to favour larger numbers (more common)
func spo2() int {
	max := 100
	min := 90

	sample := rand.Intn(max-min) + min
	return sample
}

// bp generates a random blood pressure value and returns systolic pressure and
// diastolic pressure measured in millimeters of mercury (mmHg)
// Normal resting blood pressure in an adult is 120/80 mmHg
func bp() (int, int) {
	return 0, 0
}
