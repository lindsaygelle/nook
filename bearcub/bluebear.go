package bearcub

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	bluebearBirthday = nook.Birthday{
		Day:   24,
		Month: time.June}
)

var (
	bluebearCode = nook.Code{
		Value: "cbr00"}
)

var (
	bluebearAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bluebear"}

	bluebearCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Myrtilleprune"}

	bluebearDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bluebearperzik"}

	bluebearFrenchName = nook.Name{
		Language: language.French,
		Value:    "Myrtilleprune"}

	bluebearGermanName = nook.Name{
		Language: language.German,
		Value:    "Birtefrüchtchen"}

	bluebearItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Azzurrapepè"}

	bluebearJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "グルミンキュン"}

	bluebearKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "글루민두근"}

	bluebearLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Celestemelí-melá"}

	bluebearRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Блюбеарперсик"}

	bluebearSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "娃娃怦怦"}

	bluebearSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Celestecielito"}

	bluebearTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "娃娃怦怦"}
)

var (
	bluebearName = nook.Languages{
		language.AmericanEnglish:      bluebearAmericanEnglishName,
		language.CanadianFrench:       bluebearCanadianFrenchName,
		language.Dutch:                bluebearDutchName,
		language.French:               bluebearFrenchName,
		language.German:               bluebearGermanName,
		language.Italian:              bluebearItalianName,
		language.Japanese:             bluebearJapaneseName,
		language.Korean:               bluebearKoreanName,
		language.LatinAmericanSpanish: bluebearLatinAmericanSpanishName,
		language.Russian:              bluebearRussianName,
		language.SimplifiedChinese:    bluebearSimplifiedChineseName,
		language.Spanish:              bluebearSpanishName,
		language.TraditionalChinese:   bluebearTraditionalChineseName}
)

var (
	bluebearCharacter = nook.Character{
		Animal:   Bearcub,
		Birthday: bluebearBirthday,
		Code:     bluebearCode,
		Gender:   nook.Female,
		Name:     bluebearName}
)

var (
	bluebearAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "peach"}

	bluebearCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "prune"}

	bluebearDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "perzik"}

	bluebearFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "prune"}

	bluebearGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "früchtchen"}

	bluebearItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pepè"}

	bluebearJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "キュン"}

	bluebearKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "두근"}

	bluebearLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "melí-melá"}

	bluebearRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "персик"}

	bluebearSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "怦怦"}

	bluebearSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cielito"}

	bluebearTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "怦怦"}
)

var (
	bluebearPhrase = nook.Languages{
		language.AmericanEnglish:      bluebearAmericanEnglishName,
		language.CanadianFrench:       bluebearCanadianFrenchName,
		language.Dutch:                bluebearDutchName,
		language.French:               bluebearFrenchName,
		language.German:               bluebearGermanName,
		language.Italian:              bluebearItalianName,
		language.Japanese:             bluebearJapaneseName,
		language.Korean:               bluebearKoreanName,
		language.LatinAmericanSpanish: bluebearLatinAmericanSpanishName,
		language.Russian:              bluebearRussianName,
		language.SimplifiedChinese:    bluebearSimplifiedChineseName,
		language.Spanish:              bluebearSpanishName,
		language.TraditionalChinese:   bluebearTraditionalChineseName}
)

var (
	Bluebear = nook.Villager{
		Character:   bluebearCharacter,
		Personality: nook.Peppy,
		Phrase:      bluebearPhrase}
)
