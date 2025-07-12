package turkey

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// franklinBirthday represents franklin birthday.
	franklinBirthday = nook.Birthday{
		Day:   10,
		Month: time.October}
)

var (
	// franklinCode represents franklin code.
	franklinCode = nook.Code{
		Value: "tuk"}
)

var (
	// franklinAmericanEnglishName represents franklin american english name.
	franklinAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Franklin"}

	// franklinCanadianFrenchName represents franklin canadian french name.
	franklinCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Dindou"}

	// franklinDutchName represents franklin dutch name.
	franklinDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Franklin"}

	// franklinFrenchName represents franklin french name.
	franklinFrenchName = nook.Name{
		Language: language.French,
		Value:    "Dindou"}

	// franklinGermanName represents franklin german name.
	franklinGermanName = nook.Name{
		Language: language.German,
		Value:    "Gernod"}

	// franklinItalianName represents franklin italian name.
	franklinItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cedrone"}

	// franklinJapaneseName represents franklin japanese name.
	franklinJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フランクリン"}

	// franklinKoreanName represents franklin korean name.
	franklinKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "프랭클린"}

	// franklinLatinAmericanSpanishName represents franklin latin american spanish name.
	franklinLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Guindo"}

	// franklinRussianName represents franklin russian name.
	franklinRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Франклин"}

	// franklinSimplifiedChineseName represents franklin simplified chinese name.
	franklinSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "富蓝"}

	// franklinSpanishName represents franklin spanish name.
	franklinSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Guindo"}

	// franklinTraditionalChineseName represents franklin traditional chinese name.
	franklinTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "富藍"}
)

var (
	// franklinName represents franklin name.
	franklinName = nook.Languages{
		language.AmericanEnglish:      franklinAmericanEnglishName,
		language.CanadianFrench:       franklinCanadianFrenchName,
		language.Dutch:                franklinDutchName,
		language.French:               franklinFrenchName,
		language.German:               franklinGermanName,
		language.Italian:              franklinItalianName,
		language.Japanese:             franklinJapaneseName,
		language.Korean:               franklinKoreanName,
		language.LatinAmericanSpanish: franklinLatinAmericanSpanishName,
		language.Russian:              franklinRussianName,
		language.SimplifiedChinese:    franklinSimplifiedChineseName,
		language.Spanish:              franklinSpanishName,
		language.TraditionalChinese:   franklinTraditionalChineseName}
)

var (
	// franklinCharacter represents franklin character.
	franklinCharacter = nook.Character{
		Animal:   animal.Turkey,
		Birthday: franklinBirthday,
		Code:     franklinCode,
		Key:      character.Franklin,
		Gender:   gender.Male,
		Name:     franklinName,
		Special:  true}
)

var (
	// Franklin represents franklin.
	Franklin = nook.Resident{
		Character: franklinCharacter}
)
