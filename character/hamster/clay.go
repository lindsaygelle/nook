package hamster

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
	clayBirthday = nook.Birthday{
		Day:   19,
		Month: time.October}
)

var (
	clayCode = nook.Code{
		Value: "ham05"}
)

var (
	clayAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Clay"}

	clayCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Guido"}

	clayDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Clay"}

	clayFrenchName = nook.Name{
		Language: language.French,
		Value:    "Guido"}

	clayGermanName = nook.Name{
		Language: language.German,
		Value:    "Dietmar"}

	clayItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Carmelo"}

	clayJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "どぐろう"}

	clayKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "햄둥"}

	clayLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Boliche"}

	clayRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Клэй"}

	claySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "子墨"}

	claySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Boliche"}

	clayTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "子墨"}
)

var (
	clayName = nook.Languages{
		language.AmericanEnglish:      clayAmericanEnglishName,
		language.CanadianFrench:       clayCanadianFrenchName,
		language.Dutch:                clayDutchName,
		language.French:               clayFrenchName,
		language.German:               clayGermanName,
		language.Italian:              clayItalianName,
		language.Japanese:             clayJapaneseName,
		language.Korean:               clayKoreanName,
		language.LatinAmericanSpanish: clayLatinAmericanSpanishName,
		language.Russian:              clayRussianName,
		language.SimplifiedChinese:    claySimplifiedChineseName,
		language.Spanish:              claySpanishName,
		language.TraditionalChinese:   clayTraditionalChineseName}
)

var (
	clayCharacter = nook.Character{
		Animal:   animal.Hamster,
		Birthday: clayBirthday,
		Code:     clayCode,
		Key:      character.Clay,
		Gender:   gender.Male,
		Name:     clayName}
)

var (
	clayAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "thump"}

	clayCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pindépice"}

	clayDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "boink"}

	clayFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pindépice"}

	clayGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "backi"}

	clayItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "popomf"}

	clayJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "どきどき"}

	clayKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "둥게둥게"}

	clayLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tucu-tucu"}

	clayRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "стук-стук"}

	claySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "忐忑不安"}

	claySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tucu-tucu"}

	clayTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "忐忑不安"}
)

var (
	clayPhrase = nook.Languages{
		language.AmericanEnglish:      clayAmericanEnglishName,
		language.CanadianFrench:       clayCanadianFrenchName,
		language.Dutch:                clayDutchName,
		language.French:               clayFrenchName,
		language.German:               clayGermanName,
		language.Italian:              clayItalianName,
		language.Japanese:             clayJapaneseName,
		language.Korean:               clayKoreanName,
		language.LatinAmericanSpanish: clayLatinAmericanSpanishName,
		language.Russian:              clayRussianName,
		language.SimplifiedChinese:    claySimplifiedChineseName,
		language.Spanish:              claySpanishName,
		language.TraditionalChinese:   clayTraditionalChineseName}
)

var (
	Clay = nook.Villager{
		Character:   clayCharacter,
		Personality: personality.Lazy,
		Phrase:      clayPhrase}
)
