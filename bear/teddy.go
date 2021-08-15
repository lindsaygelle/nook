package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	teddyBirthday = nook.Birthday{
		Day:   26,
		Month: time.September}
)

var (
	teddyCode = nook.Code{
		Value: "bea00"}
)

var (
	teddyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Teddy"}

	teddyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Teddygrrrrrr"}

	teddyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Teddygrommm"}

	teddyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Teddygrrrrrr"}

	teddyGermanName = nook.Name{
		Language: language.German,
		Value:    "Torstengruff"}

	teddyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Teddygruf"}

	teddyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "たいへいたですたい"}

	teddyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "병태입네다"}

	teddyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Teddygruuuf"}

	teddyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Теддибыр-быр"}

	teddySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "太平太好了"}

	teddySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Teddygruuuf"}

	teddyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "太平太好了"}
)

var (
	teddyName = nook.Languages{
		language.AmericanEnglish:      teddyAmericanEnglishName,
		language.CanadianFrench:       teddyCanadianFrenchName,
		language.Dutch:                teddyDutchName,
		language.French:               teddyFrenchName,
		language.German:               teddyGermanName,
		language.Italian:              teddyItalianName,
		language.Japanese:             teddyJapaneseName,
		language.Korean:               teddyKoreanName,
		language.LatinAmericanSpanish: teddyLatinAmericanSpanishName,
		language.Russian:              teddyRussianName,
		language.SimplifiedChinese:    teddySimplifiedChineseName,
		language.Spanish:              teddySpanishName,
		language.TraditionalChinese:   teddyTraditionalChineseName}
)

var (
	teddyCharacter = nook.Character{
		Animal:   Bear,
		Birthday: teddyBirthday,
		Code:     teddyCode,
		Gender:   nook.Male,
		Name:     teddyName}
)

var (
	teddyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "grooof"}

	teddyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grrrrrr"}

	teddyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "grommm"}

	teddyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grrrrrr"}

	teddyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "gruff"}

	teddyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "gruf"}

	teddyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですたい"}

	teddyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "입네다"}

	teddyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gruuuf"}

	teddyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "быр-быр"}

	teddySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "太好了"}

	teddySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "gruuuf"}

	teddyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "太好了"}
)

var (
	teddyPhrase = nook.Languages{
		language.AmericanEnglish:      teddyAmericanEnglishName,
		language.CanadianFrench:       teddyCanadianFrenchName,
		language.Dutch:                teddyDutchName,
		language.French:               teddyFrenchName,
		language.German:               teddyGermanName,
		language.Italian:              teddyItalianName,
		language.Japanese:             teddyJapaneseName,
		language.Korean:               teddyKoreanName,
		language.LatinAmericanSpanish: teddyLatinAmericanSpanishName,
		language.Russian:              teddyRussianName,
		language.SimplifiedChinese:    teddySimplifiedChineseName,
		language.Spanish:              teddySpanishName,
		language.TraditionalChinese:   teddyTraditionalChineseName}
)

var (
	Teddy = nook.Villager{
		Character:   teddyCharacter,
		Personality: nook.Jock,
		Phrase:      teddyPhrase}
)
