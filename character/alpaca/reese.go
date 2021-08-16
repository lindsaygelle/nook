package alpaca

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	reeseBirthday = nook.Birthday{
		Day:   5,
		Month: time.July}
)

var (
	reeseCode = nook.Code{
		Value: "alw"}
)

var (
	reeseAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Reese"}

	reeseCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Risette"}

	reeseDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Reese"}

	reeseFrenchName = nook.Name{
		Language: language.French,
		Value:    "Risette"}

	reeseGermanName = nook.Name{
		Language: language.German,
		Value:    "Rosina"}

	reeseItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Alpaca"}

	reeseJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リサ"}

	reeseKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "리사"}

	reeseLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Paca"}

	reeseRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Риз"}

	reeseSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莉咏"}

	reeseSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Paca"}

	reeseTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莉詠"}
)

var (
	reeseName = nook.Languages{
		language.AmericanEnglish:      reeseAmericanEnglishName,
		language.CanadianFrench:       reeseCanadianFrenchName,
		language.Dutch:                reeseDutchName,
		language.French:               reeseFrenchName,
		language.German:               reeseGermanName,
		language.Italian:              reeseItalianName,
		language.Japanese:             reeseJapaneseName,
		language.Korean:               reeseKoreanName,
		language.LatinAmericanSpanish: reeseLatinAmericanSpanishName,
		language.Russian:              reeseRussianName,
		language.SimplifiedChinese:    reeseSimplifiedChineseName,
		language.Spanish:              reeseSpanishName,
		language.TraditionalChinese:   reeseTraditionalChineseName}
)

var (
	reeseCharacter = nook.Character{
		Animal:   animal.Alpaca,
		Birthday: reeseBirthday,
		Code:     reeseCode,
		Key:      character.Reese,
		Gender:   gender.Female,
		Name:     reeseName}
)

var (
	Reese = nook.Resident{
		Character: reeseCharacter}
)
