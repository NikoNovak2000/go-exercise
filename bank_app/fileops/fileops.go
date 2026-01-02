package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to find file.")
	}

	balanceText := string(data)                       // helper variable, storing float64 data and converting it to string
	value, err := strconv.ParseFloat(balanceText, 64) // string conversion to float64

	if err != nil {
		return 1000, errors.New("Failed to parse stored value.")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)                  // format the balance into a string and save it into a variable, which will be passed in a byte collection
	os.WriteFile(fileName, []byte(valueText), 0644) // 0644, way of encoding file permissions --> this set means the owner of the file will be able to read and write the file, while others just read
}
