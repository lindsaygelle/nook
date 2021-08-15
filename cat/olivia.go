package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	oliviaBirthday = nook.Birthday{
		Day:   3,
		Month: time.February}
)

var (
	oliviaCode = nook.Code{
		Value: "cat03"}
)

var (
	oliviaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Olivia"}

	oliviaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Olivia"}

	oliviaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Olivia"}

	oliviaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Olivia"}

	oliviaGermanName = nook.Name{
		Language: language.German,
		Value:    "Bianca"}

	oliviaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Micina"}

	oliviaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "オリビア"}

	oliviaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "올리비아"}

	oliviaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Olivia"}

	oliviaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Оливия"}

	oliviaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "奥莉"}

	oliviaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Olivia"}

	oliviaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "奧莉"}
)

var (
	oliviaName = nook.Languages{
		language.AmericanEnglish:      oliviaAmericanEnglishName,
		language.CanadianFrench:       oliviaCanadianFrenchName,
		language.Dutch:                oliviaDutchName,
		language.French:               oliviaFrenchName,
		language.German:               oliviaGermanName,
		language.Italian:              oliviaItalianName,
		language.Japanese:             oliviaJapaneseName,
		language.Korean:               oliviaKoreanName,
		language.LatinAmericanSpanish: oliviaLatinAmericanSpanishName,
		language.Russian:              oliviaRussianName,
		language.SimplifiedChinese:    oliviaSimplifiedChineseName,
		language.Spanish:              oliviaSpanishName,
		language.TraditionalChinese:   oliviaTraditionalChineseName}
)

var (
	oliviaCharacter = nook.Character{
		Animal:   Cat,
		Birthday: oliviaBirthday,
		Code:     oliviaCode,
		Gender:   nook.Female,
		Name:     oliviaName}
)

var (
	oliviaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "purrr"}

	oliviaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "rrrrrrrrr"}

	oliviaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "prrrrrr"}

	oliviaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "rrrrrrrrr"}

	oliviaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnurr"}

	oliviaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "purrr"}

	oliviaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なんやん"}

	oliviaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "머냐"}

	oliviaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "michimiau"}

	oliviaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мур-мяу"}

	oliviaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "什么"}

	oliviaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "raspas"}

	oliviaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "什麼"}
)

var (
	oliviaPhrase = nook.Languages{
		language.AmericanEnglish:      oliviaAmericanEnglishName,
		language.CanadianFrench:       oliviaCanadianFrenchName,
		language.Dutch:                oliviaDutchName,
		language.French:               oliviaFrenchName,
		language.German:               oliviaGermanName,
		language.Italian:              oliviaItalianName,
		language.Japanese:             oliviaJapaneseName,
		language.Korean:               oliviaKoreanName,
		language.LatinAmericanSpanish: oliviaLatinAmericanSpanishName,
		language.Russian:              oliviaRussianName,
		language.SimplifiedChinese:    oliviaSimplifiedChineseName,
		language.Spanish:              oliviaSpanishName,
		language.TraditionalChinese:   oliviaTraditionalChineseName}
)

var (
	Olivia = nook.Villager{
		Character:   oliviaCharacter,
		Personality: nook.Snooty,
		Phrase:      oliviaPhrase}
)
