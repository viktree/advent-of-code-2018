package day05

import (
	"fmt"
	// "strings"
	// "adventOfCode"
)

func checkString(str string) string {
	Cache := make(map[string]string)

	for _, a := range str {
		Cache[string(a)] = string(a)
		fmt.Printf("Cache[%s]: %s \n", string(a), Cache[string(a)])
	}

	for i, a := range str {
		for _, b := range str[:i] {
			if a^b == 32 {
				Cache[string(a)+string(b)] = ""
			} else {
				Cache[string(a)+string(b)] = string(a) + string(b)
			}
			fmt.Printf("Cache[%s]: %s \n", string(a)+string(b), Cache[string(a)+string(b)])
		}
	}

	// var newString string
	for i := 3; i < len(str); i++ {
		for j := 0; i+j < len(str); j++ {
			substring := str[j:(j + i)]
			for k := 1; k < i; k++ {
				a := Cache[substring[:k]]
				b := Cache[substring[k:]]
				fmt.Println(a, a[len(a)-1], b, Cache[b])
				// fmt.Println("EE", a, a[len(a)-1]+b[0], b[1:])
				// cached := a[:len(a)-1] + Cache[string(a[len(a)-1])+string(b[0])] + b[1:]
				// fmt.Printf("HI: %s:%s | %s:%s\n", a, b, Cache[a], Cache[b])
				// if _, ok := Cache[substring]; ok {
				// 	if len(cached) < len(Cache[substring]) {
				// 		Cache[substring] = cached
				// 	}
				// } else {
				// 	Cache[substring] = cached
				// }

			}
			fmt.Printf("Cache[%s]: %s \n", substring, Cache[substring])
		}
	}

	// fmt.Println(str, Cache[str])

	return "hi"
}

func PartOne() {

	polymer := "dabAcCaCBAcCcaDA"

	changed := false

	new_polymer := polymer
	for i, unit := range new_polymer {
		var new_unit string
		if i+1 != len(new_polymer) {
			if changed {
				i = i - 1
			}
			before := new_polymer[:i]
			after := new_polymer[i+2:]
			next_unit := new_polymer[i+1]

			if byte(unit)^next_unit == 32 {
				new_unit = ""
			} else {
				new_unit = string(unit) + string(next_unit)
			}
			fmt.Println(i, before+" "+string(unit)+string(next_unit)+" "+after)
			fmt.Println(i, before+" "+new_unit+" "+after)
			polymer = new_polymer
			new_polymer = before + new_unit + after
			if polymer != new_polymer {
				changed = true
			} else {
				changed = false
			}
			fmt.Println(i, polymer)
			fmt.Println(i, new_polymer)
		}
	}
	fmt.Println(polymer)
	fmt.Println(new_polymer)

	return
	// fmt.Printf("Answer: %d\n", 42)
}

func PartTwo() {
	return
	// fmt.Printf("Answer: %d\n", 42)
}
