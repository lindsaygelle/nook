package turtle

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	leilaniBirthday = nook.Birthday{
		Day:   26,
		Month: time.September}
)

var (
	leilaniCode = nook.Code{
		Value: "kpm"}
)

var (
	leilaniAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Leilani"}

	leilaniCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Liliane"}

	leilaniDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Leilani"}

	leilaniFrenchName = nook.Name{
		Language: language.French,
		Value:    "Liliane"}

	leilaniGermanName = nook.Name{
		Language: language.German,
		Value:    "Lore"}

	leilaniItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ramona"}

	leilaniJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クーコ"}

	leilaniKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "갑순"}

	leilaniLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Marimar"}

	leilaniRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лейлани"}

	leilaniSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "沽玉"}

	leilaniSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marimar"}

	leilaniTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "沽玉"}
)

var (
	leilaniName = nook.Languages{
		language.AmericanEnglish:      leilaniAmericanEnglishName,
		language.CanadianFrench:       leilaniCanadianFrenchName,
		language.Dutch:                leilaniDutchName,
		language.French:               leilaniFrenchName,
		language.German:               leilaniGermanName,
		language.Italian:              leilaniItalianName,
		language.Japanese:             leilaniJapaneseName,
		language.Korean:               leilaniKoreanName,
		language.LatinAmericanSpanish: leilaniLatinAmericanSpanishName,
		language.Russian:              leilaniRussianName,
		language.SimplifiedChinese:    leilaniSimplifiedChineseName,
		language.Spanish:              leilaniSpanishName,
		language.TraditionalChinese:   leilaniTraditionalChineseName}
)

var (
	leilaniCharacter = nook.Character{
		Animal:   Turtle,
		Birthday: leilaniBirthday,
		Code:     leilaniCode,
		Gender:   nook.Female,
		Name:     leilaniName}
)

var (
	Leilani = nook.Resident{
		Character: leilaniCharacter}
)
