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
	// mrresettiBirthday represents mrresetti birthday.
	mrresettiBirthday = nook.Birthday{
		Day:   6,
		Month: time.April}
)

var (
	// mrresettiCode represents mrresetti code.
	mrresettiCode = nook.Code{
		Value: "mol"}
)

var (
	// mrresettiAmericanEnglishName represents mrresetti american english name.
	mrresettiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mr. Resetti"}

	// mrresettiCanadianFrenchName represents mrresetti canadian french name.
	mrresettiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Resetti"}

	// mrresettiDutchName represents mrresetti dutch name.
	mrresettiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Resetti"}

	// mrresettiFrenchName represents mrresetti french name.
	mrresettiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Resetti"}

	// mrresettiGermanName represents mrresetti german name.
	mrresettiGermanName = nook.Name{
		Language: language.German,
		Value:    "Resetti"}

	// mrresettiItalianName represents mrresetti italian name.
	mrresettiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Resetti"}

	// mrresettiJapaneseName represents mrresetti japanese name.
	mrresettiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リセットさん"}

	// mrresettiKoreanName represents mrresetti korean name.
	mrresettiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "도루묵씨"}

	// mrresettiLatinAmericanSpanishName represents mrresetti latin american spanish name.
	mrresettiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rese T."}

	// mrresettiRussianName represents mrresetti russian name.
	mrresettiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Ресетти"}

	// mrresettiSimplifiedChineseName represents mrresetti simplified chinese name.
	mrresettiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "电源叔叔"}

	// mrresettiSpanishName represents mrresetti spanish name.
	mrresettiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rese T."}

	// mrresettiTraditionalChineseName represents mrresetti traditional chinese name.
	mrresettiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "電源叔叔"}
)

var (
	// mrresettiName represents mrresetti name.
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
	// mrresettiCharacter represents mrresetti character.
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
	// MrResetti represents mr resetti.
	MrResetti = nook.Resident{
		Character: mrresettiCharacter}
)
