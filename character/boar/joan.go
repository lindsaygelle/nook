package boar

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	joanBirthday = nook.Birthday{
		Day:   8,
		Month: time.January}
)

var (
	joanCode = nook.Code{
		Value: "boa"}
)

var (
	joanAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Joan"}

	joanCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Porcella"}

	joanDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Joan"}

	joanFrenchName = nook.Name{
		Language: language.French,
		Value:    "Porcella"}

	joanGermanName = nook.Name{
		Language: language.German,
		Value:    "Sigrid"}

	joanItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Nella"}

	joanJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カブリバ"}

	joanKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "무파라"}

	joanLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Juana"}

	joanRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джоан"}

	joanSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "曹谷"}

	joanSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Juana"}

	joanTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "曹谷"}
)

var (
	joanName = nook.Languages{
		language.AmericanEnglish:      joanAmericanEnglishName,
		language.CanadianFrench:       joanCanadianFrenchName,
		language.Dutch:                joanDutchName,
		language.French:               joanFrenchName,
		language.German:               joanGermanName,
		language.Italian:              joanItalianName,
		language.Japanese:             joanJapaneseName,
		language.Korean:               joanKoreanName,
		language.LatinAmericanSpanish: joanLatinAmericanSpanishName,
		language.Russian:              joanRussianName,
		language.SimplifiedChinese:    joanSimplifiedChineseName,
		language.Spanish:              joanSpanishName,
		language.TraditionalChinese:   joanTraditionalChineseName}
)

var (
	joanCharacter = nook.Character{
		Animal:   animal.Boar,
		Birthday: joanBirthday,
		Code:     joanCode,
		Key:      character.Joan,
		Gender:   gender.Female,
		Name:     joanName}
)

var (
	Joan = nook.Resident{
		Character: joanCharacter}
)
