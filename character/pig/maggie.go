package pig

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
	maggieBirthday = nook.Birthday{
		Day:   3,
		Month: time.September}
)

var (
	maggieCode = nook.Code{
		Value: "pig10"}
)

var (
	maggieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Maggie"}

	maggieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Marjorie"}

	maggieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Maggie"}

	maggieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marjorie"}

	maggieGermanName = nook.Name{
		Language: language.German,
		Value:    "Magda"}

	maggieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Marcella"}

	maggieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マーガレット"}

	maggieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마가렛"}

	maggieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Margarí"}

	maggieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мэгги"}

	maggieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "玛格"}

	maggieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Margarí"}

	maggieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "瑪格"}
)

var (
	maggieName = nook.Languages{
		language.AmericanEnglish:      maggieAmericanEnglishName,
		language.CanadianFrench:       maggieCanadianFrenchName,
		language.Dutch:                maggieDutchName,
		language.French:               maggieFrenchName,
		language.German:               maggieGermanName,
		language.Italian:              maggieItalianName,
		language.Japanese:             maggieJapaneseName,
		language.Korean:               maggieKoreanName,
		language.LatinAmericanSpanish: maggieLatinAmericanSpanishName,
		language.Russian:              maggieRussianName,
		language.SimplifiedChinese:    maggieSimplifiedChineseName,
		language.Spanish:              maggieSpanishName,
		language.TraditionalChinese:   maggieTraditionalChineseName}
)

var (
	maggieCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: maggieBirthday,
		Code:     maggieCode,
		Key:      character.Maggie,
		Gender:   gender.Female,
		Name:     maggieName}
)

var (
	maggieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "schep"}

	maggieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grouik"}

	maggieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "knor-knor"}

	maggieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grouik"}

	maggieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kringel"}

	maggieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sissì"}

	maggieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "うん"}

	maggieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "유후"}

	maggieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "oinki"}

	maggieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрюля"}

	maggieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嗯"}

	maggieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "girasol"}

	maggieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗯"}
)

var (
	maggiePhrase = nook.Languages{
		language.AmericanEnglish:      maggieAmericanEnglishPhrase,
		language.CanadianFrench:       maggieCanadianFrenchPhrase,
		language.Dutch:                maggieDutchPhrase,
		language.French:               maggieFrenchPhrase,
		language.German:               maggieGermanPhrase,
		language.Italian:              maggieItalianPhrase,
		language.Japanese:             maggieJapanesePhrase,
		language.Korean:               maggieKoreanPhrase,
		language.LatinAmericanSpanish: maggieLatinAmericanSpanishPhrase,
		language.Russian:              maggieRussianPhrase,
		language.SimplifiedChinese:    maggieSimplifiedChinesePhrase,
		language.Spanish:              maggieSpanishPhrase,
		language.TraditionalChinese:   maggieTraditionalChinesePhrase}
)

var (
	Maggie = nook.Villager{
		Character:   maggieCharacter,
		Personality: personality.Normal,
		Phrase:      maggiePhrase}
)
