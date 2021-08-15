package alligator

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	lizBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	lizCode = nook.Code{
		Value: ""}
)

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

var (
	lizCharacter = nook.Character{
		Animal:   Alligator,
		Birthday: lizBirthday,
		Code:     lizCode,
		Gender:   nook.Female,
		Name:     lizName}
)

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

var (
	lizPhrase = nook.Languages{
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

var (
	Liz = nook.Villager{
		Character:   lizCharacter,
		Personality: nook.Normal,
		Phrase:      lizPhrase}
)
