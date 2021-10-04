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
	ayahs := make([]types.Ayah, 7000)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Trim(line, " ") == "" {
			break
		}
		
		ayah := util.BuildAyah(line)
		ayahs = append(ayahs, *ayah)
	}
}