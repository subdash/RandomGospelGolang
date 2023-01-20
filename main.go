package main

import (
	"fmt"
	"random-gospel/lib"
)

func main() {
	request := lib.RandomGospel()
	resp, err := lib.GetVerse(&request)
	if err != nil {
		panic(err)
	}
	fmt.Println(<-resp)
}
