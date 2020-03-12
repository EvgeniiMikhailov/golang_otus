package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func getRepeateTime(source string, startPos int) int {
	if startPos == len(source)-1 {
		return 1
	}
	var character = rune(source[startPos+1])
	if unicode.IsDigit(character) {
		var repeatTime, _ = strconv.Atoi(string(character))
		return repeatTime
	}
	return 1
}

func validateDigitalPlacement(source string, pos int) error {
	var err error
	if pos == 0 {
		err = ErrInvalidString
	} else {
		if unicode.IsDigit(rune(source[pos-1])) == true {
			err = ErrInvalidString
		}
	}
	return err
}

func Unpack(input string) (string, error) {
	var resultBuilder strings.Builder
	var err error
	for index, value := range input {
		if unicode.IsDigit(value) != true {
			var repeat = getRepeateTime(input, index)
			resultBuilder.WriteString(strings.Repeat(string(value), repeat))
		} else {
			err = validateDigitalPlacement(input, index)
			if err != nil {
				return "", err
			}
		}
	}
	return resultBuilder.String(), err
}
