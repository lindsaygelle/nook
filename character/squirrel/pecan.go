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
	// pecanBirthday represents Pecan's birthday.
	pecanBirthday = nook.Birthday{
		Day:   10,
		Month: time.September}
)

var (
	// pecanCode represents Pecan's unique code.
	pecanCode = nook.Code{
		Value: "squ03"}
)

var (
	// pecanAmericanEnglishName represents Pecan's name in American English.
	pecanAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pecan"}

	// pecanCanadianFrenchName represents Pecan's name in Canadian French.
	pecanCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pécan"}

	// pecanDutchName represents Pecan's name in Dutch.
	pecanDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pecan"}

	// pecanFrenchName represents Pecan's name in French.
	pecanFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pécan"}

	// pecanGermanName represents Pecan's name in German.
	pecanGermanName = nook.Name{
		Language: language.German,
		Value:    "Noisette"}

	// pecanItalianName represents Pecan's name in Italian.
	pecanItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Nocina"}

	// pecanJapaneseName represents Pecan's name in Japanese.
	pecanJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "レベッカ"}

	// pecanKoreanName represents Pecan's name in Korean.
	pecanKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "레베카"}

	// pecanLatinAmericanSpanishName represents Pecan's name in Latin American Spanish.
	pecanLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Camila"}

	// pecanRussianName represents Pecan's name in Russian.
	pecanRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пекан"}

	// pecanSimplifiedChineseName represents Pecan's name in Simplified Chinese.
	pecanSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雷贝嘉"}

	// pecanSpanishName represents Pecan's name in Spanish.
	pecanSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Camila"}

	// pecanTraditionalChineseName represents Pecan's name in Traditional Chinese.
	pecanTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雷貝嘉"}
)

var (
	// pecanName represents Pecan's name in different languages.
	pecanName = nook.Languages{
		language.AmericanEnglish:      pecanAmericanEnglishName,
		language.CanadianFrench:       pecanCanadianFrenchName,
		language.Dutch:                pecanDutchName,
		language.French:               pecanFrenchName,
		language.German:               pecanGermanName,
		language.Italian:              pecanItalianName,
		language.Japanese:             pecanJapaneseName,
		language.Korean:               pecanKoreanName,
		language.LatinAmericanSpanish: pecanLatinAmericanSpanishName,
		language.Russian:              pecanRussianName,
		language.SimplifiedChinese:    pecanSimplifiedChineseName,
		language.Spanish:              pecanSpanishName,
		language.TraditionalChinese:   pecanTraditionalChineseName}
)

var (
	// pecanCharacter represents Pecan's character information.
	pecanCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: pecanBirthday,
		Code:     pecanCode,
		Key:      character.Pecan,
		Gender:   gender.Female,
		Name:     pecanName,
		Special:  false}
)

var (
	// pecanAmericanEnglishPhrase represents Pecan's phrase in American English.
	pecanAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chipmunk"}

	// pecanCanadianFrenchPhrase represents Pecan's phrase in Canadian French.
	pecanCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "fouloulou"}

	// pecanDutchPhrase represents Pecan's phrase in Dutch.
	pecanDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "eekhoorn"}

	// pecanFrenchPhrase represents Pecan's phrase in French.
	pecanFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "fouloulou"}

	// pecanGermanPhrase represents Pecan's phrase in German.
	pecanGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "spatzl"}

	// pecanItalianPhrase represents Pecan's phrase in Italian.
	pecanItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cipì"}

	// pecanJapanesePhrase represents Pecan's phrase in Japanese.
	pecanJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "つんっ"}

	// pecanKoreanPhrase represents Pecan's phrase in Korean.
	pecanKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쯧쯧"}

	// pecanLatinAmericanSpanishPhrase represents Pecan's phrase in Latin American Spanish.
	pecanLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "piñoní"}

	// pecanRussianPhrase represents Pecan's phrase in Russian.
	pecanRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бурундук"}

	// pecanSimplifiedChinesePhrase represents Pecan's phrase in Simplified Chinese.
	pecanSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "正经"}

	// pecanSpanishPhrase represents Pecan's phrase in Spanish.
	pecanSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "piñoncito"}

	// pecanTraditionalChinesePhrase represents Pecan's phrase in Traditional Chinese.
	pecanTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "正經"}
)

var (
	// pecanPhrase represents Pecan's phrase in different languages.
	pecanPhrase = nook.Languages{
		language.AmericanEnglish:      pecanAmericanEnglishPhrase,
		language.CanadianFrench:       pecanCanadianFrenchPhrase,
		language.Dutch:                pecanDutchPhrase,
		language.French:               pecanFrenchPhrase,
		language.German:               pecanGermanPhrase,
		language.Italian:              pecanItalianPhrase,
		language.Japanese:             pecanJapanesePhrase,
		language.Korean:               pecanKoreanPhrase,
		language.LatinAmericanSpanish: pecanLatinAmericanSpanishPhrase,
		language.Russian:              pecanRussianPhrase,
		language.SimplifiedChinese:    pecanSimplifiedChinesePhrase,
		language.Spanish:              pecanSpanishPhrase,
		language.TraditionalChinese:   pecanTraditionalChinesePhrase}
)

var (
	// Pecan represents the character Pecan with her complete information.
	Pecan = nook.Villager{
		Character:   pecanCharacter,
		Personality: personality.Snooty,
		Phrase:      pecanPhrase}
)
