package mouse

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Carmenmamour"}

	carmenGermanName = nook.Name{
		Language: language.German,
		Value:    "Carmenäpfelchen"}

	carmenItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Carmenbuffetto"}

	carmenJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゼリーんとんと"}

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
		Value:    "卡门吱哧"}

	carmenSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Carmenbolilla"}

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
		Animal:   Mouse,
		Birthday: carmenBirthday,
		Code:     carmenCode,
		Gender:   nook.Female,
		Name:     carmenName}
)

var (
	carmenAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	carmenCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	carmenDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

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
		Value:    "N/A"}

	carmenLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	carmenRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	carmenSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "吱哧"}

	carmenSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bolilla"}

	carmenTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	carmenPhrase = nook.Languages{
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
	Carmen = nook.Villager{
		Character:   carmenCharacter,
		Personality: nook.Snooty,
		Phrase:      carmenPhrase}
)
