package main

import (
	"bufio"
	"fmt"
	"os"
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

	chapter_to_ayah := make(map[uint][]types.Ayah, CHAPTERS)
	build_browsable_quran(chapter_to_ayah)

	fmt.Printf("Welcome to the Qur'an Explorer\n\n")
		
	for {
		fmt.Println("Enter command: ")
		var input string
		fmt.Scanln(&input)
			
		if input == "q" {
			break
		}

		ch, err := strconv.Atoi(input)
		if err != nil {
			// TODO: Handle non integer input
			fmt.Println("Handle non integer input")
		}

		ch_ayahs, found := chapter_to_ayah[uint (ch)]
		if found {
			for _, a := range ch_ayahs {
				fmt.Printf("%d:%d %s\n", a.ChapterNumber, a.VerseNumber, a.Verse)
			}
		} else {
			fmt.Println("Invalid chapter number")
		}
		fmt.Println()
	}
}

// Given a map of uint -> []Ayah (chapter -> list of ayahs), this
// function populates the map with all the verses on each chapter.
// Due to the nature of maps looking up any verse is O(1) constant-time lookup.
func build_browsable_quran(chapter_to_ayah map[uint][]types.Ayah) {
	const filename = "quran-uthmani.txt"
	filepath, err := util.GetResource(filename)
	util.PanicMsg(err, "Could not build path for '" + filename + "'")
	
	file, err := os.Open(filepath)
	util.PanicMsg(err, "Could not open '" + filepath + "'")
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Trim(line, " ") == "" {
			break
		}
		
		ayah := util.BuildAyah(line)
		
		chapter := ayah.ChapterNumber
		u, v := chapter_to_ayah[chapter]
		if v {
			chapter_to_ayah[chapter] = append(u, *ayah)
		} else {
			ayah_arr := make([]types.Ayah, 3)
			ayah_arr = append(ayah_arr, *ayah)
			chapter_to_ayah[chapter] = ayah_arr
		}
	}
}