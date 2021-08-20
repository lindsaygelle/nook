package goat

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
	pashminaBirthday = nook.Birthday{
		Day:   26,
		Month: time.December}
)

var (
	pashminaCode = nook.Code{
		Value: "goa08"}
)

var (
	pashminaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pashmina"}

	pashminaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chavrina"}

	pashminaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pashmina"}

	pashminaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chavrina"}

	pashminaGermanName = nook.Name{
		Language: language.German,
		Value:    "Pamela"}

	pashminaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pashmina"}

	pashminaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "バーバラ"}

	pashminaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "바바라"}

	pashminaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Carla"}

	pashminaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пашмина"}

	pashminaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "芭芭拉"}

	pashminaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Carla"}

	pashminaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "芭芭拉"}
)

var (
	pashminaName = nook.Languages{
		language.AmericanEnglish:      pashminaAmericanEnglishName,
		language.CanadianFrench:       pashminaCanadianFrenchName,
		language.Dutch:                pashminaDutchName,
		language.French:               pashminaFrenchName,
		language.German:               pashminaGermanName,
		language.Italian:              pashminaItalianName,
		language.Japanese:             pashminaJapaneseName,
		language.Korean:               pashminaKoreanName,
		language.LatinAmericanSpanish: pashminaLatinAmericanSpanishName,
		language.Russian:              pashminaRussianName,
		language.SimplifiedChinese:    pashminaSimplifiedChineseName,
		language.Spanish:              pashminaSpanishName,
		language.TraditionalChinese:   pashminaTraditionalChineseName}
)

var (
	pashminaCharacter = nook.Character{
		Animal:   animal.Goat,
		Birthday: pashminaBirthday,
		Code:     pashminaCode,
		Key:      character.Pashmina,
		Gender:   gender.Female,
		Name:     pashminaName,
		Special:  false}
)

var (
	pashminaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "kidders"}

	pashminaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon moulin"}

	pashminaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "bokkig"}

	pashminaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon moulin"}

	pashminaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "heckmeck"}

	pashminaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "embeeeh"}

	pashminaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "かな"}

	pashminaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그러게"}

	pashminaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ne-eh"}

	pashminaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "козлятки"}

	pashminaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "或许芭"}

	pashminaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "muy bieeen"}

	pashminaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "或許芭"}
)

var (
	pashminaPhrase = nook.Languages{
		language.AmericanEnglish:      pashminaAmericanEnglishPhrase,
		language.CanadianFrench:       pashminaCanadianFrenchPhrase,
		language.Dutch:                pashminaDutchPhrase,
		language.French:               pashminaFrenchPhrase,
		language.German:               pashminaGermanPhrase,
		language.Italian:              pashminaItalianPhrase,
		language.Japanese:             pashminaJapanesePhrase,
		language.Korean:               pashminaKoreanPhrase,
		language.LatinAmericanSpanish: pashminaLatinAmericanSpanishPhrase,
		language.Russian:              pashminaRussianPhrase,
		language.SimplifiedChinese:    pashminaSimplifiedChinesePhrase,
		language.Spanish:              pashminaSpanishPhrase,
		language.TraditionalChinese:   pashminaTraditionalChinesePhrase}
)

var (
	Pashmina = nook.Villager{
		Character:   pashminaCharacter,
		Personality: personality.BigSister,
		Phrase:      pashminaPhrase}
)
