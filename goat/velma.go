package goat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	velmaBirthday = nook.Birthday{
		Day:   14,
		Month: time.January}
)

var (
	velmaCode = nook.Code{
		Value: "goa06"}
)

var (
	velmaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Velma"}

	velmaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Vératu vois"}

	velmaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Velmameh"}

	velmaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Vératu vois"}

	velmaGermanName = nook.Name{
		Language: language.German,
		Value:    "Wilmamecker"}

	velmaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Belardablip"}

	velmaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピティエザーマス"}

	velmaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "벨마똑바로해"}

	velmaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Erikaverááás-tú"}

	velmaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Велмаме-е-е"}

	velmaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "班长起立"}

	velmaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Erikaveráaas-tú"}

	velmaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "班長起立"}
)

var (
	velmaName = nook.Languages{
		language.AmericanEnglish:      velmaAmericanEnglishName,
		language.CanadianFrench:       velmaCanadianFrenchName,
		language.Dutch:                velmaDutchName,
		language.French:               velmaFrenchName,
		language.German:               velmaGermanName,
		language.Italian:              velmaItalianName,
		language.Japanese:             velmaJapaneseName,
		language.Korean:               velmaKoreanName,
		language.LatinAmericanSpanish: velmaLatinAmericanSpanishName,
		language.Russian:              velmaRussianName,
		language.SimplifiedChinese:    velmaSimplifiedChineseName,
		language.Spanish:              velmaSpanishName,
		language.TraditionalChinese:   velmaTraditionalChineseName}
)

var (
	velmaCharacter = nook.Character{
		Animal:   Goat,
		Birthday: velmaBirthday,
		Code:     velmaCode,
		Gender:   nook.Female,
		Name:     velmaName}
)

var (
	velmaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "blih"}

	velmaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tu vois"}

	velmaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "meh"}

	velmaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tu vois"}

	velmaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mecker"}

	velmaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "blip"}

	velmaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ザーマス"}

	velmaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "똑바로해"}

	velmaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "verááás-tú"}

	velmaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ме-е-е"}

	velmaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "起立"}

	velmaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "veráaas-tú"}

	velmaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "起立"}
)

var (
	velmaPhrase = nook.Languages{
		language.AmericanEnglish:      velmaAmericanEnglishName,
		language.CanadianFrench:       velmaCanadianFrenchName,
		language.Dutch:                velmaDutchName,
		language.French:               velmaFrenchName,
		language.German:               velmaGermanName,
		language.Italian:              velmaItalianName,
		language.Japanese:             velmaJapaneseName,
		language.Korean:               velmaKoreanName,
		language.LatinAmericanSpanish: velmaLatinAmericanSpanishName,
		language.Russian:              velmaRussianName,
		language.SimplifiedChinese:    velmaSimplifiedChineseName,
		language.Spanish:              velmaSpanishName,
		language.TraditionalChinese:   velmaTraditionalChineseName}
)

var (
	Velma = nook.Villager{
		Character:   velmaCharacter,
		Personality: nook.Snooty,
		Phrase:      velmaPhrase}
)
