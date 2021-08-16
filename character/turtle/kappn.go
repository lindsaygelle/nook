package turtle

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	kappnBirthday = nook.Birthday{
		Day:   12,
		Month: time.July}
)

var (
	kappnCode = nook.Code{
		Value: "wip/kpp"}
)

var (
	kappnAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kapp'n"}

	kappnCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Amiral"}

	kappnDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kapp'n"}

	kappnFrenchName = nook.Name{
		Language: language.French,
		Value:    "Amiral"}

	kappnGermanName = nook.Name{
		Language: language.German,
		Value:    "Käpten"}

	kappnItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Remo"}

	kappnJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "かっぺい"}

	kappnKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "갑돌"}

	kappnLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Capitán"}

	kappnRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кэпн"}

	kappnSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "航平"}

	kappnSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Capitán"}

	kappnTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "航平"}
)

var (
	kappnName = nook.Languages{
		language.AmericanEnglish:      kappnAmericanEnglishName,
		language.CanadianFrench:       kappnCanadianFrenchName,
		language.Dutch:                kappnDutchName,
		language.French:               kappnFrenchName,
		language.German:               kappnGermanName,
		language.Italian:              kappnItalianName,
		language.Japanese:             kappnJapaneseName,
		language.Korean:               kappnKoreanName,
		language.LatinAmericanSpanish: kappnLatinAmericanSpanishName,
		language.Russian:              kappnRussianName,
		language.SimplifiedChinese:    kappnSimplifiedChineseName,
		language.Spanish:              kappnSpanishName,
		language.TraditionalChinese:   kappnTraditionalChineseName}
)

var (
	kappnCharacter = nook.Character{
		Animal:   animal.Turtle,
		Birthday: kappnBirthday,
		Code:     kappnCode,
		Key:      character.Kappn,
		Gender:   gender.Male,
		Name:     kappnName}
)

var (
	Kappn = nook.Resident{
		Character: kappnCharacter}
)
