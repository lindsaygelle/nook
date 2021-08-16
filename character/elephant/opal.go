package elephant

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	opalBirthday = nook.Birthday{
		Day:   20,
		Month: time.January}
)

var (
	opalCode = nook.Code{
		Value: "elp00"}
)

var (
	opalAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Opal"}

	opalCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Opaline"}

	opalDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Opal"}

	opalFrenchName = nook.Name{
		Language: language.French,
		Value:    "Opaline"}

	opalGermanName = nook.Name{
		Language: language.German,
		Value:    "Olga"}

	opalItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Opal"}

	opalJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "オパール"}

	opalKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "오팔"}

	opalLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ópalo"}

	opalRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Опал"}

	opalSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "露露"}

	opalSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ópalo"}

	opalTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "露露"}
)

var (
	opalName = nook.Languages{
		language.AmericanEnglish:      opalAmericanEnglishName,
		language.CanadianFrench:       opalCanadianFrenchName,
		language.Dutch:                opalDutchName,
		language.French:               opalFrenchName,
		language.German:               opalGermanName,
		language.Italian:              opalItalianName,
		language.Japanese:             opalJapaneseName,
		language.Korean:               opalKoreanName,
		language.LatinAmericanSpanish: opalLatinAmericanSpanishName,
		language.Russian:              opalRussianName,
		language.SimplifiedChinese:    opalSimplifiedChineseName,
		language.Spanish:              opalSpanishName,
		language.TraditionalChinese:   opalTraditionalChineseName}
)

var (
	opalCharacter = nook.Character{
		Animal:   animal.Elephant,
		Birthday: opalBirthday,
		Code:     opalCode,
		Gender:   gender.Female,
		Name:     opalName}
)

var (
	opalAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snoot"}

	opalCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "gropif"}

	opalDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snoet"}

	opalFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "gropif"}

	opalGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tröööt"}

	opalItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "snoob"}

	opalJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヨン"}

	opalKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "땡"}

	opalLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "flip-flop"}

	opalRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хобот"}

	opalSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "勇"}

	opalSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "flip-flop"}

	opalTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "勇"}
)

var (
	opalPhrase = nook.Languages{
		language.AmericanEnglish:      opalAmericanEnglishName,
		language.CanadianFrench:       opalCanadianFrenchName,
		language.Dutch:                opalDutchName,
		language.French:               opalFrenchName,
		language.German:               opalGermanName,
		language.Italian:              opalItalianName,
		language.Japanese:             opalJapaneseName,
		language.Korean:               opalKoreanName,
		language.LatinAmericanSpanish: opalLatinAmericanSpanishName,
		language.Russian:              opalRussianName,
		language.SimplifiedChinese:    opalSimplifiedChineseName,
		language.Spanish:              opalSpanishName,
		language.TraditionalChinese:   opalTraditionalChineseName}
)

var (
	Opal = nook.Villager{
		Character:   opalCharacter,
		Personality: personality.Snooty,
		Phrase:      opalPhrase}
)
