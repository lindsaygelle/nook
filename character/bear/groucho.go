package bear

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
	grouchoBirthday = nook.Birthday{
		Day:   23,
		Month: time.October}
)

var (
	grouchoCode = nook.Code{
		Value: "bea06"}
)

var (
	grouchoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Groucho"}

	grouchoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ronchon"}

	grouchoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Groucho"}

	grouchoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ronchon"}

	grouchoGermanName = nook.Name{
		Language: language.German,
		Value:    "Ernst"}

	grouchoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Groucho"}

	grouchoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クロー"}

	grouchoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "거무틱"}

	grouchoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Groucho"}

	grouchoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Граучо"}

	grouchoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "爪爪"}

	grouchoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Groucho"}

	grouchoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "爪爪"}
)

var (
	grouchoName = nook.Languages{
		language.AmericanEnglish:      grouchoAmericanEnglishName,
		language.CanadianFrench:       grouchoCanadianFrenchName,
		language.Dutch:                grouchoDutchName,
		language.French:               grouchoFrenchName,
		language.German:               grouchoGermanName,
		language.Italian:              grouchoItalianName,
		language.Japanese:             grouchoJapaneseName,
		language.Korean:               grouchoKoreanName,
		language.LatinAmericanSpanish: grouchoLatinAmericanSpanishName,
		language.Russian:              grouchoRussianName,
		language.SimplifiedChinese:    grouchoSimplifiedChineseName,
		language.Spanish:              grouchoSpanishName,
		language.TraditionalChinese:   grouchoTraditionalChineseName}
)

var (
	grouchoCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: grouchoBirthday,
		Code:     grouchoCode,
		Key:      character.Groucho,
		Gender:   gender.Male,
		Name:     grouchoName}
)

var (
	grouchoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "grumble"}

	grouchoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grou-gnon"}

	grouchoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "brom"}

	grouchoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grou-gnon"}

	grouchoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brummel"}

	grouchoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grum grum"}

	grouchoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ワォー"}

	grouchoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어쩌라구"}

	grouchoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "jorrr"}

	grouchoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бур-бур"}

	grouchoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哇哇"}

	grouchoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "jorrr"}

	grouchoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哇哇"}
)

var (
	grouchoPhrase = nook.Languages{
		language.AmericanEnglish:      grouchoAmericanEnglishName,
		language.CanadianFrench:       grouchoCanadianFrenchName,
		language.Dutch:                grouchoDutchName,
		language.French:               grouchoFrenchName,
		language.German:               grouchoGermanName,
		language.Italian:              grouchoItalianName,
		language.Japanese:             grouchoJapaneseName,
		language.Korean:               grouchoKoreanName,
		language.LatinAmericanSpanish: grouchoLatinAmericanSpanishName,
		language.Russian:              grouchoRussianName,
		language.SimplifiedChinese:    grouchoSimplifiedChineseName,
		language.Spanish:              grouchoSpanishName,
		language.TraditionalChinese:   grouchoTraditionalChineseName}
)

var (
	Groucho = nook.Villager{
		Character:   grouchoCharacter,
		Personality: personality.Cranky,
		Phrase:      grouchoPhrase}
)
