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
	poppyBirthday = nook.Birthday{
		Day:   5,
		Month: time.August}
)

var (
	poppyCode = nook.Code{
		Value: "squ15"}
)

var (
	poppyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Poppy"}

	poppyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Irène"}

	poppyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Poppy"}

	poppyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Irène"}

	poppyGermanName = nook.Name{
		Language: language.German,
		Value:    "Trita"}

	poppyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Granella"}

	poppyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "グミ"}

	poppyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "다람"}

	poppyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Encina"}

	poppyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Поппи"}

	poppySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "软糖"}

	poppySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Encina"}

	poppyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "軟糖"}
)

var (
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
	poppyCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: poppyBirthday,
		Code:     poppyCode,
		Key:      character.Poppy,
		Gender:   gender.Female,
		Name:     poppyName}
)

var (
	poppyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nutty"}

	poppyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "abajoue"}

	poppyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "nootje"}

	poppyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "scratchy"}

	poppyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "nussi"}

	poppyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "squittole"}

	poppyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "デス"}

	poppyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "도톨도톨"}

	poppyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "nutinú"}

	poppyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "чок-чок"}

	poppySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "甜甜"}

	poppySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "albricias"}

	poppyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "甜甜"}
)

var (
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
	Poppy = nook.Villager{
		Character:   poppyCharacter,
		Personality: personality.Normal,
		Phrase:      poppyPhrase}
)
