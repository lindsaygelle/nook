package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	nibblesBirthday = nook.Birthday{
		Day:   19,
		Month: time.July}
)

var (
	nibblesCode = nook.Code{
		Value: "squ04"}
)

var (
	nibblesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nibbles"}

	nibblesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lolagrignotte"}

	nibblesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Nibblesknabbel"}

	nibblesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lolanunuche"}

	nibblesGermanName = nook.Name{
		Language: language.German,
		Value:    "Knuspiknusper"}

	nibblesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rododeatruciolo"}

	nibblesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ガリガリガジガジ"}

	nibblesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "그리미송송"}

	nibblesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Dentinatrocitrí"}

	nibblesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Ниблзорешек"}

	nibblesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咬咬咯吱咯吱"}

	nibblesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Dentinabellotita"}

	nibblesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咬咬咯吱咯吱"}
)

var (
	nibblesName = nook.Languages{
		language.AmericanEnglish:      nibblesAmericanEnglishName,
		language.CanadianFrench:       nibblesCanadianFrenchName,
		language.Dutch:                nibblesDutchName,
		language.French:               nibblesFrenchName,
		language.German:               nibblesGermanName,
		language.Italian:              nibblesItalianName,
		language.Japanese:             nibblesJapaneseName,
		language.Korean:               nibblesKoreanName,
		language.LatinAmericanSpanish: nibblesLatinAmericanSpanishName,
		language.Russian:              nibblesRussianName,
		language.SimplifiedChinese:    nibblesSimplifiedChineseName,
		language.Spanish:              nibblesSpanishName,
		language.TraditionalChinese:   nibblesTraditionalChineseName}
)

var (
	nibblesCharacter = nook.Character{
		Animal:   Squirrel,
		Birthday: nibblesBirthday,
		Code:     nibblesCode,
		Gender:   nook.Female,
		Name:     nibblesName}
)

var (
	nibblesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "niblet"}

	nibblesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grignotte"}

	nibblesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "knabbel"}

	nibblesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "nunuche"}

	nibblesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knusper"}

	nibblesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "truciolo"}

	nibblesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ガジガジ"}

	nibblesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "송송"}

	nibblesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "trocitrí"}

	nibblesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "орешек"}

	nibblesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咯吱咯吱"}

	nibblesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bellotita"}

	nibblesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咯吱咯吱"}
)

var (
	nibblesPhrase = nook.Languages{
		language.AmericanEnglish:      nibblesAmericanEnglishName,
		language.CanadianFrench:       nibblesCanadianFrenchName,
		language.Dutch:                nibblesDutchName,
		language.French:               nibblesFrenchName,
		language.German:               nibblesGermanName,
		language.Italian:              nibblesItalianName,
		language.Japanese:             nibblesJapaneseName,
		language.Korean:               nibblesKoreanName,
		language.LatinAmericanSpanish: nibblesLatinAmericanSpanishName,
		language.Russian:              nibblesRussianName,
		language.SimplifiedChinese:    nibblesSimplifiedChineseName,
		language.Spanish:              nibblesSpanishName,
		language.TraditionalChinese:   nibblesTraditionalChineseName}
)

var (
	Nibbles = nook.Villager{
		Character:   nibblesCharacter,
		Personality: nook.Peppy,
		Phrase:      nibblesPhrase}
)
