package main

import (
	"fmt"
	"strconv"
	"strings"

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
		fmt.Println()
		fmt.Print("> ")
		var input string
		fmt.Scanln(&input)
		
		if input == "q" {
			break
		}

		ch, err := strconv.ParseUint(input, 10, 32)
		if err != nil {
			// Check to see if the input has a colon
			if strings.ContainsRune(input, ':') {
				// If input contains a colon we can assume two types of inputs:
				// 1. chapter:verse
				// 2. chapter:verse_start-verse_end

				parsed_input := strings.Split(input, ":")
				if len(parsed_input) < 2 {
					fmt.Println("Input should be of type chapter:verse or chapter:verse_start-verse_end")
					continue
				}

				ch, err := strconv.ParseUint(parsed_input[0], 10, 32)
				if err != nil {
					fmt.Println("Chapter argument should be a numeric value")
					continue
				}
				var chapter_number uint = uint (ch)

				// Fetch ayahs and surah information
				ch_ayahs, found := chapter_to_ayah[chapter_number]
				if !found {
					fmt.Printf("Could not find the chapter: %d\n", chapter_number)
					continue
				}
				surah_details := surahs[chapter_number]

				// Parse verse number(s)
				verse_side := strings.Split(parsed_input[1], " ")
				parsed_verse_args := strings.Split(verse_side[0], "-")
				var start_verse, end_verse uint

				val, err := strconv.ParseUint(parsed_verse_args[0], 10, 32)
				if err != nil {
					fmt.Println("Start verse number has to be a number")
					continue
				}
				start_verse = uint (val)
				
				if len(parsed_verse_args) > 1 {
					val, err := strconv.ParseUint(parsed_verse_args[1], 10, 32)
					if err != nil {
						fmt.Println("End verse number has to be a number")
						continue
					}
					end_verse = uint (val)
				}

				if int(start_verse) > len(ch_ayahs) {
					fmt.Println("Start ayah out of bounds")
					continue
				}

				if end_verse == 0 {
					// Only start verse was provided so return a single verse for the chapter
					PrintVerses(surah_details, &ch_ayahs, int(start_verse-1), int(start_verse))
					continue
				} else {
					if end_verse < start_verse {
						fmt.Println("End verse cannot be less than start verse")
						continue
					}
					PrintVerses(surah_details, &ch_ayahs, int(start_verse-1), int(end_verse))
					continue
				}
			}

			// String input
			for ch, surah := range surahs {
				if strings.Contains(strings.ToLower(input), surah.NameLower()) {
					ayahs, found := chapter_to_ayah[ch]
					if !found {
						fmt.Println("There was an error. Please try another command")
						continue
					}
					PrintVerses(surah, &ayahs, 0, len(ayahs))
				}
			}
			continue
		} else {
			// Numeric input (i.e. chapter number)
			ch_ayahs, found := chapter_to_ayah[uint (ch)]
			if !found {
				fmt.Println("Invalid chapter number")
			}
			PrintVerses(surahs[uint (ch)], &ch_ayahs, 0, len(ch_ayahs))
			continue
		}
	}
}

// Prints all verses in the ayahs array. 
// Range is as following: [from, to) from is inclusive, and to is exclusive
func PrintVerses(surah *types.Surah, ayahs *[]*types.Ayah, from int, to int) {
	fmt.Println(surah.GetName())

	for index, ayah := range *ayahs {
		if (index < from) {
			continue
		} else if (index >= to) {
			break
		}
		fmt.Printf("%d:%d %s\n", ayah.ChapterNumber, ayah.VerseNumber, ayah.Verse)
	}
}