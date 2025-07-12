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
	// daisymaeBirthday represents daisymae birthday.
	daisymaeBirthday = nook.Birthday{
		Day:   5,
		Month: time.May}
)

var (
	// daisymaeCode represents daisymae code.
	daisymaeCode = nook.Code{
		Value: "boc"}
)

var (
	// daisymaeAmericanEnglishName represents daisymae american english name.
	daisymaeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Daisy Mae"}

	// daisymaeCanadianFrenchName represents daisymae canadian french name.
	daisymaeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Porcelette"}

	// daisymaeDutchName represents daisymae dutch name.
	daisymaeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Daisy Mae"}

	// daisymaeFrenchName represents daisymae french name.
	daisymaeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Porcelette"}

	// daisymaeGermanName represents daisymae german name.
	daisymaeGermanName = nook.Name{
		Language: language.German,
		Value:    "Jorna"}

	// daisymaeItalianName represents daisymae italian name.
	daisymaeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Brunella"}

	// daisymaeJapaneseName represents daisymae japanese name.
	daisymaeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ウリ"}

	// daisymaeKoreanName represents daisymae korean name.
	daisymaeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "무파니"}

	// daisymaeLatinAmericanSpanishName represents daisymae latin american spanish name.
	daisymaeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Juliana"}

	// daisymaeRussianName represents daisymae russian name.
	daisymaeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дейзи-Мей"}

	// daisymaeSimplifiedChineseName represents daisymae simplified chinese name.
	daisymaeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "曹卖"}

	// daisymaeSpanishName represents daisymae spanish name.
	daisymaeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Juliana"}

	// daisymaeTraditionalChineseName represents daisymae traditional chinese name.
	daisymaeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "曹賣"}
)

var (
	// daisymaeName represents daisymae name.
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
	// daisymaeCharacter represents daisymae character.
	daisymaeCharacter = nook.Character{
		Animal:   animal.Boar,
		Birthday: daisymaeBirthday,
		Code:     daisymaeCode,
		Key:      character.DaisyMae,
		Gender:   gender.Female,
		Name:     daisymaeName,
		Special:  true}
)

var (
	// DaisyMae represents daisy mae.
	DaisyMae = nook.Resident{
		Character: daisymaeCharacter}
)
