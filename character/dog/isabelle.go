package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// isabelleBirthday represents isabelle birthday.
	isabelleBirthday = nook.Birthday{
		Day:   20,
		Month: time.December}
)

var (
	// isabelleCode represents isabelle code.
	isabelleCode = nook.Code{
		Value: "sza"}
)

var (
	// isabelleAmericanEnglishName represents isabelle american english name.
	isabelleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Isabelle"}

	// isabelleCanadianFrenchName represents isabelle canadian french name.
	isabelleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Marie"}

	// isabelleDutchName represents isabelle dutch name.
	isabelleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Isabelle"}

	// isabelleFrenchName represents isabelle french name.
	isabelleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marie"}

	// isabelleGermanName represents isabelle german name.
	isabelleGermanName = nook.Name{
		Language: language.German,
		Value:    "Melinda"}

	// isabelleItalianName represents isabelle italian name.
	isabelleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fuffi"}

	// isabelleJapaneseName represents isabelle japanese name.
	isabelleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "しずえ"}

	// isabelleKoreanName represents isabelle korean name.
	isabelleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "여울"}

	// isabelleLatinAmericanSpanishName represents isabelle latin american spanish name.
	isabelleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Canela"}

	// isabelleRussianName represents isabelle russian name.
	isabelleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Изабель"}

	// isabelleSimplifiedChineseName represents isabelle simplified chinese name.
	isabelleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "西施惠"}

	// isabelleSpanishName represents isabelle spanish name.
	isabelleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Canela"}

	// isabelleTraditionalChineseName represents isabelle traditional chinese name.
	isabelleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "西施惠"}
)

var (
	// isabelleName represents isabelle name.
	isabelleName = nook.Languages{
		language.AmericanEnglish:      isabelleAmericanEnglishName,
		language.CanadianFrench:       isabelleCanadianFrenchName,
		language.Dutch:                isabelleDutchName,
		language.French:               isabelleFrenchName,
		language.German:               isabelleGermanName,
		language.Italian:              isabelleItalianName,
		language.Japanese:             isabelleJapaneseName,
		language.Korean:               isabelleKoreanName,
		language.LatinAmericanSpanish: isabelleLatinAmericanSpanishName,
		language.Russian:              isabelleRussianName,
		language.SimplifiedChinese:    isabelleSimplifiedChineseName,
		language.Spanish:              isabelleSpanishName,
		language.TraditionalChinese:   isabelleTraditionalChineseName}
)

var (
	// isabelleCharacter represents isabelle character.
	isabelleCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: isabelleBirthday,
		Code:     isabelleCode,
		Key:      character.Isabelle,
		Gender:   gender.Female,
		Name:     isabelleName,
		Special:  true}
)

var (
	// Isabelle represents isabelle.
	Isabelle = nook.Resident{
		Character: isabelleCharacter}
)
