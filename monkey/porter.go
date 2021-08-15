package monkey

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	porterBirthday = nook.Birthday{
		Day:   17,
		Month: time.April}
)

var (
	porterCode = nook.Code{
		Value: "mnk"}
)

var (
	porterAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Porter"}

	porterCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lazare"}

	porterDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Porter"}

	porterFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lazare"}

	porterGermanName = nook.Name{
		Language: language.German,
		Value:    "Flip"}

	porterItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ciufciuf"}

	porterJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "えきいんさん"}

	porterKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "역무원"}

	porterLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Estasio"}

	porterRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Портер"}

	porterSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "车长叔叔"}

	porterSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Estasio"}

	porterTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "車長叔叔"}
)

var (
	porterName = nook.Languages{
		language.AmericanEnglish:      porterAmericanEnglishName,
		language.CanadianFrench:       porterCanadianFrenchName,
		language.Dutch:                porterDutchName,
		language.French:               porterFrenchName,
		language.German:               porterGermanName,
		language.Italian:              porterItalianName,
		language.Japanese:             porterJapaneseName,
		language.Korean:               porterKoreanName,
		language.LatinAmericanSpanish: porterLatinAmericanSpanishName,
		language.Russian:              porterRussianName,
		language.SimplifiedChinese:    porterSimplifiedChineseName,
		language.Spanish:              porterSpanishName,
		language.TraditionalChinese:   porterTraditionalChineseName}
)

var (
	porterCharacter = nook.Character{
		Animal:   Monkey,
		Birthday: porterBirthday,
		Code:     porterCode,
		Gender:   nook.Male,
		Name:     porterName}
)

var (
	Porter = nook.Resident{
		Character: porterCharacter}
)
