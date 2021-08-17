package horse

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
	papiBirthday = nook.Birthday{
		Day:   10,
		Month: time.January}
)

var (
	papiCode = nook.Code{
		Value: "hrs12"}
)

var (
	papiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Papi"}

	papiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bourrico"}

	papiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Papi"}

	papiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bourrico"}

	papiGermanName = nook.Name{
		Language: language.German,
		Value:    "Friedel"}

	papiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pierino"}

	papiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "オカッピ"}

	papiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마사마"}

	papiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bayo"}

	papiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Папи"}

	papiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小冈"}

	papiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bayo"}

	papiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小岡"}
)

var (
	papiName = nook.Languages{
		language.AmericanEnglish:      papiAmericanEnglishName,
		language.CanadianFrench:       papiCanadianFrenchName,
		language.Dutch:                papiDutchName,
		language.French:               papiFrenchName,
		language.German:               papiGermanName,
		language.Italian:              papiItalianName,
		language.Japanese:             papiJapaneseName,
		language.Korean:               papiKoreanName,
		language.LatinAmericanSpanish: papiLatinAmericanSpanishName,
		language.Russian:              papiRussianName,
		language.SimplifiedChinese:    papiSimplifiedChineseName,
		language.Spanish:              papiSpanishName,
		language.TraditionalChinese:   papiTraditionalChineseName}
)

var (
	papiCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: papiBirthday,
		Code:     papiCode,
		Key:      character.Papi,
		Gender:   gender.Male,
		Name:     papiName}
)

var (
	papiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "haaay"}

	papiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "en selle"}

	papiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hooooi"}

	papiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "en selle"}

	papiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "nichja"}

	papiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "clip clop"}

	papiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だっけ"}

	papiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그런가"}

	papiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cli-cló"}

	papiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "се-е-ено"}

	papiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "听说是"}

	papiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "quepassa"}

	papiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "聽說是"}
)

var (
	papiPhrase = nook.Languages{
		language.AmericanEnglish:      papiAmericanEnglishPhrase,
		language.CanadianFrench:       papiCanadianFrenchPhrase,
		language.Dutch:                papiDutchPhrase,
		language.French:               papiFrenchPhrase,
		language.German:               papiGermanPhrase,
		language.Italian:              papiItalianPhrase,
		language.Japanese:             papiJapanesePhrase,
		language.Korean:               papiKoreanPhrase,
		language.LatinAmericanSpanish: papiLatinAmericanSpanishPhrase,
		language.Russian:              papiRussianPhrase,
		language.SimplifiedChinese:    papiSimplifiedChinesePhrase,
		language.Spanish:              papiSpanishPhrase,
		language.TraditionalChinese:   papiTraditionalChinesePhrase}
)

var (
	Papi = nook.Villager{
		Character:   papiCharacter,
		Personality: personality.Lazy,
		Phrase:      papiPhrase}
)
