package sheep

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
	// genBirthday represents Gen's birthday.
	genBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// genCode represents Gen's unique code.
	genCode = nook.Code{
		Value: ""}
)

var (
	// genAmericanEnglishName represents Gen's name in American English.
	genAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gen"}

	// genAmericanEnglishName represents Gen's name in Canadian French.
	genCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// genAmericanEnglishName represents Gen's name in Dutch.
	genDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// genAmericanEnglishName represents Gen's name in French.
	genFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// genAmericanEnglishName represents Gen's name in German.
	genGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// genAmericanEnglishName represents Gen's name in Italian.
	genItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// genAmericanEnglishName represents Gen's name in Japanese.
	genJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゲン"}

	// genAmericanEnglishName represents Gen's name in Korean.
	genKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// genAmericanEnglishName represents Gen's name in Latin American Spanish.
	genLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// genAmericanEnglishName represents Gen's name in Russian.
	genRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// genAmericanEnglishName represents Gen's name in Simplified Chinese.
	genSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// genAmericanEnglishName represents Gen's name in Spanish.
	genSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// genAmericanEnglishName represents Gen's name in Traditional Chinese.
	genTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// genName represents Gen's name in different languages.
	genName = nook.Languages{
		language.AmericanEnglish:      genAmericanEnglishName,
		language.CanadianFrench:       genCanadianFrenchName,
		language.Dutch:                genDutchName,
		language.French:               genFrenchName,
		language.German:               genGermanName,
		language.Italian:              genItalianName,
		language.Japanese:             genJapaneseName,
		language.Korean:               genKoreanName,
		language.LatinAmericanSpanish: genLatinAmericanSpanishName,
		language.Russian:              genRussianName,
		language.SimplifiedChinese:    genSimplifiedChineseName,
		language.Spanish:              genSpanishName,
		language.TraditionalChinese:   genTraditionalChineseName}
)

var (
	// genCharacter represents Gen's character information.
	genCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: genBirthday,
		Code:     genCode,
		Key:      character.Gen,
		Gender:   gender.Male,
		Name:     genName,
		Special:  false}
)

var (
	// genCanadianFrenchPhrase represents Gen's phrase in American English (empty string for now).
	genAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	// genCanadianFrenchPhrase represents Gen's phrase in Canadian French (empty string for now).
	genCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// genCanadianFrenchPhrase represents Gen's phrase in Dutch (empty string for now).
	genDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// genCanadianFrenchPhrase represents Gen's phrase in French (empty string for now).
	genFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// genCanadianFrenchPhrase represents Gen's phrase in German (empty string for now).
	genGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// genCanadianFrenchPhrase represents Gen's phrase in Italian (empty string for now).
	genItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// genCanadianFrenchPhrase represents Gen's phrase in Japanese.
	genJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "っしゃあ"}

	// genCanadianFrenchPhrase represents Gen's phrase in Korean (empty string for now).
	genKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// genCanadianFrenchPhrase represents Gen's phrase in Latin American Spanish (empty string for now).
	genLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// genCanadianFrenchPhrase represents Gen's phrase in Russian (empty string for now).
	genRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// genCanadianFrenchPhrase represents Gen's phrase in Simplified Chinese (empty string for now).
	genSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// genCanadianFrenchPhrase represents Gen's phrase in Spanish (empty string for now).
	genSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// genCanadianFrenchPhrase represents Gen's phrase in Traditional Chinese (empty string for now).
	genTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// genPhrase represents Gen's phrases in different languages.
	genPhrase = nook.Languages{
		language.AmericanEnglish:      genAmericanEnglishPhrase,
		language.CanadianFrench:       genCanadianFrenchPhrase,
		language.Dutch:                genDutchPhrase,
		language.French:               genFrenchPhrase,
		language.German:               genGermanPhrase,
		language.Italian:              genItalianPhrase,
		language.Japanese:             genJapanesePhrase,
		language.Korean:               genKoreanPhrase,
		language.LatinAmericanSpanish: genLatinAmericanSpanishPhrase,
		language.Russian:              genRussianPhrase,
		language.SimplifiedChinese:    genSimplifiedChinesePhrase,
		language.Spanish:              genSpanishPhrase,
		language.TraditionalChinese:   genTraditionalChinesePhrase}
)

var (
	// Gen represents the character Gen with his complete information.
	Gen = nook.Villager{
		Character:   genCharacter,
		Personality: personality.Jock,
		Phrase:      genPhrase}
)
