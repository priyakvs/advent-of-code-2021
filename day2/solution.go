package day2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	separator = " "
	forwardDirection = "forward"
	upDirection = "up"
	downDirection = "down"
)

type position struct {
	horizontal int
	depth int
	aim int
}

func DepthHorizontalOperationResult(commands []string, operator string) (int, error) {
	var pos position
	for _, c := range commands {
		direction, value, err := spaceSplit(c)
		if err != nil {
			return 0, fmt.Errorf("error in space split: %w", err)
		}
		switch direction {
		case forwardDirection:
			pos.horizontal += value
			pos.depth += pos.aim * value
		case upDirection:
			//pos.depth -= value (part 1)
			pos.aim -= value
		case downDirection:
			//pos.depth += value (part 1)
			pos.aim += value
		}
	}

	result, err := computeResult(pos, operator)
	if err != nil {
		return 0, fmt.Errorf("error in computing result: %w", err)
	}

	return result, nil
}

func spaceSplit(command string) (string, int, error) {
	commandSplit := strings.Split(command, separator)
	if len(commandSplit) == 0 || len(commandSplit) > 2 {
		return "", 0, errors.New("invalid command")
	}
	value, err := strconv.Atoi(commandSplit[1])
	if err != nil {
		return "", 0, fmt.Errorf("error in converting value from string to int: %w", err)
	}
	return commandSplit[0], value, nil
}

func computeResult(pos position, operator string) (int, error) {
	switch operator {
	case "+":
		return pos.horizontal + pos.depth, nil
	case "-":
		return pos.horizontal - pos.depth, nil
	case "*":
		return pos.horizontal * pos.depth, nil
	case "/":
		return pos.horizontal / pos.depth, nil
	case "%":
		return pos.horizontal % pos.depth, nil
	default:
		return 0, errors.New("invalid operation")
	}
}