package turtle

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	leilaBirthday = nook.Birthday{
		Day:   16,
		Month: time.August}
)

var (
	leilaCode = nook.Code{
		Value: "kps"}
)

var (
	leilaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Leila"}

	leilaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lila"}

	leilaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Leila"}

	leilaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lila"}

	leilaGermanName = nook.Name{
		Language: language.German,
		Value:    "Lotte"}

	leilaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Romina"}

	leilaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クク"}

	leilaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "꼬미"}

	leilaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Conchita"}

	leilaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лейла"}

	leilaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "沽沽"}

	leilaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Conchita"}

	leilaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "沽沽"}
)

var (
	leilaName = nook.Languages{
		language.AmericanEnglish:      leilaAmericanEnglishName,
		language.CanadianFrench:       leilaCanadianFrenchName,
		language.Dutch:                leilaDutchName,
		language.French:               leilaFrenchName,
		language.German:               leilaGermanName,
		language.Italian:              leilaItalianName,
		language.Japanese:             leilaJapaneseName,
		language.Korean:               leilaKoreanName,
		language.LatinAmericanSpanish: leilaLatinAmericanSpanishName,
		language.Russian:              leilaRussianName,
		language.SimplifiedChinese:    leilaSimplifiedChineseName,
		language.Spanish:              leilaSpanishName,
		language.TraditionalChinese:   leilaTraditionalChineseName}
)

var (
	leilaCharacter = nook.Character{
		Animal:   animal.Turtle,
		Birthday: leilaBirthday,
		Code:     leilaCode,
		Key:      character.Leila,
		Gender:   gender.Female,
		Name:     leilaName,
		Special:  true}
)

var (
	Leila = nook.Resident{
		Character: leilaCharacter}
)
