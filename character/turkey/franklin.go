package turkey

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	franklinBirthday = nook.Birthday{
		Day:   10,
		Month: time.October}
)

var (
	franklinCode = nook.Code{
		Value: "tuk"}
)

var (
	franklinAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Franklin"}

	franklinCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Dindou"}

	franklinDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Franklin"}

	franklinFrenchName = nook.Name{
		Language: language.French,
		Value:    "Dindou"}

	franklinGermanName = nook.Name{
		Language: language.German,
		Value:    "Gernod"}

	franklinItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cedrone"}

	franklinJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フランクリン"}

	franklinKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "프랭클린"}

	franklinLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Guindo"}

	franklinRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Франклин"}

	franklinSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "富蓝"}

	franklinSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Guindo"}

	franklinTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "富藍"}
)

var (
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
	franklinCharacter = nook.Character{
		Animal:   animal.Turkey,
		Birthday: franklinBirthday,
		Code:     franklinCode,
		Gender:   gender.Male,
		Name:     franklinName}
)

var (
	Franklin = nook.Resident{
		Character: franklinCharacter}
)
