package wolf

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
	// whitneyBirthday represents Whitney's birthday (September 17th).
	whitneyBirthday = nook.Birthday{
		Day:   17,
		Month: time.September}
)

var (
	// whitneyCode represents Whitney's unique code ("wol03").
	whitneyCode = nook.Code{
		Value: "wol03"}
)

var (
	// whitneyAmericanEnglishName represents Whitney's name in American English.
	whitneyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Whitney"}

	// whitneyCanadianFrenchName represents Whitney's name in Canadian French.
	whitneyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Blanche"}

	// whitneyDutchName represents Whitney's name in Dutch.
	whitneyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Whitney"}

	// whitneyFrenchName represents Whitney's name in French.
	whitneyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Blanche"}

	// whitneyGermanName represents Whitney's name in German.
	whitneyGermanName = nook.Name{
		Language: language.German,
		Value:    "Lupa"}

	// whitneyItalianName represents Whitney's name in Italian.
	whitneyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Bianca"}

	// whitneyJapaneseName represents Whitney's name in Japanese.
	whitneyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ビアンカ"}

	// whitneyKoreanName represents Whitney's name in Korean.
	whitneyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "비앙카"}

	// whitneyLatinAmericanSpanishName represents Whitney's name in Latin American Spanish.
	whitneyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lupe"}

	// whitneyRussianName represents Whitney's name in Russian.
	whitneyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Уитни"}

	// whitneySimplifiedChineseName represents Whitney's name in Simplified Chinese.
	whitneySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "毕安卡"}

	// whitneySpanishName represents Whitney's name in Spanish.
	whitneySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lupe"}

	// whitneyTraditionalChineseName represents Whitney's name in Traditional Chinese.
	whitneyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "畢安卡"}
)

var (
	// whitneyName represents Whitney's name in different languages.
	whitneyName = nook.Languages{
		language.AmericanEnglish:      whitneyAmericanEnglishName,
		language.CanadianFrench:       whitneyCanadianFrenchName,
		language.Dutch:                whitneyDutchName,
		language.French:               whitneyFrenchName,
		language.German:               whitneyGermanName,
		language.Italian:              whitneyItalianName,
		language.Japanese:             whitneyJapaneseName,
		language.Korean:               whitneyKoreanName,
		language.LatinAmericanSpanish: whitneyLatinAmericanSpanishName,
		language.Russian:              whitneyRussianName,
		language.SimplifiedChinese:    whitneySimplifiedChineseName,
		language.Spanish:              whitneySpanishName,
		language.TraditionalChinese:   whitneyTraditionalChineseName}
)

var (
	// whitneyCharacter represents Whitney's character information.
	whitneyCharacter = nook.Character{
		Animal:   animal.Wolf,
		Birthday: whitneyBirthday,
		Code:     whitneyCode,
		Key:      character.Whitney,
		Gender:   gender.Female,
		Name:     whitneyName,
		Special:  false}
)

var (
	// whitneyAmericanEnglishPhrase represents Whitney's phrase in American English.
	whitneyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snappy"}

	// whitneyCanadianFrenchPhrase represents Whitney's phrase in Canadian French.
	whitneyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bing bang"}

	// whitneyDutchPhrase represents Whitney's phrase in Dutch.
	whitneyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hap"}

	// whitneyFrenchPhrase represents Whitney's phrase in French.
	whitneyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "sbam"}

	// whitneyGermanPhrase represents Whitney's phrase in German.
	whitneyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "jaulll"}

	// whitneyItalianPhrase represents Whitney's phrase in Italian.
	whitneyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "snappy"}

	// whitneyJapanesePhrase represents Whitney's phrase in Japanese.
	whitneyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ステキね"}

	// whitneyKoreanPhrase represents Whitney's phrase in Korean.
	whitneyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "멋져"}

	// whitneyLatinAmericanSpanishPhrase represents Whitney's phrase in Latin American Spanish.
	whitneyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "auf-auf"}

	// whitneyRussianPhrase represents Whitney's phrase in Russian.
	whitneyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "цап"}

	// whitneySimplifiedChinesePhrase represents Whitney's phrase in Simplified Chinese.
	whitneySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "漂亮"}

	// whitneySpanishPhrase represents Whitney's phrase in Spanish.
	whitneySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "auf-auf"}

	// whitneyTraditionalChinesePhrase represents Whitney's phrase in Traditional Chinese.
	whitneyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "漂亮"}
)

var (
	// whitneyPhrase represents Whitney's phrases in different languages.
	whitneyPhrase = nook.Languages{
		language.AmericanEnglish:      whitneyAmericanEnglishPhrase,
		language.CanadianFrench:       whitneyCanadianFrenchPhrase,
		language.Dutch:                whitneyDutchPhrase,
		language.French:               whitneyFrenchPhrase,
		language.German:               whitneyGermanPhrase,
		language.Italian:              whitneyItalianPhrase,
		language.Japanese:             whitneyJapanesePhrase,
		language.Korean:               whitneyKoreanPhrase,
		language.LatinAmericanSpanish: whitneyLatinAmericanSpanishPhrase,
		language.Russian:              whitneyRussianPhrase,
		language.SimplifiedChinese:    whitneySimplifiedChinesePhrase,
		language.Spanish:              whitneySpanishPhrase,
		language.TraditionalChinese:   whitneyTraditionalChinesePhrase}
)

var (
	// Whitney represents the character Whitney with her complete information.
	Whitney = nook.Villager{
		Character:   whitneyCharacter,
		Personality: personality.Snooty,
		Phrase:      whitneyPhrase}
)
