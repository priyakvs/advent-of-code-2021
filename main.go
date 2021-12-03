package main

import (
	"./common"
	"./day1"
	"./day2"
	"fmt"
	"log"
)

var (
	filePathPrefix = "/Users/priya/personal/advent-of-code-2021/%s/seed/%s.txt"
)

func main() {
	// Day 1
	fmt.Println("*********************** Day1 ************************")
	sampleInput, err := common.ReadInputLinesAsIntegers(fmt.Sprintf(filePathPrefix, "day1", "sampleInput"))
	if err != nil {
		log.Fatal("encountered an error in getting sample data for day 1", err)
	}

	realInput, err := common.ReadInputLinesAsIntegers(fmt.Sprintf(filePathPrefix, "day1", "realInput"))
	if err != nil {
		log.Fatal("encountered an error in getting real data for day 1", err)
	}
	fmt.Printf("Number of increased measurements (Sample): %d\n", day1.NumOfIncreasedMeasurements(sampleInput))

	fmt.Printf("Number of increased measurements: %d\n", day1.NumOfIncreasedMeasurements(realInput))

	fmt.Printf("Number of increased sum (Sample), Window 3: %d\n", day1.NumberOfIncreasedSum(sampleInput, 3))

	fmt.Printf("Number of increased sum, Window 3: %d\n\n", day1.NumberOfIncreasedSum(realInput, 3))

	// Day 2
	fmt.Println("*********************** Day2 ************************")
	sampleInputSlice, err := common.ReadInputLinesAsStrings(fmt.Sprintf(filePathPrefix, "day2", "sampleInput"))
	if err != nil {
		log.Fatal("encountered an error in getting sample data for day 1", err)
	}

	realInputSlice, err := common.ReadInputLinesAsStrings(fmt.Sprintf(filePathPrefix, "day2", "realInput"))
	if err != nil {
		log.Fatal("encountered an error in getting real data for day 1", err)
	}
	opResultSample, err := day2.DepthHorizontalOperationResult(sampleInputSlice, "*")
	if err != nil {
		log.Fatal("encountered an error in computing the result for sample data (Sample)", err)
	}
	fmt.Printf("Final horizonal and depth analysis result: %d\n", opResultSample)

	opResultReal, err := day2.DepthHorizontalOperationResult(realInputSlice, "*")
	if err != nil {
		log.Fatal("encountered an error in computing the result for real data", err)
	}
	fmt.Printf("Final horizonal and depth analysis result: %d\n\n", opResultReal)
}
