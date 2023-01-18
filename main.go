package main

import (
	"fmt"
	"random-gospel/lib"
)

func main() {
	// john334 := getVerse(
	// 	Luke,
	// 	&Reading{
	// 		chapter: 3,
	// 		verses: VerseRange{begin: 3, end: 4},
	// 	},
	// )
	// fmt.Println(<-john334)

	randGosp, chapter, verseBegin, verseEnd := lib.RandomGospel()
	resp := lib.GetVerse(&lib.Reading{
		Book: randGosp,
		Chapter: chapter,
		Verses: lib.VerseRange{Begin: verseBegin, End: verseEnd},
	})
	fmt.Println(<-resp)
}
