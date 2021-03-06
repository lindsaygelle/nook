package koala

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
		Value:    ""}

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
		Value:    ""}

	huggyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mimí"}

	huggyRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	huggySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "包包"}

	huggySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mimí"}

	huggyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
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
		Animal:   animal.Koala,
		Birthday: huggyBirthday,
		Code:     huggyCode,
		Key:      character.Huggy,
		Gender:   gender.Female,
		Name:     huggyName,
		Special:  false}
)

var (
	huggyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bear"}

	huggyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "nounours"}

	huggyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

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
		Value:    ""}

	huggyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "osi-cosi"}

	huggyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	huggySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啀"}

	huggySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "osi-cosi"}

	huggyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	huggyPhrase = nook.Languages{
		language.AmericanEnglish:      huggyAmericanEnglishPhrase,
		language.CanadianFrench:       huggyCanadianFrenchPhrase,
		language.Dutch:                huggyDutchPhrase,
		language.French:               huggyFrenchPhrase,
		language.German:               huggyGermanPhrase,
		language.Italian:              huggyItalianPhrase,
		language.Japanese:             huggyJapanesePhrase,
		language.Korean:               huggyKoreanPhrase,
		language.LatinAmericanSpanish: huggyLatinAmericanSpanishPhrase,
		language.Russian:              huggyRussianPhrase,
		language.SimplifiedChinese:    huggySimplifiedChinesePhrase,
		language.Spanish:              huggySpanishPhrase,
		language.TraditionalChinese:   huggyTraditionalChinesePhrase}
)

var (
	Huggy = nook.Villager{
		Character:   huggyCharacter,
		Personality: personality.Peppy,
		Phrase:      huggyPhrase}
)
