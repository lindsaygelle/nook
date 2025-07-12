package fox

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// reddBirthday represents redd birthday.
	reddBirthday = nook.Birthday{
		Day:   18,
		Month: time.October}
)

var (
	// reddCode represents redd code.
	reddCode = nook.Code{
		Value: "fox"}
)

var (
	// reddAmericanEnglishName represents redd american english name.
	reddAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Redd"}

	// reddCanadianFrenchName represents redd canadian french name.
	reddCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rounard"}

	// reddDutchName represents redd dutch name.
	reddDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Redd"}

	// reddFrenchName represents redd french name.
	reddFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rounard"}

	// reddGermanName represents redd german name.
	reddGermanName = nook.Name{
		Language: language.German,
		Value:    "Reiner"}

	// reddItalianName represents redd italian name.
	reddItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Volpolo"}

	// reddJapaneseName represents redd japanese name.
	reddJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "つねきち"}

	// reddKoreanName represents redd korean name.
	reddKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "여욱"}

	// reddLatinAmericanSpanishName represents redd latin american spanish name.
	reddLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ladino"}

	// reddRussianName represents redd russian name.
	reddRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рэдд"}

	// reddSimplifiedChineseName represents redd simplified chinese name.
	reddSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "狐利"}

	// reddSpanishName represents redd spanish name.
	reddSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ladino"}

	// reddTraditionalChineseName represents redd traditional chinese name.
	reddTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "狐利"}
)

var (
	// reddName represents redd name.
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
	// reddCharacter represents redd character.
	reddCharacter = nook.Character{
		Animal:   animal.Fox,
		Birthday: reddBirthday,
		Code:     reddCode,
		Key:      character.Redd,
		Gender:   gender.Male,
		Name:     reddName,
		Special:  true}
)

var (
	// Redd represents redd.
	Redd = nook.Resident{
		Character: reddCharacter}
)
