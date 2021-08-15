package kangaroo

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Marciamon bébé"}

	marcyGermanName = nook.Name{
		Language: language.German,
		Value:    "Mariekerlchen"}

	marcyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Marziabebè"}

	marcyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "モジモジつきょ"}

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
		Value:    "珠珠哟"}

	marcySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marsupimpollo"}

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
		Animal:   Kangaroo,
		Birthday: marcyBirthday,
		Code:     marcyCode,
		Gender:   nook.Female,
		Name:     marcyName}
)

var (
	marcyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	marcyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	marcyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

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
		Value:    "N/A"}

	marcyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	marcyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	marcySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哟"}

	marcySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pimpollo"}

	marcyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Personality: nook.Peppy,
		Phrase:      marcyPhrase}
)
