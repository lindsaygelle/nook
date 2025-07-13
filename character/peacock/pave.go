package peacock

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// paveBirthday represents pave birthday.
	paveBirthday = nook.Birthday{
		Day:   3,
		Month: time.March}
)

var (
	// paveCode represents pave code.
	paveCode = nook.Code{
		Value: "pck"}
)

var (
	// paveAmericanEnglishName represents pave american english name.
	paveAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pavé"}

	// paveCanadianFrenchName represents pave canadian french name.
	paveCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Roberto"}

	// paveDutchName represents pave dutch name.
	paveDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pavé"}

	// paveFrenchName represents pave french name.
	paveFrenchName = nook.Name{
		Language: language.French,
		Value:    "Roberto"}

	// paveGermanName represents pave german name.
	paveGermanName = nook.Name{
		Language: language.German,
		Value:    "Pavo"}

	// paveItalianName represents pave italian name.
	paveItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pavão"}

	// paveJapaneseName represents pave japanese name.
	paveJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ベルリーナ"}

	// paveKoreanName represents pave korean name.
	paveKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "베르리나"}

	// paveLatinAmericanSpanishName represents pave latin american spanish name.
	paveLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Conga"}

	// paveRussianName represents pave russian name.
	paveRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Паве"}

	// paveSimplifiedChineseName represents pave simplified chinese name.
	paveSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿欢"}

	// paveSpanishName represents pave spanish name.
	paveSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Conga"}

	// paveTraditionalChineseName represents pave traditional chinese name.
	paveTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿歡"}
)

var (
	// paveName represents pave name.
	paveName = nook.Languages{
		language.AmericanEnglish:      paveAmericanEnglishName,
		language.CanadianFrench:       paveCanadianFrenchName,
		language.Dutch:                paveDutchName,
		language.French:               paveFrenchName,
		language.German:               paveGermanName,
		language.Italian:              paveItalianName,
		language.Japanese:             paveJapaneseName,
		language.Korean:               paveKoreanName,
		language.LatinAmericanSpanish: paveLatinAmericanSpanishName,
		language.Russian:              paveRussianName,
		language.SimplifiedChinese:    paveSimplifiedChineseName,
		language.Spanish:              paveSpanishName,
		language.TraditionalChinese:   paveTraditionalChineseName}
)

var (
	// paveCharacter represents pave character.
	paveCharacter = nook.Character{
		Animal:   animal.Peacock,
		Birthday: paveBirthday,
		Code:     paveCode,
		Key:      character.Pave,
		Gender:   gender.Male,
		Name:     paveName,
		Special:  true}
)

var (
	// Pave represents pave.
	Pave = nook.Resident{
		Character: paveCharacter}
)
