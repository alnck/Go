package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Eren Alincak"
	spacesize := 0

	for i := 0; str[i] <= '\r'; i++ {
		println("size:", i)
		spacesize = i
	}

	str1 := str[:spacesize]
	str2 := str[len(str)-spacesize:]
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
