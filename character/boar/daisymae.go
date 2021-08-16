package boar

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	daisymaeBirthday = nook.Birthday{
		Day:   5,
		Month: time.May}
)

var (
	daisymaeCode = nook.Code{
		Value: "boc"}
)

var (
	daisymaeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Daisy Mae"}

	daisymaeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Porcelette"}

	daisymaeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Daisy Mae"}

	daisymaeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Porcelette"}

	daisymaeGermanName = nook.Name{
		Language: language.German,
		Value:    "Jorna"}

	daisymaeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Brunella"}

	daisymaeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ウリ"}

	daisymaeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "무파니"}

	daisymaeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Juliana"}

	daisymaeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дейзи-Мей"}

	daisymaeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "曹卖"}

	daisymaeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Juliana"}

	daisymaeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "曹賣"}
)

var (
	daisymaeName = nook.Languages{
		language.AmericanEnglish:      daisymaeAmericanEnglishName,
		language.CanadianFrench:       daisymaeCanadianFrenchName,
		language.Dutch:                daisymaeDutchName,
		language.French:               daisymaeFrenchName,
		language.German:               daisymaeGermanName,
		language.Italian:              daisymaeItalianName,
		language.Japanese:             daisymaeJapaneseName,
		language.Korean:               daisymaeKoreanName,
		language.LatinAmericanSpanish: daisymaeLatinAmericanSpanishName,
		language.Russian:              daisymaeRussianName,
		language.SimplifiedChinese:    daisymaeSimplifiedChineseName,
		language.Spanish:              daisymaeSpanishName,
		language.TraditionalChinese:   daisymaeTraditionalChineseName}
)

var (
	daisymaeCharacter = nook.Character{
		Animal:   animal.Boar,
		Birthday: daisymaeBirthday,
		Code:     daisymaeCode,
		Key:      character.DaisyMae,
		Gender:   gender.Female,
		Name:     daisymaeName}
)

var (
	DaisyMae = nook.Resident{
		Character: daisymaeCharacter}
)
