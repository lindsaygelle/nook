package hippo

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	// luluBirthday represents lulu birthday.
	luluBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// luluCode represents lulu code.
	luluCode = nook.Code{
		Value: ""}
)

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
		Value:    "Lulu"}

	// luluGermanName represents lulu german name.
	luluGermanName = nook.Name{
		Language: language.German,
		Value:    "Lulu"}

	// luluItalianName represents lulu italian name.
	luluItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lulù"}

	// luluJapaneseName represents lulu japanese name.
	luluJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "エルニーニョ"}

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
		Value:    "萍萍"}

	// luluSpanishName represents lulu spanish name.
	luluSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lulù"}

	// luluTraditionalChineseName represents lulu traditional chinese name.
	luluTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

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

var (
	// luluCharacter represents lulu character.
	luluCharacter = nook.Character{
		Animal:   animal.Hippo,
		Birthday: luluBirthday,
		Code:     luluCode,
		Key:      character.Lulu,
		Gender:   gender.Female,
		Name:     luluName,
		Special:  false}
)

var (
	// luluAmericanEnglishPhrase represents lulu american english phrase.
	luluAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "yaaaawl"}

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
		Value:    "mon bichon"}

	// luluGermanPhrase represents lulu german phrase.
	luluGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "gääähn"}

	// luluItalianPhrase represents lulu italian phrase.
	luluItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hippurrà"}

	// luluJapanesePhrase represents lulu japanese phrase.
	luluJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ニョニョ"}

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
		Value:    "妞妞"}

	// luluSpanishPhrase represents lulu spanish phrase.
	luluSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "lilalá"}

	// luluTraditionalChinesePhrase represents lulu traditional chinese phrase.
	luluTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

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

var (
	// Lulu represents lulu.
	Lulu = nook.Villager{
		Character:   luluCharacter,
		Personality: personality.Peppy,
		Phrase:      luluPhrase}
)
