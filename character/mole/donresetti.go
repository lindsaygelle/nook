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
	// donresettiBirthday represents donresetti birthday.
	donresettiBirthday = nook.Birthday{
		Day:   1,
		Month: time.May}
)

var (
	// donresettiCode represents donresetti code.
	donresettiCode = nook.Code{
		Value: "mob"}
)

var (
	// donresettiAmericanEnglishName represents donresetti american english name.
	donresettiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Don Resetti"}

	// donresettiCanadianFrenchName represents donresetti canadian french name.
	donresettiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Don"}

	// donresettiDutchName represents donresetti dutch name.
	donresettiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Don"}

	// donresettiFrenchName represents donresetti french name.
	donresettiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Don"}

	// donresettiGermanName represents donresetti german name.
	donresettiGermanName = nook.Name{
		Language: language.German,
		Value:    "Don"}

	// donresettiItalianName represents donresetti italian name.
	donresettiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Don"}

	// donresettiJapaneseName represents donresetti japanese name.
	donresettiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ラケットさん"}

	// donresettiKoreanName represents donresetti korean name.
	donresettiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "오루묵씨"}

	// donresettiLatinAmericanSpanishName represents donresetti latin american spanish name.
	donresettiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Forma T."}

	// donresettiRussianName represents donresetti russian name.
	donresettiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дон"}

	// donresettiSimplifiedChineseName represents donresetti simplified chinese name.
	donresettiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "开关叔叔"}

	// donresettiSpanishName represents donresetti spanish name.
	donresettiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Forma T."}

	// donresettiTraditionalChineseName represents donresetti traditional chinese name.
	donresettiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "開關叔叔"}
)

var (
	// donresettiName represents donresetti name.
	donresettiName = nook.Languages{
		language.AmericanEnglish:      donresettiAmericanEnglishName,
		language.CanadianFrench:       donresettiCanadianFrenchName,
		language.Dutch:                donresettiDutchName,
		language.French:               donresettiFrenchName,
		language.German:               donresettiGermanName,
		language.Italian:              donresettiItalianName,
		language.Japanese:             donresettiJapaneseName,
		language.Korean:               donresettiKoreanName,
		language.LatinAmericanSpanish: donresettiLatinAmericanSpanishName,
		language.Russian:              donresettiRussianName,
		language.SimplifiedChinese:    donresettiSimplifiedChineseName,
		language.Spanish:              donresettiSpanishName,
		language.TraditionalChinese:   donresettiTraditionalChineseName}
)

var (
	// donresettiCharacter represents donresetti character.
	donresettiCharacter = nook.Character{
		Animal:   animal.Mole,
		Birthday: donresettiBirthday,
		Code:     donresettiCode,
		Key:      character.DonResetti,
		Gender:   gender.Male,
		Name:     donresettiName,
		Special:  true}
)

var (
	// DonResetti represents don resetti.
	DonResetti = nook.Resident{
		Character: donresettiCharacter}
)
