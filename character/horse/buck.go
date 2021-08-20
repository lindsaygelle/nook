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
	buckBirthday = nook.Birthday{
		Day:   4,
		Month: time.April}
)

var (
	buckCode = nook.Code{
		Value: "hrs00"}
)

var (
	buckAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Buck"}

	buckCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Daniel"}

	buckDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Buck"}

	buckFrenchName = nook.Name{
		Language: language.French,
		Value:    "Daniel"}

	buckGermanName = nook.Name{
		Language: language.German,
		Value:    "Rudi"}

	buckItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Trottolo"}

	buckJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヴァヤシコフ"}

	buckKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "바야시코프"}

	buckLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Trotón"}

	buckRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бак"}

	buckSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "威亚"}

	buckSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Trotón"}

	buckTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "威亞"}
)

var (
	buckName = nook.Languages{
		language.AmericanEnglish:      buckAmericanEnglishName,
		language.CanadianFrench:       buckCanadianFrenchName,
		language.Dutch:                buckDutchName,
		language.French:               buckFrenchName,
		language.German:               buckGermanName,
		language.Italian:              buckItalianName,
		language.Japanese:             buckJapaneseName,
		language.Korean:               buckKoreanName,
		language.LatinAmericanSpanish: buckLatinAmericanSpanishName,
		language.Russian:              buckRussianName,
		language.SimplifiedChinese:    buckSimplifiedChineseName,
		language.Spanish:              buckSpanishName,
		language.TraditionalChinese:   buckTraditionalChineseName}
)

var (
	buckCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: buckBirthday,
		Code:     buckCode,
		Key:      character.Buck,
		Gender:   gender.Male,
		Name:     buckName,
		Special:  false}
)

var (
	buckAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pardner"}

	buckCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "howdy"}

	buckDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "draver"}

	buckFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "taty"}

	buckGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "wiehie"}

	buckItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cowboy"}

	buckJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なんちて"}

	buckKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "농담"}

	buckLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "jopop"}

	buckRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "партнер"}

	buckSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "搞笑"}

	buckSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "zanahoria"}

	buckTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "搞笑"}
)

var (
	buckPhrase = nook.Languages{
		language.AmericanEnglish:      buckAmericanEnglishPhrase,
		language.CanadianFrench:       buckCanadianFrenchPhrase,
		language.Dutch:                buckDutchPhrase,
		language.French:               buckFrenchPhrase,
		language.German:               buckGermanPhrase,
		language.Italian:              buckItalianPhrase,
		language.Japanese:             buckJapanesePhrase,
		language.Korean:               buckKoreanPhrase,
		language.LatinAmericanSpanish: buckLatinAmericanSpanishPhrase,
		language.Russian:              buckRussianPhrase,
		language.SimplifiedChinese:    buckSimplifiedChinesePhrase,
		language.Spanish:              buckSpanishPhrase,
		language.TraditionalChinese:   buckTraditionalChinesePhrase}
)

var (
	Buck = nook.Villager{
		Character:   buckCharacter,
		Personality: personality.Jock,
		Phrase:      buckPhrase}
)
