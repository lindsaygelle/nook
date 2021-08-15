package koala

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	lymanBirthday = nook.Birthday{
		Day:   12,
		Month: time.October}
)

var (
	lymanCode = nook.Code{
		Value: "kal09"}
)

var (
	lymanAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lyman"}

	lymanCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kalyptussandwich"}

	lymanDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lymanchips"}

	lymanFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kalyptussandwich"}

	lymanGermanName = nook.Name{
		Language: language.German,
		Value:    "Pepeliptus"}

	lymanItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Nicolaponzipò"}

	lymanJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "オズモンドわ～い"}

	lymanKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "오즈먼드우와"}

	lymanLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Chipitralalá"}

	lymanRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лайманничего"}

	lymanSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "敖志明哇"}

	lymanSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Chipitrololó"}

	lymanTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "敖志明哇～"}
)

var (
	lymanName = nook.Languages{
		language.AmericanEnglish:      lymanAmericanEnglishName,
		language.CanadianFrench:       lymanCanadianFrenchName,
		language.Dutch:                lymanDutchName,
		language.French:               lymanFrenchName,
		language.German:               lymanGermanName,
		language.Italian:              lymanItalianName,
		language.Japanese:             lymanJapaneseName,
		language.Korean:               lymanKoreanName,
		language.LatinAmericanSpanish: lymanLatinAmericanSpanishName,
		language.Russian:              lymanRussianName,
		language.SimplifiedChinese:    lymanSimplifiedChineseName,
		language.Spanish:              lymanSpanishName,
		language.TraditionalChinese:   lymanTraditionalChineseName}
)

var (
	lymanCharacter = nook.Character{
		Animal:   Koala,
		Birthday: lymanBirthday,
		Code:     lymanCode,
		Gender:   nook.Male,
		Name:     lymanName}
)

var (
	lymanAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chips"}

	lymanCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "sandwich"}

	lymanDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "chips"}

	lymanFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "sandwich"}

	lymanGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "liptus"}

	lymanItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ponzipò"}

	lymanJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "わ～い"}

	lymanKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우와"}

	lymanLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tralalá"}

	lymanRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ничего"}

	lymanSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哇"}

	lymanSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "trololó"}

	lymanTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哇～"}
)

var (
	lymanPhrase = nook.Languages{
		language.AmericanEnglish:      lymanAmericanEnglishName,
		language.CanadianFrench:       lymanCanadianFrenchName,
		language.Dutch:                lymanDutchName,
		language.French:               lymanFrenchName,
		language.German:               lymanGermanName,
		language.Italian:              lymanItalianName,
		language.Japanese:             lymanJapaneseName,
		language.Korean:               lymanKoreanName,
		language.LatinAmericanSpanish: lymanLatinAmericanSpanishName,
		language.Russian:              lymanRussianName,
		language.SimplifiedChinese:    lymanSimplifiedChineseName,
		language.Spanish:              lymanSpanishName,
		language.TraditionalChinese:   lymanTraditionalChineseName}
)

var (
	Lyman = nook.Villager{
		Character:   lymanCharacter,
		Personality: nook.Jock,
		Phrase:      lymanPhrase}
)
