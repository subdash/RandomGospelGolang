package lib

import (
	"fmt"
	"math/rand"
	"time"
)

const BaseUrl = "https://bible-api.com"
const KJV = "?translation=kjv"
const (
	Matthew Gospel = "matthew"
	Mark Gospel = "mark"
	Luke Gospel = "luke"
	John Gospel = "john"
)

func RandomGospel() BibleApiRequest {
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

	return BibleApiRequest{
		Book: book,
		Chapter: chapterNum,
		Verses: VerseRange{
			Begin: beginVerse,
			End: endVerse,
		},
	}
}

func makeUrl(reading *BibleApiRequest) string {
	return fmt.Sprintf(
		"%s/%s%s",
		BaseUrl,
		makeSelection(reading),
		KJV)
}

func makeSelection(reading *BibleApiRequest) string {
	return fmt.Sprintf(
		"%s%%20%d:%d-%d",
		reading.Book,
		reading.Chapter,
		reading.Verses.Begin,
		reading.Verses.End)
}
