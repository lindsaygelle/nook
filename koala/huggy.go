package koala

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	huggyBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	huggyCode = nook.Code{
		Value: ""}
)

var (
	huggyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Huggy"}

	huggyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tina"}

	huggyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	huggyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tina"}

	huggyGermanName = nook.Name{
		Language: language.German,
		Value:    "Heike"}

	huggyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Colette"}

	huggyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ダッコ"}

	huggyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	huggyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mimí"}

	huggyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	huggySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "包包"}

	huggySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mimí"}

	huggyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	huggyName = nook.Languages{
		language.AmericanEnglish:      huggyAmericanEnglishName,
		language.CanadianFrench:       huggyCanadianFrenchName,
		language.Dutch:                huggyDutchName,
		language.French:               huggyFrenchName,
		language.German:               huggyGermanName,
		language.Italian:              huggyItalianName,
		language.Japanese:             huggyJapaneseName,
		language.Korean:               huggyKoreanName,
		language.LatinAmericanSpanish: huggyLatinAmericanSpanishName,
		language.Russian:              huggyRussianName,
		language.SimplifiedChinese:    huggySimplifiedChineseName,
		language.Spanish:              huggySpanishName,
		language.TraditionalChinese:   huggyTraditionalChineseName}
)

var (
	huggyCharacter = nook.Character{
		Animal:   Koala,
		Birthday: huggyBirthday,
		Code:     huggyCode,
		Gender:   nook.Female,
		Name:     huggyName}
)

var (
	huggyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	huggyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "nounours"}

	huggyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	huggyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "nounours"}

	huggyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bärchen"}

	huggyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "coalilà"}

	huggyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "エヘ"}

	huggyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	huggyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "osi-cosi"}

	huggyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	huggySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啀"}

	huggySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "osi-cosi"}

	huggyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	huggyPhrase = nook.Languages{
		language.AmericanEnglish:      huggyAmericanEnglishName,
		language.CanadianFrench:       huggyCanadianFrenchName,
		language.Dutch:                huggyDutchName,
		language.French:               huggyFrenchName,
		language.German:               huggyGermanName,
		language.Italian:              huggyItalianName,
		language.Japanese:             huggyJapaneseName,
		language.Korean:               huggyKoreanName,
		language.LatinAmericanSpanish: huggyLatinAmericanSpanishName,
		language.Russian:              huggyRussianName,
		language.SimplifiedChinese:    huggySimplifiedChineseName,
		language.Spanish:              huggySpanishName,
		language.TraditionalChinese:   huggyTraditionalChineseName}
)

var (
	Huggy = nook.Villager{
		Character:   huggyCharacter,
		Personality: nook.Peppy,
		Phrase:      huggyPhrase}
)
