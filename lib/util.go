package lib

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomGospel() (Gospel, int, int , int) {
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

	bookNum := r1.Intn(4)
	var book Gospel

	switch bookNum {
	case 0:
		book = Matthew
	
	case 1:
		book = Mark
	
	case 2:
		book = Luke

	case 3:
		book = John
	}

	mapping := ChapterVerseMapping[book]
	versesPerChapter := mapping.VersesPerChapter
	chapterNum := r1.Intn(mapping.NumChapters + 1) 
	beginVerse := r1.Intn(versesPerChapter[chapterNum] + 1)
	endVerse := r1.Intn(versesPerChapter[chapterNum] + 1 - beginVerse) + beginVerse

	return book, chapterNum, beginVerse, endVerse
}

func makeUrl(reading *Reading) string {
	return fmt.Sprintf(
		"%s/%s%s",
		BaseUrl,
		makeSelection(reading),
		KJV,
	)
}

func makeSelection(reading *Reading) string {
	if reading.Verses.End == 0 {
		return fmt.Sprintf("%d:%d", reading.Chapter, reading.Verses.Begin)
	}

	return fmt.Sprintf(
		"%s%%20%d:%d-%d",
		reading.Book,
		reading.Chapter,
		reading.Verses.Begin,
		reading.Verses.End)
}
