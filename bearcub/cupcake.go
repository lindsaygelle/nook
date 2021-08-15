package bearcub

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Bellahum miam"}

	cupcakeGermanName = nook.Name{
		Language: language.German,
		Value:    "Bärbelschätzchen"}

	cupcakeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Chiccazucchero"}

	cupcakeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クリスピーかしら"}

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
		Value:    "晶晶明白"}

	cupcakeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rositapastelito"}

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
		Animal:   Bearcub,
		Birthday: cupcakeBirthday,
		Code:     cupcakeCode,
		Gender:   nook.Female,
		Name:     cupcakeName}
)

var (
	cupcakeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	cupcakeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	cupcakeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

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
		Value:    "N/A"}

	cupcakeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	cupcakeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	cupcakeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "明白"}

	cupcakeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pastelito"}

	cupcakeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	cupcakePhrase = nook.Languages{
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
	Cupcake = nook.Villager{
		Character:   cupcakeCharacter,
		Personality: nook.Snooty,
		Phrase:      cupcakePhrase}
)
