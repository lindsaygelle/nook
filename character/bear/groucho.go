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

// grouchoBirthday represents Groucho's birthday.
var (
	// grouchoBirthday represents groucho birthday.
	grouchoBirthday = nook.Birthday{
		Day:   23,
		Month: time.October}
)

// grouchoCode represents Groucho's unique code.
var (
	// grouchoCode represents groucho code.
	grouchoCode = nook.Code{
		Value: "bea06"}
)

// Different names for Groucho in various languages.
var (
	// grouchoAmericanEnglishName represents groucho american english name.
	grouchoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Groucho"}

	// grouchoCanadianFrenchName represents groucho canadian french name.
	grouchoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ronchon"}

	// grouchoDutchName represents groucho dutch name.
	grouchoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Groucho"}

	// grouchoFrenchName represents groucho french name.
	grouchoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ronchon"}

	// grouchoGermanName represents groucho german name.
	grouchoGermanName = nook.Name{
		Language: language.German,
		Value:    "Ernst"}

	// grouchoItalianName represents groucho italian name.
	grouchoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Groucho"}

	// grouchoJapaneseName represents groucho japanese name.
	grouchoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クロー"}

	// grouchoKoreanName represents groucho korean name.
	grouchoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "거무틱"}

	// grouchoLatinAmericanSpanishName represents groucho latin american spanish name.
	grouchoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Groucho"}

	// grouchoRussianName represents groucho russian name.
	grouchoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Граучо"}

	// grouchoSimplifiedChineseName represents groucho simplified chinese name.
	grouchoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "爪爪"}

	// grouchoSpanishName represents groucho spanish name.
	grouchoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Groucho"}

	// grouchoTraditionalChineseName represents groucho traditional chinese name.
	grouchoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "爪爪"}
)

// grouchoName represents Groucho's name in different languages.
var (
	// grouchoName represents groucho name.
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

// grouchoCharacter represents Groucho's character information.
var (
	// grouchoCharacter represents groucho character.
	grouchoCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: grouchoBirthday,
		Code:     grouchoCode,
		Key:      character.Groucho,
		Gender:   gender.Male,
		Name:     grouchoName,
		Special:  false}
)

// Different phrases spoken by Groucho in various languages.
var (
	// grouchoAmericanEnglishPhrase represents groucho american english phrase.
	grouchoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "grumble"}

	// grouchoCanadianFrenchPhrase represents groucho canadian french phrase.
	grouchoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grou-gnon"}

	// grouchoDutchPhrase represents groucho dutch phrase.
	grouchoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "brom"}

	// grouchoFrenchPhrase represents groucho french phrase.
	grouchoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grou-gnon"}

	// grouchoGermanPhrase represents groucho german phrase.
	grouchoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brummel"}

	// grouchoItalianPhrase represents groucho italian phrase.
	grouchoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grum grum"}

	// grouchoJapanesePhrase represents groucho japanese phrase.
	grouchoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ワォー"}

	// grouchoKoreanPhrase represents groucho korean phrase.
	grouchoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어쩌라구"}

	// grouchoLatinAmericanSpanishPhrase represents groucho latin american spanish phrase.
	grouchoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "jorrr"}

	// grouchoRussianPhrase represents groucho russian phrase.
	grouchoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бур-бур"}

	// grouchoSimplifiedChinesePhrase represents groucho simplified chinese phrase.
	grouchoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哇哇"}

	// grouchoSpanishPhrase represents groucho spanish phrase.
	grouchoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "jorrr"}

	// grouchoTraditionalChinesePhrase represents groucho traditional chinese phrase.
	grouchoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哇哇"}
)

// grouchoPhrase represents Groucho's phrases in different languages.
var (
	// grouchoPhrase represents groucho phrase.
	grouchoPhrase = nook.Languages{
		language.AmericanEnglish:      grouchoAmericanEnglishPhrase,
		language.CanadianFrench:       grouchoCanadianFrenchPhrase,
		language.Dutch:                grouchoDutchPhrase,
		language.French:               grouchoFrenchPhrase,
		language.German:               grouchoGermanPhrase,
		language.Italian:              grouchoItalianPhrase,
		language.Japanese:             grouchoJapanesePhrase,
		language.Korean:               grouchoKoreanPhrase,
		language.LatinAmericanSpanish: grouchoLatinAmericanSpanishPhrase,
		language.Russian:              grouchoRussianPhrase,
		language.SimplifiedChinese:    grouchoSimplifiedChinesePhrase,
		language.Spanish:              grouchoSpanishPhrase,
		language.TraditionalChinese:   grouchoTraditionalChinesePhrase}
)

// Groucho represents the character Groucho with his complete information.
var (
	// Groucho represents groucho.
	Groucho = nook.Villager{
		Character:   grouchoCharacter,
		Personality: personality.Cranky,
		Phrase:      grouchoPhrase}
)
