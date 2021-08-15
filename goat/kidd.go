package goat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	kiddBirthday = nook.Birthday{
		Day:   28,
		Month: time.June}
)

var (
	kiddCode = nook.Code{
		Value: "goa07"}
)

var (
	kiddAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kidd"}

	kiddCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Moktarmouhai"}

	kiddDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kiddwatte"}

	kiddFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mokhtarmouhai"}

	kiddGermanName = nook.Name{
		Language: language.German,
		Value:    "Bockimehehe"}

	kiddItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Vittoriobalido"}

	kiddJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "やさおはぁ～"}

	kiddKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "염두리웁스"}

	kiddLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cabrálexein"}

	kiddRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Киддтак"}

	kiddSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "文青啊"}

	kiddSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cabrálexein"}

	kiddTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "文青啊"}
)

var (
	kiddName = nook.Languages{
		language.AmericanEnglish:      kiddAmericanEnglishName,
		language.CanadianFrench:       kiddCanadianFrenchName,
		language.Dutch:                kiddDutchName,
		language.French:               kiddFrenchName,
		language.German:               kiddGermanName,
		language.Italian:              kiddItalianName,
		language.Japanese:             kiddJapaneseName,
		language.Korean:               kiddKoreanName,
		language.LatinAmericanSpanish: kiddLatinAmericanSpanishName,
		language.Russian:              kiddRussianName,
		language.SimplifiedChinese:    kiddSimplifiedChineseName,
		language.Spanish:              kiddSpanishName,
		language.TraditionalChinese:   kiddTraditionalChineseName}
)

var (
	kiddCharacter = nook.Character{
		Animal:   Goat,
		Birthday: kiddBirthday,
		Code:     kiddCode,
		Gender:   nook.Male,
		Name:     kiddName}
)

var (
	kiddAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "wut"}

	kiddCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mouhai"}

	kiddDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "watte"}

	kiddFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mouhai"}

	kiddGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mehehe"}

	kiddItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "balido"}

	kiddJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "はぁ～"}

	kiddKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "웁스"}

	kiddLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ein"}

	kiddRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "так"}

	kiddSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啊"}

	kiddSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ein"}

	kiddTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啊"}
)

var (
	kiddPhrase = nook.Languages{
		language.AmericanEnglish:      kiddAmericanEnglishName,
		language.CanadianFrench:       kiddCanadianFrenchName,
		language.Dutch:                kiddDutchName,
		language.French:               kiddFrenchName,
		language.German:               kiddGermanName,
		language.Italian:              kiddItalianName,
		language.Japanese:             kiddJapaneseName,
		language.Korean:               kiddKoreanName,
		language.LatinAmericanSpanish: kiddLatinAmericanSpanishName,
		language.Russian:              kiddRussianName,
		language.SimplifiedChinese:    kiddSimplifiedChineseName,
		language.Spanish:              kiddSpanishName,
		language.TraditionalChinese:   kiddTraditionalChineseName}
)

var (
	Kidd = nook.Villager{
		Character:   kiddCharacter,
		Personality: nook.Smug,
		Phrase:      kiddPhrase}
)
