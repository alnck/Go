package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Eren Alincak"

	str1 := str[:4]
	str2 := str[len(str)-7:]
	str3 := fmt.Sprintf("isim: %s, Soy isim: %s", str1, str2)

	if strings.EqualFold(str1, "ErEn") {
		fmt.Println("str1 = ErEn")
	}

	if strings.HasPrefix(str2, "Alincak") {
		fmt.Println("str2 = Alincak")
	}

	fmt.Println(str)
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println(strings.ToUpper(str))
}
