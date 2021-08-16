package hedgehog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	labelBirthday = nook.Birthday{
		Day:   31,
		Month: time.October}
)

var (
	labelCode = nook.Code{
		Value: "hgc"}
)

var (
	labelAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Label"}

	labelCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tiquette"}

	labelDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Label"}

	labelFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tiquette"}

	labelGermanName = nook.Name{
		Language: language.German,
		Value:    "Minna"}

	labelItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Beatrice"}

	labelJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ことの"}

	labelKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "고숙이"}

	labelLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tili"}

	labelRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лэйбл"}

	labelSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "绵儿"}

	labelSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tili"}

	labelTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "綿兒"}
)

var (
	labelName = nook.Languages{
		language.AmericanEnglish:      labelAmericanEnglishName,
		language.CanadianFrench:       labelCanadianFrenchName,
		language.Dutch:                labelDutchName,
		language.French:               labelFrenchName,
		language.German:               labelGermanName,
		language.Italian:              labelItalianName,
		language.Japanese:             labelJapaneseName,
		language.Korean:               labelKoreanName,
		language.LatinAmericanSpanish: labelLatinAmericanSpanishName,
		language.Russian:              labelRussianName,
		language.SimplifiedChinese:    labelSimplifiedChineseName,
		language.Spanish:              labelSpanishName,
		language.TraditionalChinese:   labelTraditionalChineseName}
)

var (
	labelCharacter = nook.Character{
		Animal:   animal.Hedgehog,
		Birthday: labelBirthday,
		Code:     labelCode,
		Key:      character.Label,
		Gender:   gender.Female,
		Name:     labelName}
)

var (
	Label = nook.Resident{
		Character: labelCharacter}
)
