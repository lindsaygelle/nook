package dodo

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	orvilleBirthday = nook.Birthday{
		Day:   2,
		Month: time.October}
)

var (
	orvilleCode = nook.Code{
		Value: "dod"}
)

var (
	orvilleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Orville"}

	orvilleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Morris"}

	orvilleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Orville"}

	orvilleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Morris"}

	orvilleGermanName = nook.Name{
		Language: language.German,
		Value:    "Bodo"}

	orvilleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Dodobaldo"}

	orvilleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "モーリー"}

	orvilleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "모리"}

	orvilleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rafa"}

	orvilleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Орвилл"}

	orvilleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莫里"}

	orvilleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rafa"}

	orvilleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莫里"}
)

var (
	orvilleName = nook.Languages{
		language.AmericanEnglish:      orvilleAmericanEnglishName,
		language.CanadianFrench:       orvilleCanadianFrenchName,
		language.Dutch:                orvilleDutchName,
		language.French:               orvilleFrenchName,
		language.German:               orvilleGermanName,
		language.Italian:              orvilleItalianName,
		language.Japanese:             orvilleJapaneseName,
		language.Korean:               orvilleKoreanName,
		language.LatinAmericanSpanish: orvilleLatinAmericanSpanishName,
		language.Russian:              orvilleRussianName,
		language.SimplifiedChinese:    orvilleSimplifiedChineseName,
		language.Spanish:              orvilleSpanishName,
		language.TraditionalChinese:   orvilleTraditionalChineseName}
)

var (
	orvilleCharacter = nook.Character{
		Animal:   animal.Dodo,
		Birthday: orvilleBirthday,
		Code:     orvilleCode,
		Gender:   gender.Male,
		Name:     orvilleName}
)

var (
	Orville = nook.Resident{
		Character: orvilleCharacter}
)
