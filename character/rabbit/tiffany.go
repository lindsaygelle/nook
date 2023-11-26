package rabbit

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
	// tiffanyBirthday represents Tiffany's birthday.
	tiffanyBirthday = nook.Birthday{
		Day:   9,
		Month: time.January}
)

var (
	// tiffanyCode represents Tiffany's unique code.
	tiffanyCode = nook.Code{
		Value: "rbt07"}
)

var (
	// tiffanyAmericanEnglishName represents Tiffany's name in American English.
	tiffanyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tiffany"}

	// tiffanyCanadianFrenchName represents Tiffany's name in Canadian French.
	tiffanyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tiphaine"}

	// tiffanyDutchName represents Tiffany's name in Dutch.
	tiffanyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tiffany"}

	// tiffanyFrenchName represents Tiffany's name in French.
	tiffanyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tiphaine"}

	// tiffanyGermanName represents Tiffany's name in German.
	tiffanyGermanName = nook.Name{
		Language: language.German,
		Value:    "Michelle"}

	// tiffanyItalianName represents Tiffany's name in Italian.
	tiffanyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Stefania"}

	// tiffanyJapaneseName represents Tiffany's name in Japanese.
	tiffanyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "バズレー"}

	// tiffanyKoreanName represents Tiffany's name in Korean.
	tiffanyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "바슬레"}

	// tiffanyLatinAmericanSpanishName represents Tiffany's name in Latin American Spanish.
	tiffanyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tiffany"}

	// tiffanyRussianName represents Tiffany's name in Russian.
	tiffanyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тиффани"}

	// tiffanySimplifiedChineseName represents Tiffany's name in Simplified Chinese.
	tiffanySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "大姐头"}

	// tiffanySpanishName represents Tiffany's name in Spanish.
	tiffanySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tiffany"}

	// tiffanyTraditionalChineseName represents Tiffany's name in Traditional Chinese.
	tiffanyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "大姐頭"}
)

var (
	// tiffanyName represents Tiffany's name in different languages.
	tiffanyName = nook.Languages{
		language.AmericanEnglish:      tiffanyAmericanEnglishName,
		language.CanadianFrench:       tiffanyCanadianFrenchName,
		language.Dutch:                tiffanyDutchName,
		language.French:               tiffanyFrenchName,
		language.German:               tiffanyGermanName,
		language.Italian:              tiffanyItalianName,
		language.Japanese:             tiffanyJapaneseName,
		language.Korean:               tiffanyKoreanName,
		language.LatinAmericanSpanish: tiffanyLatinAmericanSpanishName,
		language.Russian:              tiffanyRussianName,
		language.SimplifiedChinese:    tiffanySimplifiedChineseName,
		language.Spanish:              tiffanySpanishName,
		language.TraditionalChinese:   tiffanyTraditionalChineseName}
)

var (
	// tiffanyCharacter represents Tiffany's character information.
	tiffanyCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: tiffanyBirthday,
		Code:     tiffanyCode,
		Key:      character.Tiffany,
		Gender:   gender.Female,
		Name:     tiffanyName,
		Special:  false}
)

var (
	// tiffanyAmericanEnglishPhrase represents Tiffany's phrase in American English.
	tiffanyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bunbun"}

	// tiffanyCanadianFrenchPhrase represents Tiffany's phrase in Canadian French.
	tiffanyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pinpin"}

	// tiffanyDutchPhrase represents Tiffany's phrase in Dutch.
	tiffanyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "knijnelijn"}

	// tiffanyFrenchPhrase represents Tiffany's phrase in French.
	tiffanyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pinpin"}

	// tiffanyGermanPhrase represents Tiffany's phrase in German.
	tiffanyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "öhrchen"}

	// tiffanyItalianPhrase represents Tiffany's phrase in Italian.
	tiffanyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bunbun"}

	// tiffanyJapanesePhrase represents Tiffany's phrase in Japanese.
	tiffanyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ってさ"}

	// tiffanyKoreanPhrase represents Tiffany's phrase in Korean.
	tiffanyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "라던데"}

	// tiffanyLatinAmericanSpanishPhrase represents Tiffany's phrase in Latin American Spanish.
	tiffanyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "toin-lalá"}

	// tiffanyRussianPhrase represents Tiffany's phrase in Russian.
	tiffanyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "зайчонок"}

	// tiffanySimplifiedChinesePhrase represents Tiffany's phrase in Simplified Chinese.
	tiffanySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "我说呀"}

	// tiffanySpanishPhrase represents Tiffany's phrase in Spanish.
	tiffanySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pompón"}

	// tiffanyTraditionalChinesePhrase represents Tiffany's phrase in Traditional Chinese.
	tiffanyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "我說呀"}
)

var (
	// tiffanyPhrase represents Tiffany's phrases in different languages.
	tiffanyPhrase = nook.Languages{
		language.AmericanEnglish:      tiffanyAmericanEnglishPhrase,
		language.CanadianFrench:       tiffanyCanadianFrenchPhrase,
		language.Dutch:                tiffanyDutchPhrase,
		language.French:               tiffanyFrenchPhrase,
		language.German:               tiffanyGermanPhrase,
		language.Italian:              tiffanyItalianPhrase,
		language.Japanese:             tiffanyJapanesePhrase,
		language.Korean:               tiffanyKoreanPhrase,
		language.LatinAmericanSpanish: tiffanyLatinAmericanSpanishPhrase,
		language.Russian:              tiffanyRussianPhrase,
		language.SimplifiedChinese:    tiffanySimplifiedChinesePhrase,
		language.Spanish:              tiffanySpanishPhrase,
		language.TraditionalChinese:   tiffanyTraditionalChinesePhrase}
)

var (
	// Tiffany represents the character Tiffany with her complete information.
	Tiffany = nook.Villager{
		Character:   tiffanyCharacter,
		Personality: personality.Snooty,
		Phrase:      tiffanyPhrase}
)
