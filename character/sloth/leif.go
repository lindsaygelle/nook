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
	// leifBirthday represents leif birthday.
	leifBirthday = nook.Birthday{
		Day:   8,
		Month: time.August}
)

var (
	// leifCode represents leif code.
	leifCode = nook.Code{
		Value: "slo"}
)

var (
	// leifAmericanEnglishName represents leif american english name.
	leifAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Leif"}

	// leifCanadianFrenchName represents leif canadian french name.
	leifCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Racine"}

	// leifDutchName represents leif dutch name.
	leifDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Leif"}

	// leifFrenchName represents leif french name.
	leifFrenchName = nook.Name{
		Language: language.French,
		Value:    "Racine"}

	// leifGermanName represents leif german name.
	leifGermanName = nook.Name{
		Language: language.German,
		Value:    "Gerd"}

	// leifItalianName represents leif italian name.
	leifItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Florindo"}

	// leifJapaneseName represents leif japanese name.
	leifJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "レイジ"}

	// leifKoreanName represents leif korean name.
	leifKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "늘봉"}

	// leifLatinAmericanSpanishName represents leif latin american spanish name.
	leifLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gandulio"}

	// leifRussianName represents leif russian name.
	leifRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лейф"}

	// leifSimplifiedChineseName represents leif simplified chinese name.
	leifSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "然然"}

	// leifSpanishName represents leif spanish name.
	leifSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gandulio"}

	// leifTraditionalChineseName represents leif traditional chinese name.
	leifTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "然然"}
)

var (
	// leifName represents leif name.
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
	// leifCharacter represents leif character.
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
	// Leif represents leif.
	Leif = nook.Resident{
		Character: leifCharacter}
)
