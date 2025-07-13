package pelican

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// pellyBirthday represents pelly birthday.
	pellyBirthday = nook.Birthday{
		Day:   19,
		Month: time.March}
)

var (
	// pellyCode represents pelly code.
	pellyCode = nook.Code{
		Value: "pga/plk"}
)

var (
	// pellyAmericanEnglishName represents pelly american english name.
	pellyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pelly"}

	// pellyCanadianFrenchName represents pelly canadian french name.
	pellyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Opélie"}

	// pellyDutchName represents pelly dutch name.
	pellyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pelly"}

	// pellyFrenchName represents pelly french name.
	pellyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Opélie"}

	// pellyGermanName represents pelly german name.
	pellyGermanName = nook.Name{
		Language: language.German,
		Value:    "Pelly"}

	// pellyItalianName represents pelly italian name.
	pellyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pelly"}

	// pellyJapaneseName represents pelly japanese name.
	pellyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ぺりこ"}

	// pellyKoreanName represents pelly korean name.
	pellyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "펠리"}

	// pellyLatinAmericanSpanishName represents pelly latin american spanish name.
	pellyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sol"}

	// pellyRussianName represents pelly russian name.
	pellyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пелли"}

	// pellySimplifiedChineseName represents pelly simplified chinese name.
	pellySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "宋信子"}

	// pellySpanishName represents pelly spanish name.
	pellySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sol"}

	// pellyTraditionalChineseName represents pelly traditional chinese name.
	pellyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "宋信子"}
)

var (
	// pellyName represents pelly name.
	pellyName = nook.Languages{
		language.AmericanEnglish:      pellyAmericanEnglishName,
		language.CanadianFrench:       pellyCanadianFrenchName,
		language.Dutch:                pellyDutchName,
		language.French:               pellyFrenchName,
		language.German:               pellyGermanName,
		language.Italian:              pellyItalianName,
		language.Japanese:             pellyJapaneseName,
		language.Korean:               pellyKoreanName,
		language.LatinAmericanSpanish: pellyLatinAmericanSpanishName,
		language.Russian:              pellyRussianName,
		language.SimplifiedChinese:    pellySimplifiedChineseName,
		language.Spanish:              pellySpanishName,
		language.TraditionalChinese:   pellyTraditionalChineseName}
)

var (
	// pellyCharacter represents pelly character.
	pellyCharacter = nook.Character{
		Animal:   animal.Pelican,
		Birthday: pellyBirthday,
		Code:     pellyCode,
		Key:      character.Pelly,
		Gender:   gender.Female,
		Name:     pellyName,
		Special:  true}
)

var (
	// Pelly represents pelly.
	Pelly = nook.Resident{
		Character: pellyCharacter}
)
