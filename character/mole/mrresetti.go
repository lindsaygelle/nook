package mole

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	mrresettiBirthday = nook.Birthday{
		Day:   6,
		Month: time.April}
)

var (
	mrresettiCode = nook.Code{
		Value: "mol"}
)

var (
	mrresettiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mr. Resetti"}

	mrresettiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Resetti"}

	mrresettiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Resetti"}

	mrresettiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Resetti"}

	mrresettiGermanName = nook.Name{
		Language: language.German,
		Value:    "Resetti"}

	mrresettiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Resetti"}

	mrresettiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リセットさん"}

	mrresettiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "도루묵씨"}

	mrresettiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rese T."}

	mrresettiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Ресетти"}

	mrresettiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "电源叔叔"}

	mrresettiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rese T."}

	mrresettiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "電源叔叔"}
)

var (
	mrresettiName = nook.Languages{
		language.AmericanEnglish:      mrresettiAmericanEnglishName,
		language.CanadianFrench:       mrresettiCanadianFrenchName,
		language.Dutch:                mrresettiDutchName,
		language.French:               mrresettiFrenchName,
		language.German:               mrresettiGermanName,
		language.Italian:              mrresettiItalianName,
		language.Japanese:             mrresettiJapaneseName,
		language.Korean:               mrresettiKoreanName,
		language.LatinAmericanSpanish: mrresettiLatinAmericanSpanishName,
		language.Russian:              mrresettiRussianName,
		language.SimplifiedChinese:    mrresettiSimplifiedChineseName,
		language.Spanish:              mrresettiSpanishName,
		language.TraditionalChinese:   mrresettiTraditionalChineseName}
)

var (
	mrresettiCharacter = nook.Character{
		Animal:   animal.Mole,
		Birthday: mrresettiBirthday,
		Code:     mrresettiCode,
		Key:      character.MrResetti,
		Gender:   gender.Male,
		Name:     mrresettiName,
		Special:  true}
)

var (
	MrResetti = nook.Resident{
		Character: mrresettiCharacter}
)
