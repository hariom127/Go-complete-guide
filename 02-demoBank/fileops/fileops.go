package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteValueToFile(balanceFile string, value float64) {
	var valueStr = fmt.Sprint(value)
	os.WriteFile(balanceFile, []byte(valueStr), 0644)
}

func ReadFloatFromFile(fileName string) (float64, error) {
	value, err := os.ReadFile(fileName) // _ is a special type of variable name which no need to entertain right now
	if err != nil {
		return 1000, errors.New("failed to find file")
	}

	valueStr := string(value)
	val, err := strconv.ParseFloat(valueStr, 64)

	if err != nil {
		return 1000, errors.New("failed to parse value")
	}
	return val, nil
}
