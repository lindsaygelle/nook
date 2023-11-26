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
	// nibblesBirthday represents Nibbles's birthday.
	nibblesBirthday = nook.Birthday{
		Day:   19,
		Month: time.July}
)

var (
	// nibblesCode represents Nibbles's unique code.
	nibblesCode = nook.Code{
		Value: "squ04"}
)

var (
	// nibblesAmericanEnglishName represents Nibbles's name in American English.
	nibblesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nibbles"}

	// nibblesCanadianFrenchName represents Nibbles's name in Canadian French.
	nibblesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lola"}

	// nibblesDutchName represents Nibbles's name in Dutch.
	nibblesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Nibbles"}

	// nibblesFrenchName represents Nibbles's name in French.
	nibblesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lola"}

	// nibblesGermanName represents Nibbles's name in German.
	nibblesGermanName = nook.Name{
		Language: language.German,
		Value:    "Knuspi"}

	// nibblesItalianName represents Nibbles's name in Italian.
	nibblesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rododea"}

	// nibblesJapaneseName represents Nibbles's name in Japanese.
	nibblesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ガリガリ"}

	// nibblesKoreanName represents Nibbles's name in Korean.
	nibblesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "그리미"}

	// nibblesLatinAmericanSpanishName represents Nibbles's name in Latin American Spanish.
	nibblesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Dentina"}

	// nibblesRussianName represents Nibbles's name in Russian.
	nibblesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Ниблз"}

	// nibblesSimplifiedChineseName represents Nibbles's name in Simplified Chinese.
	nibblesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咬咬"}

	// nibblesSpanishName represents Nibbles's name in Spanish.
	nibblesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Dentina"}

	// nibblesTraditionalChineseName represents Nibbles's name in Traditional Chinese.
	nibblesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咬咬"}
)

var (
	// nibblesName represents Nibbles's name in different languages.
	nibblesName = nook.Languages{
		language.AmericanEnglish:      nibblesAmericanEnglishName,
		language.CanadianFrench:       nibblesCanadianFrenchName,
		language.Dutch:                nibblesDutchName,
		language.French:               nibblesFrenchName,
		language.German:               nibblesGermanName,
		language.Italian:              nibblesItalianName,
		language.Japanese:             nibblesJapaneseName,
		language.Korean:               nibblesKoreanName,
		language.LatinAmericanSpanish: nibblesLatinAmericanSpanishName,
		language.Russian:              nibblesRussianName,
		language.SimplifiedChinese:    nibblesSimplifiedChineseName,
		language.Spanish:              nibblesSpanishName,
		language.TraditionalChinese:   nibblesTraditionalChineseName}
)

var (
	// nibblesCharacter represents Nibbles's character information.
	nibblesCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: nibblesBirthday,
		Code:     nibblesCode,
		Key:      character.Nibbles,
		Gender:   gender.Female,
		Name:     nibblesName,
		Special:  false}
)

var (
	// nibblesAmericanEnglishPhrase represents Nibbles's phrase in American English.
	nibblesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "niblet"}

	// nibblesCanadianFrenchPhrase represents Nibbles's phrase in Canadian French.
	nibblesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grignotte"}

	// nibblesDutchPhrase represents Nibbles's phrase in Dutch.
	nibblesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "knabbel"}

	// nibblesFrenchPhrase represents Nibbles's phrase in French.
	nibblesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "nunuche"}

	// nibblesGermanPhrase represents Nibbles's phrase in German.
	nibblesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knusper"}

	// nibblesItalianPhrase represents Nibbles's phrase in Italian.
	nibblesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "truciolo"}

	// nibblesJapanesePhrase represents Nibbles's phrase in Japanese.
	nibblesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ガジガジ"}

	// nibblesKoreanPhrase represents Nibbles's phrase in Korean.
	nibblesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "송송"}

	// nibblesLatinAmericanSpanishPhrase represents Nibbles's phrase in Latin American Spanish.
	nibblesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "trocitrí"}

	// nibblesRussianPhrase represents Nibbles's phrase in Russian.
	nibblesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "орешек"}

	// nibblesSimplifiedChinesePhrase represents Nibbles's phrase in Simplified Chinese.
	nibblesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咯吱咯吱"}

	// nibblesSpanishPhrase represents Nibbles's phrase in Spanish.
	nibblesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bellotita"}

	// nibblesTraditionalChinesePhrase represents Nibbles's phrase in Traditional Chinese.
	nibblesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咯吱咯吱"}
)

var (
	// nibblesPhrase represents Nibbles's phrase in different languages.
	nibblesPhrase = nook.Languages{
		language.AmericanEnglish:      nibblesAmericanEnglishPhrase,
		language.CanadianFrench:       nibblesCanadianFrenchPhrase,
		language.Dutch:                nibblesDutchPhrase,
		language.French:               nibblesFrenchPhrase,
		language.German:               nibblesGermanPhrase,
		language.Italian:              nibblesItalianPhrase,
		language.Japanese:             nibblesJapanesePhrase,
		language.Korean:               nibblesKoreanPhrase,
		language.LatinAmericanSpanish: nibblesLatinAmericanSpanishPhrase,
		language.Russian:              nibblesRussianPhrase,
		language.SimplifiedChinese:    nibblesSimplifiedChinesePhrase,
		language.Spanish:              nibblesSpanishPhrase,
		language.TraditionalChinese:   nibblesTraditionalChinesePhrase}
)

var (
	// Nibbles represents the character Nibbles with her complete information.
	Nibbles = nook.Villager{
		Character:   nibblesCharacter,
		Personality: personality.Peppy,
		Phrase:      nibblesPhrase}
)
