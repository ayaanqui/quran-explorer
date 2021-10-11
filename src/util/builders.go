package util

import (
	"bufio"
	"os"
	"strings"

	"github.com/ayaanqui/quran-explorer/src/types"
)

// Given a map of uint -> []Ayah (chapter -> list of ayahs), this
// function populates the map with all the verses on each chapter.
// Due to the nature of maps looking up any verse is O(1) constant-time lookup.
func BuildBrowsableQuran(chapter_to_ayah map[uint][]*types.Ayah) {
	const filename = "quran-uthmani.txt"
	filepath, err := GetResource(filename)
	PanicMsg(err, "Could not build path for '" + filename + "'")
	
	file, err := os.Open(filepath)
	PanicMsg(err, "Could not open '" + filepath + "'")
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Trim(line, " ") == "" {
			break
		}
		
		ayah := BuildAyah(line)
		
		chapter := ayah.ChapterNumber
		u, v := chapter_to_ayah[chapter]
		if v {
			chapter_to_ayah[chapter] = append(u, ayah)
		} else {
			ayah_arr := make([]*types.Ayah, 0, 3)
			ayah_arr = append(ayah_arr, ayah)
			chapter_to_ayah[chapter] = ayah_arr
		}
	}
}

func BuildSurahs() map[uint]*types.Surah {
	filename := "surah-names-english.txt"
	filepath, err := GetResource(filename)
	if err != nil {
		panic("Could not build path for: " + filename)
	}

	file, err := os.Open(filepath)
	if err != nil {
		panic("Could not open surahs file: " + filepath)
	}

	scanner := bufio.NewScanner(file)
	surahs := make(map[uint]*types.Surah, 114)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Trim(line, " ") == "" {
			break
		}
		surah := BuildSurah(line)
		surahs[surah.ChapterNumber] = surah
	}
	return surahs
}