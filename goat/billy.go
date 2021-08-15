package goat

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Seguinbêê quoi"}

	billyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Billynôh"}

	billyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Seguineh bââlèze"}

	billyGermanName = nook.Name{
		Language: language.German,
		Value:    "Henneso-o-ja-a-a"}

	billyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Billysorbeeelli"}

	billyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アーシンドよのう"}

	billyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "힘드러알아줘"}

	billyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Britobeerree"}

	billyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Биллитакие-сякие"}

	billySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "亚星唷喏"}

	billySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Britofenómeno"}

	billyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "亞星唷喏"}
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
		Animal:   Goat,
		Birthday: billyBirthday,
		Code:     billyCode,
		Gender:   nook.Male,
		Name:     billyName}
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
	Billy = nook.Villager{
		Character:   billyCharacter,
		Personality: nook.Jock,
		Phrase:      billyPhrase}
)
