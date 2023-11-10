package alligator

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
	// alliBirthday represents Alli's birthday (November 8th).
	alliBirthday = nook.Birthday{
		Day:   8,
		Month: time.November}
)

var (
	// alliCode represents Alli's unique code ("crd01").
	alliCode = nook.Code{
		Value: "crd01"}
)

var (
	// alliAmericanEnglishName represents Alli's name in American English.
	alliAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Alli"}

	// alliCanadianFrenchName represents Alli's name in Canadian French.
	alliCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Allie"}

	// alliDutchName represents Alli's name in Dutch.
	alliDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Alli"}

	// alliFrenchName represents Alli's name in French.
	alliFrenchName = nook.Name{
		Language: language.French,
		Value:    "Allie"}

	// alliGermanName represents Alli's name in German.
	alliGermanName = nook.Name{
		Language: language.German,
		Value:    "Ali"}

	// alliItalianName represents Alli's name in Italian.
	alliItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Alli"}

	// alliJapaneseName represents Alli's name in Japanese.
	alliJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クロコ"}

	// alliKoreanName represents Alli's name in Korean.
	alliKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "크로크"}

	// alliLatinAmericanSpanishName represents Alli's name in Latin American Spanish.
	alliLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Codrila"}

	// alliRussianName represents Alli's name in Russian.
	alliRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Элли"}

	// alliSimplifiedChineseName represents Alli's name in Simplified Chinese.
	alliSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鳄罗思"}

	// alliSpanishName represents Alli's name in Spanish.
	alliSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Codrila"}

	// alliTraditionalChineseName represents Alli's name in Traditional Chinese.
	alliTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鱷羅思"}
)

var (
	// alliName represents Alli's name in different languages.
	alliName = nook.Languages{
		language.AmericanEnglish:      alliAmericanEnglishName,
		language.CanadianFrench:       alliCanadianFrenchName,
		language.Dutch:                alliDutchName,
		language.French:               alliFrenchName,
		language.German:               alliGermanName,
		language.Italian:              alliItalianName,
		language.Japanese:             alliJapaneseName,
		language.Korean:               alliKoreanName,
		language.LatinAmericanSpanish: alliLatinAmericanSpanishName,
		language.Russian:              alliRussianName,
		language.SimplifiedChinese:    alliSimplifiedChineseName,
		language.Spanish:              alliSpanishName,
		language.TraditionalChinese:   alliTraditionalChineseName}
)

var (
	// alliCharacter represents Alli's character information.
	alliCharacter = nook.Character{
		Animal:   animal.Alligator,
		Birthday: alliBirthday,
		Code:     alliCode,
		Key:      character.Alli,
		Gender:   gender.Female,
		Name:     alliName,
		Special:  false}
)

var (
	// alliAmericanEnglishPhrase represents Alli's phrase in American English.
	alliAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "graaagh"}

	// alliCanadianFrenchPhrase represents Alli's phrase in Canadian French.
	alliCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "graaagh"}

	// alliDutchPhrase represents Alli's phrase in Dutch.
	alliDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "graaach"}

	// alliFrenchPhrase represents Alli's phrase in French.
	alliFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "graaagh"}

	// alliGermanPhrase represents Alli's phrase in German.
	alliGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "graaach"}

	// alliItalianPhrase represents Alli's phrase in Italian.
	alliItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "graaag"}

	// alliJapanesePhrase represents Alli's phrase in Japanese.
	alliJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "どすえ"}

	// alliKoreanPhrase represents Alli's phrase in Korean.
	alliKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "얘야"}

	// alliLatinAmericanSpanishPhrase represents Alli's phrase in Latin American Spanish.
	alliLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "muamú"}

	// alliRussianPhrase represents Alli's phrase in Russian.
	alliRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "гра-а-а"}

	// alliSimplifiedChinesePhrase represents Alli's phrase in Simplified Chinese.
	alliSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鳄鱼皮"}

	// alliSpanishPhrase represents Alli's phrase in Spanish.
	alliSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "lagartija"}

	// alliTraditionalChinesePhrase represents Alli's phrase in Traditional Chinese.
	alliTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鱷魚皮"}
)

var (
	// alliPhrase represents Alli's phrases in different languages.
	alliPhrase = nook.Languages{
		language.AmericanEnglish:      alliAmericanEnglishPhrase,
		language.CanadianFrench:       alliCanadianFrenchPhrase,
		language.Dutch:                alliDutchPhrase,
		language.French:               alliFrenchPhrase,
		language.German:               alliGermanPhrase,
		language.Italian:              alliItalianPhrase,
		language.Japanese:             alliJapanesePhrase,
		language.Korean:               alliKoreanPhrase,
		language.LatinAmericanSpanish: alliLatinAmericanSpanishPhrase,
		language.Russian:              alliRussianPhrase,
		language.SimplifiedChinese:    alliSimplifiedChinesePhrase,
		language.Spanish:              alliSpanishPhrase,
		language.TraditionalChinese:   alliTraditionalChinesePhrase}
)

var (
	// Alli represents the character Alli with her complete information.
	Alli = nook.Villager{
		Character:   alliCharacter,
		Personality: personality.Snooty,
		Phrase:      alliPhrase}
)
