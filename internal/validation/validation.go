package validation

import (
	"strconv"
)

func ValidateInput(input string) (int, error) {
	return strconv.Atoi(input)
}
