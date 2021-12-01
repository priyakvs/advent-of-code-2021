package day1

// NumOfIncreasedMeasurements returns the number of times a depth measurement increases from the previous measurement
func NumOfIncreasedMeasurements(measurements []int) int {
	var increasedMeasurementCount int
	for i := 0; i < len(measurements)-1; i++ {
		if measurements[i] < measurements[i+1] {
			increasedMeasurementCount++
		}
	}
	return increasedMeasurementCount
}

// NumberOfIncreasedSum returns number of times the sum of measurements
// that are larger than the previous sum in the given window
func NumberOfIncreasedSum(measurements []int, window int) int {
	var increasedSumCount int
	for i := 0; i < len(measurements)-window; i++ {
		if measurements[i] < measurements[i+window] {
			increasedSumCount++
		}
	}
	return increasedSumCount
}