package sheep

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
	pietroBirthday = nook.Birthday{
		Day:   19,
		Month: time.April}
)

var (
	pietroCode = nook.Code{
		Value: "shp13"}
)

var (
	pietroAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pietro"}

	pietroCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pietro"}

	pietroDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pietro"}

	pietroFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pietro"}

	pietroGermanName = nook.Name{
		Language: language.German,
		Value:    "Pietro"}

	pietroItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Giulivo"}

	pietroJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジュペッティ"}

	pietroKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "피엘"}

	pietroLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Kikolor"}

	pietroRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пьетро"}

	pietroSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "皮耶罗"}

	pietroSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Kikolor"}

	pietroTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "皮耶羅"}
)

var (
	pietroName = nook.Languages{
		language.AmericanEnglish:      pietroAmericanEnglishName,
		language.CanadianFrench:       pietroCanadianFrenchName,
		language.Dutch:                pietroDutchName,
		language.French:               pietroFrenchName,
		language.German:               pietroGermanName,
		language.Italian:              pietroItalianName,
		language.Japanese:             pietroJapaneseName,
		language.Korean:               pietroKoreanName,
		language.LatinAmericanSpanish: pietroLatinAmericanSpanishName,
		language.Russian:              pietroRussianName,
		language.SimplifiedChinese:    pietroSimplifiedChineseName,
		language.Spanish:              pietroSpanishName,
		language.TraditionalChinese:   pietroTraditionalChineseName}
)

var (
	pietroCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: pietroBirthday,
		Code:     pietroCode,
		Key:      character.Pietro,
		Gender:   gender.Male,
		Name:     pietroName}
)

var (
	pietroAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "honk honk"}

	pietroCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "hahaha"}

	pietroDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "alaaf"}

	pietroFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bamboche"}

	pietroGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "helau"}

	pietroItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "peco peco"}

	pietroJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "グフフ"}

	pietroKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "므흐흐"}

	pietroLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "moquimoqui"}

	pietroRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бе-бе"}

	pietroSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "丑丑"}

	pietroSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "moquimoqui"}

	pietroTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "丑丑"}
)

var (
	pietroPhrase = nook.Languages{
		language.AmericanEnglish:      pietroAmericanEnglishPhrase,
		language.CanadianFrench:       pietroCanadianFrenchPhrase,
		language.Dutch:                pietroDutchPhrase,
		language.French:               pietroFrenchPhrase,
		language.German:               pietroGermanPhrase,
		language.Italian:              pietroItalianPhrase,
		language.Japanese:             pietroJapanesePhrase,
		language.Korean:               pietroKoreanPhrase,
		language.LatinAmericanSpanish: pietroLatinAmericanSpanishPhrase,
		language.Russian:              pietroRussianPhrase,
		language.SimplifiedChinese:    pietroSimplifiedChinesePhrase,
		language.Spanish:              pietroSpanishPhrase,
		language.TraditionalChinese:   pietroTraditionalChinesePhrase}
)

var (
	Pietro = nook.Villager{
		Character:   pietroCharacter,
		Personality: personality.Smug,
		Phrase:      pietroPhrase}
)
