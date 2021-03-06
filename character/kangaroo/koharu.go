package kangaroo

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
	koharuBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	koharuCode = nook.Code{
		Value: ""}
)

var (
	koharuAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Koharu"}

	koharuCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	koharuDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	koharuFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	koharuGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	koharuItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	koharuJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "こはる"}

	koharuKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	koharuLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	koharuRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	koharuSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	koharuSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	koharuTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	koharuName = nook.Languages{
		language.AmericanEnglish:      koharuAmericanEnglishName,
		language.CanadianFrench:       koharuCanadianFrenchName,
		language.Dutch:                koharuDutchName,
		language.French:               koharuFrenchName,
		language.German:               koharuGermanName,
		language.Italian:              koharuItalianName,
		language.Japanese:             koharuJapaneseName,
		language.Korean:               koharuKoreanName,
		language.LatinAmericanSpanish: koharuLatinAmericanSpanishName,
		language.Russian:              koharuRussianName,
		language.SimplifiedChinese:    koharuSimplifiedChineseName,
		language.Spanish:              koharuSpanishName,
		language.TraditionalChinese:   koharuTraditionalChineseName}
)

var (
	koharuCharacter = nook.Character{
		Animal:   animal.Kangaroo,
		Birthday: koharuBirthday,
		Code:     koharuCode,
		Key:      character.Koharu,
		Gender:   gender.Female,
		Name:     koharuName,
		Special:  false}
)

var (
	koharuAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ちょいと"}

	koharuCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	koharuDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	koharuFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	koharuGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	koharuItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	koharuJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ちょいと"}

	koharuKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	koharuLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	koharuRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	koharuSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	koharuSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	koharuTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	koharuPhrase = nook.Languages{
		language.AmericanEnglish:      koharuAmericanEnglishPhrase,
		language.CanadianFrench:       koharuCanadianFrenchPhrase,
		language.Dutch:                koharuDutchPhrase,
		language.French:               koharuFrenchPhrase,
		language.German:               koharuGermanPhrase,
		language.Italian:              koharuItalianPhrase,
		language.Japanese:             koharuJapanesePhrase,
		language.Korean:               koharuKoreanPhrase,
		language.LatinAmericanSpanish: koharuLatinAmericanSpanishPhrase,
		language.Russian:              koharuRussianPhrase,
		language.SimplifiedChinese:    koharuSimplifiedChinesePhrase,
		language.Spanish:              koharuSpanishPhrase,
		language.TraditionalChinese:   koharuTraditionalChinesePhrase}
)

var (
	Koharu = nook.Villager{
		Character:   koharuCharacter,
		Personality: personality.Peppy,
		Phrase:      koharuPhrase}
)
