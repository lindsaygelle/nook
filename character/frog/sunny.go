package frog

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
	// sunnyBirthday represents sunny birthday.
	sunnyBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// sunnyCode represents sunny code.
	sunnyCode = nook.Code{
		Value: ""}
)

var (
	// sunnyAmericanEnglishName represents sunny american english name.
	sunnyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sunny"}

	// sunnyCanadianFrenchName represents sunny canadian french name.
	sunnyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// sunnyDutchName represents sunny dutch name.
	sunnyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// sunnyFrenchName represents sunny french name.
	sunnyFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// sunnyGermanName represents sunny german name.
	sunnyGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// sunnyItalianName represents sunny italian name.
	sunnyItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// sunnyJapaneseName represents sunny japanese name.
	sunnyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サニー"}

	// sunnyKoreanName represents sunny korean name.
	sunnyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// sunnyLatinAmericanSpanishName represents sunny latin american spanish name.
	sunnyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// sunnyRussianName represents sunny russian name.
	sunnyRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// sunnySimplifiedChineseName represents sunny simplified chinese name.
	sunnySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// sunnySpanishName represents sunny spanish name.
	sunnySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// sunnyTraditionalChineseName represents sunny traditional chinese name.
	sunnyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// sunnyName represents sunny name.
	sunnyName = nook.Languages{
		language.AmericanEnglish:      sunnyAmericanEnglishName,
		language.CanadianFrench:       sunnyCanadianFrenchName,
		language.Dutch:                sunnyDutchName,
		language.French:               sunnyFrenchName,
		language.German:               sunnyGermanName,
		language.Italian:              sunnyItalianName,
		language.Japanese:             sunnyJapaneseName,
		language.Korean:               sunnyKoreanName,
		language.LatinAmericanSpanish: sunnyLatinAmericanSpanishName,
		language.Russian:              sunnyRussianName,
		language.SimplifiedChinese:    sunnySimplifiedChineseName,
		language.Spanish:              sunnySpanishName,
		language.TraditionalChinese:   sunnyTraditionalChineseName}
)

var (
	// sunnyCharacter represents sunny character.
	sunnyCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: sunnyBirthday,
		Code:     sunnyCode,
		Key:      character.Sunny,
		Gender:   gender.Female,
		Name:     sunnyName,
		Special:  false}
)

var (
	// sunnyAmericanEnglishPhrase represents sunny american english phrase.
	sunnyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "でちょ"}

	// sunnyCanadianFrenchPhrase represents sunny canadian french phrase.
	sunnyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// sunnyDutchPhrase represents sunny dutch phrase.
	sunnyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// sunnyFrenchPhrase represents sunny french phrase.
	sunnyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// sunnyGermanPhrase represents sunny german phrase.
	sunnyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// sunnyItalianPhrase represents sunny italian phrase.
	sunnyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// sunnyJapanesePhrase represents sunny japanese phrase.
	sunnyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でちょ"}

	// sunnyKoreanPhrase represents sunny korean phrase.
	sunnyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// sunnyLatinAmericanSpanishPhrase represents sunny latin american spanish phrase.
	sunnyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// sunnyRussianPhrase represents sunny russian phrase.
	sunnyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// sunnySimplifiedChinesePhrase represents sunny simplified chinese phrase.
	sunnySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// sunnySpanishPhrase represents sunny spanish phrase.
	sunnySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// sunnyTraditionalChinesePhrase represents sunny traditional chinese phrase.
	sunnyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// sunnyPhrase represents sunny phrase.
	sunnyPhrase = nook.Languages{
		language.AmericanEnglish:      sunnyAmericanEnglishPhrase,
		language.CanadianFrench:       sunnyCanadianFrenchPhrase,
		language.Dutch:                sunnyDutchPhrase,
		language.French:               sunnyFrenchPhrase,
		language.German:               sunnyGermanPhrase,
		language.Italian:              sunnyItalianPhrase,
		language.Japanese:             sunnyJapanesePhrase,
		language.Korean:               sunnyKoreanPhrase,
		language.LatinAmericanSpanish: sunnyLatinAmericanSpanishPhrase,
		language.Russian:              sunnyRussianPhrase,
		language.SimplifiedChinese:    sunnySimplifiedChinesePhrase,
		language.Spanish:              sunnySpanishPhrase,
		language.TraditionalChinese:   sunnyTraditionalChinesePhrase}
)

var (
	// Sunny represents sunny.
	Sunny = nook.Villager{
		Character:   sunnyCharacter,
		Personality: personality.Normal,
		Phrase:      sunnyPhrase}
)
