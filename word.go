package main

import (
	"strings"
)

// split string
// n:鱼;vi:捕鱼;vt:捕鱼
func MeaningStringToMap(inputStr string) map[string]string {
	result := make(map[string]string)
	strs := strings.Split(inputStr, ";")
	for _, str := range strs {
		part := strings.Split(str, ":")
		result[part[0]] = part[1]
	}
	return result
}
