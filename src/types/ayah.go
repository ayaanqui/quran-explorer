package types

import (
	"fmt"
)

type Ayah struct {
	VerseNumber uint
	ChapterNumber uint
	Sajdah bool
	Verse string
}

type IAyah interface {
	ToString() string
}

func (a *Ayah) ToString() string {
	res := fmt.Sprintf("%d:%d %s", a.ChapterNumber, a.VerseNumber, a.Verse)
	if a.Sajdah {
		return res + " Sajdah"
	}
	return res
}