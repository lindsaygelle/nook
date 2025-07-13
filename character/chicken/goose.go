package chicken

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
	// gooseBirthday represents goose birthday.
	gooseBirthday = nook.Birthday{
		Day:   4,
		Month: time.October}
)

var (
	// gooseCode represents goose code.
	gooseCode = nook.Code{
		Value: "chn00"}
)

var (
	// gooseAmericanEnglishName represents goose american english name.
	gooseAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Goose"}

	// gooseCanadianFrenchName represents goose canadian french name.
	gooseCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pouli"}

	// gooseDutchName represents goose dutch name.
	gooseDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Goose"}

	// gooseFrenchName represents goose french name.
	gooseFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pouli"}

	// gooseGermanName represents goose german name.
	gooseGermanName = nook.Name{
		Language: language.German,
		Value:    "Konrad"}

	// gooseItalianName represents goose italian name.
	gooseItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Chiricò"}

	// gooseJapaneseName represents goose japanese name.
	gooseJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ケンタ"}

	// gooseKoreanName represents goose korean name.
	gooseKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "건태"}

	// gooseLatinAmericanSpanishName represents goose latin american spanish name.
	gooseLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gus"}

	// gooseRussianName represents goose russian name.
	gooseRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гус"}

	// gooseSimplifiedChineseName represents goose simplified chinese name.
	gooseSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "肯肯"}

	// gooseSpanishName represents goose spanish name.
	gooseSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gus"}

	// gooseTraditionalChineseName represents goose traditional chinese name.
	gooseTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "肯肯"}
)

var (
	// gooseName represents goose name.
	gooseName = nook.Languages{
		language.AmericanEnglish:      gooseAmericanEnglishName,
		language.CanadianFrench:       gooseCanadianFrenchName,
		language.Dutch:                gooseDutchName,
		language.French:               gooseFrenchName,
		language.German:               gooseGermanName,
		language.Italian:              gooseItalianName,
		language.Japanese:             gooseJapaneseName,
		language.Korean:               gooseKoreanName,
		language.LatinAmericanSpanish: gooseLatinAmericanSpanishName,
		language.Russian:              gooseRussianName,
		language.SimplifiedChinese:    gooseSimplifiedChineseName,
		language.Spanish:              gooseSpanishName,
		language.TraditionalChinese:   gooseTraditionalChineseName}
)

var (
	// gooseCharacter represents goose character.
	gooseCharacter = nook.Character{
		Animal:   animal.Chicken,
		Birthday: gooseBirthday,
		Code:     gooseCode,
		Key:      character.Goose,
		Gender:   gender.Male,
		Name:     gooseName,
		Special:  false}
)

var (
	// gooseAmericanEnglishPhrase represents goose american english phrase.
	gooseAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "buh-kay"}

	// gooseCanadianFrenchPhrase represents goose canadian french phrase.
	gooseCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bouyaka"}

	// gooseDutchPhrase represents goose dutch phrase.
	gooseDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "toktok"}

	// gooseFrenchPhrase represents goose french phrase.
	gooseFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bouyaka"}

	// gooseGermanPhrase represents goose german phrase.
	gooseGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "buaa-ka-ka"}

	// gooseItalianPhrase represents goose italian phrase.
	gooseItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cucù"}

	// gooseJapanesePhrase represents goose japanese phrase.
	gooseJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だコケ"}

	// gooseKoreanPhrase represents goose korean phrase.
	gooseKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "키득"}

	// gooseLatinAmericanSpanishPhrase represents goose latin american spanish phrase.
	gooseLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bukááá"}

	// gooseRussianPhrase represents goose russian phrase.
	gooseRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ко-кей"}

	// gooseSimplifiedChinesePhrase represents goose simplified chinese phrase.
	gooseSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咕咕"}

	// gooseSpanishPhrase represents goose spanish phrase.
	gooseSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "caracol"}

	// gooseTraditionalChinesePhrase represents goose traditional chinese phrase.
	gooseTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咕咕"}
)

var (
	// goosePhrase represents goose phrase.
	goosePhrase = nook.Languages{
		language.AmericanEnglish:      gooseAmericanEnglishPhrase,
		language.CanadianFrench:       gooseCanadianFrenchPhrase,
		language.Dutch:                gooseDutchPhrase,
		language.French:               gooseFrenchPhrase,
		language.German:               gooseGermanPhrase,
		language.Italian:              gooseItalianPhrase,
		language.Japanese:             gooseJapanesePhrase,
		language.Korean:               gooseKoreanPhrase,
		language.LatinAmericanSpanish: gooseLatinAmericanSpanishPhrase,
		language.Russian:              gooseRussianPhrase,
		language.SimplifiedChinese:    gooseSimplifiedChinesePhrase,
		language.Spanish:              gooseSpanishPhrase,
		language.TraditionalChinese:   gooseTraditionalChinesePhrase}
)

var (
	// Goose represents goose.
	Goose = nook.Villager{
		Character:   gooseCharacter,
		Personality: personality.Jock,
		Phrase:      goosePhrase}
)
