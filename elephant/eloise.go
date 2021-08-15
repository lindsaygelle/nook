package elephant

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Éloïsetooouut"}

	eloiseDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Eloisepwèèèèèp"}

	eloiseFrenchName = nook.Name{
		Language: language.French,
		Value:    "Éloïsetooouut"}

	eloiseGermanName = nook.Name{
		Language: language.German,
		Value:    "Frauketuutuuuu"}

	eloiseItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Eloisanappole"}

	eloiseJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エレフィンルン"}

	eloiseKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "엘레핀훨훨"}

	eloiseLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Eloísafruuuf"}

	eloiseRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Элоизатру-тут"}

	eloiseSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "艾勒芬鲁鲁"}

	eloiseSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Eloísaorejillas"}

	eloiseTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "艾勒芬魯魯"}
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
		Animal:   Elephant,
		Birthday: eloiseBirthday,
		Code:     eloiseCode,
		Gender:   nook.Female,
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
	Eloise = nook.Villager{
		Character:   eloiseCharacter,
		Personality: nook.Snooty,
		Phrase:      eloisePhrase}
)
