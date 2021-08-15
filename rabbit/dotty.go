package rabbit

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	dottyBirthday = nook.Birthday{
		Day:   14,
		Month: time.March}
)

var (
	dottyCode = nook.Code{
		Value: "rbt01"}
)

var (
	dottyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dotty"}

	dottyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Dorothéelapinou"}

	dottyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Dottylamprei"}

	dottyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Dorothéelapinou"}

	dottyGermanName = nook.Name{
		Language: language.German,
		Value:    "Dorouiui"}

	dottyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Dottycaroté"}

	dottyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マーサラン"}

	dottyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "서머룰루랄라"}

	dottyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Katiatoini"}

	dottyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Доттималыш"}

	dottySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "玛莎啷"}

	dottySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Katiatrompis"}

	dottyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "瑪莎啷"}
)

var (
	dottyName = nook.Languages{
		language.AmericanEnglish:      dottyAmericanEnglishName,
		language.CanadianFrench:       dottyCanadianFrenchName,
		language.Dutch:                dottyDutchName,
		language.French:               dottyFrenchName,
		language.German:               dottyGermanName,
		language.Italian:              dottyItalianName,
		language.Japanese:             dottyJapaneseName,
		language.Korean:               dottyKoreanName,
		language.LatinAmericanSpanish: dottyLatinAmericanSpanishName,
		language.Russian:              dottyRussianName,
		language.SimplifiedChinese:    dottySimplifiedChineseName,
		language.Spanish:              dottySpanishName,
		language.TraditionalChinese:   dottyTraditionalChineseName}
)

var (
	dottyCharacter = nook.Character{
		Animal:   Rabbit,
		Birthday: dottyBirthday,
		Code:     dottyCode,
		Gender:   nook.Female,
		Name:     dottyName}
)

var (
	dottyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "wee one"}

	dottyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "lapinou"}

	dottyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "lamprei"}

	dottyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "lapinou"}

	dottyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "uiui"}

	dottyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "caroté"}

	dottyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ラン"}

	dottyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "룰루랄라"}

	dottyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "toini"}

	dottyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "малыш"}

	dottySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啷"}

	dottySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "trompis"}

	dottyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啷"}
)

var (
	dottyPhrase = nook.Languages{
		language.AmericanEnglish:      dottyAmericanEnglishName,
		language.CanadianFrench:       dottyCanadianFrenchName,
		language.Dutch:                dottyDutchName,
		language.French:               dottyFrenchName,
		language.German:               dottyGermanName,
		language.Italian:              dottyItalianName,
		language.Japanese:             dottyJapaneseName,
		language.Korean:               dottyKoreanName,
		language.LatinAmericanSpanish: dottyLatinAmericanSpanishName,
		language.Russian:              dottyRussianName,
		language.SimplifiedChinese:    dottySimplifiedChineseName,
		language.Spanish:              dottySpanishName,
		language.TraditionalChinese:   dottyTraditionalChineseName}
)

var (
	Dotty = nook.Villager{
		Character:   dottyCharacter,
		Personality: nook.Peppy,
		Phrase:      dottyPhrase}
)
