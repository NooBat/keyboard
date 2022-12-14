package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetInt() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	result, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	result, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
