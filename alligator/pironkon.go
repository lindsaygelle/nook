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
		Value:    "N/A"}

	pironkonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	pironkonFrenchName = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	pironkonGermanName = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	pironkonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	pironkonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピロンコン"}

	pironkonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	pironkonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	pironkonRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	pironkonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	pironkonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	pironkonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Value:    ""}

	pironkonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	pironkonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	pironkonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	pironkonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	pironkonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	pironkonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "パヨマケ"}

	pironkonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	pironkonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	pironkonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	pironkonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	pironkonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	pironkonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
