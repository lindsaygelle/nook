package mouse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	pennyBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	pennyCode = nook.Code{
		Value: ""}
)

var (
	pennyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Penny"}

	pennyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Barbara"}

	pennyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	pennyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Barbara"}

	pennyGermanName = nook.Name{
		Language: language.German,
		Value:    "Susi"}

	pennyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Moretta"}

	pennyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ネズこ"}

	pennyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	pennyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Penélope"}

	pennyRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	pennySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小小"}

	pennySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Penélope"}

	pennyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	pennyName = nook.Languages{
		language.AmericanEnglish:      pennyAmericanEnglishName,
		language.CanadianFrench:       pennyCanadianFrenchName,
		language.Dutch:                pennyDutchName,
		language.French:               pennyFrenchName,
		language.German:               pennyGermanName,
		language.Italian:              pennyItalianName,
		language.Japanese:             pennyJapaneseName,
		language.Korean:               pennyKoreanName,
		language.LatinAmericanSpanish: pennyLatinAmericanSpanishName,
		language.Russian:              pennyRussianName,
		language.SimplifiedChinese:    pennySimplifiedChineseName,
		language.Spanish:              pennySpanishName,
		language.TraditionalChinese:   pennyTraditionalChineseName}
)

var (
	pennyCharacter = nook.Character{
		Animal:   Mouse,
		Birthday: pennyBirthday,
		Code:     pennyCode,
		Gender:   nook.Female,
		Name:     pennyName}
)

var (
	pennyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ska-WEAK"}

	pennyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "oh pétard!"}

	pennyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	pennyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "oh pétard!"}

	pennyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "piiquie"}

	pennyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "Unknown"}

	pennyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのさ"}

	pennyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	pennyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "carambita"}

	pennyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	pennySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Unknown"}

	pennySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "carambita"}

	pennyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	pennyPhrase = nook.Languages{
		language.AmericanEnglish:      pennyAmericanEnglishName,
		language.CanadianFrench:       pennyCanadianFrenchName,
		language.Dutch:                pennyDutchName,
		language.French:               pennyFrenchName,
		language.German:               pennyGermanName,
		language.Italian:              pennyItalianName,
		language.Japanese:             pennyJapaneseName,
		language.Korean:               pennyKoreanName,
		language.LatinAmericanSpanish: pennyLatinAmericanSpanishName,
		language.Russian:              pennyRussianName,
		language.SimplifiedChinese:    pennySimplifiedChineseName,
		language.Spanish:              pennySpanishName,
		language.TraditionalChinese:   pennyTraditionalChineseName}
)

var (
	Penny = nook.Villager{
		Character:   pennyCharacter,
		Personality: nook.Peppy,
		Phrase:      pennyPhrase}
)
