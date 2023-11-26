package tiger

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
	// bangleBirthday represents Bangle's birthday.
	bangleBirthday = nook.Birthday{
		Day:   27,
		Month: time.August}
)

var (
	// bangleCode represents Bangle's unique code.
	bangleCode = nook.Code{
		Value: "tig03"}
)

var (
	// bangleAmericanEnglishName represents Bangle's name in American English.
	bangleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bangle"}

	// bangleCanadianFrenchName represents Bangle's name in Canadian French.
	bangleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bengale"}

	// bangleDutchName represents Bangle's name in Dutch.
	bangleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bangle"}

	// bangleFrenchName represents Bangle's name in French.
	bangleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bengale"}

	// bangleGermanName represents Bangle's name in German.
	bangleGermanName = nook.Name{
		Language: language.German,
		Value:    "Tamara"}

	// bangleItalianName represents Bangle's name in Italian.
	bangleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tigrizia"}

	// bangleJapaneseName represents Bangle's name in Japanese.
	bangleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ルーズ"}

	// bangleKoreanName represents Bangle's name in Korean.
	bangleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "루주"}

	// bangleLatinAmericanSpanishName represents Bangle's name in Latin American Spanish.
	bangleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Felina"}

	// bangleRussianName represents Bangle's name in Russian.
	bangleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бэнгл"}

	// bangleSimplifiedChineseName represents Bangle's name in Simplified Chinese.
	bangleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "卢姿"}

	// bangleSpanishName represents Bangle's name in Spanish.
	bangleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Felina"}

	// bangleTraditionalChineseName represents Bangle's name in Traditional Chinese.
	bangleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "盧姿"}
)

var (
	// bangleName represents Bangle's name in different languages.
	bangleName = nook.Languages{
		language.AmericanEnglish:      bangleAmericanEnglishName,
		language.CanadianFrench:       bangleCanadianFrenchName,
		language.Dutch:                bangleDutchName,
		language.French:               bangleFrenchName,
		language.German:               bangleGermanName,
		language.Italian:              bangleItalianName,
		language.Japanese:             bangleJapaneseName,
		language.Korean:               bangleKoreanName,
		language.LatinAmericanSpanish: bangleLatinAmericanSpanishName,
		language.Russian:              bangleRussianName,
		language.SimplifiedChinese:    bangleSimplifiedChineseName,
		language.Spanish:              bangleSpanishName,
		language.TraditionalChinese:   bangleTraditionalChineseName}
)

var (
	// bangleCharacter represents Bangle's character information.
	bangleCharacter = nook.Character{
		Animal:   animal.Tiger,
		Birthday: bangleBirthday,
		Code:     bangleCode,
		Key:      character.Bangle,
		Gender:   gender.Female,
		Name:     bangleName,
		Special:  false}
)

var (
	// bangleAmericanEnglishPhrase represents Bangle's phrase in American English.
	bangleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "growf"}

	// bangleCanadianFrenchPhrase represents Bangle's phrase in Canadian French.
	bangleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ma souris"}

	// bangleDutchPhrase represents Bangle's phrase in Dutch.
	bangleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "grommf"}

	// bangleFrenchPhrase represents Bangle's phrase in French.
	bangleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ma souris"}

	// bangleGermanPhrase represents Bangle's phrase in German.
	bangleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnurf"}

	// bangleItalianPhrase represents Bangle's phrase in Italian.
	bangleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grouf"}

	// bangleJapanesePhrase represents Bangle's phrase in Japanese.
	bangleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのぉー"}

	// bangleKoreanPhrase represents Bangle's phrase in Korean.
	bangleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쪼옥쪽"}

	// bangleLatinAmericanSpanishPhrase represents Bangle's phrase in Latin American Spanish.
	bangleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gri-gruá"}

	// bangleRussianPhrase represents Bangle's phrase in Russian.
	bangleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "р-р-рф"}

	// bangleSimplifiedChinesePhrase represents Bangle's phrase in Simplified Chinese.
	bangleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是喔"}

	// bangleSpanishPhrase represents Bangle's phrase in Spanish.
	bangleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mi vida"}

	// bangleTraditionalChinesePhrase represents Bangle's phrase in Traditional Chinese.
	bangleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是喔"}
)

var (
	// banglePhrase represents Bangle's phrase in different languages.
	banglePhrase = nook.Languages{
		language.AmericanEnglish:      bangleAmericanEnglishPhrase,
		language.CanadianFrench:       bangleCanadianFrenchPhrase,
		language.Dutch:                bangleDutchPhrase,
		language.French:               bangleFrenchPhrase,
		language.German:               bangleGermanPhrase,
		language.Italian:              bangleItalianPhrase,
		language.Japanese:             bangleJapanesePhrase,
		language.Korean:               bangleKoreanPhrase,
		language.LatinAmericanSpanish: bangleLatinAmericanSpanishPhrase,
		language.Russian:              bangleRussianPhrase,
		language.SimplifiedChinese:    bangleSimplifiedChinesePhrase,
		language.Spanish:              bangleSpanishPhrase,
		language.TraditionalChinese:   bangleTraditionalChinesePhrase}
)

var (
	// Bangle represents the character Bangle with her complete information.
	Bangle = nook.Villager{
		Character:   bangleCharacter,
		Personality: personality.Peppy,
		Phrase:      banglePhrase}
)
