package main

import "fmt"

func main() {
	go func() {
		fmt.Println("Merhaba Dünya")
	}()

	fmt.Println("Mainden Merhaba")
}
