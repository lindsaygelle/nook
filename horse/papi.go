package horse

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Bourricoen selle"}

	papiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Papihooooi"}

	papiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bourricoen selle"}

	papiGermanName = nook.Name{
		Language: language.German,
		Value:    "Friedelnichja"}

	papiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pierinoclip clop"}

	papiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "オカッピだっけ"}

	papiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마사마그런가"}

	papiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bayocli-cló"}

	papiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Паписе-е-ено"}

	papiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小冈听说是"}

	papiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bayoquepassa"}

	papiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小岡聽說是"}
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
		Animal:   Horse,
		Birthday: papiBirthday,
		Code:     papiCode,
		Gender:   nook.Male,
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
	Papi = nook.Villager{
		Character:   papiCharacter,
		Personality: nook.Lazy,
		Phrase:      papiPhrase}
)
