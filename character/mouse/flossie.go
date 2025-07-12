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
	// flossieBirthday represents flossie birthday.
	flossieBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// flossieCode represents flossie code.
	flossieCode = nook.Code{
		Value: ""}
)

var (
	// flossieAmericanEnglishName represents flossie american english name.
	flossieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Flossie"}

	// flossieCanadianFrenchName represents flossie canadian french name.
	flossieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// flossieDutchName represents flossie dutch name.
	flossieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// flossieFrenchName represents flossie french name.
	flossieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chiquita"}

	// flossieGermanName represents flossie german name.
	flossieGermanName = nook.Name{
		Language: language.German,
		Value:    "Amalie"}

	// flossieItalianName represents flossie italian name.
	flossieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Florinda"}

	// flossieJapaneseName represents flossie japanese name.
	flossieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピスタチオ"}

	// flossieKoreanName represents flossie korean name.
	flossieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// flossieLatinAmericanSpanishName represents flossie latin american spanish name.
	flossieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// flossieRussianName represents flossie russian name.
	flossieRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// flossieSimplifiedChineseName represents flossie simplified chinese name.
	flossieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// flossieSpanishName represents flossie spanish name.
	flossieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Oregilda"}

	// flossieTraditionalChineseName represents flossie traditional chinese name.
	flossieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// flossieName represents flossie name.
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
	// flossieCharacter represents flossie character.
	flossieCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: flossieBirthday,
		Code:     flossieCode,
		Key:      character.Flossie,
		Gender:   gender.Female,
		Name:     flossieName,
		Special:  false}
)

var (
	// flossieAmericanEnglishPhrase represents flossie american english phrase.
	flossieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "squeaker"}

	// flossieCanadianFrenchPhrase represents flossie canadian french phrase.
	flossieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// flossieDutchPhrase represents flossie dutch phrase.
	flossieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// flossieFrenchPhrase represents flossie french phrase.
	flossieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "microbe"}

	// flossieGermanPhrase represents flossie german phrase.
	flossieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quietsch"}

	// flossieItalianPhrase represents flossie italian phrase.
	flossieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "fischietto"}

	// flossieJapanesePhrase represents flossie japanese phrase.
	flossieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ガハハ"}

	// flossieKoreanPhrase represents flossie korean phrase.
	flossieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// flossieLatinAmericanSpanishPhrase represents flossie latin american spanish phrase.
	flossieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// flossieRussianPhrase represents flossie russian phrase.
	flossieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// flossieSimplifiedChinesePhrase represents flossie simplified chinese phrase.
	flossieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// flossieSpanishPhrase represents flossie spanish phrase.
	flossieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "o sea"}

	// flossieTraditionalChinesePhrase represents flossie traditional chinese phrase.
	flossieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// flossiePhrase represents flossie phrase.
	flossiePhrase = nook.Languages{
		language.AmericanEnglish:      flossieAmericanEnglishPhrase,
		language.CanadianFrench:       flossieCanadianFrenchPhrase,
		language.Dutch:                flossieDutchPhrase,
		language.French:               flossieFrenchPhrase,
		language.German:               flossieGermanPhrase,
		language.Italian:              flossieItalianPhrase,
		language.Japanese:             flossieJapanesePhrase,
		language.Korean:               flossieKoreanPhrase,
		language.LatinAmericanSpanish: flossieLatinAmericanSpanishPhrase,
		language.Russian:              flossieRussianPhrase,
		language.SimplifiedChinese:    flossieSimplifiedChinesePhrase,
		language.Spanish:              flossieSpanishPhrase,
		language.TraditionalChinese:   flossieTraditionalChinesePhrase}
)

var (
	// Flossie represents flossie.
	Flossie = nook.Villager{
		Character:   flossieCharacter,
		Personality: personality.Peppy,
		Phrase:      flossiePhrase}
)
