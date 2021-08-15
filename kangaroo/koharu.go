package kangaroo

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "N/A"}

	koharuDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	koharuFrenchName = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	koharuGermanName = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	koharuItalianName = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	koharuJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "こはる"}

	koharuKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	koharuLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	koharuRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	koharuSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	koharuSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	koharuTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Animal:   Kangaroo,
		Birthday: koharuBirthday,
		Code:     koharuCode,
		Gender:   nook.Female,
		Name:     koharuName}
)

var (
	koharuAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	koharuCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	koharuDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	koharuFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	koharuGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	koharuItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	koharuJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ちょいと"}

	koharuKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	koharuLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	koharuRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	koharuSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	koharuSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	koharuTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	koharuPhrase = nook.Languages{
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
	Koharu = nook.Villager{
		Character:   koharuCharacter,
		Personality: nook.Peppy,
		Phrase:      koharuPhrase}
)
