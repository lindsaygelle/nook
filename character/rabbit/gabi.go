package rabbit

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
	gabiBirthday = nook.Birthday{
		Day:   16,
		Month: time.December}
)

var (
	gabiCode = nook.Code{
		Value: "rbt05"}
)

var (
	gabiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gabi"}

	gabiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gaby"}

	gabiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gabi"}

	gabiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gaby"}

	gabiGermanName = nook.Name{
		Language: language.German,
		Value:    "Gabi"}

	gabiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gabri"}

	gabiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ペチカ"}

	gabiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "패티카"}

	gabiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Piluca"}

	gabiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Габи"}

	gabiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "珮琪"}

	gabiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Piluca"}

	gabiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "珮琪"}
)

var (
	gabiName = nook.Languages{
		language.AmericanEnglish:      gabiAmericanEnglishName,
		language.CanadianFrench:       gabiCanadianFrenchName,
		language.Dutch:                gabiDutchName,
		language.French:               gabiFrenchName,
		language.German:               gabiGermanName,
		language.Italian:              gabiItalianName,
		language.Japanese:             gabiJapaneseName,
		language.Korean:               gabiKoreanName,
		language.LatinAmericanSpanish: gabiLatinAmericanSpanishName,
		language.Russian:              gabiRussianName,
		language.SimplifiedChinese:    gabiSimplifiedChineseName,
		language.Spanish:              gabiSpanishName,
		language.TraditionalChinese:   gabiTraditionalChineseName}
)

var (
	gabiCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: gabiBirthday,
		Code:     gabiCode,
		Key:      character.Gabi,
		Gender:   gender.Female,
		Name:     gabiName}
)

var (
	gabiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "honeybun"}

	gabiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon lièvre"}

	gabiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "honnepon"}

	gabiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon lièvre"}

	gabiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "liebling"}

	gabiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mielito"}

	gabiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だっち"}

	gabiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꺄오"}

	gabiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "reboing"}

	gabiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "зайчик мой"}

	gabiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "琪他"}

	gabiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mari"}

	gabiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "琪他"}
)

var (
	gabiPhrase = nook.Languages{
		language.AmericanEnglish:      gabiAmericanEnglishPhrase,
		language.CanadianFrench:       gabiCanadianFrenchPhrase,
		language.Dutch:                gabiDutchPhrase,
		language.French:               gabiFrenchPhrase,
		language.German:               gabiGermanPhrase,
		language.Italian:              gabiItalianPhrase,
		language.Japanese:             gabiJapanesePhrase,
		language.Korean:               gabiKoreanPhrase,
		language.LatinAmericanSpanish: gabiLatinAmericanSpanishPhrase,
		language.Russian:              gabiRussianPhrase,
		language.SimplifiedChinese:    gabiSimplifiedChinesePhrase,
		language.Spanish:              gabiSpanishPhrase,
		language.TraditionalChinese:   gabiTraditionalChinesePhrase}
)

var (
	Gabi = nook.Villager{
		Character:   gabiCharacter,
		Personality: personality.Peppy,
		Phrase:      gabiPhrase}
)
