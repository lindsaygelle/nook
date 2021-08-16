package goat

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
	chevreBirthday = nook.Birthday{
		Day:   6,
		Month: time.March}
)

var (
	chevreCode = nook.Code{
		Value: "goa00"}
)

var (
	chevreAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chevre"}

	chevreCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Biquette"}

	chevreDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chevre"}

	chevreFrenchName = nook.Name{
		Language: language.French,
		Value:    "Biquette"}

	chevreGermanName = nook.Name{
		Language: language.German,
		Value:    "Anette"}

	chevreItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Diletta"}

	chevreJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ユキ"}

	chevreKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "윤이"}

	chevreLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cabriola"}

	chevreRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Шевр"}

	chevreSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雪儿"}

	chevreSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cabriola"}

	chevreTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雪兒"}
)

var (
	chevreName = nook.Languages{
		language.AmericanEnglish:      chevreAmericanEnglishName,
		language.CanadianFrench:       chevreCanadianFrenchName,
		language.Dutch:                chevreDutchName,
		language.French:               chevreFrenchName,
		language.German:               chevreGermanName,
		language.Italian:              chevreItalianName,
		language.Japanese:             chevreJapaneseName,
		language.Korean:               chevreKoreanName,
		language.LatinAmericanSpanish: chevreLatinAmericanSpanishName,
		language.Russian:              chevreRussianName,
		language.SimplifiedChinese:    chevreSimplifiedChineseName,
		language.Spanish:              chevreSpanishName,
		language.TraditionalChinese:   chevreTraditionalChineseName}
)

var (
	chevreCharacter = nook.Character{
		Animal:   animal.Goat,
		Birthday: chevreBirthday,
		Code:     chevreCode,
		Key:      character.Chevre,
		Gender:   gender.Female,
		Name:     chevreName}
)

var (
	chevreAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "la baa"}

	chevreCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "la baa"}

	chevreDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "la mèè"}

	chevreFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "la baa"}

	chevreGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "määääh"}

	chevreItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "beeello"}

	chevreJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "っぺ"}

	chevreKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "맞아유"}

	chevreLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "beee-beee"}

	chevreRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ле ме-е-е"}

	chevreSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呗"}

	chevreSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "beee-beee"}

	chevreTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唄"}
)

var (
	chevrePhrase = nook.Languages{
		language.AmericanEnglish:      chevreAmericanEnglishName,
		language.CanadianFrench:       chevreCanadianFrenchName,
		language.Dutch:                chevreDutchName,
		language.French:               chevreFrenchName,
		language.German:               chevreGermanName,
		language.Italian:              chevreItalianName,
		language.Japanese:             chevreJapaneseName,
		language.Korean:               chevreKoreanName,
		language.LatinAmericanSpanish: chevreLatinAmericanSpanishName,
		language.Russian:              chevreRussianName,
		language.SimplifiedChinese:    chevreSimplifiedChineseName,
		language.Spanish:              chevreSpanishName,
		language.TraditionalChinese:   chevreTraditionalChineseName}
)

var (
	Chevre = nook.Villager{
		Character:   chevreCharacter,
		Personality: personality.Normal,
		Phrase:      chevrePhrase}
)
