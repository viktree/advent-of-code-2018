package day00

import (
	"fmt"
	"strconv"
)

func ProduceSum(test_case string) int64 {
	number := []int32(test_case)
	var first_digit, second_digit int32
	sum := int64(0)

	for i := 1; i < len(number) ; i++ {
		second_digit = number[i]
		first_digit = number[i - 1]
		if first_digit == second_digit {
			converted, _ := strconv.ParseInt(string(first_digit), 0, 10)
			sum += converted
		}
	}
	first_digit = number[0]
	if first_digit == second_digit {
		converted, _ := strconv.ParseInt(string(first_digit), 0, 10)
		sum += converted
	}
	fmt.Printf("Sum = %d, first_digit = %c, second_digit = %c\n", sum, first_digit, second_digit)
	return sum
}