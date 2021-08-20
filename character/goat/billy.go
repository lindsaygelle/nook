package goat

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
	billyBirthday = nook.Birthday{
		Day:   25,
		Month: time.March}
)

var (
	billyCode = nook.Code{
		Value: "goa02"}
)

var (
	billyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Billy"}

	billyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Seguin"}

	billyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Billy"}

	billyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Seguin"}

	billyGermanName = nook.Name{
		Language: language.German,
		Value:    "Hennes"}

	billyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Billy"}

	billyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アーシンド"}

	billyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "힘드러"}

	billyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Brito"}

	billyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Билли"}

	billySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "亚星"}

	billySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Brito"}

	billyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "亞星"}
)

var (
	billyName = nook.Languages{
		language.AmericanEnglish:      billyAmericanEnglishName,
		language.CanadianFrench:       billyCanadianFrenchName,
		language.Dutch:                billyDutchName,
		language.French:               billyFrenchName,
		language.German:               billyGermanName,
		language.Italian:              billyItalianName,
		language.Japanese:             billyJapaneseName,
		language.Korean:               billyKoreanName,
		language.LatinAmericanSpanish: billyLatinAmericanSpanishName,
		language.Russian:              billyRussianName,
		language.SimplifiedChinese:    billySimplifiedChineseName,
		language.Spanish:              billySpanishName,
		language.TraditionalChinese:   billyTraditionalChineseName}
)

var (
	billyCharacter = nook.Character{
		Animal:   animal.Goat,
		Birthday: billyBirthday,
		Code:     billyCode,
		Key:      character.Billy,
		Gender:   gender.Male,
		Name:     billyName,
		Special:  false}
)

var (
	billyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "dagnaabit"}

	billyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bêê quoi"}

	billyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "nôh"}

	billyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "eh bââlèze"}

	billyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "o-o-ja-a-a"}

	billyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sorbeeelli"}

	billyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "よのう"}

	billyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "알아줘"}

	billyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "beerree"}

	billyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "такие-сякие"}

	billySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唷喏"}

	billySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fenómeno"}

	billyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唷喏"}
)

var (
	billyPhrase = nook.Languages{
		language.AmericanEnglish:      billyAmericanEnglishPhrase,
		language.CanadianFrench:       billyCanadianFrenchPhrase,
		language.Dutch:                billyDutchPhrase,
		language.French:               billyFrenchPhrase,
		language.German:               billyGermanPhrase,
		language.Italian:              billyItalianPhrase,
		language.Japanese:             billyJapanesePhrase,
		language.Korean:               billyKoreanPhrase,
		language.LatinAmericanSpanish: billyLatinAmericanSpanishPhrase,
		language.Russian:              billyRussianPhrase,
		language.SimplifiedChinese:    billySimplifiedChinesePhrase,
		language.Spanish:              billySpanishPhrase,
		language.TraditionalChinese:   billyTraditionalChinesePhrase}
)

var (
	Billy = nook.Villager{
		Character:   billyCharacter,
		Personality: personality.Jock,
		Phrase:      billyPhrase}
)
