package utils

import (
	"strconv"
)

func ForceString(s interface{}) *string {
	if s == nil {
		return nil
	}
	if v, ok := s.(float64); ok {
		s2 := strconv.Itoa(int(v))
		return &s2
	}
	if v, ok := s.(string); ok {
		return &v
	}
	return nil
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
