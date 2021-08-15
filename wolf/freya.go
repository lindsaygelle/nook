package wolf

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	freyaBirthday = nook.Birthday{
		Day:   14,
		Month: time.December}
)

var (
	freyaCode = nook.Code{
		Value: "wol05"}
)

var (
	freyaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Freya"}

	freyaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Luppa"}

	freyaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Freya"}

	freyaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Luppa"}

	freyaGermanName = nook.Name{
		Language: language.German,
		Value:    "Freya"}

	freyaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Freya"}

	freyaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ツンドラ"}

	freyaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "산드라"}

	freyaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lupita"}

	freyaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фрея"}

	freyaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "冰冰"}

	freyaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lupita"}

	freyaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "冰冰"}
)

var (
	freyaName = nook.Languages{
		language.AmericanEnglish:      freyaAmericanEnglishName,
		language.CanadianFrench:       freyaCanadianFrenchName,
		language.Dutch:                freyaDutchName,
		language.French:               freyaFrenchName,
		language.German:               freyaGermanName,
		language.Italian:              freyaItalianName,
		language.Japanese:             freyaJapaneseName,
		language.Korean:               freyaKoreanName,
		language.LatinAmericanSpanish: freyaLatinAmericanSpanishName,
		language.Russian:              freyaRussianName,
		language.SimplifiedChinese:    freyaSimplifiedChineseName,
		language.Spanish:              freyaSpanishName,
		language.TraditionalChinese:   freyaTraditionalChineseName}
)

var (
	freyaCharacter = nook.Character{
		Animal:   Wolf,
		Birthday: freyaBirthday,
		Code:     freyaCode,
		Gender:   nook.Female,
		Name:     freyaName}
)

var (
	freyaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "uff da"}

	freyaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon loulou"}

	freyaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "fjord"}

	freyaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon loulou"}

	freyaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "ähmmm"}

	freyaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "uuuff"}

	freyaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのだわ"}

	freyaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "유행이야"}

	freyaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "auuuff"}

	freyaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрусть"}

	freyaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "就是说哇"}

	freyaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "auuuff"}

	freyaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "就是說哇"}
)

var (
	freyaPhrase = nook.Languages{
		language.AmericanEnglish:      freyaAmericanEnglishName,
		language.CanadianFrench:       freyaCanadianFrenchName,
		language.Dutch:                freyaDutchName,
		language.French:               freyaFrenchName,
		language.German:               freyaGermanName,
		language.Italian:              freyaItalianName,
		language.Japanese:             freyaJapaneseName,
		language.Korean:               freyaKoreanName,
		language.LatinAmericanSpanish: freyaLatinAmericanSpanishName,
		language.Russian:              freyaRussianName,
		language.SimplifiedChinese:    freyaSimplifiedChineseName,
		language.Spanish:              freyaSpanishName,
		language.TraditionalChinese:   freyaTraditionalChineseName}
)

var (
	Freya = nook.Villager{
		Character:   freyaCharacter,
		Personality: nook.Snooty,
		Phrase:      freyaPhrase}
)
