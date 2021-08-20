package pelican

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	pellyBirthday = nook.Birthday{
		Day:   19,
		Month: time.March}
)

var (
	pellyCode = nook.Code{
		Value: "pga/plk"}
)

var (
	pellyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pelly"}

	pellyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Opélie"}

	pellyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pelly"}

	pellyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Opélie"}

	pellyGermanName = nook.Name{
		Language: language.German,
		Value:    "Pelly"}

	pellyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pelly"}

	pellyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ぺりこ"}

	pellyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "펠리"}

	pellyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sol"}

	pellyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пелли"}

	pellySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "宋信子"}

	pellySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sol"}

	pellyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "宋信子"}
)

var (
	pellyName = nook.Languages{
		language.AmericanEnglish:      pellyAmericanEnglishName,
		language.CanadianFrench:       pellyCanadianFrenchName,
		language.Dutch:                pellyDutchName,
		language.French:               pellyFrenchName,
		language.German:               pellyGermanName,
		language.Italian:              pellyItalianName,
		language.Japanese:             pellyJapaneseName,
		language.Korean:               pellyKoreanName,
		language.LatinAmericanSpanish: pellyLatinAmericanSpanishName,
		language.Russian:              pellyRussianName,
		language.SimplifiedChinese:    pellySimplifiedChineseName,
		language.Spanish:              pellySpanishName,
		language.TraditionalChinese:   pellyTraditionalChineseName}
)

var (
	pellyCharacter = nook.Character{
		Animal:   animal.Pelican,
		Birthday: pellyBirthday,
		Code:     pellyCode,
		Key:      character.Pelly,
		Gender:   gender.Female,
		Name:     pellyName,
		Special:  true}
)

var (
	Pelly = nook.Resident{
		Character: pellyCharacter}
)
