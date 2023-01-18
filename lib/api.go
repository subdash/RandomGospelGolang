package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetVerse(reading *Reading) <-chan BibleApiResponse {
	ch := make(chan BibleApiResponse)
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
		var verseResp BibleApiResponse
		json.Unmarshal([]byte(body), &verseResp)

		ch <- verseResp
	}()

	return ch
}
