package anteater

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	anabelleBirthday = nook.Birthday{
		Day:   16,
		Month: time.February}
)

var (
	anabelleCode = nook.Code{
		Value: "ant03"}
)

var (
	anabelleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Anabelle"}

	anabelleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Anabelle"}

	anabelleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Anabelle"}

	anabelleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Anabelle"}

	anabelleGermanName = nook.Name{
		Language: language.German,
		Value:    "Annabell"}

	anabelleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Annalisa"}

	anabelleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "あるみ"}

	anabelleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아롱이"}

	anabelleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Anabel"}

	anabelleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Анабель"}

	anabelleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "有美"}

	anabelleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Anabel"}

	anabelleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "有美"}
)

var (
	anabelleName = nook.Languages{
		language.AmericanEnglish:      anabelleAmericanEnglishName,
		language.CanadianFrench:       anabelleCanadianFrenchName,
		language.Dutch:                anabelleDutchName,
		language.French:               anabelleFrenchName,
		language.German:               anabelleGermanName,
		language.Italian:              anabelleItalianName,
		language.Japanese:             anabelleJapaneseName,
		language.Korean:               anabelleKoreanName,
		language.LatinAmericanSpanish: anabelleLatinAmericanSpanishName,
		language.Russian:              anabelleRussianName,
		language.SimplifiedChinese:    anabelleSimplifiedChineseName,
		language.Spanish:              anabelleSpanishName,
		language.TraditionalChinese:   anabelleTraditionalChineseName}
)

var (
	anabelleCharacter = nook.Character{
		Animal:   Anteater,
		Birthday: anabelleBirthday,
		Code:     anabelleCode,
		Gender:   nook.Female,
		Name:     anabelleName}
)

var (
	anabelleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snorty"}

	anabelleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grognon"}

	anabelleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snork"}

	anabelleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grognon"}

	anabelleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "puuust"}

	anabelleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "snorty"}

	anabelleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "マジでー"}

	anabelleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "정말"}

	anabelleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fa-fiú"}

	anabelleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрум-хрум"}

	anabelleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "真的假的"}

	anabelleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "esnoink"}

	anabelleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "真的假的"}
)

var (
	anabellePhrase = nook.Languages{
		language.AmericanEnglish:      anabelleAmericanEnglishName,
		language.CanadianFrench:       anabelleCanadianFrenchName,
		language.Dutch:                anabelleDutchName,
		language.French:               anabelleFrenchName,
		language.German:               anabelleGermanName,
		language.Italian:              anabelleItalianName,
		language.Japanese:             anabelleJapaneseName,
		language.Korean:               anabelleKoreanName,
		language.LatinAmericanSpanish: anabelleLatinAmericanSpanishName,
		language.Russian:              anabelleRussianName,
		language.SimplifiedChinese:    anabelleSimplifiedChineseName,
		language.Spanish:              anabelleSpanishName,
		language.TraditionalChinese:   anabelleTraditionalChineseName}
)

var (
	Anabelle = nook.Villager{
		Character:   anabelleCharacter,
		Personality: nook.Peppy,
		Phrase:      anabellePhrase}
)
