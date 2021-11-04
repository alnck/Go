package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		go func(val int) {
			fmt.Println(val)
		}(i)
	}

	time.Sleep(time.Second * 3) //waitgroup oluşturmak yerine kısa olsun diye 3 saniye bekleme konuldu
}
