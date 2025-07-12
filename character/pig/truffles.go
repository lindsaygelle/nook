package pig

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
	// trufflesBirthday represents truffles birthday.
	trufflesBirthday = nook.Birthday{
		Day:   28,
		Month: time.July}
)

var (
	// trufflesCode represents truffles code.
	trufflesCode = nook.Code{
		Value: "pig01"}
)

var (
	// trufflesAmericanEnglishName represents truffles american english name.
	trufflesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Truffles"}

	// trufflesCanadianFrenchName represents truffles canadian french name.
	trufflesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Trufa"}

	// trufflesDutchName represents truffles dutch name.
	trufflesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Truffles"}

	// trufflesFrenchName represents truffles french name.
	trufflesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Trufa"}

	// trufflesGermanName represents truffles german name.
	trufflesGermanName = nook.Name{
		Language: language.German,
		Value:    "Luzie"}

	// trufflesItalianName represents truffles italian name.
	trufflesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Grufetta"}

	// trufflesJapaneseName represents truffles japanese name.
	trufflesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トンコ"}

	// trufflesKoreanName represents truffles korean name.
	trufflesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "탱고"}

	// trufflesLatinAmericanSpanishName represents truffles latin american spanish name.
	trufflesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Trufas"}

	// trufflesRussianName represents truffles russian name.
	trufflesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Трафлс"}

	// trufflesSimplifiedChineseName represents truffles simplified chinese name.
	trufflesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘟嘟"}

	// trufflesSpanishName represents truffles spanish name.
	trufflesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Trufas"}

	// trufflesTraditionalChineseName represents truffles traditional chinese name.
	trufflesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘟嘟"}
)

var (
	// trufflesName represents truffles name.
	trufflesName = nook.Languages{
		language.AmericanEnglish:      trufflesAmericanEnglishName,
		language.CanadianFrench:       trufflesCanadianFrenchName,
		language.Dutch:                trufflesDutchName,
		language.French:               trufflesFrenchName,
		language.German:               trufflesGermanName,
		language.Italian:              trufflesItalianName,
		language.Japanese:             trufflesJapaneseName,
		language.Korean:               trufflesKoreanName,
		language.LatinAmericanSpanish: trufflesLatinAmericanSpanishName,
		language.Russian:              trufflesRussianName,
		language.SimplifiedChinese:    trufflesSimplifiedChineseName,
		language.Spanish:              trufflesSpanishName,
		language.TraditionalChinese:   trufflesTraditionalChineseName}
)

var (
	// trufflesCharacter represents truffles character.
	trufflesCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: trufflesBirthday,
		Code:     trufflesCode,
		Key:      character.Truffles,
		Gender:   gender.Female,
		Name:     trufflesName,
		Special:  false}
)

var (
	// trufflesAmericanEnglishPhrase represents truffles american english phrase.
	trufflesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snoutie"}

	// trufflesCanadianFrenchPhrase represents truffles canadian french phrase.
	trufflesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "jambon"}

	// trufflesDutchPhrase represents truffles dutch phrase.
	trufflesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snuiter"}

	// trufflesFrenchPhrase represents truffles french phrase.
	trufflesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "jambon"}

	// trufflesGermanPhrase represents truffles german phrase.
	trufflesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnauzie"}

	// trufflesItalianPhrase represents truffles italian phrase.
	trufflesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ficcanaso"}

	// trufflesJapanesePhrase represents truffles japanese phrase.
	trufflesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だわさ"}

	// trufflesKoreanPhrase represents truffles korean phrase.
	trufflesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "거라예"}

	// trufflesLatinAmericanSpanishPhrase represents truffles latin american spanish phrase.
	trufflesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "oink-oink"}

	// trufflesRussianPhrase represents truffles russian phrase.
	trufflesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрюша"}

	// trufflesSimplifiedChinesePhrase represents truffles simplified chinese phrase.
	trufflesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哇赛"}

	// trufflesSpanishPhrase represents truffles spanish phrase.
	trufflesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "oink-oink"}

	// trufflesTraditionalChinesePhrase represents truffles traditional chinese phrase.
	trufflesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哇賽"}
)

var (
	// trufflesPhrase represents truffles phrase.
	trufflesPhrase = nook.Languages{
		language.AmericanEnglish:      trufflesAmericanEnglishPhrase,
		language.CanadianFrench:       trufflesCanadianFrenchPhrase,
		language.Dutch:                trufflesDutchPhrase,
		language.French:               trufflesFrenchPhrase,
		language.German:               trufflesGermanPhrase,
		language.Italian:              trufflesItalianPhrase,
		language.Japanese:             trufflesJapanesePhrase,
		language.Korean:               trufflesKoreanPhrase,
		language.LatinAmericanSpanish: trufflesLatinAmericanSpanishPhrase,
		language.Russian:              trufflesRussianPhrase,
		language.SimplifiedChinese:    trufflesSimplifiedChinesePhrase,
		language.Spanish:              trufflesSpanishPhrase,
		language.TraditionalChinese:   trufflesTraditionalChinesePhrase}
)

var (
	// Truffles represents truffles.
	Truffles = nook.Villager{
		Character:   trufflesCharacter,
		Personality: personality.Peppy,
		Phrase:      trufflesPhrase}
)
