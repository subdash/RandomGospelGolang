package lib

import (
	"testing"
)

func Test_makeUrl_singleVerse (t *testing.T) {
	mockReading := Reading{
		Book: John,
		Chapter: 3,
		Verses: VerseRange{
			Begin:1,
			End:0,
		},
	}
	const expectedUrl = "https://bible-api.com/john%203:1?translation=kjv"

	if url := makeUrl(&mockReading); url != expectedUrl {
		t.Fatalf(`unexpected url: %s`, url)
	}
}

func Test_makeUrl_multipleVerses (t *testing.T) {
	mockReading := Reading{
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
