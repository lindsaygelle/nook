package camel

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	saharahBirthday = nook.Birthday{
		Day:   10,
		Month: time.November}
)

var (
	saharahCode = nook.Code{
		Value: "cml"}
)

var (
	saharahAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Saharah"}

	saharahCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Sarah"}

	saharahDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Saharah"}

	saharahFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sarah"}

	saharahGermanName = nook.Name{
		Language: language.German,
		Value:    "Aziza"}

	saharahItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sahara"}

	saharahJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ローラン"}

	saharahKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "사하라"}

	saharahLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Alcatifa"}

	saharahRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сахара"}

	saharahSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "骆岚"}

	saharahSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Alcatifa"}

	saharahTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "駱嵐"}
)

var (
	saharahName = nook.Languages{
		language.AmericanEnglish:      saharahAmericanEnglishName,
		language.CanadianFrench:       saharahCanadianFrenchName,
		language.Dutch:                saharahDutchName,
		language.French:               saharahFrenchName,
		language.German:               saharahGermanName,
		language.Italian:              saharahItalianName,
		language.Japanese:             saharahJapaneseName,
		language.Korean:               saharahKoreanName,
		language.LatinAmericanSpanish: saharahLatinAmericanSpanishName,
		language.Russian:              saharahRussianName,
		language.SimplifiedChinese:    saharahSimplifiedChineseName,
		language.Spanish:              saharahSpanishName,
		language.TraditionalChinese:   saharahTraditionalChineseName}
)

var (
	saharahCharacter = nook.Character{
		Animal:   Camel,
		Birthday: saharahBirthday,
		Code:     saharahCode,
		Gender:   nook.Female,
		Name:     saharahName}
)

var (
	Saharah = nook.Resident{
		Character: saharahCharacter}
)
