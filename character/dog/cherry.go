package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	cherryBirthday = nook.Birthday{
		Day:   11,
		Month: time.May}
)

var (
	cherryCode = nook.Code{
		Value: "dog17"}
)

var (
	cherryAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cherry"}

	cherryCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Anna"}

	cherryDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cherry"}

	cherryFrenchName = nook.Name{
		Language: language.French,
		Value:    "Anna"}

	cherryGermanName = nook.Name{
		Language: language.German,
		Value:    "Bella"}

	cherryItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Amarena"}

	cherryJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ハンナ"}

	cherryKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "한나"}

	cherryLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Luna"}

	cherryRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Черри"}

	cherrySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "花娜"}

	cherrySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Luna"}

	cherryTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "花娜"}
)

var (
	cherryName = nook.Languages{
		language.AmericanEnglish:      cherryAmericanEnglishName,
		language.CanadianFrench:       cherryCanadianFrenchName,
		language.Dutch:                cherryDutchName,
		language.French:               cherryFrenchName,
		language.German:               cherryGermanName,
		language.Italian:              cherryItalianName,
		language.Japanese:             cherryJapaneseName,
		language.Korean:               cherryKoreanName,
		language.LatinAmericanSpanish: cherryLatinAmericanSpanishName,
		language.Russian:              cherryRussianName,
		language.SimplifiedChinese:    cherrySimplifiedChineseName,
		language.Spanish:              cherrySpanishName,
		language.TraditionalChinese:   cherryTraditionalChineseName}
)

var (
	cherryCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: cherryBirthday,
		Code:     cherryCode,
		Gender:   gender.Female,
		Name:     cherryName}
)

var (
	cherryAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "what what"}

	cherryCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "niniche"}

	cherryDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "wat jij"}

	cherryFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "niniche"}

	cherryGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "wedel"}

	cherryItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "arf arf"}

	cherryJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だってネ"}

	cherryKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그냥저냥"}

	cherryLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "guauchi"}

	cherryRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "что-что"}

	cherrySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "听说娜"}

	cherrySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "guauchi"}

	cherryTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "聽說娜"}
)

var (
	cherryPhrase = nook.Languages{
		language.AmericanEnglish:      cherryAmericanEnglishName,
		language.CanadianFrench:       cherryCanadianFrenchName,
		language.Dutch:                cherryDutchName,
		language.French:               cherryFrenchName,
		language.German:               cherryGermanName,
		language.Italian:              cherryItalianName,
		language.Japanese:             cherryJapaneseName,
		language.Korean:               cherryKoreanName,
		language.LatinAmericanSpanish: cherryLatinAmericanSpanishName,
		language.Russian:              cherryRussianName,
		language.SimplifiedChinese:    cherrySimplifiedChineseName,
		language.Spanish:              cherrySpanishName,
		language.TraditionalChinese:   cherryTraditionalChineseName}
)

var (
	Cherry = nook.Villager{
		Character:   cherryCharacter,
		Personality: personality.BigSister,
		Phrase:      cherryPhrase}
)
