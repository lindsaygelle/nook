package duck

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
	// fruityBirthday represents fruity birthday.
	fruityBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// fruityCode represents fruity code.
	fruityCode = nook.Code{
		Value: ""}
)

var (
	// fruityAmericanEnglishName represents fruity american english name.
	fruityAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Fruity"}

	// fruityCanadianFrenchName represents fruity canadian french name.
	fruityCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// fruityDutchName represents fruity dutch name.
	fruityDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// fruityFrenchName represents fruity french name.
	fruityFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// fruityGermanName represents fruity german name.
	fruityGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// fruityItalianName represents fruity italian name.
	fruityItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// fruityJapaneseName represents fruity japanese name.
	fruityJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フルーティー"}

	// fruityKoreanName represents fruity korean name.
	fruityKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// fruityLatinAmericanSpanishName represents fruity latin american spanish name.
	fruityLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// fruityRussianName represents fruity russian name.
	fruityRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// fruitySimplifiedChineseName represents fruity simplified chinese name.
	fruitySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// fruitySpanishName represents fruity spanish name.
	fruitySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// fruityTraditionalChineseName represents fruity traditional chinese name.
	fruityTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// fruityName represents fruity name.
	fruityName = nook.Languages{
		language.AmericanEnglish:      fruityAmericanEnglishName,
		language.CanadianFrench:       fruityCanadianFrenchName,
		language.Dutch:                fruityDutchName,
		language.French:               fruityFrenchName,
		language.German:               fruityGermanName,
		language.Italian:              fruityItalianName,
		language.Japanese:             fruityJapaneseName,
		language.Korean:               fruityKoreanName,
		language.LatinAmericanSpanish: fruityLatinAmericanSpanishName,
		language.Russian:              fruityRussianName,
		language.SimplifiedChinese:    fruitySimplifiedChineseName,
		language.Spanish:              fruitySpanishName,
		language.TraditionalChinese:   fruityTraditionalChineseName}
)

var (
	// fruityCharacter represents fruity character.
	fruityCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: fruityBirthday,
		Code:     fruityCode,
		Key:      character.Fruity,
		Gender:   gender.Male,
		Name:     fruityName,
		Special:  false}
)

var (
	// fruityAmericanEnglishPhrase represents fruity american english phrase.
	fruityAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "グァバ"}

	// fruityCanadianFrenchPhrase represents fruity canadian french phrase.
	fruityCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// fruityDutchPhrase represents fruity dutch phrase.
	fruityDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// fruityFrenchPhrase represents fruity french phrase.
	fruityFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// fruityGermanPhrase represents fruity german phrase.
	fruityGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// fruityItalianPhrase represents fruity italian phrase.
	fruityItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// fruityJapanesePhrase represents fruity japanese phrase.
	fruityJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "グァバ"}

	// fruityKoreanPhrase represents fruity korean phrase.
	fruityKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// fruityLatinAmericanSpanishPhrase represents fruity latin american spanish phrase.
	fruityLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// fruityRussianPhrase represents fruity russian phrase.
	fruityRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// fruitySimplifiedChinesePhrase represents fruity simplified chinese phrase.
	fruitySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// fruitySpanishPhrase represents fruity spanish phrase.
	fruitySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// fruityTraditionalChinesePhrase represents fruity traditional chinese phrase.
	fruityTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// fruityPhrase represents fruity phrase.
	fruityPhrase = nook.Languages{
		language.AmericanEnglish:      fruityAmericanEnglishPhrase,
		language.CanadianFrench:       fruityCanadianFrenchPhrase,
		language.Dutch:                fruityDutchPhrase,
		language.French:               fruityFrenchPhrase,
		language.German:               fruityGermanPhrase,
		language.Italian:              fruityItalianPhrase,
		language.Japanese:             fruityJapanesePhrase,
		language.Korean:               fruityKoreanPhrase,
		language.LatinAmericanSpanish: fruityLatinAmericanSpanishPhrase,
		language.Russian:              fruityRussianPhrase,
		language.SimplifiedChinese:    fruitySimplifiedChinesePhrase,
		language.Spanish:              fruitySpanishPhrase,
		language.TraditionalChinese:   fruityTraditionalChinesePhrase}
)

var (
	// Fruity represents fruity.
	Fruity = nook.Villager{
		Character:   fruityCharacter,
		Personality: personality.Jock,
		Phrase:      fruityPhrase}
)
