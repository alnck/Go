package main

import (
	"fmt"
    "strings"
)

type String *string

func main() {
	var rstr String
	var pstr string

	fmt.Println(rstr)
	fmt.Println(pstr)

	// pstr = "Eren Alincak"
	// rstr = &pstr

	// fmt.Println(rstr)
	// fmt.Println(pstr)

	// pstr = "akordiyon ve kontrbas"

	// fmt.Println(*rstr)
	// fmt.Println(pstr)

	// *rstr = "akordiyon ve kontrbas 2021"

	// fmt.Println(*rstr)
	// fmt.Println(pstr)

	// replaceStr(pstr)

	// fmt.Println(*rstr)
	// fmt.Println(pstr)

	// replaceStr(rstr)

	// fmt.Println(*rstr)
	// fmt.Println(pstr)
}

// func replaceStr(s string) {
// 	s = strings.Replace(s, "Eren", "EREN", 1)
// }

// func replaceStr(s String) {
// 	*s = strings.Replace(*s, "Eren", "EREN", 1)
// }