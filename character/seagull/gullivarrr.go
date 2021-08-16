package seagull

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	gullivarrrBirthday = nook.Birthday{
		Day:   25,
		Month: time.May}
)

var (
	gullivarrrCode = nook.Code{
		Value: "gulB"}
)

var (
	gullivarrrAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gullivarrr"}

	gullivarrrCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gullivarrr"}

	gullivarrrDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gullivarrr"}

	gullivarrrFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gullivarrr"}

	gullivarrrGermanName = nook.Name{
		Language: language.German,
		Value:    "Gullivarrr"}

	gullivarrrItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gullivarrr"}

	gullivarrrJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "かいぞく"}

	gullivarrrKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "해적"}

	gullivarrrLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gullivarrr"}

	gullivarrrRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гулливарр"}

	gullivarrrSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "海盗"}

	gullivarrrSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gullivarrr"}

	gullivarrrTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "海盜"}
)

var (
	gullivarrrName = nook.Languages{
		language.AmericanEnglish:      gullivarrrAmericanEnglishName,
		language.CanadianFrench:       gullivarrrCanadianFrenchName,
		language.Dutch:                gullivarrrDutchName,
		language.French:               gullivarrrFrenchName,
		language.German:               gullivarrrGermanName,
		language.Italian:              gullivarrrItalianName,
		language.Japanese:             gullivarrrJapaneseName,
		language.Korean:               gullivarrrKoreanName,
		language.LatinAmericanSpanish: gullivarrrLatinAmericanSpanishName,
		language.Russian:              gullivarrrRussianName,
		language.SimplifiedChinese:    gullivarrrSimplifiedChineseName,
		language.Spanish:              gullivarrrSpanishName,
		language.TraditionalChinese:   gullivarrrTraditionalChineseName}
)

var (
	gullivarrrCharacter = nook.Character{
		Animal:   animal.Seagull,
		Birthday: gullivarrrBirthday,
		Code:     gullivarrrCode,
		Gender:   gender.Male,
		Name:     gullivarrrName}
)

var (
	Gullivarrr = nook.Resident{
		Character: gullivarrrCharacter}
)
