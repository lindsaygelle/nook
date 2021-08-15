package sheep

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Charlènetricoti"}

	muffyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Muffynachtvacht"}

	muffyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Charlènetricoti"}

	muffyGermanName = nook.Name{
		Language: language.German,
		Value:    "Marionknuffi"}

	muffyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Morenalanadà"}

	muffyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フリルおつかれ"}

	muffyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "프릴수고"}

	muffyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Muarélanavá"}

	muffyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Маффитравинка"}

	muffySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "彭澎辛苦了"}

	muffySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Muarélanas"}

	muffyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "彭澎辛苦了"}
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
		Animal:   Sheep,
		Birthday: muffyBirthday,
		Code:     muffyCode,
		Gender:   nook.Female,
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
	Muffy = nook.Villager{
		Character:   muffyCharacter,
		Personality: nook.BigSister,
		Phrase:      muffyPhrase}
)
