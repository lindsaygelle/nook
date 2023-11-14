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
	// sallyBirthday represents Sally's birthday (June 19th).
	sallyBirthday = nook.Birthday{
		Day:   19,
		Month: time.June}
)

var (
	// sallyCode represents Sally's unique code ("squ07").
	sallyCode = nook.Code{
		Value: "squ07"}
)

var (
	// sallyAmericanEnglishName represents Sally's name in American English.
	sallyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sally"}

	// sallyCanadianFrenchName represents Sally's name in Canadian French.
	sallyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Damia"}

	// sallyDutchName represents Sally's name in Dutch.
	sallyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sally"}

	// sallyFrenchName represents Sally's name in French.
	sallyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Damia"}

	// sallyGermanName represents Sally's name in German.
	sallyGermanName = nook.Name{
		Language: language.German,
		Value:    "Hanne"}

	// sallyItalianName represents Sally's name in Italian.
	sallyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rossella"}

	// sallyJapaneseName represents Sally's name in Japanese.
	sallyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ララミー"}

	// sallyKoreanName represents Sally's name in Korean.
	sallyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "라라미"}

	// sallyLatinAmericanSpanishName represents Sally's name in Latin American Spanish.
	sallyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Praliné"}

	// sallyRussianName represents Sally's name in Russian.
	sallyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Салли"}

	// sallySimplifiedChineseName represents Sally's name in Simplified Chinese.
	sallySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "拉拉米"}

	// sallySpanishName represents Sally's name in Spanish.
	sallySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Praliné"}

	// sallyTraditionalChineseName represents Sally's name in Traditional Chinese.
	sallyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "拉拉米"}
)

var (
	// sallyName represents Sally's name in different languages.
	sallyName = nook.Languages{
		language.AmericanEnglish:      sallyAmericanEnglishName,
		language.CanadianFrench:       sallyCanadianFrenchName,
		language.Dutch:                sallyDutchName,
		language.French:               sallyFrenchName,
		language.German:               sallyGermanName,
		language.Italian:              sallyItalianName,
		language.Japanese:             sallyJapaneseName,
		language.Korean:               sallyKoreanName,
		language.LatinAmericanSpanish: sallyLatinAmericanSpanishName,
		language.Russian:              sallyRussianName,
		language.SimplifiedChinese:    sallySimplifiedChineseName,
		language.Spanish:              sallySpanishName,
		language.TraditionalChinese:   sallyTraditionalChineseName}
)

var (
	// sallyCharacter represents Sally's character information.
	sallyCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: sallyBirthday,
		Code:     sallyCode,
		Key:      character.Sally,
		Gender:   gender.Female,
		Name:     sallyName,
		Special:  false}
)

var (
	// sallyAmericanEnglishPhrase represents Sally's phrase in American English.
	sallyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nutmeg"}

	// sallyCanadianFrenchPhrase represents Sally's phrase in Canadian French.
	sallyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tac"}

	// sallyDutchPhrase represents Sally's phrase in Dutch.
	sallyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "muskaatje"}

	// sallyFrenchPhrase represents Sally's phrase in French.
	sallyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tac"}

	// sallyGermanPhrase represents Sally's phrase in German.
	sallyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "muskat"}

	// sallyItalianPhrase represents Sally's phrase in Italian.
	sallyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "nocciolina"}

	// sallyJapanesePhrase represents Sally's phrase in Japanese.
	sallyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ったら"}

	// sallyKoreanPhrase represents Sally's phrase in Korean.
	sallyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "까꿍"}

	// sallyLatinAmericanSpanishPhrase represents Sally's phrase in Latin American Spanish.
	sallyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "raminí"}

	// sallyRussianPhrase represents Sally's phrase in Russian.
	sallyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мускатик"}

	// sallySimplifiedChinesePhrase represents Sally's phrase in Simplified Chinese.
	sallySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "要是"}

	// sallySpanishPhrase represents Sally's phrase in Spanish.
	sallySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ramita"}

	// sallyTraditionalChinesePhrase represents Sally's phrase in Traditional Chinese.
	sallyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "要是"}
)

var (
	// sallyPhrase represents Sally's phrases in different languages.
	sallyPhrase = nook.Languages{
		language.AmericanEnglish:      sallyAmericanEnglishPhrase,
		language.CanadianFrench:       sallyCanadianFrenchPhrase,
		language.Dutch:                sallyDutchPhrase,
		language.French:               sallyFrenchPhrase,
		language.German:               sallyGermanPhrase,
		language.Italian:              sallyItalianPhrase,
		language.Japanese:             sallyJapanesePhrase,
		language.Korean:               sallyKoreanPhrase,
		language.LatinAmericanSpanish: sallyLatinAmericanSpanishPhrase,
		language.Russian:              sallyRussianPhrase,
		language.SimplifiedChinese:    sallySimplifiedChinesePhrase,
		language.Spanish:              sallySpanishPhrase,
		language.TraditionalChinese:   sallyTraditionalChinesePhrase}
)

var (
	// Sally represents the character Sally with her complete information.
	Sally = nook.Villager{
		Character:   sallyCharacter,
		Personality: personality.Normal,
		Phrase:      sallyPhrase}
)
