package sloth

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	leifBirthday = nook.Birthday{
		Day:   8,
		Month: time.August}
)

var (
	leifCode = nook.Code{
		Value: "slo"}
)

var (
	leifAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Leif"}

	leifCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Racine"}

	leifDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Leif"}

	leifFrenchName = nook.Name{
		Language: language.French,
		Value:    "Racine"}

	leifGermanName = nook.Name{
		Language: language.German,
		Value:    "Gerd"}

	leifItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Florindo"}

	leifJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "レイジ"}

	leifKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "늘봉"}

	leifLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gandulio"}

	leifRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лейф"}

	leifSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "然然"}

	leifSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gandulio"}

	leifTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "然然"}
)

var (
	leifName = nook.Languages{
		language.AmericanEnglish:      leifAmericanEnglishName,
		language.CanadianFrench:       leifCanadianFrenchName,
		language.Dutch:                leifDutchName,
		language.French:               leifFrenchName,
		language.German:               leifGermanName,
		language.Italian:              leifItalianName,
		language.Japanese:             leifJapaneseName,
		language.Korean:               leifKoreanName,
		language.LatinAmericanSpanish: leifLatinAmericanSpanishName,
		language.Russian:              leifRussianName,
		language.SimplifiedChinese:    leifSimplifiedChineseName,
		language.Spanish:              leifSpanishName,
		language.TraditionalChinese:   leifTraditionalChineseName}
)

var (
	leifCharacter = nook.Character{
		Animal:   animal.Sloth,
		Birthday: leifBirthday,
		Code:     leifCode,
		Key:      character.Leif,
		Gender:   gender.Male,
		Name:     leifName,
		Special:  true}
)

var (
	Leif = nook.Resident{
		Character: leifCharacter}
)
