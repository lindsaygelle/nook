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
	// woolioBirthday represents Woolio's birthday (Month and Day values are placeholders).
	woolioBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// woolioCode represents Woolio's unique code (empty string for now).
	woolioCode = nook.Code{
		Value: ""}
)

var (
	// woolioAmericanEnglishName represents Woolio's name in American English.
	woolioAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Woolio"}

	// woolioCanadianFrenchName represents Woolio's name in Canadian French (empty string for now).
	woolioCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// woolioDutchName represents Woolio's name in Dutch (empty string for now).
	woolioDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// woolioFrenchName represents Woolio's name in French.
	woolioFrenchName = nook.Name{
		Language: language.French,
		Value:    "Moumoute"}

	// woolioGermanName represents Woolio's name in German.
	woolioGermanName = nook.Name{
		Language: language.German,
		Value:    "Wulle"}

	// woolioItalianName represents Woolio's name in Italian.
	woolioItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Moero"}

	// woolioJapaneseName represents Woolio's name in Japanese.
	woolioJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ホシオ"}

	// woolioKoreanName represents Woolio's name in Korean (empty string for now).
	woolioKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// woolioLatinAmericanSpanishName represents Woolio's name in Latin American Spanish (empty string for now).
	woolioLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// woolioRussianName represents Woolio's name in Russian (empty string for now).
	woolioRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// woolioSimplifiedChineseName represents Woolio's name in Simplified Chinese.
	woolioSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "星星"}

	// woolioSpanishName represents Woolio's name in Spanish.
	woolioSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Esquilo"}

	// woolioTraditionalChineseName represents Woolio's name in Traditional Chinese (empty string for now).
	woolioTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// woolioName represents Woolio's name in different languages.
	woolioName = nook.Languages{
		language.AmericanEnglish:      woolioAmericanEnglishName,
		language.CanadianFrench:       woolioCanadianFrenchName,
		language.Dutch:                woolioDutchName,
		language.French:               woolioFrenchName,
		language.German:               woolioGermanName,
		language.Italian:              woolioItalianName,
		language.Japanese:             woolioJapaneseName,
		language.Korean:               woolioKoreanName,
		language.LatinAmericanSpanish: woolioLatinAmericanSpanishName,
		language.Russian:              woolioRussianName,
		language.SimplifiedChinese:    woolioSimplifiedChineseName,
		language.Spanish:              woolioSpanishName,
		language.TraditionalChinese:   woolioTraditionalChineseName}
)

var (
	// woolioCharacter represents Woolio's character information.
	woolioCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: woolioBirthday,
		Code:     woolioCode,
		Key:      character.Woolio,
		Gender:   gender.Male,
		Name:     woolioName,
		Special:  false}
)

var (
	// woolioAmericanEnglishPhrase represents Woolio's phrase in American English.
	woolioAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "biz-aaa"}

	// woolioCanadianFrenchPhrase represents Woolio's phrase in Canadian French (empty string for now).
	woolioCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// woolioDutchPhrase represents Woolio's phrase in Dutch (empty string for now).
	woolioDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// woolioFrenchPhrase represents Woolio's phrase in French.
	woolioFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bêêêê quoi"}

	// woolioGermanPhrase represents Woolio's phrase in German.
	woolioGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bizääää"}

	// woolioItalianPhrase represents Woolio's phrase in Italian.
	woolioItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sbaallo"}

	// woolioJapanesePhrase represents Woolio's phrase in Japanese.
	woolioJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヨロシク"}

	// woolioKoreanPhrase represents Woolio's phrase in Korean (empty string for now).
	woolioKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// woolioLatinAmericanSpanishPhrase represents Woolio's phrase in Latin American Spanish (empty string for now).
	woolioLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// woolioRussianPhrase represents Woolio's phrase in Russian (empty string for now).
	woolioRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// woolioSimplifiedChinesePhrase represents Woolio's phrase in Simplified Chinese.
	woolioSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "帮忙"}

	// woolioSpanishPhrase represents Woolio's phrase in Spanish.
	woolioSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "paaasssa"}

	// woolioTraditionalChinesePhrase represents Woolio's phrase in Traditional Chinese (empty string for now).
	woolioTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// woolioPhrase represents Woolio's phrases in different languages.
	woolioPhrase = nook.Languages{
		language.AmericanEnglish:      woolioAmericanEnglishPhrase,
		language.CanadianFrench:       woolioCanadianFrenchPhrase,
		language.Dutch:                woolioDutchPhrase,
		language.French:               woolioFrenchPhrase,
		language.German:               woolioGermanPhrase,
		language.Italian:              woolioItalianPhrase,
		language.Japanese:             woolioJapanesePhrase,
		language.Korean:               woolioKoreanPhrase,
		language.LatinAmericanSpanish: woolioLatinAmericanSpanishPhrase,
		language.Russian:              woolioRussianPhrase,
		language.SimplifiedChinese:    woolioSimplifiedChinesePhrase,
		language.Spanish:              woolioSpanishPhrase,
		language.TraditionalChinese:   woolioTraditionalChinesePhrase}
)

var (
	// Woolio represents the character Woolio with his complete information.
	Woolio = nook.Villager{
		Character:   woolioCharacter,
		Personality: personality.Jock,
		Phrase:      woolioPhrase}
)
