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

// chowBirthday represents Chow's birthday.
var (
	chowBirthday = nook.Birthday{
		Day:   22,
		Month: time.July}
)

// chowCode represents Chow's unique code.
var (
	chowCode = nook.Code{
		Value: "bea03"}
)

// Different names for Chow in various languages.
var (
	chowAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chow"}

	chowCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chulin"}

	chowDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chow"}

	chowFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chulin"}

	chowGermanName = nook.Name{
		Language: language.German,
		Value:    "Chang"}

	chowItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Chowchow"}

	chowJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チャウヤン"}

	chowKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "츄양"}

	chowLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pando"}

	chowRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Чау"}

	chowSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "朝阳"}

	chowSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pando"}

	chowTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "朝陽"}
)

// chowName represents Chow's name in different languages.
var (
	chowName = nook.Languages{
		language.AmericanEnglish:      chowAmericanEnglishName,
		language.CanadianFrench:       chowCanadianFrenchName,
		language.Dutch:                chowDutchName,
		language.French:               chowFrenchName,
		language.German:               chowGermanName,
		language.Italian:              chowItalianName,
		language.Japanese:             chowJapaneseName,
		language.Korean:               chowKoreanName,
		language.LatinAmericanSpanish: chowLatinAmericanSpanishName,
		language.Russian:              chowRussianName,
		language.SimplifiedChinese:    chowSimplifiedChineseName,
		language.Spanish:              chowSpanishName,
		language.TraditionalChinese:   chowTraditionalChineseName}
)

// chowCharacter represents Chow's character information.
var (
	chowCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: chowBirthday,
		Code:     chowCode,
		Key:      character.Chow,
		Gender:   gender.Male,
		Name:     chowName,
		Special:  false}
)

// Different phrases spoken by Chow in various languages.
var (
	chowAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "aiya"}

	chowCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "lala"}

	chowDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kiai"}

	chowFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "lala"}

	chowGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hiijaa"}

	chowItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ehilà"}

	chowJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アルヨ"}

	chowKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그럼"}

	chowLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grinchu"}

	chowRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ай-яй"}

	chowSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "有喔"}

	chowSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ándale-oso"}

	chowTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "有喔"}
)

// chowPhrase represents Chow's phrases in different languages.
var (
	chowPhrase = nook.Languages{
		language.AmericanEnglish:      chowAmericanEnglishPhrase,
		language.CanadianFrench:       chowCanadianFrenchPhrase,
		language.Dutch:                chowDutchPhrase,
		language.French:               chowFrenchPhrase,
		language.German:               chowGermanPhrase,
		language.Italian:              chowItalianPhrase,
		language.Japanese:             chowJapanesePhrase,
		language.Korean:               chowKoreanPhrase,
		language.LatinAmericanSpanish: chowLatinAmericanSpanishPhrase,
		language.Russian:              chowRussianPhrase,
		language.SimplifiedChinese:    chowSimplifiedChinesePhrase,
		language.Spanish:              chowSpanishPhrase,
		language.TraditionalChinese:   chowTraditionalChinesePhrase}
)

// Chow represents the character Chow with his complete information.
var (
	Chow = nook.Villager{
		Character:   chowCharacter,
		Personality: personality.Cranky,
		Phrase:      chowPhrase}
)
