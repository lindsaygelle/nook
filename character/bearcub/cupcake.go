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
	cupcakeBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	cupcakeCode = nook.Code{
		Value: ""}
)

var (
	cupcakeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cupcake"}

	cupcakeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	cupcakeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	cupcakeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bella"}

	cupcakeGermanName = nook.Name{
		Language: language.German,
		Value:    "Bärbel"}

	cupcakeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Chicca"}

	cupcakeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クリスピー"}

	cupcakeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	cupcakeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	cupcakeRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	cupcakeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "晶晶"}

	cupcakeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rosita"}

	cupcakeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	cupcakeName = nook.Languages{
		language.AmericanEnglish:      cupcakeAmericanEnglishName,
		language.CanadianFrench:       cupcakeCanadianFrenchName,
		language.Dutch:                cupcakeDutchName,
		language.French:               cupcakeFrenchName,
		language.German:               cupcakeGermanName,
		language.Italian:              cupcakeItalianName,
		language.Japanese:             cupcakeJapaneseName,
		language.Korean:               cupcakeKoreanName,
		language.LatinAmericanSpanish: cupcakeLatinAmericanSpanishName,
		language.Russian:              cupcakeRussianName,
		language.SimplifiedChinese:    cupcakeSimplifiedChineseName,
		language.Spanish:              cupcakeSpanishName,
		language.TraditionalChinese:   cupcakeTraditionalChineseName}
)

var (
	cupcakeCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: cupcakeBirthday,
		Code:     cupcakeCode,
		Key:      character.Cupcake,
		Gender:   gender.Female,
		Name:     cupcakeName}
)

var (
	cupcakeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sugar pie"}

	cupcakeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	cupcakeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	cupcakeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "hum miam"}

	cupcakeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schätzchen"}

	cupcakeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "zucchero"}

	cupcakeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "かしら"}

	cupcakeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	cupcakeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	cupcakeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	cupcakeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "明白"}

	cupcakeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pastelito"}

	cupcakeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	cupcakePhrase = nook.Languages{
		language.AmericanEnglish:      cupcakeAmericanEnglishPhrase,
		language.CanadianFrench:       cupcakeCanadianFrenchPhrase,
		language.Dutch:                cupcakeDutchPhrase,
		language.French:               cupcakeFrenchPhrase,
		language.German:               cupcakeGermanPhrase,
		language.Italian:              cupcakeItalianPhrase,
		language.Japanese:             cupcakeJapanesePhrase,
		language.Korean:               cupcakeKoreanPhrase,
		language.LatinAmericanSpanish: cupcakeLatinAmericanSpanishPhrase,
		language.Russian:              cupcakeRussianPhrase,
		language.SimplifiedChinese:    cupcakeSimplifiedChinesePhrase,
		language.Spanish:              cupcakeSpanishPhrase,
		language.TraditionalChinese:   cupcakeTraditionalChinesePhrase}
)

var (
	Cupcake = nook.Villager{
		Character:   cupcakeCharacter,
		Personality: personality.Snooty,
		Phrase:      cupcakePhrase}
)
