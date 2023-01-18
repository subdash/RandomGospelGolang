package lib

import (
	"fmt"
	"testing"
)

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
