package lib

import (
	"testing"
)

var randNum int

func init() {
	randIntN = func(n int) int {
		return randNum
	}
}

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

var expectedMatthew = BibleApiRequest{
	Book: Matthew,
	Chapter: 0,
	Verses: VerseRange{
		Begin: 0,
		End: 0,
	},
}
var expectedMark = BibleApiRequest{
	Book: Mark,
	Chapter: 1,
	Verses: VerseRange{
		Begin: 1,
		End: 2,
	},
}
var expectedLuke = BibleApiRequest{
	Book: Luke,
	Chapter: 2,
	Verses: VerseRange{
		Begin: 2,
		End: 4,
	},
}
var expectedJohn = BibleApiRequest{
	Book: John,
	Chapter: 3,
	Verses: VerseRange{
		Begin: 3,
		End: 6,
	},
}
func Test_RandomGospel(t *testing.T) {
	randNum = 0
	expectedResults := []BibleApiRequest{
		expectedMatthew,
		expectedMark,
		expectedLuke,
		expectedJohn,
	}

	for randNum < 4 {
		if gospel := RandomGospel(); gospel != expectedResults[randNum] {
			t.Fatalf(`Unexpected random gospel result: %s`, gospel)
		}
		randNum++
	}
}
