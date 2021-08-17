package horse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
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
		Value:    "Savana"}

	savannahDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Savannah"}

	savannahFrenchName = nook.Name{
		Language: language.French,
		Value:    "Savana"}

	savannahGermanName = nook.Name{
		Language: language.German,
		Value:    "Zara"}

	savannahItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Savannah"}

	savannahJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サバンナ"}

	savannahKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "사반나"}

	savannahLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sabana"}

	savannahRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Саванна"}

	savannahSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "草斑娜"}

	savannahSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sabana"}

	savannahTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "草斑娜"}
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
		Animal:   animal.Horse,
		Birthday: savannahBirthday,
		Code:     savannahCode,
		Key:      character.Savannah,
		Gender:   gender.Female,
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
		language.AmericanEnglish:      savannahAmericanEnglishPhrase,
		language.CanadianFrench:       savannahCanadianFrenchPhrase,
		language.Dutch:                savannahDutchPhrase,
		language.French:               savannahFrenchPhrase,
		language.German:               savannahGermanPhrase,
		language.Italian:              savannahItalianPhrase,
		language.Japanese:             savannahJapanesePhrase,
		language.Korean:               savannahKoreanPhrase,
		language.LatinAmericanSpanish: savannahLatinAmericanSpanishPhrase,
		language.Russian:              savannahRussianPhrase,
		language.SimplifiedChinese:    savannahSimplifiedChinesePhrase,
		language.Spanish:              savannahSpanishPhrase,
		language.TraditionalChinese:   savannahTraditionalChinesePhrase}
)

var (
	Savannah = nook.Villager{
		Character:   savannahCharacter,
		Personality: personality.Normal,
		Phrase:      savannahPhrase}
)
