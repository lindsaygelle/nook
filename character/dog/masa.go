package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	masaBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	masaCode = nook.Code{
		Value: ""}
)

var (
	masaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Masa"}

	masaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	masaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	masaFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	masaGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	masaItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	masaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マサ"}

	masaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	masaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	masaRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	masaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	masaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	masaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	masaName = nook.Languages{
		language.AmericanEnglish:      masaAmericanEnglishName,
		language.CanadianFrench:       masaCanadianFrenchName,
		language.Dutch:                masaDutchName,
		language.French:               masaFrenchName,
		language.German:               masaGermanName,
		language.Italian:              masaItalianName,
		language.Japanese:             masaJapaneseName,
		language.Korean:               masaKoreanName,
		language.LatinAmericanSpanish: masaLatinAmericanSpanishName,
		language.Russian:              masaRussianName,
		language.SimplifiedChinese:    masaSimplifiedChineseName,
		language.Spanish:              masaSpanishName,
		language.TraditionalChinese:   masaTraditionalChineseName}
)

var (
	masaCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: masaBirthday,
		Code:     masaCode,
		Gender:   gender.Male,
		Name:     masaName}
)

var (
	masaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "てもんだ"}

	masaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	masaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	masaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	masaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	masaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	masaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "てもんだ"}

	masaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	masaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	masaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	masaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	masaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	masaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	masaPhrase = nook.Languages{
		language.AmericanEnglish:      masaAmericanEnglishName,
		language.CanadianFrench:       masaCanadianFrenchName,
		language.Dutch:                masaDutchName,
		language.French:               masaFrenchName,
		language.German:               masaGermanName,
		language.Italian:              masaItalianName,
		language.Japanese:             masaJapaneseName,
		language.Korean:               masaKoreanName,
		language.LatinAmericanSpanish: masaLatinAmericanSpanishName,
		language.Russian:              masaRussianName,
		language.SimplifiedChinese:    masaSimplifiedChineseName,
		language.Spanish:              masaSpanishName,
		language.TraditionalChinese:   masaTraditionalChineseName}
)

var (
	Masa = nook.Villager{
		Character:   masaCharacter,
		Personality: personality.Jock,
		Phrase:      masaPhrase}
)
