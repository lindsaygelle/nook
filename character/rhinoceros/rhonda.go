package rhinoceros

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
	// rhondaBirthday represents Rhonda's birthday (January 24th).
	rhondaBirthday = nook.Birthday{
		Day:   24,
		Month: time.January}
)

var (
	// rhondaCode represents Rhonda's unique code ("rhn01").
	rhondaCode = nook.Code{
		Value: "rhn01"}
)

var (
	// rhondaAmericanEnglishName represents Rhonda's name in American English.
	rhondaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rhonda"}

	// rhondaCanadianFrenchName represents Rhonda's name in Canadian French.
	rhondaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Renée"}

	// rhondaDutchName represents Rhonda's name in Dutch.
	rhondaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rhonda"}

	// rhondaFrenchName represents Rhonda's name in French.
	rhondaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Renée"}

	// rhondaGermanName represents Rhonda's name in German.
	rhondaGermanName = nook.Name{
		Language: language.German,
		Value:    "Regina"}

	// rhondaItalianName represents Rhonda's name in Italian.
	rhondaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Roby"}

	// rhondaJapaneseName represents Rhonda's name in Japanese.
	rhondaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ユメコ"}

	// rhondaKoreanName represents Rhonda's name in Korean.
	rhondaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "론다"}

	// rhondaLatinAmericanSpanishName represents Rhonda's name in Latin American Spanish.
	rhondaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ronda"}

	// rhondaRussianName represents Rhonda's name in Russian.
	rhondaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Ронда"}

	// rhondaSimplifiedChineseName represents Rhonda's name in Simplified Chinese.
	rhondaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "梦梦"}

	// rhondaSpanishName represents Rhonda's name in Spanish.
	rhondaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ronda"}

	// rhondaTraditionalChineseName represents Rhonda's name in Traditional Chinese.
	rhondaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "夢夢"}
)

var (
	// rhondaName represents Rhonda's name in different languages.
	rhondaName = nook.Languages{
		language.AmericanEnglish:      rhondaAmericanEnglishName,
		language.CanadianFrench:       rhondaCanadianFrenchName,
		language.Dutch:                rhondaDutchName,
		language.French:               rhondaFrenchName,
		language.German:               rhondaGermanName,
		language.Italian:              rhondaItalianName,
		language.Japanese:             rhondaJapaneseName,
		language.Korean:               rhondaKoreanName,
		language.LatinAmericanSpanish: rhondaLatinAmericanSpanishName,
		language.Russian:              rhondaRussianName,
		language.SimplifiedChinese:    rhondaSimplifiedChineseName,
		language.Spanish:              rhondaSpanishName,
		language.TraditionalChinese:   rhondaTraditionalChineseName}
)

var (
	// rhondaCharacter represents Rhonda's character information.
	rhondaCharacter = nook.Character{
		Animal:   animal.Rhinoceros,
		Birthday: rhondaBirthday,
		Code:     rhondaCode,
		Key:      character.Rhonda,
		Gender:   gender.Female,
		Name:     rhondaName,
		Special:  false}
)

var (
	// rhondaAmericanEnglishPhrase represents Rhonda's phrase in American English.
	rhondaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bigfoot"}

	// rhondaCanadianFrenchPhrase represents Rhonda's phrase in Canadian French.
	rhondaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grand pied"}

	// rhondaDutchPhrase represents Rhonda's phrase in Dutch.
	rhondaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "grootvoet"}

	// rhondaFrenchPhrase represents Rhonda's phrase in French.
	rhondaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grand pied"}

	// rhondaGermanPhrase represents Rhonda's phrase in German.
	rhondaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "stampf"}

	// rhondaItalianPhrase represents Rhonda's phrase in Italian.
	rhondaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "piedone"}

	// rhondaJapanesePhrase represents Rhonda's phrase in Japanese.
	rhondaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ンフ"}

	// rhondaKoreanPhrase represents Rhonda's phrase in Korean.
	rhondaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "훗"}

	// rhondaLatinAmericanSpanishPhrase represents Rhonda's phrase in Latin American Spanish.
	rhondaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "jip-jip"}

	// rhondaRussianPhrase represents Rhonda's phrase in Russian.
	rhondaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "босс"}

	// rhondaSimplifiedChinesePhrase represents Rhonda's phrase in Simplified Chinese.
	rhondaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嗯噗"}

	// rhondaSpanishPhrase represents Rhonda's phrase in Spanish.
	rhondaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "jip-jip"}

	// rhondaTraditionalChinesePhrase represents Rhonda's phrase in Traditional Chinese.
	rhondaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗯噗"}
)

var (
	// rhondaPhrase represents Rhonda's phrase in different languages.
	rhondaPhrase = nook.Languages{
		language.AmericanEnglish:      rhondaAmericanEnglishPhrase,
		language.CanadianFrench:       rhondaCanadianFrenchPhrase,
		language.Dutch:                rhondaDutchPhrase,
		language.French:               rhondaFrenchPhrase,
		language.German:               rhondaGermanPhrase,
		language.Italian:              rhondaItalianPhrase,
		language.Japanese:             rhondaJapanesePhrase,
		language.Korean:               rhondaKoreanPhrase,
		language.LatinAmericanSpanish: rhondaLatinAmericanSpanishPhrase,
		language.Russian:              rhondaRussianPhrase,
		language.SimplifiedChinese:    rhondaSimplifiedChinesePhrase,
		language.Spanish:              rhondaSpanishPhrase,
		language.TraditionalChinese:   rhondaTraditionalChinesePhrase}
)

var (
	// Rhonda represents the character Rhonda with her complete information.
	Rhonda = nook.Villager{
		Character:   rhondaCharacter,
		Personality: personality.Normal,
		Phrase:      rhondaPhrase}
)
