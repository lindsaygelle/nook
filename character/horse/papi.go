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
	// papiBirthday represents papi birthday.
	papiBirthday = nook.Birthday{
		Day:   10,
		Month: time.January}
)

var (
	// papiCode represents papi code.
	papiCode = nook.Code{
		Value: "hrs12"}
)

var (
	// papiAmericanEnglishName represents papi american english name.
	papiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Papi"}

	// papiCanadianFrenchName represents papi canadian french name.
	papiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bourrico"}

	// papiDutchName represents papi dutch name.
	papiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Papi"}

	// papiFrenchName represents papi french name.
	papiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bourrico"}

	// papiGermanName represents papi german name.
	papiGermanName = nook.Name{
		Language: language.German,
		Value:    "Friedel"}

	// papiItalianName represents papi italian name.
	papiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pierino"}

	// papiJapaneseName represents papi japanese name.
	papiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "オカッピ"}

	// papiKoreanName represents papi korean name.
	papiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마사마"}

	// papiLatinAmericanSpanishName represents papi latin american spanish name.
	papiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bayo"}

	// papiRussianName represents papi russian name.
	papiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Папи"}

	// papiSimplifiedChineseName represents papi simplified chinese name.
	papiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小冈"}

	// papiSpanishName represents papi spanish name.
	papiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bayo"}

	// papiTraditionalChineseName represents papi traditional chinese name.
	papiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小岡"}
)

var (
	// papiName represents papi name.
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
	// papiCharacter represents papi character.
	papiCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: papiBirthday,
		Code:     papiCode,
		Key:      character.Papi,
		Gender:   gender.Male,
		Name:     papiName,
		Special:  false}
)

var (
	// papiAmericanEnglishPhrase represents papi american english phrase.
	papiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "haaay"}

	// papiCanadianFrenchPhrase represents papi canadian french phrase.
	papiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "en selle"}

	// papiDutchPhrase represents papi dutch phrase.
	papiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hooooi"}

	// papiFrenchPhrase represents papi french phrase.
	papiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "en selle"}

	// papiGermanPhrase represents papi german phrase.
	papiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "nichja"}

	// papiItalianPhrase represents papi italian phrase.
	papiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "clip clop"}

	// papiJapanesePhrase represents papi japanese phrase.
	papiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だっけ"}

	// papiKoreanPhrase represents papi korean phrase.
	papiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그런가"}

	// papiLatinAmericanSpanishPhrase represents papi latin american spanish phrase.
	papiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cli-cló"}

	// papiRussianPhrase represents papi russian phrase.
	papiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "се-е-ено"}

	// papiSimplifiedChinesePhrase represents papi simplified chinese phrase.
	papiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "听说是"}

	// papiSpanishPhrase represents papi spanish phrase.
	papiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "quepassa"}

	// papiTraditionalChinesePhrase represents papi traditional chinese phrase.
	papiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "聽說是"}
)

var (
	// papiPhrase represents papi phrase.
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
	// Papi represents papi.
	Papi = nook.Villager{
		Character:   papiCharacter,
		Personality: personality.Lazy,
		Phrase:      papiPhrase}
)
