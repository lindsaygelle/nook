package gorilla

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
	// yodelBirthday represents yodel birthday.
	yodelBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// yodelCode represents yodel code.
	yodelCode = nook.Code{
		Value: ""}
)

var (
	// yodelAmericanEnglishName represents yodel american english name.
	yodelAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Yodel"}

	// yodelCanadianFrenchName represents yodel canadian french name.
	yodelCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// yodelDutchName represents yodel dutch name.
	yodelDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// yodelFrenchName represents yodel french name.
	yodelFrenchName = nook.Name{
		Language: language.French,
		Value:    "Thomas"}

	// yodelGermanName represents yodel german name.
	yodelGermanName = nook.Name{
		Language: language.German,
		Value:    "Jürgen"}

	// yodelItalianName represents yodel italian name.
	yodelItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Yodel"}

	// yodelJapaneseName represents yodel japanese name.
	yodelJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヨーデル"}

	// yodelKoreanName represents yodel korean name.
	yodelKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// yodelLatinAmericanSpanishName represents yodel latin american spanish name.
	yodelLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// yodelRussianName represents yodel russian name.
	yodelRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// yodelSimplifiedChineseName represents yodel simplified chinese name.
	yodelSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// yodelSpanishName represents yodel spanish name.
	yodelSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Asdrúbal"}

	// yodelTraditionalChineseName represents yodel traditional chinese name.
	yodelTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// yodelName represents yodel name.
	yodelName = nook.Languages{
		language.AmericanEnglish:      yodelAmericanEnglishName,
		language.CanadianFrench:       yodelCanadianFrenchName,
		language.Dutch:                yodelDutchName,
		language.French:               yodelFrenchName,
		language.German:               yodelGermanName,
		language.Italian:              yodelItalianName,
		language.Japanese:             yodelJapaneseName,
		language.Korean:               yodelKoreanName,
		language.LatinAmericanSpanish: yodelLatinAmericanSpanishName,
		language.Russian:              yodelRussianName,
		language.SimplifiedChinese:    yodelSimplifiedChineseName,
		language.Spanish:              yodelSpanishName,
		language.TraditionalChinese:   yodelTraditionalChineseName}
)

var (
	// yodelCharacter represents yodel character.
	yodelCharacter = nook.Character{
		Animal:   animal.Gorilla,
		Birthday: yodelBirthday,
		Code:     yodelCode,
		Key:      character.Yodel,
		Gender:   gender.Male,
		Name:     yodelName,
		Special:  false}
)

var (
	// yodelAmericanEnglishPhrase represents yodel american english phrase.
	yodelAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "odelay"}

	// yodelCanadianFrenchPhrase represents yodel canadian french phrase.
	yodelCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// yodelDutchPhrase represents yodel dutch phrase.
	yodelDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// yodelFrenchPhrase represents yodel french phrase.
	yodelFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "yolahitoou"}

	// yodelGermanPhrase represents yodel german phrase.
	yodelGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "jodeli"}

	// yodelItalianPhrase represents yodel italian phrase.
	yodelItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ullallà"}

	// yodelJapanesePhrase represents yodel japanese phrase.
	yodelJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヨホホー"}

	// yodelKoreanPhrase represents yodel korean phrase.
	yodelKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// yodelLatinAmericanSpanishPhrase represents yodel latin american spanish phrase.
	yodelLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// yodelRussianPhrase represents yodel russian phrase.
	yodelRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// yodelSimplifiedChinesePhrase represents yodel simplified chinese phrase.
	yodelSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// yodelSpanishPhrase represents yodel spanish phrase.
	yodelSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "joya"}

	// yodelTraditionalChinesePhrase represents yodel traditional chinese phrase.
	yodelTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// yodelPhrase represents yodel phrase.
	yodelPhrase = nook.Languages{
		language.AmericanEnglish:      yodelAmericanEnglishPhrase,
		language.CanadianFrench:       yodelCanadianFrenchPhrase,
		language.Dutch:                yodelDutchPhrase,
		language.French:               yodelFrenchPhrase,
		language.German:               yodelGermanPhrase,
		language.Italian:              yodelItalianPhrase,
		language.Japanese:             yodelJapanesePhrase,
		language.Korean:               yodelKoreanPhrase,
		language.LatinAmericanSpanish: yodelLatinAmericanSpanishPhrase,
		language.Russian:              yodelRussianPhrase,
		language.SimplifiedChinese:    yodelSimplifiedChinesePhrase,
		language.Spanish:              yodelSpanishPhrase,
		language.TraditionalChinese:   yodelTraditionalChinesePhrase}
)

var (
	// Yodel represents yodel.
	Yodel = nook.Villager{
		Character:   yodelCharacter,
		Personality: personality.Lazy,
		Phrase:      yodelPhrase}
)
