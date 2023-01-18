package lib

import (
	"testing"
)

func Test_makeUrl (t *testing.T) {
	mockReading := BibleApiRequest{
		Book: John,
		Chapter: 3,
		Verses: VerseRange{
			Begin:1,
			End:2,
		},
	}
	const expectedUrl = "https://bible-api.com/john%203:1-2?translation=kjv"

	if url := makeUrl(&mockReading); url != expectedUrl {
		t.Fatalf(`unexpected url: %s`, url)
	}
}
