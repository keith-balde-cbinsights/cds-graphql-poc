package utils

import (
	"fmt"
	"strconv"
)

func StringToInt(str *string) (int, error) {
	return strconv.Atoi(*str)
}

func ConvertStringsToInts(strSlice []*string) ([]int, error) {
	intSlice := []int{}

	for _, str := range strSlice {
		idInt, err := StringToInt(str)

		if err != nil {
			return nil, fmt.Errorf("failed to convert id to int: %w", err)
		}
		intSlice = append(intSlice, idInt)
	}

	return intSlice, nil
}

func ConvertIntsToStrings(intSlice []int) ([]*string, error) {
	strSlice := []*string{}

	for _, id := range intSlice {
		convertedId := strconv.Itoa(id)
		strSlice = append(strSlice, &convertedId)
	}

	return strSlice, nil
}
