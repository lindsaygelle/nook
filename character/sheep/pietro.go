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
	// pietroBirthday represents Pietro's birthday.
	pietroBirthday = nook.Birthday{
		Day:   19,
		Month: time.April}
)

var (
	// pietroCode represents Pietro's unique code.
	pietroCode = nook.Code{
		Value: "shp13"}
)

var (
	// pietroAmericanEnglishName represents Pietro's name in American English.
	pietroAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pietro"}

	// pietroCanadianFrenchName represents Pietro's name in Canadian French.
	pietroCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pietro"}

	// pietroDutchName represents Pietro's name in Dutch.
	pietroDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pietro"}

	// pietroFrenchName represents Pietro's name in French.
	pietroFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pietro"}

	// pietroGermanName represents Pietro's name in German.
	pietroGermanName = nook.Name{
		Language: language.German,
		Value:    "Pietro"}

	// pietroItalianName represents Pietro's name in Italian.
	pietroItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Giulivo"}

	// pietroJapaneseName represents Pietro's name in Japanese.
	pietroJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジュペッティ"}

	// pietroKoreanName represents Pietro's name in Korean.
	pietroKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "피엘"}

	// pietroLatinAmericanSpanishName represents Pietro's name in Latin American Spanish.
	pietroLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Kikolor"}

	// pietroRussianName represents Pietro's name in Russian.
	pietroRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пьетро"}

	// pietroSimplifiedChineseName represents Pietro's name in Simplified Chinese.
	pietroSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "皮耶罗"}

	// pietroSpanishName represents Pietro's name in Spanish.
	pietroSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Kikolor"}

	// pietroTraditionalChineseName represents Pietro's name in Traditional Chinese.
	pietroTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "皮耶羅"}
)

var (
	// pietroName represents Pietro's name in different languages.
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
	// pietroCharacter represents Pietro's character information.
	pietroCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: pietroBirthday,
		Code:     pietroCode,
		Key:      character.Pietro,
		Gender:   gender.Male,
		Name:     pietroName,
		Special:  false}
)

var (
	// pietroAmericanEnglishPhrase represents Pietro's phrase in American English.
	pietroAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "honk honk"}

	// pietroCanadianFrenchPhrase represents Pietro's phrase in Canadian French.
	pietroCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "hahaha"}

	// pietroDutchPhrase represents Pietro's phrase in Dutch.
	pietroDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "alaaf"}

	// pietroFrenchPhrase represents Pietro's phrase in French.
	pietroFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bamboche"}

	// pietroGermanPhrase represents Pietro's phrase in German.
	pietroGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "helau"}

	// pietroItalianPhrase represents Pietro's phrase in Italian.
	pietroItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "peco peco"}

	// pietroJapanesePhrase represents Pietro's phrase in Japanese.
	pietroJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "グフフ"}

	// pietroKoreanPhrase represents Pietro's phrase in Korean.
	pietroKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "므흐흐"}

	// pietroLatinAmericanSpanishPhrase represents Pietro's phrase in Latin American Spanish.
	pietroLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "moquimoqui"}

	// pietroRussianPhrase represents Pietro's phrase in Russian.
	pietroRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бе-бе"}

	// pietroSimplifiedChinesePhrase represents Pietro's phrase in Simplified Chinese.
	pietroSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "丑丑"}

	// pietroSpanishPhrase represents Pietro's phrase in Spanish.
	pietroSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "moquimoqui"}

	// pietroTraditionalChinesePhrase represents Pietro's phrase in Traditional Chinese.
	pietroTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "丑丑"}
)

var (
	// pietroPhrase represents Pietro's phrases in different languages.
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
	// Pietro represents the character Pietro with his complete information.
	Pietro = nook.Villager{
		Character:   pietroCharacter,
		Personality: personality.Smug,
		Phrase:      pietroPhrase}
)
