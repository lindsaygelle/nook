package game

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// dongwuSenlin is the common reference for Dongwu Senlin.
	dongwuSenlin = "DongwuSenlin"
)

var (
	// dongwuSenlinNameAmericanEnglish represents Dongwu Senlin's name in American English.
	dongwuSenlinNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dongwu Senlin",
	}
)

var (
	// dongwuSenlinName contains the localized names of Dongwu Senlin.
	dongwuSenlinName = nook.Languages{
		language.AmericanEnglish: dongwuSenlinNameAmericanEnglish,
	}
)

var (
	// DongwuSenlin represents Dongwu Senlin.
	DongwuSenlin = nook.Game{
		Key:          nook.Key(dongwuSenlin),
		Name:         dongwuSenlinName,
		ReleaseOrder: 5,
	}
)
