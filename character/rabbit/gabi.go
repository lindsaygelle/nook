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
	// gabiBirthday represents Gabi's birthday.
	gabiBirthday = nook.Birthday{
		Day:   16,
		Month: time.December}
)

var (
	// gabiCode represents Gabi's unique code.
	gabiCode = nook.Code{
		Value: "rbt05"}
)

var (
	// gabiAmericanEnglishName represents Gabi's name in American English.
	gabiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gabi"}

	// gabiCanadianFrenchName represents Gabi's name in Canadian French.
	gabiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gaby"}

	// gabiDutchName represents Gabi's name in Dutch.
	gabiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gabi"}

	// gabiFrenchName represents Gabi's name in French.
	gabiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gaby"}

	// gabiGermanName represents Gabi's name in German.
	gabiGermanName = nook.Name{
		Language: language.German,
		Value:    "Gabi"}

	// gabiItalianName represents Gabi's name in Italian.
	gabiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gabri"}

	// gabiJapaneseName represents Gabi's name in Japanese.
	gabiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ペチカ"}

	// gabiKoreanName represents Gabi's name in Korean.
	gabiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "패티카"}

	// gabiLatinAmericanSpanishName represents Gabi's name in Latin American Spanish.
	gabiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Piluca"}

	// gabiRussianName represents Gabi's name in Russian.
	gabiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Габи"}

	// gabiSimplifiedChineseName represents Gabi's name in Simplified Chinese.
	gabiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "珮琪"}

	// gabiSpanishName represents Gabi's name in Spanish.
	gabiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Piluca"}

	// gabiTraditionalChineseName represents Gabi's name in Traditional Chinese.
	gabiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "珮琪"}
)

var (
	// gabiName represents Gabi's name in different languages.
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
	// gabiCharacter represents Gabi's character information.
	gabiCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: gabiBirthday,
		Code:     gabiCode,
		Key:      character.Gabi,
		Gender:   gender.Female,
		Name:     gabiName,
		Special:  false}
)

var (
	// gabiAmericanEnglishPhrase represents Gabi's phrase in American English.
	gabiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "honeybun"}

	// gabiCanadianFrenchPhrase represents Gabi's phrase in Canadian French.
	gabiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon lièvre"}

	// gabiDutchPhrase represents Gabi's phrase in Dutch.
	gabiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "honnepon"}

	// gabiFrenchPhrase represents Gabi's phrase in French.
	gabiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon lièvre"}

	// gabiGermanPhrase represents Gabi's phrase in German.
	gabiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "liebling"}

	// gabiItalianPhrase represents Gabi's phrase in Italian.
	gabiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mielito"}

	// gabiJapanesePhrase represents Gabi's phrase in Japanese.
	gabiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だっち"}

	// gabiKoreanPhrase represents Gabi's phrase in Korean.
	gabiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꺄오"}

	// gabiLatinAmericanSpanishPhrase represents Gabi's phrase in Latin American Spanish.
	gabiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "reboing"}

	// gabiRussianPhrase represents Gabi's phrase in Russian.
	gabiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "зайчик мой"}

	// gabiSimplifiedChinesePhrase represents Gabi's phrase in Simplified Chinese.
	gabiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "琪他"}

	// gabiSpanishPhrase represents Gabi's phrase in Spanish.
	gabiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mari"}

	// gabiTraditionalChinesePhrase represents Gabi's phrase in Traditional Chinese.
	gabiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "琪他"}
)

var (
	// gabiPhrase represents Gabi's phrases in different languages.
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
	// Gabi represents the character Gabi with her complete information.
	Gabi = nook.Villager{
		Character:   gabiCharacter,
		Personality: personality.Peppy,
		Phrase:      gabiPhrase}
)
