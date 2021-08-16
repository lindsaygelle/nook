package hamster

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	hamphreyBirthday = nook.Birthday{
		Day:   25,
		Month: time.February}
)

var (
	hamphreyCode = nook.Code{
		Value: "ham07"}
)

var (
	hamphreyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hamphrey"}

	hamphreyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Charles"}

	hamphreyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hamphrey"}

	hamphreyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Charles"}

	hamphreyGermanName = nook.Name{
		Language: language.German,
		Value:    "Heinrich"}

	hamphreyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Nerone"}

	hamphreyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハムジ"}

	hamphreyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "햄쥐"}

	hamphreyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Arsenio"}

	hamphreyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Хэмфри"}

	hamphreySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小仓"}

	hamphreySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Arsenio"}

	hamphreyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小倉"}
)

var (
	hamphreyName = nook.Languages{
		language.AmericanEnglish:      hamphreyAmericanEnglishName,
		language.CanadianFrench:       hamphreyCanadianFrenchName,
		language.Dutch:                hamphreyDutchName,
		language.French:               hamphreyFrenchName,
		language.German:               hamphreyGermanName,
		language.Italian:              hamphreyItalianName,
		language.Japanese:             hamphreyJapaneseName,
		language.Korean:               hamphreyKoreanName,
		language.LatinAmericanSpanish: hamphreyLatinAmericanSpanishName,
		language.Russian:              hamphreyRussianName,
		language.SimplifiedChinese:    hamphreySimplifiedChineseName,
		language.Spanish:              hamphreySpanishName,
		language.TraditionalChinese:   hamphreyTraditionalChineseName}
)

var (
	hamphreyCharacter = nook.Character{
		Animal:   animal.Hamster,
		Birthday: hamphreyBirthday,
		Code:     hamphreyCode,
		Gender:   gender.Male,
		Name:     hamphreyName}
)

var (
	hamphreyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snort"}

	hamphreyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "gna gna"}

	hamphreyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snuifsnuif"}

	hamphreyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "balivernes"}

	hamphreyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "squiep"}

	hamphreyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "spatapumf"}

	hamphreyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "カーッ"}

	hamphreyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "카악"}

	hamphreyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "noquenó"}

	hamphreyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "фырк"}

	hamphreySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "怒"}

	hamphreySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "andaquenó"}

	hamphreyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "怒"}
)

var (
	hamphreyPhrase = nook.Languages{
		language.AmericanEnglish:      hamphreyAmericanEnglishName,
		language.CanadianFrench:       hamphreyCanadianFrenchName,
		language.Dutch:                hamphreyDutchName,
		language.French:               hamphreyFrenchName,
		language.German:               hamphreyGermanName,
		language.Italian:              hamphreyItalianName,
		language.Japanese:             hamphreyJapaneseName,
		language.Korean:               hamphreyKoreanName,
		language.LatinAmericanSpanish: hamphreyLatinAmericanSpanishName,
		language.Russian:              hamphreyRussianName,
		language.SimplifiedChinese:    hamphreySimplifiedChineseName,
		language.Spanish:              hamphreySpanishName,
		language.TraditionalChinese:   hamphreyTraditionalChineseName}
)

var (
	Hamphrey = nook.Villager{
		Character:   hamphreyCharacter,
		Personality: personality.Cranky,
		Phrase:      hamphreyPhrase}
)
