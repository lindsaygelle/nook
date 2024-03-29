package anteater

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// luluBirthday represents Lulu's birthday.
var (
	luluBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

// luluCode represents Lulu's unique code.
var (
	luluCode = nook.Code{
		Value: ""}
)

// Different names for Lulu in various languages.
var (
	luluAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lulu"}

	luluCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	luluDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	luluFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	luluGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	luluItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	luluJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ルル"}

	luluKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	luluLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	luluRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	luluSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	luluSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	luluTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

// luluName represents Lulu's name in different languages.
var (
	luluName = nook.Languages{
		language.AmericanEnglish:      luluAmericanEnglishName,
		language.CanadianFrench:       luluCanadianFrenchName,
		language.Dutch:                luluDutchName,
		language.French:               luluFrenchName,
		language.German:               luluGermanName,
		language.Italian:              luluItalianName,
		language.Japanese:             luluJapaneseName,
		language.Korean:               luluKoreanName,
		language.LatinAmericanSpanish: luluLatinAmericanSpanishName,
		language.Russian:              luluRussianName,
		language.SimplifiedChinese:    luluSimplifiedChineseName,
		language.Spanish:              luluSpanishName,
		language.TraditionalChinese:   luluTraditionalChineseName}
)

// luluCharacter represents Lulu's character information.
var (
	luluCharacter = nook.Character{
		Animal:   animal.Anteater,
		Birthday: luluBirthday,
		Code:     luluCode,
		Key:      character.Lulu,
		Gender:   gender.Female,
		Name:     luluName,
		Special:  false}
)

// Different phrases spoken by Lulu in various languages.
var (
	luluAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "あらあら"}

	luluCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	luluDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	luluFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	luluGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	luluItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	luluJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "あらあら"}

	luluKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	luluLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	luluRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	luluSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	luluSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	luluTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

// luluPhrase represents Lulu's phrases in different languages.
var (
	luluPhrase = nook.Languages{
		language.AmericanEnglish:      luluAmericanEnglishPhrase,
		language.CanadianFrench:       luluCanadianFrenchPhrase,
		language.Dutch:                luluDutchPhrase,
		language.French:               luluFrenchPhrase,
		language.German:               luluGermanPhrase,
		language.Italian:              luluItalianPhrase,
		language.Japanese:             luluJapanesePhrase,
		language.Korean:               luluKoreanPhrase,
		language.LatinAmericanSpanish: luluLatinAmericanSpanishPhrase,
		language.Russian:              luluRussianPhrase,
		language.SimplifiedChinese:    luluSimplifiedChinesePhrase,
		language.Spanish:              luluSpanishPhrase,
		language.TraditionalChinese:   luluTraditionalChinesePhrase}
)

// Lulu represents the character Lulu with her complete information.
var (
	Lulu = nook.Villager{
		Character:   luluCharacter,
		Personality: personality.Snooty,
		Phrase:      luluPhrase}
)
