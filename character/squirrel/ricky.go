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
	// rickyBirthday represents Ricky's birthday.
	rickyBirthday = nook.Birthday{
		Day:   14,
		Month: time.September}
)

var (
	// rickyCode represents Ricky's unique code.
	rickyCode = nook.Code{
		Value: "squ10"}
)

var (
	// rickyAmericanEnglishName represents Ricky's name in American English.
	rickyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ricky"}

	// rickyCanadianFrenchName represents Ricky's name in Canadian French.
	rickyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rocky"}

	// rickyDutchName represents Ricky's name in Dutch.
	rickyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ricky"}

	// rickyFrenchName represents Ricky's name in French.
	rickyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rocky"}

	// rickyGermanName represents Ricky's name in German.
	rickyGermanName = nook.Name{
		Language: language.German,
		Value:    "Ronny"}

	// rickyItalianName represents Ricky's name in Italian.
	rickyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sberlo"}

	// rickyJapaneseName represents Ricky's name in Japanese.
	rickyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カジロウ"}

	// rickyKoreanName represents Ricky's name in Korean.
	rickyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "갈가리"}

	// rickyLatinAmericanSpanishName represents Ricky's name in Latin American Spanish.
	rickyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Altramuz"}

	// rickyRussianName represents Ricky's name in Russian.
	rickyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рики"}

	// rickySimplifiedChineseName represents Ricky's name in Simplified Chinese.
	rickySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘉治"}

	// rickySpanishName represents Ricky's name in Spanish.
	rickySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Altramuz"}

	// rickyTraditionalChineseName represents Ricky's name in Traditional Chinese.
	rickyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘉治"}
)

var (
	// rickyName represents Ricky's name in different languages.
	rickyName = nook.Languages{
		language.AmericanEnglish:      rickyAmericanEnglishName,
		language.CanadianFrench:       rickyCanadianFrenchName,
		language.Dutch:                rickyDutchName,
		language.French:               rickyFrenchName,
		language.German:               rickyGermanName,
		language.Italian:              rickyItalianName,
		language.Japanese:             rickyJapaneseName,
		language.Korean:               rickyKoreanName,
		language.LatinAmericanSpanish: rickyLatinAmericanSpanishName,
		language.Russian:              rickyRussianName,
		language.SimplifiedChinese:    rickySimplifiedChineseName,
		language.Spanish:              rickySpanishName,
		language.TraditionalChinese:   rickyTraditionalChineseName}
)

var (
	// rickyCharacter represents Ricky's character information.
	rickyCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: rickyBirthday,
		Code:     rickyCode,
		Key:      character.Ricky,
		Gender:   gender.Male,
		Name:     rickyName,
		Special:  false}
)

var (
	// rickyAmericanEnglishPhrase represents Ricky's phrase in American English.
	rickyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nutcase"}

	// rickyCanadianFrenchPhrase represents Ricky's phrase in Canadian French.
	rickyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "dingo"}

	// rickyDutchPhrase represents Ricky's phrase in Dutch.
	rickyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mafnoot"}

	// rickyFrenchPhrase represents Ricky's phrase in French.
	rickyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "dingo"}

	// rickyGermanPhrase represents Ricky's phrase in German.
	rickyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knorz"}

	// rickyItalianPhrase represents Ricky's phrase in Italian.
	rickyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "nocciola"}

	// rickyJapanesePhrase represents Ricky's phrase in Japanese.
	rickyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "キッ"}

	// rickyKoreanPhrase represents Ricky's phrase in Korean.
	rickyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "사각사각"}

	// rickyLatinAmericanSpanishPhrase represents Ricky's phrase in Latin American Spanish.
	rickyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "crac-crac"}

	// rickyRussianPhrase represents Ricky's phrase in Russian.
	rickyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "арахис"}

	// rickySimplifiedChinesePhrase represents Ricky's phrase in Simplified Chinese.
	rickySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "炯炯"}

	// rickySpanishPhrase represents Ricky's phrase in Spanish.
	rickySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cáscaras"}

	// rickyTraditionalChinesePhrase represents Ricky's phrase in Traditional Chinese.
	rickyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "炯炯"}
)

var (
	// rickyPhrase represents Ricky's phrases in different languages.
	rickyPhrase = nook.Languages{
		language.AmericanEnglish:      rickyAmericanEnglishPhrase,
		language.CanadianFrench:       rickyCanadianFrenchPhrase,
		language.Dutch:                rickyDutchPhrase,
		language.French:               rickyFrenchPhrase,
		language.German:               rickyGermanPhrase,
		language.Italian:              rickyItalianPhrase,
		language.Japanese:             rickyJapanesePhrase,
		language.Korean:               rickyKoreanPhrase,
		language.LatinAmericanSpanish: rickyLatinAmericanSpanishPhrase,
		language.Russian:              rickyRussianPhrase,
		language.SimplifiedChinese:    rickySimplifiedChinesePhrase,
		language.Spanish:              rickySpanishName,
		language.TraditionalChinese:   rickyTraditionalChinesePhrase}
)

var (
	// Ricky represents the character Ricky with his complete information.
	Ricky = nook.Villager{
		Character:   rickyCharacter,
		Personality: personality.Cranky,
		Phrase:      rickyPhrase}
)
