package fizzbuzz

import (
	"strconv"
	"strings"
)

func convert(raw, int1, int2 int, s1, s2 string) string {
	if raw%(int1*int2) == 0 {
		return strings.Join([]string{s1, s2}, " ")
	}

	if raw%int1 == 0 {
		return s1
	}

	if raw%int2 == 0 {
		return s2
	}

	return strconv.Itoa(raw)
}

func Generate(int1 int, int2 int, limit int, s1 string, s2 string) []string {
	result := make([]string, limit)

	for i := 0; i < limit; i++ {
		result[i] = convert(i+1, int1, int2, s1, s2)
	}

	return result
}
