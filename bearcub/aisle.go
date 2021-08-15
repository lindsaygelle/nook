package bearcub

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	aisleBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	aisleCode = nook.Code{
		Value: ""}
)

var (
	aisleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Aisle"}

	aisleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	aisleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	aisleFrenchName = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	aisleGermanName = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	aisleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	aisleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アイル"}

	aisleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	aisleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	aisleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	aisleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	aisleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	aisleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	aisleName = nook.Languages{
		language.AmericanEnglish:      aisleAmericanEnglishName,
		language.CanadianFrench:       aisleCanadianFrenchName,
		language.Dutch:                aisleDutchName,
		language.French:               aisleFrenchName,
		language.German:               aisleGermanName,
		language.Italian:              aisleItalianName,
		language.Japanese:             aisleJapaneseName,
		language.Korean:               aisleKoreanName,
		language.LatinAmericanSpanish: aisleLatinAmericanSpanishName,
		language.Russian:              aisleRussianName,
		language.SimplifiedChinese:    aisleSimplifiedChineseName,
		language.Spanish:              aisleSpanishName,
		language.TraditionalChinese:   aisleTraditionalChineseName}
)

var (
	aisleCharacter = nook.Character{
		Animal:   Bearcub,
		Birthday: aisleBirthday,
		Code:     aisleCode,
		Gender:   nook.Male,
		Name:     aisleName}
)

var (
	aisleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	aisleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	aisleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	aisleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	aisleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	aisleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	aisleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "あーあ"}

	aisleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	aisleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	aisleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	aisleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	aisleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	aisleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	aislePhrase = nook.Languages{
		language.AmericanEnglish:      aisleAmericanEnglishName,
		language.CanadianFrench:       aisleCanadianFrenchName,
		language.Dutch:                aisleDutchName,
		language.French:               aisleFrenchName,
		language.German:               aisleGermanName,
		language.Italian:              aisleItalianName,
		language.Japanese:             aisleJapaneseName,
		language.Korean:               aisleKoreanName,
		language.LatinAmericanSpanish: aisleLatinAmericanSpanishName,
		language.Russian:              aisleRussianName,
		language.SimplifiedChinese:    aisleSimplifiedChineseName,
		language.Spanish:              aisleSpanishName,
		language.TraditionalChinese:   aisleTraditionalChineseName}
)

var (
	Aisle = nook.Villager{
		Character:   aisleCharacter,
		Personality: nook.Lazy,
		Phrase:      aislePhrase}
)
