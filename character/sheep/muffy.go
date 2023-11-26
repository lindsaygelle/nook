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
	// muffyBirthday represents Muffy's birthday.
	muffyBirthday = nook.Birthday{
		Day:   14,
		Month: time.February}
)

var (
	// muffyCode represents Muffy's unique code.
	muffyCode = nook.Code{
		Value: "shp12"}
)

var (
	// muffyAmericanEnglishName represents Muffy's name in American English.
	muffyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Muffy"}

	// muffyCanadianFrenchName represents Muffy's name in Canadian French.
	muffyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Charlène"}

	// muffyDutchName represents Muffy's name in Dutch.
	muffyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Muffy"}

	// muffyFrenchName represents Muffy's name in French.
	muffyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Charlène"}

	// muffyGermanName represents Muffy's name in German.
	muffyGermanName = nook.Name{
		Language: language.German,
		Value:    "Marion"}

	// muffyItalianName represents Muffy's name in Italian.
	muffyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Morena"}

	// muffyJapaneseName represents Muffy's name in Japanese.
	muffyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フリル"}

	// muffyKoreanName represents Muffy's name in Korean.
	muffyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "프릴"}

	// muffyLatinAmericanSpanishName represents Muffy's name in Latin American Spanish.
	muffyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Muaré"}

	// muffyRussianName represents Muffy's name in Russian.
	muffyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Маффи"}

	// muffySimplifiedChineseName represents Muffy's name in Simplified Chinese.
	muffySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "彭澎"}

	// muffySpanishName represents Muffy's name in Spanish.
	muffySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Muaré"}

	// muffyTraditionalChineseName represents Muffy's name in Traditional Chinese.
	muffyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "彭澎"}
)

var (
	// muffyName represents Muffy's name in different languages.
	muffyName = nook.Languages{
		language.AmericanEnglish:      muffyAmericanEnglishName,
		language.CanadianFrench:       muffyCanadianFrenchName,
		language.Dutch:                muffyDutchName,
		language.French:               muffyFrenchName,
		language.German:               muffyGermanName,
		language.Italian:              muffyItalianName,
		language.Japanese:             muffyJapaneseName,
		language.Korean:               muffyKoreanName,
		language.LatinAmericanSpanish: muffyLatinAmericanSpanishName,
		language.Russian:              muffyRussianName,
		language.SimplifiedChinese:    muffySimplifiedChineseName,
		language.Spanish:              muffySpanishName,
		language.TraditionalChinese:   muffyTraditionalChineseName}
)

var (
	// muffyCharacter represents Muffy's character information.
	muffyCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: muffyBirthday,
		Code:     muffyCode,
		Key:      character.Muffy,
		Gender:   gender.Female,
		Name:     muffyName,
		Special:  false}
)

var (
	// muffyAmericanEnglishPhrase represents Muffy's phrase in American English.
	muffyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nightshade"}

	// muffyCanadianFrenchPhrase represents Muffy's phrase in Canadian French.
	muffyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tricoti"}

	// muffyDutchPhrase represents Muffy's phrase in Dutch.
	muffyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "nachtvacht"}

	// muffyFrenchPhrase represents Muffy's phrase in French.
	muffyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tricoti"}

	// muffyGermanPhrase represents Muffy's phrase in German.
	muffyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knuffi"}

	// muffyItalianPhrase represents Muffy's phrase in Italian.
	muffyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "lanadà"}

	// muffyJapanesePhrase represents Muffy's phrase in Japanese.
	muffyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "おつかれ"}

	// muffyKoreanPhrase represents Muffy's phrase in Korean.
	muffyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "수고"}

	// muffyLatinAmericanSpanishPhrase represents Muffy's phrase in Latin American Spanish.
	muffyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "lanavá"}

	// muffyRussianPhrase represents Muffy's phrase in Russian.
	muffyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "травинка"}

	// muffySimplifiedChinesePhrase represents Muffy's phrase in Simplified Chinese.
	muffySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "辛苦了"}

	// muffySpanishPhrase represents Muffy's phrase in Spanish.
	muffySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "lanas"}

	// muffyTraditionalChinesePhrase represents Muffy's phrase in Traditional Chinese.
	muffyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "辛苦了"}
)

var (
	// muffyPhrase represents Muffy's phrases in different languages.
	muffyPhrase = nook.Languages{
		language.AmericanEnglish:      muffyAmericanEnglishPhrase,
		language.CanadianFrench:       muffyCanadianFrenchPhrase,
		language.Dutch:                muffyDutchPhrase,
		language.French:               muffyFrenchPhrase,
		language.German:               muffyGermanPhrase,
		language.Italian:              muffyItalianPhrase,
		language.Japanese:             muffyJapanesePhrase,
		language.Korean:               muffyKoreanPhrase,
		language.LatinAmericanSpanish: muffyLatinAmericanSpanishPhrase,
		language.Russian:              muffyRussianPhrase,
		language.SimplifiedChinese:    muffySimplifiedChinesePhrase,
		language.Spanish:              muffySpanishPhrase,
		language.TraditionalChinese:   muffyTraditionalChinesePhrase}
)

var (
	// Muffy represents the character Muffy with her complete information.
	Muffy = nook.Villager{
		Character:   muffyCharacter,
		Personality: personality.BigSister,
		Phrase:      muffyPhrase}
)
