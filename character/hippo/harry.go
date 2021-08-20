package hippo

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
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
		Value:    "Bob"}

	harryDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Harry"}

	harryFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bob"}

	harryGermanName = nook.Name{
		Language: language.German,
		Value:    "Jürgen"}

	harryItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ercolino"}

	harryJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "オリバー"}

	harryKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "올리버"}

	harryLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Harpo"}

	harryRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гарри"}

	harrySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "欧世豪"}

	harrySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Harpo"}

	harryTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "歐世豪"}
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
		Animal:   animal.Hippo,
		Birthday: harryBirthday,
		Code:     harryCode,
		Key:      character.Harry,
		Gender:   gender.Male,
		Name:     harryName,
		Special:  false}
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
		language.AmericanEnglish:      harryAmericanEnglishPhrase,
		language.CanadianFrench:       harryCanadianFrenchPhrase,
		language.Dutch:                harryDutchPhrase,
		language.French:               harryFrenchPhrase,
		language.German:               harryGermanPhrase,
		language.Italian:              harryItalianPhrase,
		language.Japanese:             harryJapanesePhrase,
		language.Korean:               harryKoreanPhrase,
		language.LatinAmericanSpanish: harryLatinAmericanSpanishPhrase,
		language.Russian:              harryRussianPhrase,
		language.SimplifiedChinese:    harrySimplifiedChinesePhrase,
		language.Spanish:              harrySpanishPhrase,
		language.TraditionalChinese:   harryTraditionalChinesePhrase}
)

var (
	Harry = nook.Villager{
		Character:   harryCharacter,
		Personality: personality.Cranky,
		Phrase:      harryPhrase}
)
