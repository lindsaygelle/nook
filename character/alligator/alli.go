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

// alliBirthday represents Alli's birthday (November 8th).
var (
	alliBirthday = nook.Birthday{
		Day:   8,
		Month: time.November}
)

// alliCode represents Alli's unique code ("crd01").
var (
	alliCode = nook.Code{
		Value: "crd01"}
)

// Different names for Alli in various languages.
var (
	alliAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Alli"}

	alliCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Allie"}

	alliDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Alli"}

	alliFrenchName = nook.Name{
		Language: language.French,
		Value:    "Allie"}

	alliGermanName = nook.Name{
		Language: language.German,
		Value:    "Ali"}

	alliItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Alli"}

	alliJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クロコ"}

	alliKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "크로크"}

	alliLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Codrila"}

	alliRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Элли"}

	alliSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鳄罗思"}

	alliSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Codrila"}

	alliTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鱷羅思"}
)

// alliName represents Alli's name in different languages.
var (
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

// alliCharacter represents Alli's character information.
var (
	alliCharacter = nook.Character{
		Animal:   animal.Alligator,
		Birthday: alliBirthday,
		Code:     alliCode,
		Key:      character.Alli,
		Gender:   gender.Female,
		Name:     alliName,
		Special:  false}
)

// Different phrases spoken by Alli in various languages.
var (
	alliAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "graaagh"}

	alliCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "graaagh"}

	alliDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "graaach"}

	alliFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "graaagh"}

	alliGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "graaach"}

	alliItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "graaag"}

	alliJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "どすえ"}

	alliKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "얘야"}

	alliLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "muamú"}

	alliRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "гра-а-а"}

	alliSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鳄鱼皮"}

	alliSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "lagartija"}

	alliTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鱷魚皮"}
)

// alliPhrase represents Alli's phrases in different languages.
var (
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

// Alli represents the character Alli with her complete information.
var (
	Alli = nook.Villager{
		Character:   alliCharacter,
		Personality: personality.Snooty,
		Phrase:      alliPhrase}
)
