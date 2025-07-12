package cat

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
	// meowBirthday represents meow birthday.
	meowBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// meowCode represents meow code.
	meowCode = nook.Code{
		Value: ""}
)

var (
	// meowAmericanEnglishName represents meow american english name.
	meowAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Meow"}

	// meowCanadianFrenchName represents meow canadian french name.
	meowCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// meowDutchName represents meow dutch name.
	meowDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// meowFrenchName represents meow french name.
	meowFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// meowGermanName represents meow german name.
	meowGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// meowItalianName represents meow italian name.
	meowItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// meowJapaneseName represents meow japanese name.
	meowJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ミャウ"}

	// meowKoreanName represents meow korean name.
	meowKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// meowLatinAmericanSpanishName represents meow latin american spanish name.
	meowLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// meowRussianName represents meow russian name.
	meowRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// meowSimplifiedChineseName represents meow simplified chinese name.
	meowSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// meowSpanishName represents meow spanish name.
	meowSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// meowTraditionalChineseName represents meow traditional chinese name.
	meowTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// meowName represents meow name.
	meowName = nook.Languages{
		language.AmericanEnglish:      meowAmericanEnglishName,
		language.CanadianFrench:       meowCanadianFrenchName,
		language.Dutch:                meowDutchName,
		language.French:               meowFrenchName,
		language.German:               meowGermanName,
		language.Italian:              meowItalianName,
		language.Japanese:             meowJapaneseName,
		language.Korean:               meowKoreanName,
		language.LatinAmericanSpanish: meowLatinAmericanSpanishName,
		language.Russian:              meowRussianName,
		language.SimplifiedChinese:    meowSimplifiedChineseName,
		language.Spanish:              meowSpanishName,
		language.TraditionalChinese:   meowTraditionalChineseName}
)

var (
	// meowCharacter represents meow character.
	meowCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: meowBirthday,
		Code:     meowCode,
		Key:      character.Meow,
		Gender:   gender.Female,
		Name:     meowName,
		Special:  false}
)

var (
	// meowAmericanEnglishPhrase represents meow american english phrase.
	meowAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ミャウ"}

	// meowCanadianFrenchPhrase represents meow canadian french phrase.
	meowCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// meowDutchPhrase represents meow dutch phrase.
	meowDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// meowFrenchPhrase represents meow french phrase.
	meowFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// meowGermanPhrase represents meow german phrase.
	meowGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// meowItalianPhrase represents meow italian phrase.
	meowItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// meowJapanesePhrase represents meow japanese phrase.
	meowJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ミャウ"}

	// meowKoreanPhrase represents meow korean phrase.
	meowKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// meowLatinAmericanSpanishPhrase represents meow latin american spanish phrase.
	meowLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// meowRussianPhrase represents meow russian phrase.
	meowRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// meowSimplifiedChinesePhrase represents meow simplified chinese phrase.
	meowSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// meowSpanishPhrase represents meow spanish phrase.
	meowSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// meowTraditionalChinesePhrase represents meow traditional chinese phrase.
	meowTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// meowPhrase represents meow phrase.
	meowPhrase = nook.Languages{
		language.AmericanEnglish:      meowAmericanEnglishPhrase,
		language.CanadianFrench:       meowCanadianFrenchPhrase,
		language.Dutch:                meowDutchPhrase,
		language.French:               meowFrenchPhrase,
		language.German:               meowGermanPhrase,
		language.Italian:              meowItalianPhrase,
		language.Japanese:             meowJapanesePhrase,
		language.Korean:               meowKoreanPhrase,
		language.LatinAmericanSpanish: meowLatinAmericanSpanishPhrase,
		language.Russian:              meowRussianPhrase,
		language.SimplifiedChinese:    meowSimplifiedChinesePhrase,
		language.Spanish:              meowSpanishPhrase,
		language.TraditionalChinese:   meowTraditionalChinesePhrase}
)

var (
	// Meow represents meow.
	Meow = nook.Villager{
		Character:   meowCharacter,
		Personality: personality.Peppy,
		Phrase:      meowPhrase}
)
