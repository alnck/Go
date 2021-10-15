package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Users struct {
	Ad		string 'json:"Eren"' //tag
	Soyad 	string 'json:"Alincak"'
	Takipci []Users
	//Begeni []struct {
	// 	Tarih time.Time
	//}
}

type Begen struct {
	Tarih time.Time
}

func main() {
	k := Users{
		Ad: 	"Eren",
		Soyad:	"Alincak",
	}

	arr, _ := json.Marshal(k)
	fmt.Println(string(arr))

	fmt.Println(k)
}
