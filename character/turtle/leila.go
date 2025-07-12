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
	// leilaBirthday represents leila birthday.
	leilaBirthday = nook.Birthday{
		Day:   16,
		Month: time.August}
)

var (
	// leilaCode represents leila code.
	leilaCode = nook.Code{
		Value: "kps"}
)

var (
	// leilaAmericanEnglishName represents leila american english name.
	leilaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Leila"}

	// leilaCanadianFrenchName represents leila canadian french name.
	leilaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lila"}

	// leilaDutchName represents leila dutch name.
	leilaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Leila"}

	// leilaFrenchName represents leila french name.
	leilaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lila"}

	// leilaGermanName represents leila german name.
	leilaGermanName = nook.Name{
		Language: language.German,
		Value:    "Lotte"}

	// leilaItalianName represents leila italian name.
	leilaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Romina"}

	// leilaJapaneseName represents leila japanese name.
	leilaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クク"}

	// leilaKoreanName represents leila korean name.
	leilaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "꼬미"}

	// leilaLatinAmericanSpanishName represents leila latin american spanish name.
	leilaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Conchita"}

	// leilaRussianName represents leila russian name.
	leilaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лейла"}

	// leilaSimplifiedChineseName represents leila simplified chinese name.
	leilaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "沽沽"}

	// leilaSpanishName represents leila spanish name.
	leilaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Conchita"}

	// leilaTraditionalChineseName represents leila traditional chinese name.
	leilaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "沽沽"}
)

var (
	// leilaName represents leila name.
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
	// leilaCharacter represents leila character.
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
	// Leila represents leila.
	Leila = nook.Resident{
		Character: leilaCharacter}
)
