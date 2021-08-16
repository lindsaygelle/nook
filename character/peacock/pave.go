package peacock

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	paveBirthday = nook.Birthday{
		Day:   3,
		Month: time.March}
)

var (
	paveCode = nook.Code{
		Value: "pck"}
)

var (
	paveAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pavé"}

	paveCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Roberto"}

	paveDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pavé"}

	paveFrenchName = nook.Name{
		Language: language.French,
		Value:    "Roberto"}

	paveGermanName = nook.Name{
		Language: language.German,
		Value:    "Pavo"}

	paveItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pavão"}

	paveJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ベルリーナ"}

	paveKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "베르리나"}

	paveLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Conga"}

	paveRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Паве"}

	paveSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿欢"}

	paveSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Conga"}

	paveTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿歡"}
)

var (
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
	paveCharacter = nook.Character{
		Animal:   animal.Peacock,
		Birthday: paveBirthday,
		Code:     paveCode,
		Gender:   gender.Male,
		Name:     paveName}
)

var (
	Pave = nook.Resident{
		Character: paveCharacter}
)
