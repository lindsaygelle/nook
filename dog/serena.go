package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	serenaBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	serenaCode = nook.Code{
		Value: "gds"}
)

var (
	serenaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Serena"}

	serenaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Unknown"}

	serenaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	serenaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Unknown"}

	serenaGermanName = nook.Name{
		Language: language.German,
		Value:    "Divahua"}

	serenaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Candea"}

	serenaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "めがみさま"}

	serenaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "Unknown"}

	serenaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Divahua"}

	serenaRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	serenaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	serenaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Divahua"}

	serenaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	serenaName = nook.Languages{
		language.AmericanEnglish:      serenaAmericanEnglishName,
		language.CanadianFrench:       serenaCanadianFrenchName,
		language.Dutch:                serenaDutchName,
		language.French:               serenaFrenchName,
		language.German:               serenaGermanName,
		language.Italian:              serenaItalianName,
		language.Japanese:             serenaJapaneseName,
		language.Korean:               serenaKoreanName,
		language.LatinAmericanSpanish: serenaLatinAmericanSpanishName,
		language.Russian:              serenaRussianName,
		language.SimplifiedChinese:    serenaSimplifiedChineseName,
		language.Spanish:              serenaSpanishName,
		language.TraditionalChinese:   serenaTraditionalChineseName}
)

var (
	serenaCharacter = nook.Character{
		Animal:   Dog,
		Birthday: serenaBirthday,
		Code:     serenaCode,
		Gender:   nook.Female,
		Name:     serenaName}
)

var (
	Serena = nook.Resident{
		Character: serenaCharacter}
)
