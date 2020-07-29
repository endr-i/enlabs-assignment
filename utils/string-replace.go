package utils

import "strings"

func StringReplace(s string, replaces map[string]string) string {
	result := s
	for oldS, newS := range replaces {
		result = strings.Replace(result, oldS, newS, -1)
	}
	return result
}
