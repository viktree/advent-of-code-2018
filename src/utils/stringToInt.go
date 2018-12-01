package utils

import (
	"fmt"
	"os"
	"strconv"
)

func StrToInt(str string) int {
	number, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return number
}
