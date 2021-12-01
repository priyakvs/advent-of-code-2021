package main

import (
	"./day1"
	"./seed"
	"fmt"
)

func main() {
	fmt.Printf("Number of increased measurements (Sample): %d\n", day1.NumOfIncreasedMeasurements(seed.GetSampleData()))

	fmt.Printf("Number of increased measurements: %d\n", day1.NumOfIncreasedMeasurements(seed.GetRealData()))

	fmt.Printf("Number of increased sum (Sample), Window 3: %d\n", day1.NumberOfIncreasedSum(seed.GetSampleData(), 3))

	fmt.Printf("Number of increased sum, Window 3: %d\n", day1.NumberOfIncreasedSum(seed.GetRealData(), 3))
}
