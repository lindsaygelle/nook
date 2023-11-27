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
	// tashaBirthday represents Tasha's birthday.
	tashaBirthday = nook.Birthday{
		Day:   30,
		Month: time.November}
)

var (
	// tashaCode represents Tasha's unique code.
	tashaCode = nook.Code{
		Value: "squ13"}
)

var (
	// tashaAmericanEnglishName represents Tasha's name in American English.
	tashaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tasha"}

	// tashaCanadianFrenchName represents Tasha's name in Canadian French.
	tashaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Nadeige"}

	// tashaDutchName represents Tasha's name in Dutch.
	tashaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tasha"}

	// tashaFrenchName represents Tasha's name in French.
	tashaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Nadeige"}

	// tashaGermanName represents Tasha's name in German.
	tashaGermanName = nook.Name{
		Language: language.German,
		Value:    "Natalja"}

	// tashaItalianName represents Tasha's name in Italian.
	tashaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Teresa"}

	// tashaJapaneseName represents Tasha's name in Japanese.
	tashaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ナターシャ"}

	// tashaKoreanName represents Tasha's name in Korean.
	tashaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "나타샤"}

	// tashaLatinAmericanSpanishName represents Tasha's name in Latin American Spanish.
	tashaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tania"}

	// tashaRussianName represents Tasha's name in Russian.
	tashaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Таша"}

	// tashaSimplifiedChineseName represents Tasha's name in Simplified Chinese.
	tashaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "娜塔丽"}

	// tashaSpanishName represents Tasha's name in Spanish.
	tashaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tania"}

	// tashaTraditionalChineseName represents Tasha's name in Traditional Chinese.
	tashaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "娜塔麗"}
)

var (
	// tashaName represents Tasha's name in different languages.
	tashaName = nook.Languages{
		language.AmericanEnglish:      tashaAmericanEnglishName,
		language.CanadianFrench:       tashaCanadianFrenchName,
		language.Dutch:                tashaDutchName,
		language.French:               tashaFrenchName,
		language.German:               tashaGermanName,
		language.Italian:              tashaItalianName,
		language.Japanese:             tashaJapaneseName,
		language.Korean:               tashaKoreanName,
		language.LatinAmericanSpanish: tashaLatinAmericanSpanishName,
		language.Russian:              tashaRussianName,
		language.SimplifiedChinese:    tashaSimplifiedChineseName,
		language.Spanish:              tashaSpanishName,
		language.TraditionalChinese:   tashaTraditionalChineseName}
)

var (
	// tashaCharacter represents Tasha's character information.
	tashaCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: tashaBirthday,
		Code:     tashaCode,
		Key:      character.Tasha,
		Gender:   gender.Female,
		Name:     tashaName,
		Special:  false}
)

var (

	// tashaAmericanEnglishPhrase represents Tasha's phrase in American English.
	tashaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nice nice"}

	// tashaCanadianFrenchPhrase represents Tasha's phrase in Canadian French.
	tashaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ptigri"}

	// tashaDutchPhrase represents Tasha's phrase in Dutch.
	tashaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "jaja"}

	// tashaFrenchPhrase represents Tasha's phrase in French.
	tashaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ptigri"}

	// tashaGermanPhrase represents Tasha's phrase in German.
	tashaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tjaja"}

	// tashaItalianPhrase represents Tasha's phrase in Italian.
	tashaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "munch"}

	// tashaJapanesePhrase represents Tasha's phrase in Japanese.
	tashaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やるわね"}

	// tashaKoreanPhrase represents Tasha's phrase in Korean.
	tashaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "굿굿"}

	// tashaLatinAmericanSpanishPhrase represents Tasha's phrase in Latin American Spanish.
	tashaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bi-buá"}

	// tashaRussianPhrase represents Tasha's phrase in Russian.
	tashaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ладно"}

	// tashaSimplifiedChinesePhrase represents Tasha's phrase in Simplified Chinese.
	tashaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "做得好"}

	// tashaSpanishPhrase represents Tasha's phrase in Spanish.
	tashaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pincel"}

	// tashaTraditionalChinesePhrase represents Tasha's phrase in Traditional Chinese.
	tashaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "做得好"}
)

var (
	// tashaPhrase represents Tasha's phrase in different languages.
	tashaPhrase = nook.Languages{
		language.AmericanEnglish:      tashaAmericanEnglishPhrase,
		language.CanadianFrench:       tashaCanadianFrenchPhrase,
		language.Dutch:                tashaDutchPhrase,
		language.French:               tashaFrenchPhrase,
		language.German:               tashaGermanPhrase,
		language.Italian:              tashaItalianPhrase,
		language.Japanese:             tashaJapanesePhrase,
		language.Korean:               tashaKoreanPhrase,
		language.LatinAmericanSpanish: tashaLatinAmericanSpanishPhrase,
		language.Russian:              tashaRussianPhrase,
		language.SimplifiedChinese:    tashaSimplifiedChinesePhrase,
		language.Spanish:              tashaSpanishPhrase,
		language.TraditionalChinese:   tashaTraditionalChinesePhrase}
)

var (
	// Tasha represents the character Tasha with her complete information.
	Tasha = nook.Villager{
		Character:   tashaCharacter,
		Personality: personality.Snooty,
		Phrase:      tashaPhrase}
)
