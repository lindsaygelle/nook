package goat

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Chavrinamon moulin"}

	pashminaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pashminabokkig"}

	pashminaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chavrinamon moulin"}

	pashminaGermanName = nook.Name{
		Language: language.German,
		Value:    "Pamelaheckmeck"}

	pashminaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pashminaembeeeh"}

	pashminaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "バーバラかな"}

	pashminaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "바바라그러게"}

	pashminaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Carlane-eh"}

	pashminaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пашминакозлятки"}

	pashminaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "芭芭拉或许芭"}

	pashminaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Carlamuy bieeen"}

	pashminaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "芭芭拉或許芭"}
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
		Animal:   Goat,
		Birthday: pashminaBirthday,
		Code:     pashminaCode,
		Gender:   nook.Female,
		Name:     pashminaName}
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
	Pashmina = nook.Villager{
		Character:   pashminaCharacter,
		Personality: nook.BigSister,
		Phrase:      pashminaPhrase}
)
