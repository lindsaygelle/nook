package dodo

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	wilburBirthday = nook.Birthday{
		Day:   17,
		Month: time.December}
)

var (
	wilburCode = nook.Code{
		Value: "doc"}
)

var (
	wilburAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Wilbur"}

	wilburCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rodrigue"}

	wilburDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Wilbur"}

	wilburFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rodrigue"}

	wilburGermanName = nook.Name{
		Language: language.German,
		Value:    "Udo"}

	wilburItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Dodoardo"}

	wilburJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ロドリー"}

	wilburKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "로드리"}

	wilburLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rodri"}

	wilburRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Уилбур"}

	wilburSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "陆德里"}

	wilburSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rodri"}

	wilburTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "陸德里"}
)

var (
	wilburName = nook.Languages{
		language.AmericanEnglish:      wilburAmericanEnglishName,
		language.CanadianFrench:       wilburCanadianFrenchName,
		language.Dutch:                wilburDutchName,
		language.French:               wilburFrenchName,
		language.German:               wilburGermanName,
		language.Italian:              wilburItalianName,
		language.Japanese:             wilburJapaneseName,
		language.Korean:               wilburKoreanName,
		language.LatinAmericanSpanish: wilburLatinAmericanSpanishName,
		language.Russian:              wilburRussianName,
		language.SimplifiedChinese:    wilburSimplifiedChineseName,
		language.Spanish:              wilburSpanishName,
		language.TraditionalChinese:   wilburTraditionalChineseName}
)

var (
	wilburCharacter = nook.Character{
		Animal:   animal.Dodo,
		Birthday: wilburBirthday,
		Code:     wilburCode,
		Gender:   gender.Male,
		Name:     wilburName}
)

var (
	Wilbur = nook.Resident{
		Character: wilburCharacter}
)
