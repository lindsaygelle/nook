package bearcub

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
		Value:    ""}

	aisleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	aisleFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	aisleGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	aisleItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	aisleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アイル"}

	aisleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	aisleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	aisleRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	aisleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	aisleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	aisleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
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
		Animal:   animal.Bearcub,
		Birthday: aisleBirthday,
		Code:     aisleCode,
		Key:      character.Aisle,
		Gender:   gender.Male,
		Name:     aisleName}
)

var (
	aisleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "あーあ"}

	aisleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	aisleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	aisleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	aisleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	aisleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	aisleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "あーあ"}

	aisleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	aisleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	aisleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	aisleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	aisleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	aisleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	aislePhrase = nook.Languages{
		language.AmericanEnglish:      aisleAmericanEnglishPhrase,
		language.CanadianFrench:       aisleCanadianFrenchPhrase,
		language.Dutch:                aisleDutchPhrase,
		language.French:               aisleFrenchPhrase,
		language.German:               aisleGermanPhrase,
		language.Italian:              aisleItalianPhrase,
		language.Japanese:             aisleJapanesePhrase,
		language.Korean:               aisleKoreanPhrase,
		language.LatinAmericanSpanish: aisleLatinAmericanSpanishPhrase,
		language.Russian:              aisleRussianPhrase,
		language.SimplifiedChinese:    aisleSimplifiedChinesePhrase,
		language.Spanish:              aisleSpanishPhrase,
		language.TraditionalChinese:   aisleTraditionalChinesePhrase}
)

var (
	Aisle = nook.Villager{
		Character:   aisleCharacter,
		Personality: personality.Lazy,
		Phrase:      aislePhrase}
)
