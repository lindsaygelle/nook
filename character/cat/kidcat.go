package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	kidcatBirthday = nook.Birthday{
		Day:   1,
		Month: time.August}
)

var (
	kidcatCode = nook.Code{
		Value: "cat10"}
)

var (
	kidcatAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kid Cat"}

	kidcatCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Câlin"}

	kidcatDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kid Cat"}

	kidcatFrenchName = nook.Name{
		Language: language.French,
		Value:    "Câlin"}

	kidcatGermanName = nook.Name{
		Language: language.German,
		Value:    "Pit"}

	kidcatItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Zampa"}

	kidcatJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "１ごう"}

	kidcatKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "1호"}

	kidcatLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gatomán"}

	kidcatRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кид Кэт"}

	kidcatSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿一"}

	kidcatSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gatomán"}

	kidcatTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿一"}
)

var (
	kidcatName = nook.Languages{
		language.AmericanEnglish:      kidcatAmericanEnglishName,
		language.CanadianFrench:       kidcatCanadianFrenchName,
		language.Dutch:                kidcatDutchName,
		language.French:               kidcatFrenchName,
		language.German:               kidcatGermanName,
		language.Italian:              kidcatItalianName,
		language.Japanese:             kidcatJapaneseName,
		language.Korean:               kidcatKoreanName,
		language.LatinAmericanSpanish: kidcatLatinAmericanSpanishName,
		language.Russian:              kidcatRussianName,
		language.SimplifiedChinese:    kidcatSimplifiedChineseName,
		language.Spanish:              kidcatSpanishName,
		language.TraditionalChinese:   kidcatTraditionalChineseName}
)

var (
	kidcatCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: kidcatBirthday,
		Code:     kidcatCode,
		Gender:   gender.Male,
		Name:     kidcatName}
)

var (
	kidcatAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "psst"}

	kidcatCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "psst"}

	kidcatDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "psst"}

	kidcatFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "psst"}

	kidcatGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "maunz"}

	kidcatItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "psst"}

	kidcatJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "とぉっ"}

	kidcatKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "얍"}

	kidcatLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "psst"}

	kidcatRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "шш"}

	kidcatSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喝"}

	kidcatSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "soldado"}

	kidcatTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喝"}
)

var (
	kidcatPhrase = nook.Languages{
		language.AmericanEnglish:      kidcatAmericanEnglishName,
		language.CanadianFrench:       kidcatCanadianFrenchName,
		language.Dutch:                kidcatDutchName,
		language.French:               kidcatFrenchName,
		language.German:               kidcatGermanName,
		language.Italian:              kidcatItalianName,
		language.Japanese:             kidcatJapaneseName,
		language.Korean:               kidcatKoreanName,
		language.LatinAmericanSpanish: kidcatLatinAmericanSpanishName,
		language.Russian:              kidcatRussianName,
		language.SimplifiedChinese:    kidcatSimplifiedChineseName,
		language.Spanish:              kidcatSpanishName,
		language.TraditionalChinese:   kidcatTraditionalChineseName}
)

var (
	KidCat = nook.Villager{
		Character:   kidcatCharacter,
		Personality: personality.Jock,
		Phrase:      kidcatPhrase}
)
