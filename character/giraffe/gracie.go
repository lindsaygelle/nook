package giraffe

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	gracieBirthday = nook.Birthday{
		Day:   14,
		Month: time.November}
)

var (
	gracieCode = nook.Code{
		Value: "grf"}
)

var (
	gracieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gracie"}

	gracieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Carla"}

	gracieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gracie"}

	gracieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Carla"}

	gracieGermanName = nook.Name{
		Language: language.German,
		Value:    "Grazia"}

	gracieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Griffa"}

	gracieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "グレース"}

	gracieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "그레이스"}

	gracieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Graciela"}

	gracieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Грейси"}

	gracieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "葛瑞斯"}

	gracieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Graciela"}

	gracieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "葛瑞斯"}
)

var (
	gracieName = nook.Languages{
		language.AmericanEnglish:      gracieAmericanEnglishName,
		language.CanadianFrench:       gracieCanadianFrenchName,
		language.Dutch:                gracieDutchName,
		language.French:               gracieFrenchName,
		language.German:               gracieGermanName,
		language.Italian:              gracieItalianName,
		language.Japanese:             gracieJapaneseName,
		language.Korean:               gracieKoreanName,
		language.LatinAmericanSpanish: gracieLatinAmericanSpanishName,
		language.Russian:              gracieRussianName,
		language.SimplifiedChinese:    gracieSimplifiedChineseName,
		language.Spanish:              gracieSpanishName,
		language.TraditionalChinese:   gracieTraditionalChineseName}
)

var (
	gracieCharacter = nook.Character{
		Animal:   animal.Giraffe,
		Birthday: gracieBirthday,
		Code:     gracieCode,
		Key:      character.Gracie,
		Gender:   gender.Female,
		Name:     gracieName,
		Special:  true}
)

var (
	Gracie = nook.Resident{
		Character: gracieCharacter}
)
