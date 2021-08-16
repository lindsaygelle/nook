package duck

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	mollyBirthday = nook.Birthday{
		Day:   7,
		Month: time.March}
)

var (
	mollyCode = nook.Code{
		Value: "duk16"}
)

var (
	mollyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Molly"}

	mollyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Molly"}

	mollyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Molly"}

	mollyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Molly"}

	mollyGermanName = nook.Name{
		Language: language.German,
		Value:    "Monika"}

	mollyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Molly"}

	mollyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カモミ"}

	mollyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "귀오미"}

	mollyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Deira"}

	mollyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Молли"}

	mollySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "亚美"}

	mollySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Deira"}

	mollyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "亞美"}
)

var (
	mollyName = nook.Languages{
		language.AmericanEnglish:      mollyAmericanEnglishName,
		language.CanadianFrench:       mollyCanadianFrenchName,
		language.Dutch:                mollyDutchName,
		language.French:               mollyFrenchName,
		language.German:               mollyGermanName,
		language.Italian:              mollyItalianName,
		language.Japanese:             mollyJapaneseName,
		language.Korean:               mollyKoreanName,
		language.LatinAmericanSpanish: mollyLatinAmericanSpanishName,
		language.Russian:              mollyRussianName,
		language.SimplifiedChinese:    mollySimplifiedChineseName,
		language.Spanish:              mollySpanishName,
		language.TraditionalChinese:   mollyTraditionalChineseName}
)

var (
	mollyCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: mollyBirthday,
		Code:     mollyCode,
		Gender:   gender.Female,
		Name:     mollyName}
)

var (
	mollyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "quackidee"}

	mollyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "canénette"}

	mollyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "topper"}

	mollyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "canénette"}

	mollyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "krümel"}

	mollyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "coin coin"}

	mollyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "そうかも"}

	mollyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그럴지도"}

	mollyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cuáquidi"}

	mollyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "крякульки"}

	mollySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "说不定鸭"}

	mollySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "deverdá"}

	mollyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "說不定鴨"}
)

var (
	mollyPhrase = nook.Languages{
		language.AmericanEnglish:      mollyAmericanEnglishName,
		language.CanadianFrench:       mollyCanadianFrenchName,
		language.Dutch:                mollyDutchName,
		language.French:               mollyFrenchName,
		language.German:               mollyGermanName,
		language.Italian:              mollyItalianName,
		language.Japanese:             mollyJapaneseName,
		language.Korean:               mollyKoreanName,
		language.LatinAmericanSpanish: mollyLatinAmericanSpanishName,
		language.Russian:              mollyRussianName,
		language.SimplifiedChinese:    mollySimplifiedChineseName,
		language.Spanish:              mollySpanishName,
		language.TraditionalChinese:   mollyTraditionalChineseName}
)

var (
	Molly = nook.Villager{
		Character:   mollyCharacter,
		Personality: personality.Normal,
		Phrase:      mollyPhrase}
)
