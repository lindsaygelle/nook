package mouse

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
	mooseBirthday = nook.Birthday{
		Day:   13,
		Month: time.September}
)

var (
	mooseCode = nook.Code{
		Value: "mus14"}
)

var (
	mooseAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Moose"}

	mooseCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Joachim"}

	mooseDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Moose"}

	mooseFrenchName = nook.Name{
		Language: language.French,
		Value:    "Joachim"}

	mooseGermanName = nook.Name{
		Language: language.German,
		Value:    "Mausbert"}

	mooseItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Aldo"}

	mooseJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピン"}

	mooseKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "핑"}

	mooseLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sato"}

	mooseRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Муз"}

	mooseSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿聘"}

	mooseSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sato"}

	mooseTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿聘"}
)

var (
	mooseName = nook.Languages{
		language.AmericanEnglish:      mooseAmericanEnglishName,
		language.CanadianFrench:       mooseCanadianFrenchName,
		language.Dutch:                mooseDutchName,
		language.French:               mooseFrenchName,
		language.German:               mooseGermanName,
		language.Italian:              mooseItalianName,
		language.Japanese:             mooseJapaneseName,
		language.Korean:               mooseKoreanName,
		language.LatinAmericanSpanish: mooseLatinAmericanSpanishName,
		language.Russian:              mooseRussianName,
		language.SimplifiedChinese:    mooseSimplifiedChineseName,
		language.Spanish:              mooseSpanishName,
		language.TraditionalChinese:   mooseTraditionalChineseName}
)

var (
	mooseCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: mooseBirthday,
		Code:     mooseCode,
		Key:      character.Moose,
		Gender:   gender.Male,
		Name:     mooseName,
		Special:  false}
)

var (
	mooseAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "shorty"}

	mooseCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bééléé"}

	mooseDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kleintje"}

	mooseFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bééléé"}

	mooseGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "käääse"}

	mooseItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quink"}

	mooseJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "にゃろ"}

	mooseKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "짜샤"}

	mooseLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "atiza"}

	mooseRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "коротышка"}

	mooseSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "可恶"}

	mooseSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "atiza"}

	mooseTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "可惡"}
)

var (
	moosePhrase = nook.Languages{
		language.AmericanEnglish:      mooseAmericanEnglishPhrase,
		language.CanadianFrench:       mooseCanadianFrenchPhrase,
		language.Dutch:                mooseDutchPhrase,
		language.French:               mooseFrenchPhrase,
		language.German:               mooseGermanPhrase,
		language.Italian:              mooseItalianPhrase,
		language.Japanese:             mooseJapanesePhrase,
		language.Korean:               mooseKoreanPhrase,
		language.LatinAmericanSpanish: mooseLatinAmericanSpanishPhrase,
		language.Russian:              mooseRussianPhrase,
		language.SimplifiedChinese:    mooseSimplifiedChinesePhrase,
		language.Spanish:              mooseSpanishPhrase,
		language.TraditionalChinese:   mooseTraditionalChinesePhrase}
)

var (
	Moose = nook.Villager{
		Character:   mooseCharacter,
		Personality: personality.Jock,
		Phrase:      moosePhrase}
)
