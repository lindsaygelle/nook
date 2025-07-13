package elephant

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
	// tuckerBirthday represents tucker birthday.
	tuckerBirthday = nook.Birthday{
		Day:   7,
		Month: time.September}
)

var (
	// tuckerCode represents tucker code.
	tuckerCode = nook.Code{
		Value: "elp09"}
)

var (
	// tuckerAmericanEnglishName represents tucker american english name.
	tuckerAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tucker"}

	// tuckerCanadianFrenchName represents tucker canadian french name.
	tuckerCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Barry"}

	// tuckerDutchName represents tucker dutch name.
	tuckerDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tucker"}

	// tuckerFrenchName represents tucker french name.
	tuckerFrenchName = nook.Name{
		Language: language.French,
		Value:    "Barry"}

	// tuckerGermanName represents tucker german name.
	tuckerGermanName = nook.Name{
		Language: language.German,
		Value:    "Thorsten"}

	// tuckerItalianName represents tucker italian name.
	tuckerItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sventolo"}

	// tuckerJapaneseName represents tucker japanese name.
	tuckerJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "はじめ"}

	// tuckerKoreanName represents tucker korean name.
	tuckerKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "맘모"}

	// tuckerLatinAmericanSpanishName represents tucker latin american spanish name.
	tuckerLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pirolo"}

	// tuckerRussianName represents tucker russian name.
	tuckerRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Такер"}

	// tuckerSimplifiedChineseName represents tucker simplified chinese name.
	tuckerSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿原"}

	// tuckerSpanishName represents tucker spanish name.
	tuckerSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pirolo"}

	// tuckerTraditionalChineseName represents tucker traditional chinese name.
	tuckerTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿原"}
)

var (
	// tuckerName represents tucker name.
	tuckerName = nook.Languages{
		language.AmericanEnglish:      tuckerAmericanEnglishName,
		language.CanadianFrench:       tuckerCanadianFrenchName,
		language.Dutch:                tuckerDutchName,
		language.French:               tuckerFrenchName,
		language.German:               tuckerGermanName,
		language.Italian:              tuckerItalianName,
		language.Japanese:             tuckerJapaneseName,
		language.Korean:               tuckerKoreanName,
		language.LatinAmericanSpanish: tuckerLatinAmericanSpanishName,
		language.Russian:              tuckerRussianName,
		language.SimplifiedChinese:    tuckerSimplifiedChineseName,
		language.Spanish:              tuckerSpanishName,
		language.TraditionalChinese:   tuckerTraditionalChineseName}
)

var (
	// tuckerCharacter represents tucker character.
	tuckerCharacter = nook.Character{
		Animal:   animal.Elephant,
		Birthday: tuckerBirthday,
		Code:     tuckerCode,
		Key:      character.Tucker,
		Gender:   gender.Male,
		Name:     tuckerName,
		Special:  false}
)

var (
	// tuckerAmericanEnglishPhrase represents tucker american english phrase.
	tuckerAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "fuzzers"}

	// tuckerCanadianFrenchPhrase represents tucker canadian french phrase.
	tuckerCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mammouth"}

	// tuckerDutchPhrase represents tucker dutch phrase.
	tuckerDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pluisie"}

	// tuckerFrenchPhrase represents tucker french phrase.
	tuckerFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mammouth"}

	// tuckerGermanPhrase represents tucker german phrase.
	tuckerGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "puuuuuh"}

	// tuckerItalianPhrase represents tucker italian phrase.
	tuckerItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "barrir"}

	// tuckerJapanesePhrase represents tucker japanese phrase.
	tuckerJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "もじゃ"}

	// tuckerKoreanPhrase represents tucker korean phrase.
	tuckerKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "뿌뿌"}

	// tuckerLatinAmericanSpanishPhrase represents tucker latin american spanish phrase.
	tuckerLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "purupú"}

	// tuckerRussianPhrase represents tucker russian phrase.
	tuckerRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бивень"}

	// tuckerSimplifiedChinesePhrase represents tucker simplified chinese phrase.
	tuckerSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "毛毛"}

	// tuckerSpanishPhrase represents tucker spanish phrase.
	tuckerSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "purupú"}

	// tuckerTraditionalChinesePhrase represents tucker traditional chinese phrase.
	tuckerTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "毛毛"}
)

var (
	// tuckerPhrase represents tucker phrase.
	tuckerPhrase = nook.Languages{
		language.AmericanEnglish:      tuckerAmericanEnglishPhrase,
		language.CanadianFrench:       tuckerCanadianFrenchPhrase,
		language.Dutch:                tuckerDutchPhrase,
		language.French:               tuckerFrenchPhrase,
		language.German:               tuckerGermanPhrase,
		language.Italian:              tuckerItalianPhrase,
		language.Japanese:             tuckerJapanesePhrase,
		language.Korean:               tuckerKoreanPhrase,
		language.LatinAmericanSpanish: tuckerLatinAmericanSpanishPhrase,
		language.Russian:              tuckerRussianPhrase,
		language.SimplifiedChinese:    tuckerSimplifiedChinesePhrase,
		language.Spanish:              tuckerSpanishPhrase,
		language.TraditionalChinese:   tuckerTraditionalChinesePhrase}
)

var (
	// Tucker represents tucker.
	Tucker = nook.Villager{
		Character:   tuckerCharacter,
		Personality: personality.Lazy,
		Phrase:      tuckerPhrase}
)
