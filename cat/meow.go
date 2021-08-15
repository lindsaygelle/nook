package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "N/A"}

	meowDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	meowFrenchName = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	meowGermanName = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	meowItalianName = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	meowJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ミャウ"}

	meowKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	meowLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	meowRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	meowSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	meowSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	meowTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Animal:   Cat,
		Birthday: meowBirthday,
		Code:     meowCode,
		Gender:   nook.Female,
		Name:     meowName}
)

var (
	meowAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	meowCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	meowDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	meowFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	meowGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	meowItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	meowJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ミャウ"}

	meowKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	meowLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	meowRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	meowSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	meowSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	meowTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Personality: nook.Peppy,
		Phrase:      meowPhrase}
)
