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

var (
	klausBirthday = nook.Birthday{
		Day:   31,
		Month: time.March}
)

var (
	klausCode = nook.Code{
		Value: "bea14"}
)

var (
	klausAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Klaus"}

	klausCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Klaus"}

	klausDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Klaus"}

	klausFrenchName = nook.Name{
		Language: language.French,
		Value:    "Klaus"}

	klausGermanName = nook.Name{
		Language: language.German,
		Value:    "Grischa"}

	klausItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Maciste"}

	klausJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クマロス"}

	klausKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "곰도로스"}

	klausLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gruñerto"}

	klausRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Клаус"}

	klausSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "熊战士"}

	klausSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gruñerto"}

	klausTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "熊戰士"}
)

var (
	klausName = nook.Languages{
		language.AmericanEnglish:      klausAmericanEnglishName,
		language.CanadianFrench:       klausCanadianFrenchName,
		language.Dutch:                klausDutchName,
		language.French:               klausFrenchName,
		language.German:               klausGermanName,
		language.Italian:              klausItalianName,
		language.Japanese:             klausJapaneseName,
		language.Korean:               klausKoreanName,
		language.LatinAmericanSpanish: klausLatinAmericanSpanishName,
		language.Russian:              klausRussianName,
		language.SimplifiedChinese:    klausSimplifiedChineseName,
		language.Spanish:              klausSpanishName,
		language.TraditionalChinese:   klausTraditionalChineseName}
)

var (
	klausCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: klausBirthday,
		Code:     klausCode,
		Key:      character.Klaus,
		Gender:   gender.Male,
		Name:     klausName}
)

var (
	klausAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "strudel"}

	klausCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "achtung"}

	klausDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "strudel"}

	klausFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "achtung"}

	klausGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "gruaaa"}

	klausItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bailar"}

	klausJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "オパー"}

	klausKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "라저"}

	klausLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "achurf"}

	klausRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "штрудель"}

	klausSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Over"}

	klausSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ricitos"}

	klausTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "Over"}
)

var (
	klausPhrase = nook.Languages{
		language.AmericanEnglish:      klausAmericanEnglishPhrase,
		language.CanadianFrench:       klausCanadianFrenchPhrase,
		language.Dutch:                klausDutchPhrase,
		language.French:               klausFrenchPhrase,
		language.German:               klausGermanPhrase,
		language.Italian:              klausItalianPhrase,
		language.Japanese:             klausJapanesePhrase,
		language.Korean:               klausKoreanPhrase,
		language.LatinAmericanSpanish: klausLatinAmericanSpanishPhrase,
		language.Russian:              klausRussianPhrase,
		language.SimplifiedChinese:    klausSimplifiedChinesePhrase,
		language.Spanish:              klausSpanishPhrase,
		language.TraditionalChinese:   klausTraditionalChinesePhrase}
)

var (
	Klaus = nook.Villager{
		Character:   klausCharacter,
		Personality: personality.Smug,
		Phrase:      klausPhrase}
)
