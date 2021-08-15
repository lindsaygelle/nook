package horse

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Danielhowdy"}

	buckDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Buckdraver"}

	buckFrenchName = nook.Name{
		Language: language.French,
		Value:    "Danieltaty"}

	buckGermanName = nook.Name{
		Language: language.German,
		Value:    "Rudiwiehie"}

	buckItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Trottolocowboy"}

	buckJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヴァヤシコフなんちて"}

	buckKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "바야시코프농담"}

	buckLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Trotónjopop"}

	buckRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бакпартнер"}

	buckSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "威亚搞笑"}

	buckSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Trotónzanahoria"}

	buckTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "威亞搞笑"}
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
		Animal:   Horse,
		Birthday: buckBirthday,
		Code:     buckCode,
		Gender:   nook.Male,
		Name:     buckName}
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
	Buck = nook.Villager{
		Character:   buckCharacter,
		Personality: nook.Jock,
		Phrase:      buckPhrase}
)
