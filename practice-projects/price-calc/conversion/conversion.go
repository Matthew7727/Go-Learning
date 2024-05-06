package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	var floats []float64

	for _, string := range strings{
		floatVal, err := strconv.ParseFloat(string, 64)
		if err != nil {
			return nil, errors.New("Faied to convert into float")
		}
		floats = append(floats, floatVal )
	}

	return floats, nil
}