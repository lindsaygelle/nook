package mouse

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
	flossieBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	flossieCode = nook.Code{
		Value: ""}
)

var (
	flossieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Flossie"}

	flossieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	flossieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	flossieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chiquita"}

	flossieGermanName = nook.Name{
		Language: language.German,
		Value:    "Amalie"}

	flossieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Florinda"}

	flossieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピスタチオ"}

	flossieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	flossieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	flossieRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	flossieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	flossieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Oregilda"}

	flossieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	flossieName = nook.Languages{
		language.AmericanEnglish:      flossieAmericanEnglishName,
		language.CanadianFrench:       flossieCanadianFrenchName,
		language.Dutch:                flossieDutchName,
		language.French:               flossieFrenchName,
		language.German:               flossieGermanName,
		language.Italian:              flossieItalianName,
		language.Japanese:             flossieJapaneseName,
		language.Korean:               flossieKoreanName,
		language.LatinAmericanSpanish: flossieLatinAmericanSpanishName,
		language.Russian:              flossieRussianName,
		language.SimplifiedChinese:    flossieSimplifiedChineseName,
		language.Spanish:              flossieSpanishName,
		language.TraditionalChinese:   flossieTraditionalChineseName}
)

var (
	flossieCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: flossieBirthday,
		Code:     flossieCode,
		Key:      character.Flossie,
		Gender:   gender.Female,
		Name:     flossieName}
)

var (
	flossieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "squeaker"}

	flossieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	flossieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	flossieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "microbe"}

	flossieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quietsch"}

	flossieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "fischietto"}

	flossieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ガハハ"}

	flossieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	flossieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	flossieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	flossieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	flossieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "o sea"}

	flossieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	flossiePhrase = nook.Languages{
		language.AmericanEnglish:      flossieAmericanEnglishName,
		language.CanadianFrench:       flossieCanadianFrenchName,
		language.Dutch:                flossieDutchName,
		language.French:               flossieFrenchName,
		language.German:               flossieGermanName,
		language.Italian:              flossieItalianName,
		language.Japanese:             flossieJapaneseName,
		language.Korean:               flossieKoreanName,
		language.LatinAmericanSpanish: flossieLatinAmericanSpanishName,
		language.Russian:              flossieRussianName,
		language.SimplifiedChinese:    flossieSimplifiedChineseName,
		language.Spanish:              flossieSpanishName,
		language.TraditionalChinese:   flossieTraditionalChineseName}
)

var (
	Flossie = nook.Villager{
		Character:   flossieCharacter,
		Personality: personality.Peppy,
		Phrase:      flossiePhrase}
)
