package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		fmt.Println("Gorotine Merhaba")
		wg.Done()
	}()

	wg.Wait() //blocking

	fmt.Println("Mainden Merhaba")
}
