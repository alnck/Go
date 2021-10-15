package main

import (
	"fmt"
	"time"
)

func main() {
	slc1 := []int{1, 2, 3}
	slc2 := []int{}

	for i, value := range slc1 {
		fmt.Printf("index: %d value: %d\n", i, value)
	}

	for i := 0; i < 10; i++ {
		slc2 = append(slc2, i)
	}

	// fmt.Println(slc1)
	// fmt.Println(slc2)
	// fmt.Printf("slc1 len:%d cap:%d\n", len(slc1), cap(slc1))
	// fmt.Printf("slc2 len:%d cap:%d\n", len(slc2), cap(slc2))

	c := time.After(5 * time.Second)

	for {
		b := false

		select {
		case <-c:
			b = true
		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
		}

		if b {
			break
		}
	}
}
