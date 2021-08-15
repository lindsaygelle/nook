package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	champagneBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	champagneCode = nook.Code{
		Value: ""}
)

var (
	champagneAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Champagne"}

	champagneCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	champagneDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	champagneFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	champagneGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	champagneItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	champagneJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シャンパンったく"}

	champagneKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	champagneLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	champagneRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	champagneSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	champagneSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	champagneTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	champagneName = nook.Languages{
		language.AmericanEnglish:      champagneAmericanEnglishName,
		language.CanadianFrench:       champagneCanadianFrenchName,
		language.Dutch:                champagneDutchName,
		language.French:               champagneFrenchName,
		language.German:               champagneGermanName,
		language.Italian:              champagneItalianName,
		language.Japanese:             champagneJapaneseName,
		language.Korean:               champagneKoreanName,
		language.LatinAmericanSpanish: champagneLatinAmericanSpanishName,
		language.Russian:              champagneRussianName,
		language.SimplifiedChinese:    champagneSimplifiedChineseName,
		language.Spanish:              champagneSpanishName,
		language.TraditionalChinese:   champagneTraditionalChineseName}
)

var (
	champagneCharacter = nook.Character{
		Animal:   Dog,
		Birthday: champagneBirthday,
		Code:     champagneCode,
		Gender:   nook.Male,
		Name:     champagneName}
)

var (
	champagneAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	champagneCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	champagneDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	champagneFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	champagneGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	champagneItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	champagneJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ったく"}

	champagneKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	champagneLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	champagneRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	champagneSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	champagneSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	champagneTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	champagnePhrase = nook.Languages{
		language.AmericanEnglish:      champagneAmericanEnglishName,
		language.CanadianFrench:       champagneCanadianFrenchName,
		language.Dutch:                champagneDutchName,
		language.French:               champagneFrenchName,
		language.German:               champagneGermanName,
		language.Italian:              champagneItalianName,
		language.Japanese:             champagneJapaneseName,
		language.Korean:               champagneKoreanName,
		language.LatinAmericanSpanish: champagneLatinAmericanSpanishName,
		language.Russian:              champagneRussianName,
		language.SimplifiedChinese:    champagneSimplifiedChineseName,
		language.Spanish:              champagneSpanishName,
		language.TraditionalChinese:   champagneTraditionalChineseName}
)

var (
	Champagne = nook.Villager{
		Character:   champagneCharacter,
		Personality: nook.Cranky,
		Phrase:      champagnePhrase}
)
