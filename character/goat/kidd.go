package goat

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
		Value:    "Moktar"}

	kiddDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kidd"}

	kiddFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mokhtar"}

	kiddGermanName = nook.Name{
		Language: language.German,
		Value:    "Bocki"}

	kiddItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Vittorio"}

	kiddJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "やさお"}

	kiddKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "염두리"}

	kiddLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cabrálex"}

	kiddRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кидд"}

	kiddSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "文青"}

	kiddSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cabrálex"}

	kiddTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "文青"}
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
		Animal:   animal.Goat,
		Birthday: kiddBirthday,
		Code:     kiddCode,
		Key:      character.Kidd,
		Gender:   gender.Male,
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
		language.AmericanEnglish:      kiddAmericanEnglishPhrase,
		language.CanadianFrench:       kiddCanadianFrenchPhrase,
		language.Dutch:                kiddDutchPhrase,
		language.French:               kiddFrenchPhrase,
		language.German:               kiddGermanPhrase,
		language.Italian:              kiddItalianPhrase,
		language.Japanese:             kiddJapanesePhrase,
		language.Korean:               kiddKoreanPhrase,
		language.LatinAmericanSpanish: kiddLatinAmericanSpanishPhrase,
		language.Russian:              kiddRussianPhrase,
		language.SimplifiedChinese:    kiddSimplifiedChinesePhrase,
		language.Spanish:              kiddSpanishPhrase,
		language.TraditionalChinese:   kiddTraditionalChinesePhrase}
)

var (
	Kidd = nook.Villager{
		Character:   kiddCharacter,
		Personality: personality.Smug,
		Phrase:      kiddPhrase}
)
