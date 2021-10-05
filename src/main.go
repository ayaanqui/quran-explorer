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

func build_browsable_quran(chapter_map map[uint][]types.Ayah) {
	const filename = "quran-simple.txt"
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
		u, v := chapter_map[chapter]
		if v {
			chapter_map[chapter] = append(u, *ayah)
		} else {
			ayah_arr := make([]types.Ayah, 3)
			ayah_arr = append(ayah_arr, *ayah)
			chapter_map[chapter] = ayah_arr
		}
	}	
}

	
func main() {
	// const VERSES = 6236
	const CHAPTERS = 114

	chapter_map := make(map[uint][]types.Ayah, CHAPTERS)
	build_browsable_quran(chapter_map)

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

		ch_ayahs, found := chapter_map[uint (ch)]
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