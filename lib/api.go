package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type BibleApiRequest struct {
	Book Gospel
	Chapter int
	Verses VerseRange
}

type VerseRange struct {
	Begin int
	End int
}

func (request BibleApiRequest) String() string {
	return fmt.Sprintf(
		"Book: %s, Chapter: %d, Verses: %d, %d",
		request.Book,
		request.Chapter,
		request.Verses.Begin,
		request.Verses.End)
}

func GetVerse(reading *BibleApiRequest) <-chan bibleApiResponse {
	ch := make(chan bibleApiResponse)
	url := makeUrl(reading)

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
		var verseResp bibleApiResponse
		json.Unmarshal([]byte(body), &verseResp)

		ch <- verseResp
	}()

	return ch
}

type bibleApiResponse struct {
	Reference string `json:"reference"`
	Verses []bibleApiVerse `json:"verses"`
	Text string `json:"text"`
	TranslationId string `json:"translation_id"`
	TranslationName string `json:"translation_name"`
	TranslationNote string `json:"translation_note"`
}

type bibleApiVerse struct {
	BookId string `json:"book_id"`
	BookName string `json:"book_name"`
	Chapter int `json:"chapter"`
	Verse int `json:"verse"`
	Text string `json:"text"`
}

func (response bibleApiResponse) String() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		response.Reference,
		response.Text)
}
