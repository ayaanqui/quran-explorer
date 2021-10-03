package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ayaanqui/quran-explorer/src/util"
)

func main() {
	fmt.Println("Welcome to the Qur'an Explorer")

	const filename = "quran-simple.txt1"
	filepath, err := util.GetResource(filename)
	util.PanicMsg(err, "Could not build path for '" + filename + "'")

	file, err := os.Open(filepath)
	util.PanicMsg(err, "Could not open '" + filepath + "'")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fmt.Println("Finished!")
}