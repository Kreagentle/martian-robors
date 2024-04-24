package parsing

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseMap(input string) (int, int, error) {
	parts := strings.Fields(input)

	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid input format: expected two space-separated components")
	}

	num1, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing first number: %v", err)
	}

	num2, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing second number: %v", err)
	}

	return num1, num2, nil
}

func ParseRobot(input string) (int, int, int, error) {
	parts := strings.Fields(input)

	if len(parts) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid input format: expected three space-separated components")
	}

	num1, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error parsing first number: %v", err)
	}

	num2, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error parsing second number: %v", err)
	}

	runeInput := []rune(parts[2])
	if len(runeInput) != 1 {
		return 0, 0, 0, fmt.Errorf("invalid rune input: expected a single rune")
	}
	char := runeInput[0]

	if _, ok := AngleMap[char]; !ok {
		return 0, 0, 0, fmt.Errorf("invalid rune input: expected a single rune")
	}

	return num1, num2, AngleMap[char].(int), nil
}
