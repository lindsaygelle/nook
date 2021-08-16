package kangaroo

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	marcyBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	marcyCode = nook.Code{
		Value: ""}
)

var (
	marcyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Marcy"}

	marcyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	marcyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	marcyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marcia"}

	marcyGermanName = nook.Name{
		Language: language.German,
		Value:    "Marie"}

	marcyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Marzia"}

	marcyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "モジモジ"}

	marcyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	marcyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	marcyRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	marcySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "珠珠"}

	marcySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marsu"}

	marcyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	marcyName = nook.Languages{
		language.AmericanEnglish:      marcyAmericanEnglishName,
		language.CanadianFrench:       marcyCanadianFrenchName,
		language.Dutch:                marcyDutchName,
		language.French:               marcyFrenchName,
		language.German:               marcyGermanName,
		language.Italian:              marcyItalianName,
		language.Japanese:             marcyJapaneseName,
		language.Korean:               marcyKoreanName,
		language.LatinAmericanSpanish: marcyLatinAmericanSpanishName,
		language.Russian:              marcyRussianName,
		language.SimplifiedChinese:    marcySimplifiedChineseName,
		language.Spanish:              marcySpanishName,
		language.TraditionalChinese:   marcyTraditionalChineseName}
)

var (
	marcyCharacter = nook.Character{
		Animal:   animal.Kangaroo,
		Birthday: marcyBirthday,
		Code:     marcyCode,
		Gender:   gender.Female,
		Name:     marcyName}
)

var (
	marcyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "young 'un"}

	marcyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	marcyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	marcyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon bébé"}

	marcyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kerlchen"}

	marcyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bebè"}

	marcyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "つきょ"}

	marcyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	marcyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	marcyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	marcySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哟"}

	marcySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pimpollo"}

	marcyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	marcyPhrase = nook.Languages{
		language.AmericanEnglish:      marcyAmericanEnglishName,
		language.CanadianFrench:       marcyCanadianFrenchName,
		language.Dutch:                marcyDutchName,
		language.French:               marcyFrenchName,
		language.German:               marcyGermanName,
		language.Italian:              marcyItalianName,
		language.Japanese:             marcyJapaneseName,
		language.Korean:               marcyKoreanName,
		language.LatinAmericanSpanish: marcyLatinAmericanSpanishName,
		language.Russian:              marcyRussianName,
		language.SimplifiedChinese:    marcySimplifiedChineseName,
		language.Spanish:              marcySpanishName,
		language.TraditionalChinese:   marcyTraditionalChineseName}
)

var (
	Marcy = nook.Villager{
		Character:   marcyCharacter,
		Personality: personality.Peppy,
		Phrase:      marcyPhrase}
)
