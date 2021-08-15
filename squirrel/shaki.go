package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "N/A"}

	shakiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	shakiFrenchName = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	shakiGermanName = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	shakiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	shakiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シャキッ"}

	shakiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	shakiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	shakiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	shakiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	shakiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	shakiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Animal:   Squirrel,
		Birthday: shakiBirthday,
		Code:     shakiCode,
		Gender:   nook.Female,
		Name:     shakiName}
)

var (
	Shaki = nook.Resident{
		Character: shakiCharacter}
)
