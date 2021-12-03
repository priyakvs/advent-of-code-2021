package common

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadInputLinesAsStrings(filename string) (result []string, err error){
	if filename == "" {
		return nil, errors.New("filename cannot be empty")
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error when reading the file: %w", err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		result = append(result, line)
	}

	return result, nil
}

func ReadInputLinesAsIntegers(filename string) (result []int, err error) {
	if filename == "" {
		return nil, errors.New("filename cannot be empty")
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error when reading the file: %w", err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		value, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("error in converting string to int: %w", err)
		}
		result = append(result, value)
	}

	return result, nil
}