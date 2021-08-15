package otter

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	pascalBirthday = nook.Birthday{
		Day:   19,
		Month: time.July}
)

var (
	pascalCode = nook.Code{
		Value: "seo"}
)

var (
	pascalAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pascal"}

	pascalCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pascal"}

	pascalDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pascal"}

	pascalFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pascal"}

	pascalGermanName = nook.Name{
		Language: language.German,
		Value:    "Johannes"}

	pascalItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pasqualo"}

	pascalJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ラコスケ"}

	pascalKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "해탈한"}

	pascalLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pascal"}

	pascalRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Паскаль"}

	pascalSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿獭"}

	pascalSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pascal"}

	pascalTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿獺"}
)

var (
	pascalName = nook.Languages{
		language.AmericanEnglish:      pascalAmericanEnglishName,
		language.CanadianFrench:       pascalCanadianFrenchName,
		language.Dutch:                pascalDutchName,
		language.French:               pascalFrenchName,
		language.German:               pascalGermanName,
		language.Italian:              pascalItalianName,
		language.Japanese:             pascalJapaneseName,
		language.Korean:               pascalKoreanName,
		language.LatinAmericanSpanish: pascalLatinAmericanSpanishName,
		language.Russian:              pascalRussianName,
		language.SimplifiedChinese:    pascalSimplifiedChineseName,
		language.Spanish:              pascalSpanishName,
		language.TraditionalChinese:   pascalTraditionalChineseName}
)

var (
	pascalCharacter = nook.Character{
		Animal:   Otter,
		Birthday: pascalBirthday,
		Code:     pascalCode,
		Gender:   nook.Male,
		Name:     pascalName}
)

var (
	Pascal = nook.Resident{
		Character: pascalCharacter}
)
