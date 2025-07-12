package bearcub

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
	// bluebearBirthday represents bluebear birthday.
	bluebearBirthday = nook.Birthday{
		Day:   24,
		Month: time.June}
)

var (
	// bluebearCode represents bluebear code.
	bluebearCode = nook.Code{
		Value: "cbr00"}
)

var (
	// bluebearAmericanEnglishName represents bluebear american english name.
	bluebearAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bluebear"}

	// bluebearCanadianFrenchName represents bluebear canadian french name.
	bluebearCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Myrtille"}

	// bluebearDutchName represents bluebear dutch name.
	bluebearDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bluebear"}

	// bluebearFrenchName represents bluebear french name.
	bluebearFrenchName = nook.Name{
		Language: language.French,
		Value:    "Myrtille"}

	// bluebearGermanName represents bluebear german name.
	bluebearGermanName = nook.Name{
		Language: language.German,
		Value:    "Birte"}

	// bluebearItalianName represents bluebear italian name.
	bluebearItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Azzurra"}

	// bluebearJapaneseName represents bluebear japanese name.
	bluebearJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "グルミン"}

	// bluebearKoreanName represents bluebear korean name.
	bluebearKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "글루민"}

	// bluebearLatinAmericanSpanishName represents bluebear latin american spanish name.
	bluebearLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Celeste"}

	// bluebearRussianName represents bluebear russian name.
	bluebearRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Блюбеар"}

	// bluebearSimplifiedChineseName represents bluebear simplified chinese name.
	bluebearSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "娃娃"}

	// bluebearSpanishName represents bluebear spanish name.
	bluebearSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Celeste"}

	// bluebearTraditionalChineseName represents bluebear traditional chinese name.
	bluebearTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "娃娃"}
)

var (
	// bluebearName represents bluebear name.
	bluebearName = nook.Languages{
		language.AmericanEnglish:      bluebearAmericanEnglishName,
		language.CanadianFrench:       bluebearCanadianFrenchName,
		language.Dutch:                bluebearDutchName,
		language.French:               bluebearFrenchName,
		language.German:               bluebearGermanName,
		language.Italian:              bluebearItalianName,
		language.Japanese:             bluebearJapaneseName,
		language.Korean:               bluebearKoreanName,
		language.LatinAmericanSpanish: bluebearLatinAmericanSpanishName,
		language.Russian:              bluebearRussianName,
		language.SimplifiedChinese:    bluebearSimplifiedChineseName,
		language.Spanish:              bluebearSpanishName,
		language.TraditionalChinese:   bluebearTraditionalChineseName}
)

var (
	// bluebearCharacter represents bluebear character.
	bluebearCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: bluebearBirthday,
		Code:     bluebearCode,
		Key:      character.Bluebear,
		Gender:   gender.Female,
		Name:     bluebearName,
		Special:  false}
)

var (
	// bluebearAmericanEnglishPhrase represents bluebear american english phrase.
	bluebearAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "peach"}

	// bluebearCanadianFrenchPhrase represents bluebear canadian french phrase.
	bluebearCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "prune"}

	// bluebearDutchPhrase represents bluebear dutch phrase.
	bluebearDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "perzik"}

	// bluebearFrenchPhrase represents bluebear french phrase.
	bluebearFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "prune"}

	// bluebearGermanPhrase represents bluebear german phrase.
	bluebearGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "früchtchen"}

	// bluebearItalianPhrase represents bluebear italian phrase.
	bluebearItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pepè"}

	// bluebearJapanesePhrase represents bluebear japanese phrase.
	bluebearJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "キュン"}

	// bluebearKoreanPhrase represents bluebear korean phrase.
	bluebearKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "두근"}

	// bluebearLatinAmericanSpanishPhrase represents bluebear latin american spanish phrase.
	bluebearLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "melí-melá"}

	// bluebearRussianPhrase represents bluebear russian phrase.
	bluebearRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "персик"}

	// bluebearSimplifiedChinesePhrase represents bluebear simplified chinese phrase.
	bluebearSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "怦怦"}

	// bluebearSpanishPhrase represents bluebear spanish phrase.
	bluebearSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cielito"}

	// bluebearTraditionalChinesePhrase represents bluebear traditional chinese phrase.
	bluebearTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "怦怦"}
)

var (
	// bluebearPhrase represents bluebear phrase.
	bluebearPhrase = nook.Languages{
		language.AmericanEnglish:      bluebearAmericanEnglishPhrase,
		language.CanadianFrench:       bluebearCanadianFrenchPhrase,
		language.Dutch:                bluebearDutchPhrase,
		language.French:               bluebearFrenchPhrase,
		language.German:               bluebearGermanPhrase,
		language.Italian:              bluebearItalianPhrase,
		language.Japanese:             bluebearJapanesePhrase,
		language.Korean:               bluebearKoreanPhrase,
		language.LatinAmericanSpanish: bluebearLatinAmericanSpanishPhrase,
		language.Russian:              bluebearRussianPhrase,
		language.SimplifiedChinese:    bluebearSimplifiedChinesePhrase,
		language.Spanish:              bluebearSpanishPhrase,
		language.TraditionalChinese:   bluebearTraditionalChinesePhrase}
)

var (
	// Bluebear represents bluebear.
	Bluebear = nook.Villager{
		Character:   bluebearCharacter,
		Personality: personality.Peppy,
		Phrase:      bluebearPhrase}
)
