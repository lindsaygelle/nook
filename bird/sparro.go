package bird

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	sparroBirthday = nook.Birthday{
		Day:   20,
		Month: time.November}
)

var (
	sparroCode = nook.Code{
		Value: "brd18"}
)

var (
	sparroAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sparro"}

	sparroCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Dariuspiaaaaf"}

	sparroDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sparrowreed"}

	sparroFrenchName = nook.Name{
		Language: language.French,
		Value:    "Dariuspiaaaaf"}

	sparroGermanName = nook.Name{
		Language: language.German,
		Value:    "Nestorbirtz"}

	sparroItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Piumardoadiós"}

	sparroJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ちゅんのすけだチュン"}

	sparroKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "춘섭섭섭"}

	sparroLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jaimechirpi"}

	sparroRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Спарротак-так"}

	sparroSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "周之翔啁啁"}

	sparroSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jaimegorrión"}

	sparroTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "周之翔啁啁"}
)

var (
	sparroName = nook.Languages{
		language.AmericanEnglish:      sparroAmericanEnglishName,
		language.CanadianFrench:       sparroCanadianFrenchName,
		language.Dutch:                sparroDutchName,
		language.French:               sparroFrenchName,
		language.German:               sparroGermanName,
		language.Italian:              sparroItalianName,
		language.Japanese:             sparroJapaneseName,
		language.Korean:               sparroKoreanName,
		language.LatinAmericanSpanish: sparroLatinAmericanSpanishName,
		language.Russian:              sparroRussianName,
		language.SimplifiedChinese:    sparroSimplifiedChineseName,
		language.Spanish:              sparroSpanishName,
		language.TraditionalChinese:   sparroTraditionalChineseName}
)

var (
	sparroCharacter = nook.Character{
		Animal:   Bird,
		Birthday: sparroBirthday,
		Code:     sparroCode,
		Gender:   nook.Male,
		Name:     sparroName}
)

var (
	sparroAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "like whoa"}

	sparroCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "piaaaaf"}

	sparroDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "wreed"}

	sparroFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "piaaaaf"}

	sparroGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "birtz"}

	sparroItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "adiós"}

	sparroJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だチュン"}

	sparroKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "섭섭"}

	sparroLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chirpi"}

	sparroRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "так-так"}

	sparroSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啁啁"}

	sparroSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "gorrión"}

	sparroTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啁啁"}
)

var (
	sparroPhrase = nook.Languages{
		language.AmericanEnglish:      sparroAmericanEnglishName,
		language.CanadianFrench:       sparroCanadianFrenchName,
		language.Dutch:                sparroDutchName,
		language.French:               sparroFrenchName,
		language.German:               sparroGermanName,
		language.Italian:              sparroItalianName,
		language.Japanese:             sparroJapaneseName,
		language.Korean:               sparroKoreanName,
		language.LatinAmericanSpanish: sparroLatinAmericanSpanishName,
		language.Russian:              sparroRussianName,
		language.SimplifiedChinese:    sparroSimplifiedChineseName,
		language.Spanish:              sparroSpanishName,
		language.TraditionalChinese:   sparroTraditionalChineseName}
)

var (
	Sparro = nook.Villager{
		Character:   sparroCharacter,
		Personality: nook.Jock,
		Phrase:      sparroPhrase}
)
