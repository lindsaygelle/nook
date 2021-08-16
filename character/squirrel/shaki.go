package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	shakiBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	shakiCode = nook.Code{
		Value: "xsq"}
)

var (
	shakiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Shaki"}

	shakiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	shakiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	shakiFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	shakiGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	shakiItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	shakiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シャキッ"}

	shakiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	shakiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	shakiRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	shakiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	shakiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	shakiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	shakiName = nook.Languages{
		language.AmericanEnglish:      shakiAmericanEnglishName,
		language.CanadianFrench:       shakiCanadianFrenchName,
		language.Dutch:                shakiDutchName,
		language.French:               shakiFrenchName,
		language.German:               shakiGermanName,
		language.Italian:              shakiItalianName,
		language.Japanese:             shakiJapaneseName,
		language.Korean:               shakiKoreanName,
		language.LatinAmericanSpanish: shakiLatinAmericanSpanishName,
		language.Russian:              shakiRussianName,
		language.SimplifiedChinese:    shakiSimplifiedChineseName,
		language.Spanish:              shakiSpanishName,
		language.TraditionalChinese:   shakiTraditionalChineseName}
)

var (
	shakiCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: shakiBirthday,
		Code:     shakiCode,
		Key:      character.Shaki,
		Gender:   gender.Female,
		Name:     shakiName}
)

var (
	Shaki = nook.Resident{
		Character: shakiCharacter}
)
