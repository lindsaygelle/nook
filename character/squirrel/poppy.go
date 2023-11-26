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
	// poppyBirthday represents Poppy's birthday.
	poppyBirthday = nook.Birthday{
		Day:   5,
		Month: time.August}
)

var (
	// poppyCode represents Poppy's unique code.
	poppyCode = nook.Code{
		Value: "squ15"}
)

var (
	// poppyAmericanEnglishName represents Poppy's name in American English.
	poppyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Poppy"}

	// poppyCanadianFrenchName represents Poppy's name in Canadian French.
	poppyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Irène"}

	// poppyDutchName represents Poppy's name in Dutch.
	poppyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Poppy"}

	// poppyFrenchName represents Poppy's name in French.
	poppyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Irène"}

	// poppyGermanName represents Poppy's name in German.
	poppyGermanName = nook.Name{
		Language: language.German,
		Value:    "Trita"}

	// poppyItalianName represents Poppy's name in Italian.
	poppyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Granella"}

	// poppyJapaneseName represents Poppy's name in Japanese.
	poppyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "グミ"}

	// poppyKoreanName represents Poppy's name in Korean.
	poppyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "다람"}

	// poppyLatinAmericanSpanishName represents Poppy's name in Latin American Spanish.
	poppyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Encina"}

	// poppyRussianName represents Poppy's name in Russian.
	poppyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Поппи"}

	// poppySimplifiedChineseName represents Poppy's name in Simplified Chinese.
	poppySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "软糖"}

	// poppySpanishName represents Poppy's name in Spanish.
	poppySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Encina"}

	// poppyTraditionalChineseName represents Poppy's name in Traditional Chinese.
	poppyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "軟糖"}
)

var (
	// poppyName represents Poppy's name in different languages.
	poppyName = nook.Languages{
		language.AmericanEnglish:      poppyAmericanEnglishName,
		language.CanadianFrench:       poppyCanadianFrenchName,
		language.Dutch:                poppyDutchName,
		language.French:               poppyFrenchName,
		language.German:               poppyGermanName,
		language.Italian:              poppyItalianName,
		language.Japanese:             poppyJapaneseName,
		language.Korean:               poppyKoreanName,
		language.LatinAmericanSpanish: poppyLatinAmericanSpanishName,
		language.Russian:              poppyRussianName,
		language.SimplifiedChinese:    poppySimplifiedChineseName,
		language.Spanish:              poppySpanishName,
		language.TraditionalChinese:   poppyTraditionalChineseName}
)

var (
	// poppyCharacter represents Poppy's character information.
	poppyCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: poppyBirthday,
		Code:     poppyCode,
		Key:      character.Poppy,
		Gender:   gender.Female,
		Name:     poppyName,
		Special:  false}
)

var (
	// poppyAmericanEnglishPhrase represents Poppy's phrase in American English.
	poppyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nutty"}

	// poppyCanadianFrenchPhrase represents Poppy's phrase in Canadian French.
	poppyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "abajoue"}

	// poppyDutchPhrase represents Poppy's phrase in Dutch.
	poppyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "nootje"}

	// poppyFrenchPhrase represents Poppy's phrase in French.
	poppyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "scratchy"}

	// poppyGermanPhrase represents Poppy's phrase in German.
	poppyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "nussi"}

	// poppyItalianPhrase represents Poppy's phrase in Italian.
	poppyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "squittole"}

	// poppyJapanesePhrase represents Poppy's phrase in Japanese.
	poppyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "デス"}

	// poppyKoreanPhrase represents Poppy's phrase in Korean.
	poppyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "도톨도톨"}

	// poppyLatinAmericanSpanishPhrase represents Poppy's phrase in Latin American Spanish.
	poppyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "nutinú"}

	// poppyRussianPhrase represents Poppy's phrase in Russian.
	poppyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "чок-чок"}

	// poppySimplifiedChinesePhrase represents Poppy's phrase in Simplified Chinese.
	poppySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "甜甜"}

	// poppySpanishPhrase represents Poppy's phrase in Spanish.
	poppySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "albricias"}

	// poppyTraditionalChinesePhrase represents Poppy's phrase in Traditional Chinese.
	poppyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "甜甜"}
)

var (
	// poppyPhrase represents Poppy's phrases in different languages.
	poppyPhrase = nook.Languages{
		language.AmericanEnglish:      poppyAmericanEnglishPhrase,
		language.CanadianFrench:       poppyCanadianFrenchPhrase,
		language.Dutch:                poppyDutchPhrase,
		language.French:               poppyFrenchPhrase,
		language.German:               poppyGermanPhrase,
		language.Italian:              poppyItalianPhrase,
		language.Japanese:             poppyJapanesePhrase,
		language.Korean:               poppyKoreanPhrase,
		language.LatinAmericanSpanish: poppyLatinAmericanSpanishPhrase,
		language.Russian:              poppyRussianPhrase,
		language.SimplifiedChinese:    poppySimplifiedChinesePhrase,
		language.Spanish:              poppySpanishPhrase,
		language.TraditionalChinese:   poppyTraditionalChinesePhrase}
)

var (
	// Poppy represents the character Poppy with her complete information.
	Poppy = nook.Villager{
		Character:   poppyCharacter,
		Personality: personality.Normal,
		Phrase:      poppyPhrase}
)
