package lib

import "fmt"

// TODO: move to...???
type Gospel string


// TODO: move these to api.go
type Reading struct {
	Book Gospel
	Chapter int
	Verses VerseRange
}

type VerseRange struct {
	Begin int
	End int
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

func (verseResp BibleApiResponse) String() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		verseResp.Reference,
		verseResp.Text)
}
