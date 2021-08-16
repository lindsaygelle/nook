package pelican

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	peteBirthday = nook.Birthday{
		Day:   8,
		Month: time.March}
)

var (
	peteCode = nook.Code{
		Value: "plb/plo"}
)

var (
	peteAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pete"}

	peteCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Antoine"}

	peteDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pete"}

	peteFrenchName = nook.Name{
		Language: language.French,
		Value:    "Antoine"}

	peteGermanName = nook.Name{
		Language: language.German,
		Value:    "Peter"}

	peteItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tino"}

	peteJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ぺりお"}

	peteKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "펠리오"}

	peteLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Carturo"}

	peteRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пит"}

	peteSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "程信雄"}

	peteSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Carturo"}

	peteTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "程信雄"}
)

var (
	peteName = nook.Languages{
		language.AmericanEnglish:      peteAmericanEnglishName,
		language.CanadianFrench:       peteCanadianFrenchName,
		language.Dutch:                peteDutchName,
		language.French:               peteFrenchName,
		language.German:               peteGermanName,
		language.Italian:              peteItalianName,
		language.Japanese:             peteJapaneseName,
		language.Korean:               peteKoreanName,
		language.LatinAmericanSpanish: peteLatinAmericanSpanishName,
		language.Russian:              peteRussianName,
		language.SimplifiedChinese:    peteSimplifiedChineseName,
		language.Spanish:              peteSpanishName,
		language.TraditionalChinese:   peteTraditionalChineseName}
)

var (
	peteCharacter = nook.Character{
		Animal:   animal.Pelican,
		Birthday: peteBirthday,
		Code:     peteCode,
		Gender:   gender.Male,
		Name:     peteName}
)

var (
	Pete = nook.Resident{
		Character: peteCharacter}
)
