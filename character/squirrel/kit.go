package squirrel

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
	// kitBirthday represents Kit's birthday.
	kitBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// kitCode represents Kit's unique code.
	kitCode = nook.Code{
		Value: ""}
)

var (
	// kitAmericanEnglishName represents Kit's name in American English.
	kitAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kit"}

	// kitCanadianFrenchName represents Kit's name in Canadian French.
	kitCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// kitDutchName represents Kit's name in Dutch.
	kitDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// kitFrenchName represents Kit's name in French.
	kitFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// kitGermanName represents Kit's name in AGerman.
	kitGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// kitItalianName represents Kit's name in Italian.
	kitItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// kitJapaneseName represents Kit's name in Japanese.
	kitJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キット"}

	// kitKoreanName represents Kit's name in Korean.
	kitKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// kitLatinAmericanSpanishName represents Kit's name in Latin American Spanish.
	kitLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// kitRussianName represents Kit's name in Russian.
	kitRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// kitSimplifiedChineseName represents Kit's name in Simplified Chinese.
	kitSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// kitSpanishName represents Kit's name in Spanish.
	kitSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// kitTraditionalChineseName represents Kit's name in Traditional Chinese.
	kitTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// kitName represents Kit's name in different languages.
	kitName = nook.Languages{
		language.AmericanEnglish:      kitAmericanEnglishName,
		language.CanadianFrench:       kitCanadianFrenchName,
		language.Dutch:                kitDutchName,
		language.French:               kitFrenchName,
		language.German:               kitGermanName,
		language.Italian:              kitItalianName,
		language.Japanese:             kitJapaneseName,
		language.Korean:               kitKoreanName,
		language.LatinAmericanSpanish: kitLatinAmericanSpanishName,
		language.Russian:              kitRussianName,
		language.SimplifiedChinese:    kitSimplifiedChineseName,
		language.Spanish:              kitSpanishName,
		language.TraditionalChinese:   kitTraditionalChineseName}
)

var (
	// kitCharacter represents Kit's character information.
	kitCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: kitBirthday,
		Code:     kitCode,
		Key:      character.Kit,
		Gender:   gender.Male,
		Name:     kitName,
		Special:  false}
)

var (
	// kitAmericanEnglishPhrase represents Kit's phrase in American English.
	kitAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "だキョ"}

	// kitAmericanEnglishPhrase represents Kit's phrase in Canadian French.
	kitCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// kitAmericanEnglishPhrase represents Kit's phrase in Dutch.
	kitDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// kitAmericanEnglishPhrase represents Kit's phrase in French.
	kitFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// kitAmericanEnglishPhrase represents Kit's phrase in German.
	kitGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// kitAmericanEnglishPhrase represents Kit's phrase in Italian.
	kitItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// kitAmericanEnglishPhrase represents Kit's phrase in Japanese.
	kitJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だキョ"}

	// kitAmericanEnglishPhrase represents Kit's phrase in Korean.
	kitKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// kitAmericanEnglishPhrase represents Kit's phrase in Latin American Spanish.
	kitLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// kitAmericanEnglishPhrase represents Kit's phrase in Russian.
	kitRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// kitAmericanEnglishPhrase represents Kit's phrase in Simplified Chinese.
	kitSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// kitAmericanEnglishPhrase represents Kit's phrase in Spanish.
	kitSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// kitAmericanEnglishPhrase represents Kit's phrase in Traditional Chinese.
	kitTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// kitPhrase represents Kit's phrases in different languages.
	kitPhrase = nook.Languages{
		language.AmericanEnglish:      kitAmericanEnglishPhrase,
		language.CanadianFrench:       kitCanadianFrenchPhrase,
		language.Dutch:                kitDutchPhrase,
		language.French:               kitFrenchPhrase,
		language.German:               kitGermanPhrase,
		language.Italian:              kitItalianPhrase,
		language.Japanese:             kitJapanesePhrase,
		language.Korean:               kitKoreanPhrase,
		language.LatinAmericanSpanish: kitLatinAmericanSpanishPhrase,
		language.Russian:              kitRussianPhrase,
		language.SimplifiedChinese:    kitSimplifiedChinesePhrase,
		language.Spanish:              kitSpanishPhrase,
		language.TraditionalChinese:   kitTraditionalChinesePhrase}
)

var (
	// Kit represents the character Kit with his complete information.
	Kit = nook.Villager{
		Character:   kitCharacter,
		Personality: personality.Jock,
		Phrase:      kitPhrase}
)
