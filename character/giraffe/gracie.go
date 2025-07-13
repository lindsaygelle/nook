package giraffe

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// gracieBirthday represents gracie birthday.
	gracieBirthday = nook.Birthday{
		Day:   14,
		Month: time.November}
)

var (
	// gracieCode represents gracie code.
	gracieCode = nook.Code{
		Value: "grf"}
)

var (
	// gracieAmericanEnglishName represents gracie american english name.
	gracieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gracie"}

	// gracieCanadianFrenchName represents gracie canadian french name.
	gracieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Carla"}

	// gracieDutchName represents gracie dutch name.
	gracieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gracie"}

	// gracieFrenchName represents gracie french name.
	gracieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Carla"}

	// gracieGermanName represents gracie german name.
	gracieGermanName = nook.Name{
		Language: language.German,
		Value:    "Grazia"}

	// gracieItalianName represents gracie italian name.
	gracieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Griffa"}

	// gracieJapaneseName represents gracie japanese name.
	gracieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "グレース"}

	// gracieKoreanName represents gracie korean name.
	gracieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "그레이스"}

	// gracieLatinAmericanSpanishName represents gracie latin american spanish name.
	gracieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Graciela"}

	// gracieRussianName represents gracie russian name.
	gracieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Грейси"}

	// gracieSimplifiedChineseName represents gracie simplified chinese name.
	gracieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "葛瑞斯"}

	// gracieSpanishName represents gracie spanish name.
	gracieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Graciela"}

	// gracieTraditionalChineseName represents gracie traditional chinese name.
	gracieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "葛瑞斯"}
)

var (
	// gracieName represents gracie name.
	gracieName = nook.Languages{
		language.AmericanEnglish:      gracieAmericanEnglishName,
		language.CanadianFrench:       gracieCanadianFrenchName,
		language.Dutch:                gracieDutchName,
		language.French:               gracieFrenchName,
		language.German:               gracieGermanName,
		language.Italian:              gracieItalianName,
		language.Japanese:             gracieJapaneseName,
		language.Korean:               gracieKoreanName,
		language.LatinAmericanSpanish: gracieLatinAmericanSpanishName,
		language.Russian:              gracieRussianName,
		language.SimplifiedChinese:    gracieSimplifiedChineseName,
		language.Spanish:              gracieSpanishName,
		language.TraditionalChinese:   gracieTraditionalChineseName}
)

var (
	// gracieCharacter represents gracie character.
	gracieCharacter = nook.Character{
		Animal:   animal.Giraffe,
		Birthday: gracieBirthday,
		Code:     gracieCode,
		Key:      character.Gracie,
		Gender:   gender.Female,
		Name:     gracieName,
		Special:  true}
)

var (
	// Gracie represents gracie.
	Gracie = nook.Resident{
		Character: gracieCharacter}
)
