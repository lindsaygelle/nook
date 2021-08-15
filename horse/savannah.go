package horse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	savannahBirthday = nook.Birthday{
		Day:   25,
		Month: time.January}
)

var (
	savannahCode = nook.Code{
		Value: "hrs02"}
)

var (
	savannahAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Savannah"}

	savannahCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Savanawimbowé"}

	savannahDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Savannahpony"}

	savannahFrenchName = nook.Name{
		Language: language.French,
		Value:    "Savanawimbowé"}

	savannahGermanName = nook.Name{
		Language: language.German,
		Value:    "Zarawichichi"}

	savannahItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Savannahhip hip"}

	savannahJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サバンナってば"}

	savannahKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "사반나맞아요"}

	savannahLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sabanaazuquita"}

	savannahRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Саваннаи все"}

	savannahSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "草斑娜说到"}

	savannahSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sabanaazucarillo"}

	savannahTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "草斑娜說到"}
)

var (
	savannahName = nook.Languages{
		language.AmericanEnglish:      savannahAmericanEnglishName,
		language.CanadianFrench:       savannahCanadianFrenchName,
		language.Dutch:                savannahDutchName,
		language.French:               savannahFrenchName,
		language.German:               savannahGermanName,
		language.Italian:              savannahItalianName,
		language.Japanese:             savannahJapaneseName,
		language.Korean:               savannahKoreanName,
		language.LatinAmericanSpanish: savannahLatinAmericanSpanishName,
		language.Russian:              savannahRussianName,
		language.SimplifiedChinese:    savannahSimplifiedChineseName,
		language.Spanish:              savannahSpanishName,
		language.TraditionalChinese:   savannahTraditionalChineseName}
)

var (
	savannahCharacter = nook.Character{
		Animal:   Horse,
		Birthday: savannahBirthday,
		Code:     savannahCode,
		Gender:   nook.Female,
		Name:     savannahName}
)

var (
	savannahAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "y'all"}

	savannahCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "wimbowé"}

	savannahDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pony"}

	savannahFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "wimbowé"}

	savannahGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "wichichi"}

	savannahItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hip hip"}

	savannahJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ってば"}

	savannahKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "맞아요"}

	savannahLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "azuquita"}

	savannahRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "и все"}

	savannahSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "说到"}

	savannahSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "azucarillo"}

	savannahTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "說到"}
)

var (
	savannahPhrase = nook.Languages{
		language.AmericanEnglish:      savannahAmericanEnglishName,
		language.CanadianFrench:       savannahCanadianFrenchName,
		language.Dutch:                savannahDutchName,
		language.French:               savannahFrenchName,
		language.German:               savannahGermanName,
		language.Italian:              savannahItalianName,
		language.Japanese:             savannahJapaneseName,
		language.Korean:               savannahKoreanName,
		language.LatinAmericanSpanish: savannahLatinAmericanSpanishName,
		language.Russian:              savannahRussianName,
		language.SimplifiedChinese:    savannahSimplifiedChineseName,
		language.Spanish:              savannahSpanishName,
		language.TraditionalChinese:   savannahTraditionalChineseName}
)

var (
	Savannah = nook.Villager{
		Character:   savannahCharacter,
		Personality: nook.Normal,
		Phrase:      savannahPhrase}
)
