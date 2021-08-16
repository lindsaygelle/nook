package fox

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	reddBirthday = nook.Birthday{
		Day:   18,
		Month: time.October}
)

var (
	reddCode = nook.Code{
		Value: "fox"}
)

var (
	reddAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Redd"}

	reddCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rounard"}

	reddDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Redd"}

	reddFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rounard"}

	reddGermanName = nook.Name{
		Language: language.German,
		Value:    "Reiner"}

	reddItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Volpolo"}

	reddJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "つねきち"}

	reddKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "여욱"}

	reddLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ladino"}

	reddRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рэдд"}

	reddSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "狐利"}

	reddSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ladino"}

	reddTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "狐利"}
)

var (
	reddName = nook.Languages{
		language.AmericanEnglish:      reddAmericanEnglishName,
		language.CanadianFrench:       reddCanadianFrenchName,
		language.Dutch:                reddDutchName,
		language.French:               reddFrenchName,
		language.German:               reddGermanName,
		language.Italian:              reddItalianName,
		language.Japanese:             reddJapaneseName,
		language.Korean:               reddKoreanName,
		language.LatinAmericanSpanish: reddLatinAmericanSpanishName,
		language.Russian:              reddRussianName,
		language.SimplifiedChinese:    reddSimplifiedChineseName,
		language.Spanish:              reddSpanishName,
		language.TraditionalChinese:   reddTraditionalChineseName}
)

var (
	reddCharacter = nook.Character{
		Animal:   animal.Fox,
		Birthday: reddBirthday,
		Code:     reddCode,
		Gender:   gender.Male,
		Name:     reddName}
)

var (
	Redd = nook.Resident{
		Character: reddCharacter}
)
