package kangaroo

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
	// marcyBirthday represents marcy birthday.
	marcyBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// marcyCode represents marcy code.
	marcyCode = nook.Code{
		Value: ""}
)

var (
	// marcyAmericanEnglishName represents marcy american english name.
	marcyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Marcy"}

	// marcyCanadianFrenchName represents marcy canadian french name.
	marcyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// marcyDutchName represents marcy dutch name.
	marcyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// marcyFrenchName represents marcy french name.
	marcyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marcia"}

	// marcyGermanName represents marcy german name.
	marcyGermanName = nook.Name{
		Language: language.German,
		Value:    "Marie"}

	// marcyItalianName represents marcy italian name.
	marcyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Marzia"}

	// marcyJapaneseName represents marcy japanese name.
	marcyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "モジモジ"}

	// marcyKoreanName represents marcy korean name.
	marcyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// marcyLatinAmericanSpanishName represents marcy latin american spanish name.
	marcyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// marcyRussianName represents marcy russian name.
	marcyRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// marcySimplifiedChineseName represents marcy simplified chinese name.
	marcySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "珠珠"}

	// marcySpanishName represents marcy spanish name.
	marcySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marsu"}

	// marcyTraditionalChineseName represents marcy traditional chinese name.
	marcyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// marcyName represents marcy name.
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
	// marcyCharacter represents marcy character.
	marcyCharacter = nook.Character{
		Animal:   animal.Kangaroo,
		Birthday: marcyBirthday,
		Code:     marcyCode,
		Key:      character.Marcy,
		Gender:   gender.Female,
		Name:     marcyName,
		Special:  false}
)

var (
	// marcyAmericanEnglishPhrase represents marcy american english phrase.
	marcyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "young 'un"}

	// marcyCanadianFrenchPhrase represents marcy canadian french phrase.
	marcyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// marcyDutchPhrase represents marcy dutch phrase.
	marcyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// marcyFrenchPhrase represents marcy french phrase.
	marcyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon bébé"}

	// marcyGermanPhrase represents marcy german phrase.
	marcyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kerlchen"}

	// marcyItalianPhrase represents marcy italian phrase.
	marcyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bebè"}

	// marcyJapanesePhrase represents marcy japanese phrase.
	marcyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "つきょ"}

	// marcyKoreanPhrase represents marcy korean phrase.
	marcyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// marcyLatinAmericanSpanishPhrase represents marcy latin american spanish phrase.
	marcyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// marcyRussianPhrase represents marcy russian phrase.
	marcyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// marcySimplifiedChinesePhrase represents marcy simplified chinese phrase.
	marcySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哟"}

	// marcySpanishPhrase represents marcy spanish phrase.
	marcySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pimpollo"}

	// marcyTraditionalChinesePhrase represents marcy traditional chinese phrase.
	marcyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// marcyPhrase represents marcy phrase.
	marcyPhrase = nook.Languages{
		language.AmericanEnglish:      marcyAmericanEnglishPhrase,
		language.CanadianFrench:       marcyCanadianFrenchPhrase,
		language.Dutch:                marcyDutchPhrase,
		language.French:               marcyFrenchPhrase,
		language.German:               marcyGermanPhrase,
		language.Italian:              marcyItalianPhrase,
		language.Japanese:             marcyJapanesePhrase,
		language.Korean:               marcyKoreanPhrase,
		language.LatinAmericanSpanish: marcyLatinAmericanSpanishPhrase,
		language.Russian:              marcyRussianPhrase,
		language.SimplifiedChinese:    marcySimplifiedChinesePhrase,
		language.Spanish:              marcySpanishPhrase,
		language.TraditionalChinese:   marcyTraditionalChinesePhrase}
)

var (
	// Marcy represents marcy.
	Marcy = nook.Villager{
		Character:   marcyCharacter,
		Personality: personality.Peppy,
		Phrase:      marcyPhrase}
)
