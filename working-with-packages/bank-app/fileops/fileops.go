package fileops

import (
	"fmt"
	"strconv"
	"errors"
	"os"
)


func WriteFloatToFile(value float64, filename string) {
	valueTxt := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueTxt), 0644)
}

func GetFloatFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return 10000, errors.New("failed to find file")
	}

	valueTxt := string(data)
	floatValue, err := strconv.ParseFloat(valueTxt, 64)

	if err != nil {
		return 10000, errors.New("failed to parse file")
	}
	
	return floatValue, nil 
}
