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
	hopkinsBirthday = nook.Birthday{
		Day:   11,
		Month: time.March}
)

var (
	hopkinsCode = nook.Code{
		Value: "rbt14"}
)

var (
	hopkinsAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hopkins"}

	hopkinsCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Grignote"}

	hopkinsDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hopkins"}

	hopkinsFrenchName = nook.Name{
		Language: language.French,
		Value:    "Grignote"}

	hopkinsGermanName = nook.Name{
		Language: language.German,
		Value:    "Poldi"}

	hopkinsItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Azeglio"}

	hopkinsJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "プースケ"}

	hopkinsKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "홉킨스"}

	hopkinsLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Saltiago"}

	hopkinsRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Хопкинс"}

	hopkinsSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "风杰"}

	hopkinsSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Saltiago"}

	hopkinsTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "風杰"}
)

var (
	hopkinsName = nook.Languages{
		language.AmericanEnglish:      hopkinsAmericanEnglishName,
		language.CanadianFrench:       hopkinsCanadianFrenchName,
		language.Dutch:                hopkinsDutchName,
		language.French:               hopkinsFrenchName,
		language.German:               hopkinsGermanName,
		language.Italian:              hopkinsItalianName,
		language.Japanese:             hopkinsJapaneseName,
		language.Korean:               hopkinsKoreanName,
		language.LatinAmericanSpanish: hopkinsLatinAmericanSpanishName,
		language.Russian:              hopkinsRussianName,
		language.SimplifiedChinese:    hopkinsSimplifiedChineseName,
		language.Spanish:              hopkinsSpanishName,
		language.TraditionalChinese:   hopkinsTraditionalChineseName}
)

var (
	hopkinsCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: hopkinsBirthday,
		Code:     hopkinsCode,
		Key:      character.Hopkins,
		Gender:   gender.Male,
		Name:     hopkinsName}
)

var (
	hopkinsAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "thumper"}

	hopkinsCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pfui"}

	hopkinsDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "stamper"}

	hopkinsFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pfui"}

	hopkinsGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "puschel"}

	hopkinsItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "siiigh"}

	hopkinsJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ぷぅ"}

	hopkinsKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "뿌우"}

	hopkinsLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fiiiiú"}

	hopkinsRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "топ-топ"}

	hopkinsSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "风"}

	hopkinsSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "vida extra"}

	hopkinsTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "風"}
)

var (
	hopkinsPhrase = nook.Languages{
		language.AmericanEnglish:      hopkinsAmericanEnglishName,
		language.CanadianFrench:       hopkinsCanadianFrenchName,
		language.Dutch:                hopkinsDutchName,
		language.French:               hopkinsFrenchName,
		language.German:               hopkinsGermanName,
		language.Italian:              hopkinsItalianName,
		language.Japanese:             hopkinsJapaneseName,
		language.Korean:               hopkinsKoreanName,
		language.LatinAmericanSpanish: hopkinsLatinAmericanSpanishName,
		language.Russian:              hopkinsRussianName,
		language.SimplifiedChinese:    hopkinsSimplifiedChineseName,
		language.Spanish:              hopkinsSpanishName,
		language.TraditionalChinese:   hopkinsTraditionalChineseName}
)

var (
	Hopkins = nook.Villager{
		Character:   hopkinsCharacter,
		Personality: personality.Lazy,
		Phrase:      hopkinsPhrase}
)
