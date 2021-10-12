package types

import (
	"fmt"
	"strings"
)

type Surah struct {
	ChapterNumber uint
	Name string
	Prefix string
}

type ISurah interface {
	GetName() string
}

func (s *Surah) GetName() string {
	if s.Prefix == "" {
		return fmt.Sprintf("Surat %s", s.Name)
	}
	return fmt.Sprintf("Surat %s-%s", s.Prefix, s.Name)
}

func (s *Surah) NameLower() string {
	return strings.ToLower(s.Name)
}