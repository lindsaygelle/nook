package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Émeraudesproiiing"}

	emeraldGermanName = nook.Name{
		Language: language.German,
		Value:    "Emmaboioioing"}

	emeraldItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Smeraldaboing"}

	emeraldJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ケロミだケロ"}

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
		Value:    "呱呱咕咕"}

	emeraldSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Espeesproing"}

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
		Animal:   Frog,
		Birthday: emeraldBirthday,
		Code:     emeraldCode,
		Gender:   nook.Female,
		Name:     emeraldName}
)

var (
	emeraldAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	emeraldCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	emeraldDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

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
		Value:    "N/A"}

	emeraldLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	emeraldRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	emeraldSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咕咕"}

	emeraldSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "esproing"}

	emeraldTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	emeraldPhrase = nook.Languages{
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
	Emerald = nook.Villager{
		Character:   emeraldCharacter,
		Personality: nook.Normal,
		Phrase:      emeraldPhrase}
)
