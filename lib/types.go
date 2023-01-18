package lib

// TODO: move to...???
type Gospel string


// TODO: move these to api.go
type BibleApiRequest struct {
	Book Gospel
	Chapter int
	Verses VerseRange
}

type VerseRange struct {
	Begin int
	End int
}
