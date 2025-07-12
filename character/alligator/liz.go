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
	// lizBirthday represents liz birthday.
	lizBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

// lizCode represents Liz's unique code.
var (
	// lizCode represents liz code.
	lizCode = nook.Code{
		Value: ""}
)

// Different names for Liz in various languages.
var (
	// lizAmericanEnglishName represents liz american english name.
	lizAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Liz"}

	// lizCanadianFrenchName represents liz canadian french name.
	lizCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lisette"}

	// lizDutchName represents liz dutch name.
	lizDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// lizFrenchName represents liz french name.
	lizFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lisette"}

	// lizGermanName represents liz german name.
	lizGermanName = nook.Name{
		Language: language.German,
		Value:    "Ruth"}

	// lizItalianName represents liz italian name.
	lizItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Drilla"}

	// lizJapaneseName represents liz japanese name.
	lizJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ストロベリー"}

	// lizKoreanName represents liz korean name.
	lizKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// lizLatinAmericanSpanishName represents liz latin american spanish name.
	lizLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Natalia"}

	// lizRussianName represents liz russian name.
	lizRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// lizSimplifiedChineseName represents liz simplified chinese name.
	lizSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "贝蒂"}

	// lizSpanishName represents liz spanish name.
	lizSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Natalia"}

	// lizTraditionalChineseName represents liz traditional chinese name.
	lizTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

// lizName represents Liz's name in different languages.
var (
	// lizName represents liz name.
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
	// lizCharacter represents liz character.
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
	// lizAmericanEnglishPhrase represents liz american english phrase.
	lizAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "groonch"}

	// lizCanadianFrenchPhrase represents liz canadian french phrase.
	lizCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "groonch"}

	// lizDutchPhrase represents liz dutch phrase.
	lizDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// lizFrenchPhrase represents liz french phrase.
	lizFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "groonch"}

	// lizGermanPhrase represents liz german phrase.
	lizGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grrruuuua"}

	// lizItalianPhrase represents liz italian phrase.
	lizItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grunch"}

	// lizJapanesePhrase represents liz japanese phrase.
	lizJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なの"}

	// lizKoreanPhrase represents liz korean phrase.
	lizKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// lizLatinAmericanSpanishPhrase represents liz latin american spanish phrase.
	lizLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñec-ñec"}

	// lizRussianPhrase represents liz russian phrase.
	lizRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// lizSimplifiedChinesePhrase represents liz simplified chinese phrase.
	lizSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嗯呐"}

	// lizSpanishPhrase represents liz spanish phrase.
	lizSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ñec-ñec"}

	// lizTraditionalChinesePhrase represents liz traditional chinese phrase.
	lizTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

// lizPhrase represents Liz's phrases in different languages.
var (
	// lizPhrase represents liz phrase.
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
	// Liz represents liz.
	Liz = nook.Villager{
		Character:   lizCharacter,
		Personality: personality.Normal,
		Phrase:      lizPhrase}
)
