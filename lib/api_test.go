package lib

import (
	"fmt"
	"testing"
)

func Test_BibleApiRequest_stringer(t *testing.T) {
	mockRequest := BibleApiRequest{
		Book: John,
		Chapter: 3,
		Verses: VerseRange{
			Begin: 1,
			End: 3,
		},
	}
	expected := fmt.Sprintf(
		"Book: %s, Chapter: %d, Verses: %d, %d",
		John,
		3,
		1,
		3)

	if asString := fmt.Sprintf("%s", mockRequest); asString != expected {
		t.Fatalf(`Unexpected text: %s`, asString)
	}
}


func Test_bibleApiResponse_stringer(t *testing.T) {
	mockReference := "John 4:4-11"
	mockText := "And he must needs go through Samaria...(truncated)"
	mockResponse := bibleApiResponse{
		Reference: mockReference,
		Text: mockText,
	}
	expected := fmt.Sprintf("%s\n\n%s", mockReference, mockText)

	if asString := fmt.Sprintf("%s", mockResponse); asString != expected {
		t.Fatalf(`Unexpected text: %s`, asString)
	}
}
