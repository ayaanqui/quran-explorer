package main

import (
	"fmt"
	"strconv"

	"github.com/ayaanqui/quran-explorer/src/types"
	"github.com/ayaanqui/quran-explorer/src/util"
)
	
func main() {
	// Qur'an facts
	// TOTAL VERSES/AYAHS: 6,236
	// TOTAL CHAPTERS: 114
	// TOTAL SAJDAHS: 15

	const CHAPTERS = 114

	chapter_to_ayah := make(map[uint][]*types.Ayah, CHAPTERS)
	util.BuildBrowsableQuran(chapter_to_ayah)

	surahs := util.BuildSurahs()

	fmt.Printf("Welcome to the Qur'an Explorer\n\n")
		
	for {
		fmt.Print("Enter command: ")
		var input string
		fmt.Scanln(&input)
			
		if input == "q" {
			break
		}

		ch, err := strconv.ParseUint(input, 10, 32)
		if err != nil {
			// TODO: Handle non integer input
			fmt.Println("Handle non integer input")
		}

		ch_ayahs, found := chapter_to_ayah[uint (ch)]
		if found {
			fmt.Println(surahs[uint (ch)].GetName())
			for _, a := range ch_ayahs {
				fmt.Printf("%d:%d %s\n", a.ChapterNumber, a.VerseNumber, a.Verse)
			}
		} else {
			fmt.Println("Invalid chapter number")
		}
		fmt.Println()
	}
}