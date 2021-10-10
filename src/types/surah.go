package types

import "fmt"

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