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
	// hopkinsBirthday represents Hopkins's birthday.
	hopkinsBirthday = nook.Birthday{
		Day:   11,
		Month: time.March}
)

var (
	// hopkinsCode represents Hopkins's unique code ("rbt14").
	hopkinsCode = nook.Code{
		Value: "rbt14"}
)

var (
	// hopkinsAmericanEnglishName represents Hopkins's name in American English.
	hopkinsAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hopkins"}

	// hopkinsCanadianFrenchName represents Hopkins's name in Canadian French.
	hopkinsCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Grignote"}

	// hopkinsDutchName represents Hopkins's name in Dutch.
	hopkinsDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hopkins"}

	// hopkinsFrenchName represents Hopkins's name in French.
	hopkinsFrenchName = nook.Name{
		Language: language.French,
		Value:    "Grignote"}

	// hopkinsGermanName represents Hopkins's name in German.
	hopkinsGermanName = nook.Name{
		Language: language.German,
		Value:    "Poldi"}

	// hopkinsItalianName represents Hopkins's name in Italian.
	hopkinsItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Azeglio"}

	// hopkinsJapaneseName represents Hopkins's name in Japanese.
	hopkinsJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "プースケ"}

	// hopkinsKoreanName represents Hopkins's name in Korean.
	hopkinsKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "홉킨스"}

	// hopkinsLatinAmericanSpanishName represents Hopkins's name in Latin American Spanish.
	hopkinsLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Saltiago"}

	// hopkinsRussianName represents Hopkins's name in Russian.
	hopkinsRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Хопкинс"}

	// hopkinsSimplifiedChineseName represents Hopkins's name in Simplified Chinese.
	hopkinsSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "风杰"}

	// hopkinsSpanishName represents Hopkins's name in Spanish.
	hopkinsSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Saltiago"}

	// hopkinsTraditionalChineseName represents Hopkins's name in Traditional Chinese.
	hopkinsTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "風杰"}
)

var (
	// hopkinsName represents Hopkins's name in different languages.
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
	// hopkinsCharacter represents Hopkins's character information.
	hopkinsCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: hopkinsBirthday,
		Code:     hopkinsCode,
		Key:      character.Hopkins,
		Gender:   gender.Male,
		Name:     hopkinsName,
		Special:  false}
)

var (
	// hopkinsAmericanEnglishPhrase represents Hopkins's phrase in American English.
	hopkinsAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "thumper"}

	// hopkinsCanadianFrenchPhrase represents Hopkins's phrase in Canadian French.
	hopkinsCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pfui"}

	// hopkinsDutchPhrase represents Hopkins's phrase in Dutch.
	hopkinsDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "stamper"}

	// hopkinsFrenchPhrase represents Hopkins's phrase in French.
	hopkinsFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pfui"}

	// hopkinsGermanPhrase represents Hopkins's phrase in German.
	hopkinsGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "puschel"}

	// hopkinsItalianPhrase represents Hopkins's phrase in Italian.
	hopkinsItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "siiigh"}

	// hopkinsJapanesePhrase represents Hopkins's phrase in Japanese.
	hopkinsJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ぷぅ"}

	// hopkinsKoreanPhrase represents Hopkins's phrase in Korean.
	hopkinsKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "뿌우"}

	// hopkinsLatinAmericanSpanishPhrase represents Hopkins's phrase in Latin American Spanish.
	hopkinsLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fiiiiú"}

	// hopkinsRussianPhrase represents Hopkins's phrase in Russian.
	hopkinsRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "топ-топ"}

	// hopkinsSimplifiedChinesePhrase represents Hopkins's phrase in Simplified Chinese.
	hopkinsSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "风"}

	// hopkinsSpanishPhrase represents Hopkins's phrase in Spanish.
	hopkinsSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "vida extra"}

	// hopkinsTraditionalChinesePhrase represents Hopkins's phrase in Traditional Chinese.
	hopkinsTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "風"}
)

var (
	// hopkinsPhrase represents Hopkins's phrases in different languages.
	hopkinsPhrase = nook.Languages{
		language.AmericanEnglish:      hopkinsAmericanEnglishPhrase,
		language.CanadianFrench:       hopkinsCanadianFrenchPhrase,
		language.Dutch:                hopkinsDutchPhrase,
		language.French:               hopkinsFrenchPhrase,
		language.German:               hopkinsGermanPhrase,
		language.Italian:              hopkinsItalianPhrase,
		language.Japanese:             hopkinsJapanesePhrase,
		language.Korean:               hopkinsKoreanPhrase,
		language.LatinAmericanSpanish: hopkinsLatinAmericanSpanishPhrase,
		language.Russian:              hopkinsRussianPhrase,
		language.SimplifiedChinese:    hopkinsSimplifiedChinesePhrase,
		language.Spanish:              hopkinsSpanishPhrase,
		language.TraditionalChinese:   hopkinsTraditionalChinesePhrase}
)

var (
	// Hopkins represents the character Hopkins with his complete information.
	Hopkins = nook.Villager{
		Character:   hopkinsCharacter,
		Personality: personality.Lazy,
		Phrase:      hopkinsPhrase}
)
