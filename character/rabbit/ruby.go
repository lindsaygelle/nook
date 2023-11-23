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
	// rubyBirthday represents Ruby's birthday.
	rubyBirthday = nook.Birthday{
		Day:   25,
		Month: time.December}
)

var (
	// rubyCode represents Ruby's unique code ("rbt09").
	rubyCode = nook.Code{
		Value: "rbt09"}
)

var (
	// rubyAmericanEnglishName represents Ruby's name in American English.
	rubyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ruby"}

	// rubyCanadianFrenchName represents Ruby's name in Canadian French.
	rubyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rubis"}

	// rubyDutchName represents Ruby's name in Dutch.
	rubyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ruby"}

	// rubyFrenchName represents Ruby's name in French.
	rubyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rubis"}

	// rubyGermanName represents Ruby's name in German.
	rubyGermanName = nook.Name{
		Language: language.German,
		Value:    "Rubina"}

	// rubyItalianName represents Ruby's name in Italian.
	rubyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rubina"}

	// rubyJapaneseName represents Ruby's name in Japanese.
	rubyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ルナ"}

	// rubyKoreanName represents Ruby's name in Korean.
	rubyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "루나"}

	// rubyLatinAmericanSpanishName represents Ruby's name in Latin American Spanish.
	rubyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rubí"}

	// rubyRussianName represents Ruby's name in Russian.
	rubyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Руби"}

	// rubySimplifiedChineseName represents Ruby's name in Simplified Chinese.
	rubySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "月兔"}

	// rubySpanishName represents Ruby's name in Spanish.
	rubySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rubí"}

	// rubyTraditionalChineseName represents Ruby's name in Traditional Chinese.
	rubyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "月兔"}
)

var (
	// rubyName represents Ruby's name in different languages.
	rubyName = nook.Languages{
		language.AmericanEnglish:      rubyAmericanEnglishName,
		language.CanadianFrench:       rubyCanadianFrenchName,
		language.Dutch:                rubyDutchName,
		language.French:               rubyFrenchName,
		language.German:               rubyGermanName,
		language.Italian:              rubyItalianName,
		language.Japanese:             rubyJapaneseName,
		language.Korean:               rubyKoreanName,
		language.LatinAmericanSpanish: rubyLatinAmericanSpanishName,
		language.Russian:              rubyRussianName,
		language.SimplifiedChinese:    rubySimplifiedChineseName,
		language.Spanish:              rubySpanishName,
		language.TraditionalChinese:   rubyTraditionalChineseName}
)

var (
	// rubyCharacter represents Ruby's character information.
	rubyCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: rubyBirthday,
		Code:     rubyCode,
		Key:      character.Ruby,
		Gender:   gender.Female,
		Name:     rubyName,
		Special:  false}
)

var (
	// rubyAmericanEnglishPhrase represents Ruby's phrase in American English.
	rubyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "li'l ears"}

	// rubyCanadianFrenchPhrase represents Ruby's phrase in Canadian French.
	rubyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "n'oreille"}

	// rubyDutchPhrase represents Ruby's phrase in Dutch.
	rubyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hangoor"}

	// rubyFrenchPhrase represents Ruby's phrase in French.
	rubyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "n'oreille"}

	// rubyGermanPhrase represents Ruby's phrase in German.
	rubyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "paffpaff"}

	// rubyItalianPhrase represents Ruby's phrase in Italian.
	rubyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "orecchine"}

	// rubyJapanesePhrase represents Ruby's phrase in Japanese.
	rubyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "フツーに"}

	// rubyKoreanPhrase represents Ruby's phrase in Korean.
	rubyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "보통이지"}

	// rubyLatinAmericanSpanishPhrase represents Ruby's phrase in Latin American Spanish.
	rubyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "orejí"}

	// rubyRussianPhrase represents Ruby's phrase in Russian.
	rubyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ушки"}

	// rubySimplifiedChinesePhrase represents Ruby's phrase in Simplified Chinese.
	rubySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "普通"}

	// rubySpanishPhrase represents Ruby's phrase in Spanish.
	rubySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "orejitas"}

	// rubyTraditionalChinesePhrase represents Ruby's phrase in Traditional Chinese.
	rubyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "普通"}
)

var (
	// rubyPhrase represents Ruby's phrases in different languages.
	rubyPhrase = nook.Languages{
		language.AmericanEnglish:      rubyAmericanEnglishPhrase,
		language.CanadianFrench:       rubyCanadianFrenchPhrase,
		language.Dutch:                rubyDutchPhrase,
		language.French:               rubyFrenchPhrase,
		language.German:               rubyGermanPhrase,
		language.Italian:              rubyItalianPhrase,
		language.Japanese:             rubyJapanesePhrase,
		language.Korean:               rubyKoreanPhrase,
		language.LatinAmericanSpanish: rubyLatinAmericanSpanishPhrase,
		language.Russian:              rubyRussianPhrase,
		language.SimplifiedChinese:    rubySimplifiedChinesePhrase,
		language.Spanish:              rubySpanishPhrase,
		language.TraditionalChinese:   rubyTraditionalChinesePhrase}
)

var (
	// Ruby represents the character Ruby with her complete information.
	Ruby = nook.Villager{
		Character:   rubyCharacter,
		Personality: personality.Peppy,
		Phrase:      rubyPhrase}
)
