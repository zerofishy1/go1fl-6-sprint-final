package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoDetectAndConvert(input string) (string, error) {
	if input == "" {
		return "", nil
	}

	temp := input
	morseChars := []string{".", "-", " ", "\n", "\r", "\t"}
	for _, char := range morseChars {
		temp = strings.ReplaceAll(temp, char, "")
	}

	if len(temp) > 0 {
		result := morse.ToMorse(input)
		return result, nil
	} else {
		result := morse.ToText(input)
		return result, nil
	}
}
