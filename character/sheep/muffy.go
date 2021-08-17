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
	muffyBirthday = nook.Birthday{
		Day:   14,
		Month: time.February}
)

var (
	muffyCode = nook.Code{
		Value: "shp12"}
)

var (
	muffyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Muffy"}

	muffyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Charlène"}

	muffyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Muffy"}

	muffyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Charlène"}

	muffyGermanName = nook.Name{
		Language: language.German,
		Value:    "Marion"}

	muffyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Morena"}

	muffyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フリル"}

	muffyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "프릴"}

	muffyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Muaré"}

	muffyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Маффи"}

	muffySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "彭澎"}

	muffySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Muaré"}

	muffyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "彭澎"}
)

var (
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
	muffyCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: muffyBirthday,
		Code:     muffyCode,
		Key:      character.Muffy,
		Gender:   gender.Female,
		Name:     muffyName}
)

var (
	muffyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nightshade"}

	muffyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tricoti"}

	muffyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "nachtvacht"}

	muffyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tricoti"}

	muffyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knuffi"}

	muffyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "lanadà"}

	muffyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "おつかれ"}

	muffyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "수고"}

	muffyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "lanavá"}

	muffyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "травинка"}

	muffySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "辛苦了"}

	muffySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "lanas"}

	muffyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "辛苦了"}
)

var (
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
	Muffy = nook.Villager{
		Character:   muffyCharacter,
		Personality: personality.BigSister,
		Phrase:      muffyPhrase}
)
