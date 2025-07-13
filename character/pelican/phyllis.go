package pelican

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// phyllisBirthday represents phyllis birthday.
	phyllisBirthday = nook.Birthday{
		Day:   21,
		Month: time.November}
)

var (
	// phyllisCode represents phyllis code.
	phyllisCode = nook.Code{
		Value: "pgb/plm"}
)

var (
	// phyllisAmericanEnglishName represents phyllis american english name.
	phyllisAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Phyllis"}

	// phyllisCanadianFrenchName represents phyllis canadian french name.
	phyllisCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Élisabec"}

	// phyllisDutchName represents phyllis dutch name.
	phyllisDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Phyllis"}

	// phyllisFrenchName represents phyllis french name.
	phyllisFrenchName = nook.Name{
		Language: language.French,
		Value:    "Élisabec"}

	// phyllisGermanName represents phyllis german name.
	phyllisGermanName = nook.Name{
		Language: language.German,
		Value:    "Peggy"}

	// phyllisItalianName represents phyllis italian name.
	phyllisItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Polly"}

	// phyllisJapaneseName represents phyllis japanese name.
	phyllisJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ぺりみ"}

	// phyllisKoreanName represents phyllis korean name.
	phyllisKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "펠리미"}

	// phyllisLatinAmericanSpanishName represents phyllis latin american spanish name.
	phyllisLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Estrella"}

	// phyllisRussianName represents phyllis russian name.
	phyllisRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Филлис"}

	// phyllisSimplifiedChineseName represents phyllis simplified chinese name.
	phyllisSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "宋信美"}

	// phyllisSpanishName represents phyllis spanish name.
	phyllisSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Estrella"}

	// phyllisTraditionalChineseName represents phyllis traditional chinese name.
	phyllisTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "宋信美"}
)

var (
	// phyllisName represents phyllis name.
	phyllisName = nook.Languages{
		language.AmericanEnglish:      phyllisAmericanEnglishName,
		language.CanadianFrench:       phyllisCanadianFrenchName,
		language.Dutch:                phyllisDutchName,
		language.French:               phyllisFrenchName,
		language.German:               phyllisGermanName,
		language.Italian:              phyllisItalianName,
		language.Japanese:             phyllisJapaneseName,
		language.Korean:               phyllisKoreanName,
		language.LatinAmericanSpanish: phyllisLatinAmericanSpanishName,
		language.Russian:              phyllisRussianName,
		language.SimplifiedChinese:    phyllisSimplifiedChineseName,
		language.Spanish:              phyllisSpanishName,
		language.TraditionalChinese:   phyllisTraditionalChineseName}
)

var (
	// phyllisCharacter represents phyllis character.
	phyllisCharacter = nook.Character{
		Animal:   animal.Pelican,
		Birthday: phyllisBirthday,
		Code:     phyllisCode,
		Key:      character.Phyllis,
		Gender:   gender.Female,
		Name:     phyllisName,
		Special:  true}
)

var (
	// Phyllis represents phyllis.
	Phyllis = nook.Resident{
		Character: phyllisCharacter}
)
