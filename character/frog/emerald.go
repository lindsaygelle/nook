package frog

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
	emeraldBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	emeraldCode = nook.Code{
		Value: ""}
)

var (
	emeraldAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Emerald"}

	emeraldCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	emeraldDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	emeraldFrenchName = nook.Name{
		Language: language.French,
		Value:    "Émeraude"}

	emeraldGermanName = nook.Name{
		Language: language.German,
		Value:    "Emma"}

	emeraldItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Smeralda"}

	emeraldJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ケロミ"}

	emeraldKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	emeraldLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	emeraldRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	emeraldSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呱呱"}

	emeraldSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Espe"}

	emeraldTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	emeraldName = nook.Languages{
		language.AmericanEnglish:      emeraldAmericanEnglishName,
		language.CanadianFrench:       emeraldCanadianFrenchName,
		language.Dutch:                emeraldDutchName,
		language.French:               emeraldFrenchName,
		language.German:               emeraldGermanName,
		language.Italian:              emeraldItalianName,
		language.Japanese:             emeraldJapaneseName,
		language.Korean:               emeraldKoreanName,
		language.LatinAmericanSpanish: emeraldLatinAmericanSpanishName,
		language.Russian:              emeraldRussianName,
		language.SimplifiedChinese:    emeraldSimplifiedChineseName,
		language.Spanish:              emeraldSpanishName,
		language.TraditionalChinese:   emeraldTraditionalChineseName}
)

var (
	emeraldCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: emeraldBirthday,
		Code:     emeraldCode,
		Key:      character.Emerald,
		Gender:   gender.Female,
		Name:     emeraldName}
)

var (
	emeraldAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sproing"}

	emeraldCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	emeraldDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	emeraldFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "sproiiing"}

	emeraldGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "boioioing"}

	emeraldItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "boing"}

	emeraldJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だケロ"}

	emeraldKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	emeraldLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	emeraldRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	emeraldSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咕咕"}

	emeraldSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "esproing"}

	emeraldTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	emeraldPhrase = nook.Languages{
		language.AmericanEnglish:      emeraldAmericanEnglishPhrase,
		language.CanadianFrench:       emeraldCanadianFrenchPhrase,
		language.Dutch:                emeraldDutchPhrase,
		language.French:               emeraldFrenchPhrase,
		language.German:               emeraldGermanPhrase,
		language.Italian:              emeraldItalianPhrase,
		language.Japanese:             emeraldJapanesePhrase,
		language.Korean:               emeraldKoreanPhrase,
		language.LatinAmericanSpanish: emeraldLatinAmericanSpanishPhrase,
		language.Russian:              emeraldRussianPhrase,
		language.SimplifiedChinese:    emeraldSimplifiedChinesePhrase,
		language.Spanish:              emeraldSpanishPhrase,
		language.TraditionalChinese:   emeraldTraditionalChinesePhrase}
)

var (
	Emerald = nook.Villager{
		Character:   emeraldCharacter,
		Personality: personality.Normal,
		Phrase:      emeraldPhrase}
)
