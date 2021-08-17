package elephant

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
	eloiseBirthday = nook.Birthday{
		Day:   8,
		Month: time.December}
)

var (
	eloiseCode = nook.Code{
		Value: "elp03"}
)

var (
	eloiseAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Eloise"}

	eloiseCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Éloïse"}

	eloiseDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Eloise"}

	eloiseFrenchName = nook.Name{
		Language: language.French,
		Value:    "Éloïse"}

	eloiseGermanName = nook.Name{
		Language: language.German,
		Value:    "Frauke"}

	eloiseItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Eloisa"}

	eloiseJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エレフィン"}

	eloiseKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "엘레핀"}

	eloiseLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Eloísa"}

	eloiseRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Элоиза"}

	eloiseSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "艾勒芬"}

	eloiseSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Eloísa"}

	eloiseTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "艾勒芬"}
)

var (
	eloiseName = nook.Languages{
		language.AmericanEnglish:      eloiseAmericanEnglishName,
		language.CanadianFrench:       eloiseCanadianFrenchName,
		language.Dutch:                eloiseDutchName,
		language.French:               eloiseFrenchName,
		language.German:               eloiseGermanName,
		language.Italian:              eloiseItalianName,
		language.Japanese:             eloiseJapaneseName,
		language.Korean:               eloiseKoreanName,
		language.LatinAmericanSpanish: eloiseLatinAmericanSpanishName,
		language.Russian:              eloiseRussianName,
		language.SimplifiedChinese:    eloiseSimplifiedChineseName,
		language.Spanish:              eloiseSpanishName,
		language.TraditionalChinese:   eloiseTraditionalChineseName}
)

var (
	eloiseCharacter = nook.Character{
		Animal:   animal.Elephant,
		Birthday: eloiseBirthday,
		Code:     eloiseCode,
		Key:      character.Eloise,
		Gender:   gender.Female,
		Name:     eloiseName}
)

var (
	eloiseAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "tooooot"}

	eloiseCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tooouut"}

	eloiseDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pwèèèèèp"}

	eloiseFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tooouut"}

	eloiseGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tuutuuuu"}

	eloiseItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "nappole"}

	eloiseJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ルン"}

	eloiseKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "훨훨"}

	eloiseLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fruuuf"}

	eloiseRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тру-тут"}

	eloiseSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鲁鲁"}

	eloiseSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "orejillas"}

	eloiseTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "魯魯"}
)

var (
	eloisePhrase = nook.Languages{
		language.AmericanEnglish:      eloiseAmericanEnglishPhrase,
		language.CanadianFrench:       eloiseCanadianFrenchPhrase,
		language.Dutch:                eloiseDutchPhrase,
		language.French:               eloiseFrenchPhrase,
		language.German:               eloiseGermanPhrase,
		language.Italian:              eloiseItalianPhrase,
		language.Japanese:             eloiseJapanesePhrase,
		language.Korean:               eloiseKoreanPhrase,
		language.LatinAmericanSpanish: eloiseLatinAmericanSpanishPhrase,
		language.Russian:              eloiseRussianPhrase,
		language.SimplifiedChinese:    eloiseSimplifiedChinesePhrase,
		language.Spanish:              eloiseSpanishPhrase,
		language.TraditionalChinese:   eloiseTraditionalChinesePhrase}
)

var (
	Eloise = nook.Villager{
		Character:   eloiseCharacter,
		Personality: personality.Snooty,
		Phrase:      eloisePhrase}
)
