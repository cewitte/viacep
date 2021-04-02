package cepcleaners

import (
	"regexp"
	"strings"
)

func ExtractNumbers(number string) string {
	re := regexp.MustCompile("[0-9]+")
	numbers := re.FindAllString(number, -1)
	return strings.TrimSpace(strings.Join(numbers, ""))
}
