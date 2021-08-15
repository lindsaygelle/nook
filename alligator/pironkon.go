package alligator

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	pironkonBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	pironkonCode = nook.Code{
		Value: ""}
)

var (
	pironkonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pironkon"}

	pironkonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	pironkonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	pironkonFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	pironkonGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	pironkonItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	pironkonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピロンコン"}

	pironkonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	pironkonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	pironkonRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	pironkonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	pironkonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	pironkonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	pironkonName = nook.Languages{
		language.AmericanEnglish:      pironkonAmericanEnglishName,
		language.CanadianFrench:       pironkonCanadianFrenchName,
		language.Dutch:                pironkonDutchName,
		language.French:               pironkonFrenchName,
		language.German:               pironkonGermanName,
		language.Italian:              pironkonItalianName,
		language.Japanese:             pironkonJapaneseName,
		language.Korean:               pironkonKoreanName,
		language.LatinAmericanSpanish: pironkonLatinAmericanSpanishName,
		language.Russian:              pironkonRussianName,
		language.SimplifiedChinese:    pironkonSimplifiedChineseName,
		language.Spanish:              pironkonSpanishName,
		language.TraditionalChinese:   pironkonTraditionalChineseName}
)

var (
	pironkonCharacter = nook.Character{
		Animal:   Alligator,
		Birthday: pironkonBirthday,
		Code:     pironkonCode,
		Gender:   nook.Male,
		Name:     pironkonName}
)

var (
	pironkonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "パヨマケ"}

	pironkonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	pironkonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	pironkonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	pironkonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	pironkonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	pironkonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "パヨマケ"}

	pironkonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	pironkonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	pironkonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	pironkonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	pironkonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	pironkonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	pironkonPhrase = nook.Languages{
		language.AmericanEnglish:      pironkonAmericanEnglishName,
		language.CanadianFrench:       pironkonCanadianFrenchName,
		language.Dutch:                pironkonDutchName,
		language.French:               pironkonFrenchName,
		language.German:               pironkonGermanName,
		language.Italian:              pironkonItalianName,
		language.Japanese:             pironkonJapaneseName,
		language.Korean:               pironkonKoreanName,
		language.LatinAmericanSpanish: pironkonLatinAmericanSpanishName,
		language.Russian:              pironkonRussianName,
		language.SimplifiedChinese:    pironkonSimplifiedChineseName,
		language.Spanish:              pironkonSpanishName,
		language.TraditionalChinese:   pironkonTraditionalChineseName}
)

var (
	Pironkon = nook.Villager{
		Character:   pironkonCharacter,
		Personality: nook.Lazy,
		Phrase:      pironkonPhrase}
)
