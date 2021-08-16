package gyroid

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	lloidBirthday = nook.Birthday{
		Day:   28,
		Month: time.August}
)

var (
	lloidCode = nook.Code{
		Value: "hnw"}
)

var (
	lloidAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lloid"}

	lloidCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gyroïde"}

	lloidDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lloid"}

	lloidFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gyroïde"}

	lloidGermanName = nook.Name{
		Language: language.German,
		Value:    "Gyroid"}

	lloidItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gironio"}

	lloidJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハニワくん"}

	lloidKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "토용군"}

	lloidLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Giroide"}

	lloidRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Ллойд"}

	lloidSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "俑俑"}

	lloidSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Giroide"}

	lloidTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "俑俑"}
)

var (
	lloidName = nook.Languages{
		language.AmericanEnglish:      lloidAmericanEnglishName,
		language.CanadianFrench:       lloidCanadianFrenchName,
		language.Dutch:                lloidDutchName,
		language.French:               lloidFrenchName,
		language.German:               lloidGermanName,
		language.Italian:              lloidItalianName,
		language.Japanese:             lloidJapaneseName,
		language.Korean:               lloidKoreanName,
		language.LatinAmericanSpanish: lloidLatinAmericanSpanishName,
		language.Russian:              lloidRussianName,
		language.SimplifiedChinese:    lloidSimplifiedChineseName,
		language.Spanish:              lloidSpanishName,
		language.TraditionalChinese:   lloidTraditionalChineseName}
)

var (
	lloidCharacter = nook.Character{
		Animal:   animal.Gyroid,
		Birthday: lloidBirthday,
		Code:     lloidCode,
		Gender:   gender.Male,
		Name:     lloidName}
)

var (
	Lloid = nook.Resident{
		Character: lloidCharacter}
)
