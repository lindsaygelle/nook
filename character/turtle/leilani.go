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
	// leilaniBirthday represents leilani birthday.
	leilaniBirthday = nook.Birthday{
		Day:   26,
		Month: time.September}
)

var (
	// leilaniCode represents leilani code.
	leilaniCode = nook.Code{
		Value: "kpm"}
)

var (
	// leilaniAmericanEnglishName represents leilani american english name.
	leilaniAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Leilani"}

	// leilaniCanadianFrenchName represents leilani canadian french name.
	leilaniCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Liliane"}

	// leilaniDutchName represents leilani dutch name.
	leilaniDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Leilani"}

	// leilaniFrenchName represents leilani french name.
	leilaniFrenchName = nook.Name{
		Language: language.French,
		Value:    "Liliane"}

	// leilaniGermanName represents leilani german name.
	leilaniGermanName = nook.Name{
		Language: language.German,
		Value:    "Lore"}

	// leilaniItalianName represents leilani italian name.
	leilaniItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ramona"}

	// leilaniJapaneseName represents leilani japanese name.
	leilaniJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クーコ"}

	// leilaniKoreanName represents leilani korean name.
	leilaniKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "갑순"}

	// leilaniLatinAmericanSpanishName represents leilani latin american spanish name.
	leilaniLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Marimar"}

	// leilaniRussianName represents leilani russian name.
	leilaniRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лейлани"}

	// leilaniSimplifiedChineseName represents leilani simplified chinese name.
	leilaniSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "沽玉"}

	// leilaniSpanishName represents leilani spanish name.
	leilaniSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marimar"}

	// leilaniTraditionalChineseName represents leilani traditional chinese name.
	leilaniTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "沽玉"}
)

var (
	// leilaniName represents leilani name.
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
	// leilaniCharacter represents leilani character.
	leilaniCharacter = nook.Character{
		Animal:   animal.Turtle,
		Birthday: leilaniBirthday,
		Code:     leilaniCode,
		Key:      character.Leilani,
		Gender:   gender.Female,
		Name:     leilaniName,
		Special:  true}
)

var (
	// Leilani represents leilani.
	Leilani = nook.Resident{
		Character: leilaniCharacter}
)
