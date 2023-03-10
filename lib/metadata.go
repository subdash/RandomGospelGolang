package lib

type gospelMetadata struct {
	NumChapters int
	VersesPerChapter chapterVerseMap
}
type chapterVerseMap map[int]int

var ChapterVerseMapping = map[Gospel]gospelMetadata{
	Matthew: matthewMetadata,
	Mark: markMetadata,
	Luke: lukeMetadata,
	John: johnMetadata,
}

var matthewMetadata = gospelMetadata{
	NumChapters: 28,
	VersesPerChapter: chapterVerseMap{
		1:25,
		2:23,
		3:17,
		4:25,
		5:48,
		6:34,
		7:29,
		8:34,
		9:38,
		10:42,
		11:30,
		12:50,
		13:58,
		14:36,
		15:39,
		16:28,
		17:27,
		18:35,
		19:30,
		20:34,
		21:46,
		22:46,
		23:39,
		24:51,
		25:46,
		26:75,
		27:66,
		28:20,
	},
}

var markMetadata = gospelMetadata{
	NumChapters: 16,
	VersesPerChapter: chapterVerseMap{
		1:45,
		2:28,
		3:35,
		4:41,
		5:43,
		6:56,
		7:37,
		8:38,
		9:50,
		10:52,
		11:33,
		12:44,
		13:37,
		14:72,
		15:47,
		16:20,
	},
}

var lukeMetadata = gospelMetadata{
	NumChapters: 24,
	VersesPerChapter: chapterVerseMap{
		1:80,
		2:52,
		3:38,
		4:44,
		5:39,
		6:49,
		7:50,
		8:56,
		9:62,
		10:42,
		11:54,
		12:59,
		13:35,
		14:35,
		15:32,
		16:31,
		17:37,
		18:43,
		19:48,
		20:47,
		21:38,
		22:71,
		23:56,
		24:53,
	},
}

var johnMetadata = gospelMetadata{
	NumChapters: 21,
	VersesPerChapter: chapterVerseMap{
		1:51,
		2:25,
		3:36,
		4:54,
		5:47,
		6:71,
		7:53,
		8:59,
		9:41,
		10:42,
		11:57,
		12:50,
		13:38,
		14:31,
		15:27,
		16:33,
		17:26,
		18:40,
		19:42,
		20:31,
		21:25,
	},
}
