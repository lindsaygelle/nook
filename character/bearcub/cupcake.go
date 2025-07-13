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
	// cupcakeBirthday represents cupcake birthday.
	cupcakeBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// cupcakeCode represents cupcake code.
	cupcakeCode = nook.Code{
		Value: ""}
)

var (
	// cupcakeAmericanEnglishName represents cupcake american english name.
	cupcakeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cupcake"}

	// cupcakeCanadianFrenchName represents cupcake canadian french name.
	cupcakeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// cupcakeDutchName represents cupcake dutch name.
	cupcakeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// cupcakeFrenchName represents cupcake french name.
	cupcakeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bella"}

	// cupcakeGermanName represents cupcake german name.
	cupcakeGermanName = nook.Name{
		Language: language.German,
		Value:    "Bärbel"}

	// cupcakeItalianName represents cupcake italian name.
	cupcakeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Chicca"}

	// cupcakeJapaneseName represents cupcake japanese name.
	cupcakeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クリスピー"}

	// cupcakeKoreanName represents cupcake korean name.
	cupcakeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// cupcakeLatinAmericanSpanishName represents cupcake latin american spanish name.
	cupcakeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// cupcakeRussianName represents cupcake russian name.
	cupcakeRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// cupcakeSimplifiedChineseName represents cupcake simplified chinese name.
	cupcakeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "晶晶"}

	// cupcakeSpanishName represents cupcake spanish name.
	cupcakeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rosita"}

	// cupcakeTraditionalChineseName represents cupcake traditional chinese name.
	cupcakeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// cupcakeName represents cupcake name.
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
	// cupcakeCharacter represents cupcake character.
	cupcakeCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: cupcakeBirthday,
		Code:     cupcakeCode,
		Key:      character.Cupcake,
		Gender:   gender.Female,
		Name:     cupcakeName,
		Special:  false}
)

var (
	// cupcakeAmericanEnglishPhrase represents cupcake american english phrase.
	cupcakeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sugar pie"}

	// cupcakeCanadianFrenchPhrase represents cupcake canadian french phrase.
	cupcakeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// cupcakeDutchPhrase represents cupcake dutch phrase.
	cupcakeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// cupcakeFrenchPhrase represents cupcake french phrase.
	cupcakeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "hum miam"}

	// cupcakeGermanPhrase represents cupcake german phrase.
	cupcakeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schätzchen"}

	// cupcakeItalianPhrase represents cupcake italian phrase.
	cupcakeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "zucchero"}

	// cupcakeJapanesePhrase represents cupcake japanese phrase.
	cupcakeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "かしら"}

	// cupcakeKoreanPhrase represents cupcake korean phrase.
	cupcakeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// cupcakeLatinAmericanSpanishPhrase represents cupcake latin american spanish phrase.
	cupcakeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// cupcakeRussianPhrase represents cupcake russian phrase.
	cupcakeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// cupcakeSimplifiedChinesePhrase represents cupcake simplified chinese phrase.
	cupcakeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "明白"}

	// cupcakeSpanishPhrase represents cupcake spanish phrase.
	cupcakeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pastelito"}

	// cupcakeTraditionalChinesePhrase represents cupcake traditional chinese phrase.
	cupcakeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// cupcakePhrase represents cupcake phrase.
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
	// Cupcake represents cupcake.
	Cupcake = nook.Villager{
		Character:   cupcakeCharacter,
		Personality: personality.Snooty,
		Phrase:      cupcakePhrase}
)
