package mouse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	penelopeBirthday = nook.Birthday{
		Day:   5,
		Month: time.February}
)

var (
	penelopeCode = nook.Code{
		Value: "mus17"}
)

var (
	penelopeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Penelope"}

	penelopeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Missy"}

	penelopeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Penelope"}

	penelopeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Missy"}

	penelopeGermanName = nook.Name{
		Language: language.German,
		Value:    "Penelope"}

	penelopeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Penelope"}

	penelopeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チューこ"}

	penelopeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "찍순이"}

	penelopeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Adela"}

	penelopeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Адела"}

	penelopeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小啾"}

	penelopeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Adela"}

	penelopeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小啾"}
)

var (
	penelopeName = nook.Languages{
		language.AmericanEnglish:      penelopeAmericanEnglishName,
		language.CanadianFrench:       penelopeCanadianFrenchName,
		language.Dutch:                penelopeDutchName,
		language.French:               penelopeFrenchName,
		language.German:               penelopeGermanName,
		language.Italian:              penelopeItalianName,
		language.Japanese:             penelopeJapaneseName,
		language.Korean:               penelopeKoreanName,
		language.LatinAmericanSpanish: penelopeLatinAmericanSpanishName,
		language.Russian:              penelopeRussianName,
		language.SimplifiedChinese:    penelopeSimplifiedChineseName,
		language.Spanish:              penelopeSpanishName,
		language.TraditionalChinese:   penelopeTraditionalChineseName}
)

var (
	penelopeCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: penelopeBirthday,
		Code:     penelopeCode,
		Gender:   gender.Female,
		Name:     penelopeName}
)

var (
	penelopeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "oh bow"}

	penelopeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grugru"}

	penelopeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "strikje"}

	penelopeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grugru"}

	penelopeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bling"}

	penelopeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cheeese"}

	penelopeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でちゅの"}

	penelopeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그러찍"}

	penelopeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "brillilí"}

	penelopeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бантик"}

	penelopeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啾啾"}

	penelopeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "migas"}

	penelopeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啾啾"}
)

var (
	penelopePhrase = nook.Languages{
		language.AmericanEnglish:      penelopeAmericanEnglishName,
		language.CanadianFrench:       penelopeCanadianFrenchName,
		language.Dutch:                penelopeDutchName,
		language.French:               penelopeFrenchName,
		language.German:               penelopeGermanName,
		language.Italian:              penelopeItalianName,
		language.Japanese:             penelopeJapaneseName,
		language.Korean:               penelopeKoreanName,
		language.LatinAmericanSpanish: penelopeLatinAmericanSpanishName,
		language.Russian:              penelopeRussianName,
		language.SimplifiedChinese:    penelopeSimplifiedChineseName,
		language.Spanish:              penelopeSpanishName,
		language.TraditionalChinese:   penelopeTraditionalChineseName}
)

var (
	Penelope = nook.Villager{
		Character:   penelopeCharacter,
		Personality: personality.Peppy,
		Phrase:      penelopePhrase}
)
