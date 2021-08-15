package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "N/A"}

	masaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	masaFrenchName = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	masaGermanName = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	masaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	masaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マサ"}

	masaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	masaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	masaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	masaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	masaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	masaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Animal:   Dog,
		Birthday: masaBirthday,
		Code:     masaCode,
		Gender:   nook.Male,
		Name:     masaName}
)

var (
	masaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	masaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	masaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	masaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	masaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	masaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	masaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "てもんだ"}

	masaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	masaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	masaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	masaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	masaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	masaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Personality: nook.Jock,
		Phrase:      masaPhrase}
)
