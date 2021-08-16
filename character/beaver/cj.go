package beaver

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	cjBirthday = nook.Birthday{
		Day:   7,
		Month: time.March}
)

var (
	cjCode = nook.Code{
		Value: "bey"}
)

var (
	cjAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "C.J."}

	cjCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pollux"}

	cjDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "C.J."}

	cjFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pollux"}

	cjGermanName = nook.Name{
		Language: language.German,
		Value:    "Lomeus"}

	cjItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Castorino"}

	cjJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジャスティン"}

	cjKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "저스틴"}

	cjLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "CJ"}

	cjRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Си-Джей"}

	cjSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "俞司廷"}

	cjSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "CJ"}

	cjTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "俞司廷"}
)

var (
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
	cjCharacter = nook.Character{
		Animal:   animal.Beaver,
		Birthday: cjBirthday,
		Code:     cjCode,
		Gender:   gender.Male,
		Name:     cjName}
)

var (
	CJ = nook.Resident{
		Character: cjCharacter}
)
