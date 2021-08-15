package rabbit

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Gabymon lièvre"}

	gabiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gabihonnepon"}

	gabiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gabymon lièvre"}

	gabiGermanName = nook.Name{
		Language: language.German,
		Value:    "Gabiliebling"}

	gabiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gabrimielito"}

	gabiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ペチカだっち"}

	gabiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "패티카꺄오"}

	gabiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pilucareboing"}

	gabiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Габизайчик мой"}

	gabiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "珮琪琪他"}

	gabiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pilucamari"}

	gabiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "珮琪琪他"}
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
		Animal:   Rabbit,
		Birthday: gabiBirthday,
		Code:     gabiCode,
		Gender:   nook.Female,
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
	Gabi = nook.Villager{
		Character:   gabiCharacter,
		Personality: nook.Peppy,
		Phrase:      gabiPhrase}
)
