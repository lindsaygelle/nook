package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// blancaBirthday represents blanca birthday.
	blancaBirthday = nook.Birthday{
		Day:   8,
		Month: time.February}
)

var (
	// blancaCode represents blanca code.
	blancaCode = nook.Code{
		Value: "mka"}
)

var (
	// blancaAmericanEnglishName represents blanca american english name.
	blancaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Blanca"}

	// blancaCanadianFrenchName represents blanca canadian french name.
	blancaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Blanca"}

	// blancaDutchName represents blanca dutch name.
	blancaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Blanca"}

	// blancaFrenchName represents blanca french name.
	blancaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Blanca"}

	// blancaGermanName represents blanca german name.
	blancaGermanName = nook.Name{
		Language: language.German,
		Value:    "Blanka"}

	// blancaItalianName represents blanca italian name.
	blancaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Blanca"}

	// blancaJapaneseName represents blanca japanese name.
	blancaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "あやしいネコ"}

	// blancaKoreanName represents blanca korean name.
	blancaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아트고양이"}

	// blancaLatinAmericanSpanishName represents blanca latin american spanish name.
	blancaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Blanca"}

	// blancaRussianName represents blanca russian name.
	blancaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бланка"}

	// blancaSimplifiedChineseName represents blanca simplified chinese name.
	blancaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "无脸猫"}

	// blancaSpanishName represents blanca spanish name.
	blancaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Blanca"}

	// blancaTraditionalChineseName represents blanca traditional chinese name.
	blancaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "無臉貓"}
)

var (
	// blancaName represents blanca name.
	blancaName = nook.Languages{
		language.AmericanEnglish:      blancaAmericanEnglishName,
		language.CanadianFrench:       blancaCanadianFrenchName,
		language.Dutch:                blancaDutchName,
		language.French:               blancaFrenchName,
		language.German:               blancaGermanName,
		language.Italian:              blancaItalianName,
		language.Japanese:             blancaJapaneseName,
		language.Korean:               blancaKoreanName,
		language.LatinAmericanSpanish: blancaLatinAmericanSpanishName,
		language.Russian:              blancaRussianName,
		language.SimplifiedChinese:    blancaSimplifiedChineseName,
		language.Spanish:              blancaSpanishName,
		language.TraditionalChinese:   blancaTraditionalChineseName}
)

var (
	// blancaCharacter represents blanca character.
	blancaCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: blancaBirthday,
		Code:     blancaCode,
		Key:      character.Blanca,
		Gender:   gender.Female,
		Name:     blancaName,
		Special:  true}
)

var (
	// Blanca represents blanca.
	Blanca = nook.Resident{
		Character: blancaCharacter}
)
