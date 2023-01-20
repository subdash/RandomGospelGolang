package lib

import (
	"errors"
	"fmt"
	"net/http"
	// "net/http/httptest"
	"testing"
)


var mockResponse http.Response

var mockApiError error

func init() {
	HttpGET = func(url string) (resp *http.Response, err error) {
		return &mockResponse, mockApiError
	}
}

var mockRequest = BibleApiRequest{
	Book: John,
	Chapter: 3,
	Verses: VerseRange{
		Begin: 1,
		End: 3,
	},
}

func Test_BibleApiRequest_stringer(t *testing.T) {
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

// TODO: find out why test does not always pass with coverage
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

// func Test_GetVerse(t *testing.T) {
// 	mockApiResponse := 1
// 	mockApiError = nil
// 	expectedResponse := bibleApiResponse{
// 		Reference: "John 6:17-18",
// 		Verses: []bibleApiVerse{
// 			bibleApiVerse{
// 				BookId: "JHN",
// 				BookName: "John",
// 				Chapter: 6,
// 				Verse: 17,
// 				Text: "And entered into a ship...",
// 			},
// 			bibleApiVerse{
// 				BookId: "JHN",
// 				BookName: "John",
// 				Chapter: 6,
// 				Verse: 18,
// 				Text: "And the sea arose by reason of a great wind that blew.",
// 			},
// 		},
// 	}
// 	ch := make(chan bibleApiResponse)
// 	ch <- expectedResponse
	
// 	if response := GetVerse(&mockRequest); response != ch {
// 		t.Fatalf(`%s`, &response)
// 	}
// }

func Test_GetVerse_error(t *testing.T) {
	mockApiError = errors.New("Something went wrong")

	if _, err := GetVerse(&mockRequest); err != nil {
		t.Fatalf(`GetVerse should return error`)
	}
}