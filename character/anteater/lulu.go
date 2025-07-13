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
	// luluBirthday represents lulu birthday.
	luluBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

// luluCode represents Lulu's unique code.
var (
	// luluCode represents lulu code.
	luluCode = nook.Code{
		Value: ""}
)

// Different names for Lulu in various languages.
var (
	// luluAmericanEnglishName represents lulu american english name.
	luluAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lulu"}

	// luluCanadianFrenchName represents lulu canadian french name.
	luluCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// luluDutchName represents lulu dutch name.
	luluDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// luluFrenchName represents lulu french name.
	luluFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// luluGermanName represents lulu german name.
	luluGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// luluItalianName represents lulu italian name.
	luluItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// luluJapaneseName represents lulu japanese name.
	luluJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ルル"}

	// luluKoreanName represents lulu korean name.
	luluKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// luluLatinAmericanSpanishName represents lulu latin american spanish name.
	luluLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// luluRussianName represents lulu russian name.
	luluRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// luluSimplifiedChineseName represents lulu simplified chinese name.
	luluSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// luluSpanishName represents lulu spanish name.
	luluSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// luluTraditionalChineseName represents lulu traditional chinese name.
	luluTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

// luluName represents Lulu's name in different languages.
var (
	// luluName represents lulu name.
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
	// luluCharacter represents lulu character.
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
	// luluAmericanEnglishPhrase represents lulu american english phrase.
	luluAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "あらあら"}

	// luluCanadianFrenchPhrase represents lulu canadian french phrase.
	luluCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// luluDutchPhrase represents lulu dutch phrase.
	luluDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// luluFrenchPhrase represents lulu french phrase.
	luluFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// luluGermanPhrase represents lulu german phrase.
	luluGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// luluItalianPhrase represents lulu italian phrase.
	luluItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// luluJapanesePhrase represents lulu japanese phrase.
	luluJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "あらあら"}

	// luluKoreanPhrase represents lulu korean phrase.
	luluKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// luluLatinAmericanSpanishPhrase represents lulu latin american spanish phrase.
	luluLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// luluRussianPhrase represents lulu russian phrase.
	luluRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// luluSimplifiedChinesePhrase represents lulu simplified chinese phrase.
	luluSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// luluSpanishPhrase represents lulu spanish phrase.
	luluSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// luluTraditionalChinesePhrase represents lulu traditional chinese phrase.
	luluTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

// luluPhrase represents Lulu's phrases in different languages.
var (
	// luluPhrase represents lulu phrase.
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
	// Lulu represents lulu.
	Lulu = nook.Villager{
		Character:   luluCharacter,
		Personality: personality.Snooty,
		Phrase:      luluPhrase}
)
