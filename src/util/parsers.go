package util

import (
	"strconv"
	"strings"

	"github.com/ayaanqui/quran-explorer/src/types"
)

// Expected verse format: 114|6|مِنَ الْجِنَّةِ وَالنَّاسِ [۩]
// Order: chapter number, verse number, ayah or verse, and the symbol ۩ indicating a sajdah.
// Return: (chapter number, verse number, sajdah, verse)
func ParseRawAyah(raw_verse string) (uint, uint, bool, string) {
	parsed_ayah := strings.Split(raw_verse, "|")
	if len(parsed_ayah) < 3 {
		panic("Invalid ayah form: " + raw_verse)
	}
	
	chapter_number, err := strconv.ParseUint(parsed_ayah[0], 10, 32)
	if err != nil {
		PanicMsg(err, "Could not covert chapter number it uint: " + raw_verse + ": " + parsed_ayah[2])
	}

	verse_number, err := strconv.ParseUint(parsed_ayah[1], 10, 32)
	if err != nil {
		PanicMsg(err, "Could not covert verse number it uint: " + raw_verse + ": " + parsed_ayah[1])
	}
	verse_raw := parsed_ayah[2]
	sajdah := false

	if strings.Contains(verse_raw, "۩") {
		verse_raw = verse_raw[0:len(verse_raw)-2]
		sajdah = true
	}

	return uint (chapter_number), uint (verse_number), sajdah, verse_raw
}

// Expected verse format: 114|6|مِنَ الْجِنَّةِ وَالنَّاسِ [۩]
// Order: chapter number, verse number, ayah/verse, and the symbol ۩ indicating a sajdah.
func BuildAyah(raw_verse string) *types.Ayah {
	ch_num, verse_num, sajdah, verse := ParseRawAyah(raw_verse)
	ayah := types.Ayah{
		VerseNumber: verse_num, 
		ChapterNumber: ch_num, 
		Sajdah: sajdah, 
		Verse: verse,
	}
	return &ayah
}

// Expected line format: 15|Hijr|Al|. 
// With the order: chapter number, surah name, surah prefix (can be empty. Ex. 11|Hud||)
func BuildSurah(line string) *types.Surah {
	parsed_line := strings.Split(line, "|")
	if len(parsed_line) < 3 {
		panic("Could not parse surah line: " + line)
	}

	ch_number, err := strconv.ParseUint(parsed_line[0], 10, 32)
	if err != nil {
		panic("Could not convert string to number: " + parsed_line[0])
	}

	surah := types.Surah {
		ChapterNumber: uint (ch_number),
		Name: parsed_line[1],
		Prefix: parsed_line[2],
	}
	return &surah
}