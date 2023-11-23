package rabbit

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
	// tobyBirthday represents Toby's birthday.
	tobyBirthday = nook.Birthday{
		Day:   10,
		Month: time.July}
)

var (
	// tobyCode represents Toby's unique code ("rbt20").
	tobyCode = nook.Code{
		Value: "rbt20"}
)

var (
	// tobyAmericanEnglishName represents Toby's name in American English.
	tobyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Toby"}

	// tobyCanadianFrenchName represents Toby's name in Canadian French.
	tobyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Toby"}

	// tobyDutchName represents Toby's name in Dutch.
	tobyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Toby"}

	// tobyFrenchName represents Toby's name in French.
	tobyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Toby"}

	// tobyGermanName represents Toby's name in German.
	tobyGermanName = nook.Name{
		Language: language.German,
		Value:    "Toby"}

	// tobyItalianName represents Toby's name in Italian.
	tobyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Toby"}

	// tobyJapaneseName represents Toby's name in Japanese.
	tobyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トビー"}

	// tobyKoreanName represents Toby's name in Korean.
	tobyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "토비"}

	// tobyLatinAmericanSpanishName represents Toby's name in Latin American Spanish.
	tobyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Toby"}

	// tobyRussianName represents Toby's name in Russian.
	tobyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тоби"}

	// tobySimplifiedChineseName represents Toby's name in Simplified Chinese.
	tobySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咚比"}

	// tobySpanishName represents Toby's name in Spanish.
	tobySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Toby"}

	// tobyTraditionalChineseName represents Toby's name in Traditional Chinese.
	tobyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咚比"}
)

var (
	// tobyName represents Toby's name in different languages.
	tobyName = nook.Languages{
		language.AmericanEnglish:      tobyAmericanEnglishName,
		language.CanadianFrench:       tobyCanadianFrenchName,
		language.Dutch:                tobyDutchName,
		language.French:               tobyFrenchName,
		language.German:               tobyGermanName,
		language.Italian:              tobyItalianName,
		language.Japanese:             tobyJapaneseName,
		language.Korean:               tobyKoreanName,
		language.LatinAmericanSpanish: tobyLatinAmericanSpanishName,
		language.Russian:              tobyRussianName,
		language.SimplifiedChinese:    tobySimplifiedChineseName,
		language.Spanish:              tobySpanishName,
		language.TraditionalChinese:   tobyTraditionalChineseName}
)

var (
	// tobyCharacter represents Toby's character information.
	tobyCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: tobyBirthday,
		Code:     tobyCode,
		Key:      character.Toby,
		Gender:   gender.Male,
		Name:     tobyName,
		Special:  false}
)

var (
	// tobyAmericanEnglishPhrase represents Toby's phrase in American English.
	tobyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ribbit"}

	// tobyCanadianFrenchPhrase represents Toby's phrase in Canadian French.
	tobyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "kérocarot"}

	// tobyDutchPhrase represents Toby's phrase in Dutch.
	tobyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kikk'r"}

	// tobyFrenchPhrase represents Toby's phrase in French.
	tobyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "kérocarot"}

	// tobyGermanPhrase represents Toby's phrase in German.
	tobyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kerokero"}

	// tobyItalianPhrase represents Toby's phrase in Italian.
	tobyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "kerokero"}

	// tobyJapanesePhrase represents Toby's phrase in Japanese.
	tobyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "けろけろ"}

	// tobyKoreanPhrase represents Toby's phrase in Korean.
	tobyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "케로케로"}

	// tobyLatinAmericanSpanishPhrase represents Toby's phrase in Latin American Spanish.
	tobyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "keroppi"}

	// tobyRussianPhrase represents Toby's phrase in Russian.
	tobyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "зайцеквак"}

	// tobySimplifiedChinesePhrase represents Toby's phrase in Simplified Chinese.
	tobySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蛙蛙"}

	// tobySpanishPhrase represents Toby's phrase in Spanish.
	tobySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "keroppi"}

	// tobyTraditionalChinesePhrase represents Toby's phrase in Traditional Chinese.
	tobyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蛙蛙"}
)

var (
	// tobyPhrase represents Toby's phrases in different languages.
	tobyPhrase = nook.Languages{
		language.AmericanEnglish:      tobyAmericanEnglishPhrase,
		language.CanadianFrench:       tobyCanadianFrenchPhrase,
		language.Dutch:                tobyDutchPhrase,
		language.French:               tobyFrenchPhrase,
		language.German:               tobyGermanPhrase,
		language.Italian:              tobyItalianPhrase,
		language.Japanese:             tobyJapanesePhrase,
		language.Korean:               tobyKoreanPhrase,
		language.LatinAmericanSpanish: tobyLatinAmericanSpanishPhrase,
		language.Russian:              tobyRussianPhrase,
		language.SimplifiedChinese:    tobySimplifiedChinesePhrase,
		language.Spanish:              tobySpanishPhrase,
		language.TraditionalChinese:   tobyTraditionalChinesePhrase}
)

var (
	// Toby represents the character Toby with his complete information.
	Toby = nook.Villager{
		Character:   tobyCharacter,
		Personality: personality.Smug,
		Phrase:      tobyPhrase}
)
