package bear

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
		Value:    "Teddy"}

	teddyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Teddy"}

	teddyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Teddy"}

	teddyGermanName = nook.Name{
		Language: language.German,
		Value:    "Torsten"}

	teddyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Teddy"}

	teddyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "たいへいた"}

	teddyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "병태"}

	teddyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Teddy"}

	teddyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тедди"}

	teddySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "太平"}

	teddySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Teddy"}

	teddyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "太平"}
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
		Animal:   animal.Bear,
		Birthday: teddyBirthday,
		Code:     teddyCode,
		Key:      character.Teddy,
		Gender:   gender.Male,
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
		Personality: personality.Jock,
		Phrase:      teddyPhrase}
)
