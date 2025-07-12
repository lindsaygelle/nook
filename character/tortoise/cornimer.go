package tortoise

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// cornimerBirthday represents cornimer birthday.
	cornimerBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// cornimerCode represents cornimer code.
	cornimerCode = nook.Code{
		Value: "dnk"}
)

var (
	// cornimerAmericanEnglishName represents cornimer american english name.
	cornimerAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cornimer"}

	// cornimerCanadianFrenchName represents cornimer canadian french name.
	cornimerCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cornimer"}

	// cornimerDutchName represents cornimer dutch name.
	cornimerDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// cornimerFrenchName represents cornimer french name.
	cornimerFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cornimer"}

	// cornimerGermanName represents cornimer german name.
	cornimerGermanName = nook.Name{
		Language: language.German,
		Value:    "Eichenm."}

	// cornimerItalianName represents cornimer italian name.
	cornimerItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ghiandy"}

	// cornimerJapaneseName represents cornimer japanese name.
	cornimerJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ドンどんぐり"}

	// cornimerKoreanName represents cornimer korean name.
	cornimerKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "도토리옹"}

	// cornimerLatinAmericanSpanishName represents cornimer latin american spanish name.
	cornimerLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Córnimer"}

	// cornimerRussianName represents cornimer russian name.
	cornimerRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// cornimerSimplifiedChineseName represents cornimer simplified chinese name.
	cornimerSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// cornimerSpanishName represents cornimer spanish name.
	cornimerSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Córnimer"}

	// cornimerTraditionalChineseName represents cornimer traditional chinese name.
	cornimerTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// cornimerName represents cornimer name.
	cornimerName = nook.Languages{
		language.AmericanEnglish:      cornimerAmericanEnglishName,
		language.CanadianFrench:       cornimerCanadianFrenchName,
		language.Dutch:                cornimerDutchName,
		language.French:               cornimerFrenchName,
		language.German:               cornimerGermanName,
		language.Italian:              cornimerItalianName,
		language.Japanese:             cornimerJapaneseName,
		language.Korean:               cornimerKoreanName,
		language.LatinAmericanSpanish: cornimerLatinAmericanSpanishName,
		language.Russian:              cornimerRussianName,
		language.SimplifiedChinese:    cornimerSimplifiedChineseName,
		language.Spanish:              cornimerSpanishName,
		language.TraditionalChinese:   cornimerTraditionalChineseName}
)

var (
	// cornimerCharacter represents cornimer character.
	cornimerCharacter = nook.Character{
		Animal:   animal.Tortoise,
		Birthday: cornimerBirthday,
		Code:     cornimerCode,
		Key:      character.Cornimer,
		Gender:   gender.Male,
		Name:     cornimerName,
		Special:  true}
)

var (
	// Cornimer represents cornimer.
	Cornimer = nook.Resident{
		Character: cornimerCharacter}
)
