package main

import (
	"fmt"
	"log"
	"random-gospel/lib"
)

func main() {
	request := lib.RandomGospel()
	resp, err := lib.GetVerse(&request)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(<-resp)
}
