package horse

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
	annaliseBirthday = nook.Birthday{
		Day:   2,
		Month: time.December}
)

var (
	annaliseCode = nook.Code{
		Value: "hrs09"}
)

var (
	annaliseAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Annalise"}

	annaliseCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Âne-lise"}

	annaliseDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Annalise"}

	annaliseFrenchName = nook.Name{
		Language: language.French,
		Value:    "Âne-lise"}

	annaliseGermanName = nook.Name{
		Language: language.German,
		Value:    "Annerose"}

	annaliseItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lisa"}

	annaliseJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シルブプレ"}

	annaliseKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "실부플레"}

	annaliseLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Isadora"}

	annaliseRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лиза"}

	annaliseSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "舒芙蕾"}

	annaliseSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Isadora"}

	annaliseTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "舒芙蕾"}
)

var (
	annaliseName = nook.Languages{
		language.AmericanEnglish:      annaliseAmericanEnglishName,
		language.CanadianFrench:       annaliseCanadianFrenchName,
		language.Dutch:                annaliseDutchName,
		language.French:               annaliseFrenchName,
		language.German:               annaliseGermanName,
		language.Italian:              annaliseItalianName,
		language.Japanese:             annaliseJapaneseName,
		language.Korean:               annaliseKoreanName,
		language.LatinAmericanSpanish: annaliseLatinAmericanSpanishName,
		language.Russian:              annaliseRussianName,
		language.SimplifiedChinese:    annaliseSimplifiedChineseName,
		language.Spanish:              annaliseSpanishName,
		language.TraditionalChinese:   annaliseTraditionalChineseName}
)

var (
	annaliseCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: annaliseBirthday,
		Code:     annaliseCode,
		Key:      character.Annalise,
		Gender:   gender.Female,
		Name:     annaliseName}
)

var (
	annaliseAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nipper"}

	annaliseCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "hippique"}

	annaliseDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "veulen"}

	annaliseFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "galapia"}

	annaliseGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mausezahn"}

	annaliseItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "palio"}

	annaliseJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "サヴァ"}

	annaliseKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "봉쥬르"}

	annaliseLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "nih-nah"}

	annaliseRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кусь-кусь"}

	annaliseSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "舒服"}

	annaliseSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "criatura"}

	annaliseTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "舒服"}
)

var (
	annalisePhrase = nook.Languages{
		language.AmericanEnglish:      annaliseAmericanEnglishPhrase,
		language.CanadianFrench:       annaliseCanadianFrenchPhrase,
		language.Dutch:                annaliseDutchPhrase,
		language.French:               annaliseFrenchPhrase,
		language.German:               annaliseGermanPhrase,
		language.Italian:              annaliseItalianPhrase,
		language.Japanese:             annaliseJapanesePhrase,
		language.Korean:               annaliseKoreanPhrase,
		language.LatinAmericanSpanish: annaliseLatinAmericanSpanishPhrase,
		language.Russian:              annaliseRussianPhrase,
		language.SimplifiedChinese:    annaliseSimplifiedChinesePhrase,
		language.Spanish:              annaliseSpanishPhrase,
		language.TraditionalChinese:   annaliseTraditionalChinesePhrase}
)

var (
	Annalise = nook.Villager{
		Character:   annaliseCharacter,
		Personality: personality.Snooty,
		Phrase:      annalisePhrase}
)
