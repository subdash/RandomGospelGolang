package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Gospel string
type DidError bool
type Reading struct {
	chapter int
	verses VerseRange
}

type VerseRange struct {
	begin int
	end int
}

type BibleApiResponse struct {
	Reference string `json:"reference"`
	Verses []BibleApiVerse `json:"verses"`
	Text string `json:"text"`
	TranslationId string `json:"translation_id"`
	TranslationName string `json:"translation_name"`
	TranslationNote string `json:"translation_note"`
}

type BibleApiVerse struct {
	BookId string `json:"book_id"`
	BookName string `json:"book_name"`
	Chapter int `json:"chapter"`
	Verse int `json:"verse"`
	Text string `json:"text"`
}

const (
	Matthew Gospel = "matthew"
	Mark Gospel = "mark"
	Luke Gospel = "luke"
	John Gospel = "john"
)

const BaseUrl = "https://bible-api.com"
const KJV = "?translation=kjv"

func getVerse(book Gospel, reading *Reading) <-chan BibleApiResponse {
	ch := make(chan BibleApiResponse)
	url := makeUrl(book, reading)

	go func() {
		defer close(ch)
		fmt.Printf("___FETCHING URL___: %s\n\n", url)
		resp, err := http.Get(url)

		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}
		var verseResp BibleApiResponse
		json.Unmarshal([]byte(body), &verseResp)

		ch <- verseResp
	}()

	return ch
}

func makeUrl(book Gospel, reading *Reading) string {
	return fmt.Sprintf(
		"%s/%s:%s%s",
		BaseUrl,
		book,
		makeSelection(reading),
		KJV,
	)
}

func makeSelection(reading *Reading) string {
	if reading.verses.end == 0 {
		return fmt.Sprintf("%d:%d", reading.chapter, reading.verses.begin)
	}

	return fmt.Sprintf(
		"%d:%d-%d",
		reading.chapter,
		reading.verses.begin,
		reading.verses.end)
}

func (verseResp BibleApiResponse) String() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		verseResp.Reference,
		verseResp.Text)
}

func randomGospel() (Gospel, int, int , int) {
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	var book Gospel
	bookNum := r1.Intn(4)
	chapterNum := (r1.Intn(10)) + 1
	beginVerse := r1.Intn(10)
	endVerse := beginVerse + (r1.Intn(10))

	switch bookNum {
	case 0:
		book = Matthew
	
	case 1:
		book = Mark
	
	case 2:
		book = Mark

	case 3:
		book = Mark
	}

	return book, chapterNum, beginVerse, endVerse
}

func main() {
	// john334 := getVerse(
	// 	Luke,
	// 	&Reading{
	// 		chapter: 3,
	// 		verses: VerseRange{begin: 3, end: 4},
	// 	},
	// )
	// fmt.Println(<-john334)

	randGosp, chapter, verseBegin, verseEnd := randomGospel()
	resp := getVerse(randGosp, &Reading{
		chapter: chapter,
		verses: VerseRange{begin: verseBegin, end: verseEnd},
	})
	fmt.Println(<-resp)
}
