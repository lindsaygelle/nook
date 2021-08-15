package hippo

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	harryBirthday = nook.Birthday{
		Day:   7,
		Month: time.January}
)

var (
	harryCode = nook.Code{
		Value: "hip08"}
)

var (
	harryAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Harry"}

	harryCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bobmille"}

	harryDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Harrystrandgast"}

	harryFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bobmille"}

	harryGermanName = nook.Name{
		Language: language.German,
		Value:    "Jürgenhöhöhö"}

	harryItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ercolinouhuhuh"}

	harryJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "オリバーでっせ"}

	harryKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "올리버냉큼오슈"}

	harryLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Harpoho-hop"}

	harryRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гаррилежебока"}

	harrySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "欧世豪世世"}

	harrySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Harpoalgas"}

	harryTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "歐世豪世世"}
)

var (
	harryName = nook.Languages{
		language.AmericanEnglish:      harryAmericanEnglishName,
		language.CanadianFrench:       harryCanadianFrenchName,
		language.Dutch:                harryDutchName,
		language.French:               harryFrenchName,
		language.German:               harryGermanName,
		language.Italian:              harryItalianName,
		language.Japanese:             harryJapaneseName,
		language.Korean:               harryKoreanName,
		language.LatinAmericanSpanish: harryLatinAmericanSpanishName,
		language.Russian:              harryRussianName,
		language.SimplifiedChinese:    harrySimplifiedChineseName,
		language.Spanish:              harrySpanishName,
		language.TraditionalChinese:   harryTraditionalChineseName}
)

var (
	harryCharacter = nook.Character{
		Animal:   Hippo,
		Birthday: harryBirthday,
		Code:     harryCode,
		Gender:   nook.Male,
		Name:     harryName}
)

var (
	harryAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "beach bum"}

	harryCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mille"}

	harryDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "strandgast"}

	harryFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mille"}

	harryGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "höhöhö"}

	harryItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "uhuhuh"}

	harryJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でっせ"}

	harryKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "냉큼오슈"}

	harryLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ho-hop"}

	harryRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "лежебока"}

	harrySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "世世"}

	harrySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "algas"}

	harryTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "世世"}
)

var (
	harryPhrase = nook.Languages{
		language.AmericanEnglish:      harryAmericanEnglishName,
		language.CanadianFrench:       harryCanadianFrenchName,
		language.Dutch:                harryDutchName,
		language.French:               harryFrenchName,
		language.German:               harryGermanName,
		language.Italian:              harryItalianName,
		language.Japanese:             harryJapaneseName,
		language.Korean:               harryKoreanName,
		language.LatinAmericanSpanish: harryLatinAmericanSpanishName,
		language.Russian:              harryRussianName,
		language.SimplifiedChinese:    harrySimplifiedChineseName,
		language.Spanish:              harrySpanishName,
		language.TraditionalChinese:   harryTraditionalChineseName}
)

var (
	Harry = nook.Villager{
		Character:   harryCharacter,
		Personality: nook.Cranky,
		Phrase:      harryPhrase}
)
