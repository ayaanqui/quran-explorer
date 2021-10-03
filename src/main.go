package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ayaanqui/quran-explorer/src/types"
	"github.com/ayaanqui/quran-explorer/src/util"
)

func main() {
	fmt.Printf("Welcome to the Qur'an Explorer\n\n")

	const filename = "quran-simple.txt"
	filepath, err := util.GetResource(filename)
	util.PanicMsg(err, "Could not build path for '" + filename + "'")

	file, err := os.Open(filepath)
	util.PanicMsg(err, "Could not open '" + filepath + "'")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arabic_ayahs := make([]types.Ayah, 7000)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Trim(line, " ") == "" {
			break
		}

		ch, v, sajdah, verse := util.ParseRawAyah(line)

		fmt.Printf("%d:%d %s | Sajdah: %v\n", ch, v, verse, sajdah)
		
		ayah := util.BuildAyah(line)
		arabic_ayahs = append(arabic_ayahs, *ayah)
	}
}