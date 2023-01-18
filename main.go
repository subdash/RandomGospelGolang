package main

import (
	"fmt"
	"random-gospel/lib"
)

func main() {
	request := lib.RandomGospel()
	resp := lib.GetVerse(&request)
	fmt.Println(<-resp)
}
