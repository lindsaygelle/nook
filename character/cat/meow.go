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
	meowBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	meowCode = nook.Code{
		Value: ""}
)

var (
	meowAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Meow"}

	meowCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	meowDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	meowFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	meowGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	meowItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	meowJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ミャウ"}

	meowKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	meowLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	meowRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	meowSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	meowSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	meowTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
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
	meowCharacter = nook.Character{
		Animal:   animal.Cat,
		Birthday: meowBirthday,
		Code:     meowCode,
		Key:      character.Meow,
		Gender:   gender.Female,
		Name:     meowName}
)

var (
	meowAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ミャウ"}

	meowCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	meowDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	meowFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	meowGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	meowItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	meowJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ミャウ"}

	meowKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	meowLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	meowRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	meowSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	meowSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	meowTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	meowPhrase = nook.Languages{
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
	Meow = nook.Villager{
		Character:   meowCharacter,
		Personality: personality.Peppy,
		Phrase:      meowPhrase}
)
