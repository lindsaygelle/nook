package squirrel

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
	rickyBirthday = nook.Birthday{
		Day:   14,
		Month: time.September}
)

var (
	rickyCode = nook.Code{
		Value: "squ10"}
)

var (
	rickyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ricky"}

	rickyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rocky"}

	rickyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ricky"}

	rickyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rocky"}

	rickyGermanName = nook.Name{
		Language: language.German,
		Value:    "Ronny"}

	rickyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sberlo"}

	rickyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カジロウ"}

	rickyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "갈가리"}

	rickyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Altramuz"}

	rickyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рики"}

	rickySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘉治"}

	rickySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Altramuz"}

	rickyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘉治"}
)

var (
	rickyName = nook.Languages{
		language.AmericanEnglish:      rickyAmericanEnglishName,
		language.CanadianFrench:       rickyCanadianFrenchName,
		language.Dutch:                rickyDutchName,
		language.French:               rickyFrenchName,
		language.German:               rickyGermanName,
		language.Italian:              rickyItalianName,
		language.Japanese:             rickyJapaneseName,
		language.Korean:               rickyKoreanName,
		language.LatinAmericanSpanish: rickyLatinAmericanSpanishName,
		language.Russian:              rickyRussianName,
		language.SimplifiedChinese:    rickySimplifiedChineseName,
		language.Spanish:              rickySpanishName,
		language.TraditionalChinese:   rickyTraditionalChineseName}
)

var (
	rickyCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: rickyBirthday,
		Code:     rickyCode,
		Key:      character.Ricky,
		Gender:   gender.Male,
		Name:     rickyName}
)

var (
	rickyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nutcase"}

	rickyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "dingo"}

	rickyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mafnoot"}

	rickyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "dingo"}

	rickyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knorz"}

	rickyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "nocciola"}

	rickyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "キッ"}

	rickyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "사각사각"}

	rickyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "crac-crac"}

	rickyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "арахис"}

	rickySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "炯炯"}

	rickySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cáscaras"}

	rickyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "炯炯"}
)

var (
	rickyPhrase = nook.Languages{
		language.AmericanEnglish:      rickyAmericanEnglishPhrase,
		language.CanadianFrench:       rickyCanadianFrenchPhrase,
		language.Dutch:                rickyDutchPhrase,
		language.French:               rickyFrenchPhrase,
		language.German:               rickyGermanPhrase,
		language.Italian:              rickyItalianPhrase,
		language.Japanese:             rickyJapanesePhrase,
		language.Korean:               rickyKoreanPhrase,
		language.LatinAmericanSpanish: rickyLatinAmericanSpanishPhrase,
		language.Russian:              rickyRussianPhrase,
		language.SimplifiedChinese:    rickySimplifiedChinesePhrase,
		language.Spanish:              rickySpanishPhrase,
		language.TraditionalChinese:   rickyTraditionalChinesePhrase}
)

var (
	Ricky = nook.Villager{
		Character:   rickyCharacter,
		Personality: personality.Cranky,
		Phrase:      rickyPhrase}
)
