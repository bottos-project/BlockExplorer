package common

import (
	"strconv"
)

// string to int
func String2Int(str string) (int, error) {
	return strconv.Atoi(str)
}
