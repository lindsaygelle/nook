package chicken

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	gooseBirthday = nook.Birthday{
		Day:   4,
		Month: time.October}
)

var (
	gooseCode = nook.Code{
		Value: "chn00"}
)

var (
	gooseAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Goose"}

	gooseCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Poulibouyaka"}

	gooseDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Goosetoktok"}

	gooseFrenchName = nook.Name{
		Language: language.French,
		Value:    "Poulibouyaka"}

	gooseGermanName = nook.Name{
		Language: language.German,
		Value:    "Konradbuaa-ka-ka"}

	gooseItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Chiricòcucù"}

	gooseJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ケンタだコケ"}

	gooseKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "건태키득"}

	gooseLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gusbukááá"}

	gooseRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гуско-кей"}

	gooseSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "肯肯咕咕"}

	gooseSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Guscaracol"}

	gooseTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "肯肯咕咕"}
)

var (
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
	gooseCharacter = nook.Character{
		Animal:   Chicken,
		Birthday: gooseBirthday,
		Code:     gooseCode,
		Gender:   nook.Male,
		Name:     gooseName}
)

var (
	gooseAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "buh-kay"}

	gooseCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bouyaka"}

	gooseDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "toktok"}

	gooseFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bouyaka"}

	gooseGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "buaa-ka-ka"}

	gooseItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cucù"}

	gooseJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だコケ"}

	gooseKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "키득"}

	gooseLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bukááá"}

	gooseRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ко-кей"}

	gooseSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咕咕"}

	gooseSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "caracol"}

	gooseTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咕咕"}
)

var (
	goosePhrase = nook.Languages{
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
	Goose = nook.Villager{
		Character:   gooseCharacter,
		Personality: nook.Jock,
		Phrase:      goosePhrase}
)
