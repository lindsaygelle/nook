package alligator

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// lizBirthday represents Liz's birthday.
var (
	lizBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

// lizCode represents Liz's unique code.
var (
	lizCode = nook.Code{
		Value: ""}
)

// Different names for Liz in various languages.
var (
	lizAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Liz"}

	lizCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lisette"}

	lizDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	lizFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lisette"}

	lizGermanName = nook.Name{
		Language: language.German,
		Value:    "Ruth"}

	lizItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Drilla"}

	lizJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ストロベリー"}

	lizKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	lizLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Natalia"}

	lizRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	lizSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "贝蒂"}

	lizSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Natalia"}

	lizTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

// lizName represents Liz's name in different languages.
var (
	lizName = nook.Languages{
		language.AmericanEnglish:      lizAmericanEnglishName,
		language.CanadianFrench:       lizCanadianFrenchName,
		language.Dutch:                lizDutchName,
		language.French:               lizFrenchName,
		language.German:               lizGermanName,
		language.Italian:              lizItalianName,
		language.Japanese:             lizJapaneseName,
		language.Korean:               lizKoreanName,
		language.LatinAmericanSpanish: lizLatinAmericanSpanishName,
		language.Russian:              lizRussianName,
		language.SimplifiedChinese:    lizSimplifiedChineseName,
		language.Spanish:              lizSpanishName,
		language.TraditionalChinese:   lizTraditionalChineseName}
)

// lizCharacter represents Liz's character information.
var (
	lizCharacter = nook.Character{
		Animal:   animal.Alligator,
		Birthday: lizBirthday,
		Code:     lizCode,
		Key:      character.Liz,
		Gender:   gender.Female,
		Name:     lizName,
		Special:  false}
)

// Different phrases spoken by Liz in various languages.
var (
	lizAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "groonch"}

	lizCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "groonch"}

	lizDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	lizFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "groonch"}

	lizGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grrruuuua"}

	lizItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grunch"}

	lizJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なの"}

	lizKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	lizLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñec-ñec"}

	lizRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	lizSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嗯呐"}

	lizSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ñec-ñec"}

	lizTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

// lizPhrase represents Liz's phrases in different languages.
var (
	lizPhrase = nook.Languages{
		language.AmericanEnglish:      lizAmericanEnglishPhrase,
		language.CanadianFrench:       lizCanadianFrenchPhrase,
		language.Dutch:                lizDutchPhrase,
		language.French:               lizFrenchPhrase,
		language.German:               lizGermanPhrase,
		language.Italian:              lizItalianPhrase,
		language.Japanese:             lizJapanesePhrase,
		language.Korean:               lizKoreanPhrase,
		language.LatinAmericanSpanish: lizLatinAmericanSpanishPhrase,
		language.Russian:              lizRussianPhrase,
		language.SimplifiedChinese:    lizSimplifiedChinesePhrase,
		language.Spanish:              lizSpanishPhrase,
		language.TraditionalChinese:   lizTraditionalChinesePhrase}
)

// Liz represents the character Liz with her complete information.
var (
	Liz = nook.Villager{
		Character:   lizCharacter,
		Personality: personality.Normal,
		Phrase:      lizPhrase}
)
