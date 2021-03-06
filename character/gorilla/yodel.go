package gorilla

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
	yodelBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	yodelCode = nook.Code{
		Value: ""}
)

var (
	yodelAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Yodel"}

	yodelCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	yodelDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	yodelFrenchName = nook.Name{
		Language: language.French,
		Value:    "Thomas"}

	yodelGermanName = nook.Name{
		Language: language.German,
		Value:    "Jürgen"}

	yodelItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Yodel"}

	yodelJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヨーデル"}

	yodelKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	yodelLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	yodelRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	yodelSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	yodelSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Asdrúbal"}

	yodelTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	yodelName = nook.Languages{
		language.AmericanEnglish:      yodelAmericanEnglishName,
		language.CanadianFrench:       yodelCanadianFrenchName,
		language.Dutch:                yodelDutchName,
		language.French:               yodelFrenchName,
		language.German:               yodelGermanName,
		language.Italian:              yodelItalianName,
		language.Japanese:             yodelJapaneseName,
		language.Korean:               yodelKoreanName,
		language.LatinAmericanSpanish: yodelLatinAmericanSpanishName,
		language.Russian:              yodelRussianName,
		language.SimplifiedChinese:    yodelSimplifiedChineseName,
		language.Spanish:              yodelSpanishName,
		language.TraditionalChinese:   yodelTraditionalChineseName}
)

var (
	yodelCharacter = nook.Character{
		Animal:   animal.Gorilla,
		Birthday: yodelBirthday,
		Code:     yodelCode,
		Key:      character.Yodel,
		Gender:   gender.Male,
		Name:     yodelName,
		Special:  false}
)

var (
	yodelAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "odelay"}

	yodelCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	yodelDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	yodelFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "yolahitoou"}

	yodelGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "jodeli"}

	yodelItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ullallà"}

	yodelJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヨホホー"}

	yodelKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	yodelLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	yodelRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	yodelSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	yodelSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "joya"}

	yodelTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	yodelPhrase = nook.Languages{
		language.AmericanEnglish:      yodelAmericanEnglishPhrase,
		language.CanadianFrench:       yodelCanadianFrenchPhrase,
		language.Dutch:                yodelDutchPhrase,
		language.French:               yodelFrenchPhrase,
		language.German:               yodelGermanPhrase,
		language.Italian:              yodelItalianPhrase,
		language.Japanese:             yodelJapanesePhrase,
		language.Korean:               yodelKoreanPhrase,
		language.LatinAmericanSpanish: yodelLatinAmericanSpanishPhrase,
		language.Russian:              yodelRussianPhrase,
		language.SimplifiedChinese:    yodelSimplifiedChinesePhrase,
		language.Spanish:              yodelSpanishPhrase,
		language.TraditionalChinese:   yodelTraditionalChinesePhrase}
)

var (
	Yodel = nook.Villager{
		Character:   yodelCharacter,
		Personality: personality.Lazy,
		Phrase:      yodelPhrase}
)
