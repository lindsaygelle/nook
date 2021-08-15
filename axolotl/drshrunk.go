package axolotl

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	drshrunkBirthday = nook.Birthday{
		Day:   2,
		Month: time.January}
)

var (
	drshrunkCode = nook.Code{
		Value: "upa"}
)

var (
	drshrunkAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dr. Shrunk"}

	drshrunkCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Ciboulot"}

	drshrunkDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Shrunk"}

	drshrunkFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ciboulot"}

	drshrunkGermanName = nook.Name{
		Language: language.German,
		Value:    "Samselt"}

	drshrunkItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Strizzo"}

	drshrunkJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ししょー"}

	drshrunkKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "스승"}

	drshrunkLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Dr. Sito"}

	drshrunkRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Шранк"}

	drshrunkSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "笑匠"}

	drshrunkSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Dr. Sito"}

	drshrunkTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "笑匠"}
)

var (
	drshrunkName = nook.Languages{
		language.AmericanEnglish:      drshrunkAmericanEnglishName,
		language.CanadianFrench:       drshrunkCanadianFrenchName,
		language.Dutch:                drshrunkDutchName,
		language.French:               drshrunkFrenchName,
		language.German:               drshrunkGermanName,
		language.Italian:              drshrunkItalianName,
		language.Japanese:             drshrunkJapaneseName,
		language.Korean:               drshrunkKoreanName,
		language.LatinAmericanSpanish: drshrunkLatinAmericanSpanishName,
		language.Russian:              drshrunkRussianName,
		language.SimplifiedChinese:    drshrunkSimplifiedChineseName,
		language.Spanish:              drshrunkSpanishName,
		language.TraditionalChinese:   drshrunkTraditionalChineseName}
)

var (
	drshrunkCharacter = nook.Character{
		Animal:   Axolotl,
		Birthday: drshrunkBirthday,
		Code:     drshrunkCode,
		Gender:   nook.Male,
		Name:     drshrunkName}
)

var (
	DrShrunk = nook.Resident{
		Character: drshrunkCharacter}
)
