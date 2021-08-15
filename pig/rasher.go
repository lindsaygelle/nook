package pig

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	rasherBirthday = nook.Birthday{
		Day:   7,
		Month: time.April}
)

var (
	rasherCode = nook.Code{
		Value: "pig02"}
)

var (
	rasherAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rasher"}

	rasherCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Salamiporcelet"}

	rasherDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rasherbig"}

	rasherFrenchName = nook.Name{
		Language: language.French,
		Value:    "Salamiporcelet"}

	rasherGermanName = nook.Name{
		Language: language.German,
		Value:    "Edegronk"}

	rasherItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bronciosnoooink"}

	rasherJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "グレオまんねん"}

	rasherKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "글레이꾸엑"}

	rasherLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Curtisesnoink"}

	rasherRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рашерхряк"}

	rasherSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "古烈万年"}

	rasherSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Curtisporcino"}

	rasherTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "古烈萬年"}
)

var (
	rasherName = nook.Languages{
		language.AmericanEnglish:      rasherAmericanEnglishName,
		language.CanadianFrench:       rasherCanadianFrenchName,
		language.Dutch:                rasherDutchName,
		language.French:               rasherFrenchName,
		language.German:               rasherGermanName,
		language.Italian:              rasherItalianName,
		language.Japanese:             rasherJapaneseName,
		language.Korean:               rasherKoreanName,
		language.LatinAmericanSpanish: rasherLatinAmericanSpanishName,
		language.Russian:              rasherRussianName,
		language.SimplifiedChinese:    rasherSimplifiedChineseName,
		language.Spanish:              rasherSpanishName,
		language.TraditionalChinese:   rasherTraditionalChineseName}
)

var (
	rasherCharacter = nook.Character{
		Animal:   Pig,
		Birthday: rasherBirthday,
		Code:     rasherCode,
		Gender:   nook.Male,
		Name:     rasherName}
)

var (
	rasherAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "swine"}

	rasherCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "porcelet"}

	rasherDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "big"}

	rasherFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "porcelet"}

	rasherGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "gronk"}

	rasherItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "snoooink"}

	rasherJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "まんねん"}

	rasherKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꾸엑"}

	rasherLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "esnoink"}

	rasherRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хряк"}

	rasherSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "万年"}

	rasherSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "porcino"}

	rasherTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "萬年"}
)

var (
	rasherPhrase = nook.Languages{
		language.AmericanEnglish:      rasherAmericanEnglishName,
		language.CanadianFrench:       rasherCanadianFrenchName,
		language.Dutch:                rasherDutchName,
		language.French:               rasherFrenchName,
		language.German:               rasherGermanName,
		language.Italian:              rasherItalianName,
		language.Japanese:             rasherJapaneseName,
		language.Korean:               rasherKoreanName,
		language.LatinAmericanSpanish: rasherLatinAmericanSpanishName,
		language.Russian:              rasherRussianName,
		language.SimplifiedChinese:    rasherSimplifiedChineseName,
		language.Spanish:              rasherSpanishName,
		language.TraditionalChinese:   rasherTraditionalChineseName}
)

var (
	Rasher = nook.Villager{
		Character:   rasherCharacter,
		Personality: nook.Cranky,
		Phrase:      rasherPhrase}
)
