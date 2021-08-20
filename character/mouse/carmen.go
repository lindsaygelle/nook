package mouse

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
	carmenBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	carmenCode = nook.Code{
		Value: ""}
)

var (
	carmenAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Carmen"}

	carmenCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	carmenDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	carmenFrenchName = nook.Name{
		Language: language.French,
		Value:    "Carmen"}

	carmenGermanName = nook.Name{
		Language: language.German,
		Value:    "Carmen"}

	carmenItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Carmen"}

	carmenJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゼリー"}

	carmenKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	carmenLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	carmenRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	carmenSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "卡门"}

	carmenSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Carmen"}

	carmenTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	carmenName = nook.Languages{
		language.AmericanEnglish:      carmenAmericanEnglishName,
		language.CanadianFrench:       carmenCanadianFrenchName,
		language.Dutch:                carmenDutchName,
		language.French:               carmenFrenchName,
		language.German:               carmenGermanName,
		language.Italian:              carmenItalianName,
		language.Japanese:             carmenJapaneseName,
		language.Korean:               carmenKoreanName,
		language.LatinAmericanSpanish: carmenLatinAmericanSpanishName,
		language.Russian:              carmenRussianName,
		language.SimplifiedChinese:    carmenSimplifiedChineseName,
		language.Spanish:              carmenSpanishName,
		language.TraditionalChinese:   carmenTraditionalChineseName}
)

var (
	carmenCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: carmenBirthday,
		Code:     carmenCode,
		Key:      character.Carmen,
		Gender:   gender.Female,
		Name:     carmenName,
		Special:  false}
)

var (
	carmenAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bumpkin"}

	carmenCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	carmenDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	carmenFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mamour"}

	carmenGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "äpfelchen"}

	carmenItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "buffetto"}

	carmenJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "んとんと"}

	carmenKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	carmenLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	carmenRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	carmenSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "吱哧"}

	carmenSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bolilla"}

	carmenTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	carmenPhrase = nook.Languages{
		language.AmericanEnglish:      carmenAmericanEnglishPhrase,
		language.CanadianFrench:       carmenCanadianFrenchPhrase,
		language.Dutch:                carmenDutchPhrase,
		language.French:               carmenFrenchPhrase,
		language.German:               carmenGermanPhrase,
		language.Italian:              carmenItalianPhrase,
		language.Japanese:             carmenJapanesePhrase,
		language.Korean:               carmenKoreanPhrase,
		language.LatinAmericanSpanish: carmenLatinAmericanSpanishPhrase,
		language.Russian:              carmenRussianPhrase,
		language.SimplifiedChinese:    carmenSimplifiedChinesePhrase,
		language.Spanish:              carmenSpanishPhrase,
		language.TraditionalChinese:   carmenTraditionalChinesePhrase}
)

var (
	Carmen = nook.Villager{
		Character:   carmenCharacter,
		Personality: personality.Snooty,
		Phrase:      carmenPhrase}
)
