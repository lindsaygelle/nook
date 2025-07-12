package beaver

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// cjBirthday represents cj birthday.
	cjBirthday = nook.Birthday{
		Day:   7,
		Month: time.March}
)

var (
	// cjCode represents cj code.
	cjCode = nook.Code{
		Value: "bey"}
)

var (
	// cjAmericanEnglishName represents cj american english name.
	cjAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "C.J."}

	// cjCanadianFrenchName represents cj canadian french name.
	cjCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pollux"}

	// cjDutchName represents cj dutch name.
	cjDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "C.J."}

	// cjFrenchName represents cj french name.
	cjFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pollux"}

	// cjGermanName represents cj german name.
	cjGermanName = nook.Name{
		Language: language.German,
		Value:    "Lomeus"}

	// cjItalianName represents cj italian name.
	cjItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Castorino"}

	// cjJapaneseName represents cj japanese name.
	cjJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジャスティン"}

	// cjKoreanName represents cj korean name.
	cjKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "저스틴"}

	// cjLatinAmericanSpanishName represents cj latin american spanish name.
	cjLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "CJ"}

	// cjRussianName represents cj russian name.
	cjRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Си-Джей"}

	// cjSimplifiedChineseName represents cj simplified chinese name.
	cjSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "俞司廷"}

	// cjSpanishName represents cj spanish name.
	cjSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "CJ"}

	// cjTraditionalChineseName represents cj traditional chinese name.
	cjTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "俞司廷"}
)

var (
	// cjName represents cj name.
	cjName = nook.Languages{
		language.AmericanEnglish:      cjAmericanEnglishName,
		language.CanadianFrench:       cjCanadianFrenchName,
		language.Dutch:                cjDutchName,
		language.French:               cjFrenchName,
		language.German:               cjGermanName,
		language.Italian:              cjItalianName,
		language.Japanese:             cjJapaneseName,
		language.Korean:               cjKoreanName,
		language.LatinAmericanSpanish: cjLatinAmericanSpanishName,
		language.Russian:              cjRussianName,
		language.SimplifiedChinese:    cjSimplifiedChineseName,
		language.Spanish:              cjSpanishName,
		language.TraditionalChinese:   cjTraditionalChineseName}
)

var (
	// cjCharacter represents cj character.
	cjCharacter = nook.Character{
		Animal:   animal.Beaver,
		Birthday: cjBirthday,
		Code:     cjCode,
		Key:      character.CJ,
		Gender:   gender.Male,
		Name:     cjName,
		Special:  true}
)

var (
	// CJ represents c j.
	CJ = nook.Resident{
		Character: cjCharacter}
)
