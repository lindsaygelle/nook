package gyroid

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// lloidBirthday represents lloid birthday.
	lloidBirthday = nook.Birthday{
		Day:   28,
		Month: time.August}
)

var (
	// lloidCode represents lloid code.
	lloidCode = nook.Code{
		Value: "hnw"}
)

var (
	// lloidAmericanEnglishName represents lloid american english name.
	lloidAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lloid"}

	// lloidCanadianFrenchName represents lloid canadian french name.
	lloidCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gyroïde"}

	// lloidDutchName represents lloid dutch name.
	lloidDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lloid"}

	// lloidFrenchName represents lloid french name.
	lloidFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gyroïde"}

	// lloidGermanName represents lloid german name.
	lloidGermanName = nook.Name{
		Language: language.German,
		Value:    "Gyroid"}

	// lloidItalianName represents lloid italian name.
	lloidItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gironio"}

	// lloidJapaneseName represents lloid japanese name.
	lloidJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハニワくん"}

	// lloidKoreanName represents lloid korean name.
	lloidKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "토용군"}

	// lloidLatinAmericanSpanishName represents lloid latin american spanish name.
	lloidLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Giroide"}

	// lloidRussianName represents lloid russian name.
	lloidRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Ллойд"}

	// lloidSimplifiedChineseName represents lloid simplified chinese name.
	lloidSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "俑俑"}

	// lloidSpanishName represents lloid spanish name.
	lloidSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Giroide"}

	// lloidTraditionalChineseName represents lloid traditional chinese name.
	lloidTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "俑俑"}
)

var (
	// lloidName represents lloid name.
	lloidName = nook.Languages{
		language.AmericanEnglish:      lloidAmericanEnglishName,
		language.CanadianFrench:       lloidCanadianFrenchName,
		language.Dutch:                lloidDutchName,
		language.French:               lloidFrenchName,
		language.German:               lloidGermanName,
		language.Italian:              lloidItalianName,
		language.Japanese:             lloidJapaneseName,
		language.Korean:               lloidKoreanName,
		language.LatinAmericanSpanish: lloidLatinAmericanSpanishName,
		language.Russian:              lloidRussianName,
		language.SimplifiedChinese:    lloidSimplifiedChineseName,
		language.Spanish:              lloidSpanishName,
		language.TraditionalChinese:   lloidTraditionalChineseName}
)

var (
	// lloidCharacter represents lloid character.
	lloidCharacter = nook.Character{
		Animal:   animal.Gyroid,
		Birthday: lloidBirthday,
		Code:     lloidCode,
		Key:      character.Lloid,
		Gender:   gender.Male,
		Name:     lloidName,
		Special:  true}
)

var (
	// Lloid represents lloid.
	Lloid = nook.Resident{
		Character: lloidCharacter}
)
