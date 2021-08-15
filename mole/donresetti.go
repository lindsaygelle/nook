package mole

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	donresettiBirthday = nook.Birthday{
		Day:   1,
		Month: time.May}
)

var (
	donresettiCode = nook.Code{
		Value: "mob"}
)

var (
	donresettiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Don Resetti"}

	donresettiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Don"}

	donresettiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Don"}

	donresettiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Don"}

	donresettiGermanName = nook.Name{
		Language: language.German,
		Value:    "Don"}

	donresettiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Don"}

	donresettiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ラケットさん"}

	donresettiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "오루묵씨"}

	donresettiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Forma T."}

	donresettiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дон"}

	donresettiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "开关叔叔"}

	donresettiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Forma T."}

	donresettiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "開關叔叔"}
)

var (
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
	donresettiCharacter = nook.Character{
		Animal:   Mole,
		Birthday: donresettiBirthday,
		Code:     donresettiCode,
		Gender:   nook.Male,
		Name:     donresettiName}
)

var (
	DonResetti = nook.Resident{
		Character: donresettiCharacter}
)
