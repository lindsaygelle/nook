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
	rhondaBirthday = nook.Birthday{
		Day:   24,
		Month: time.January}
)

var (
	rhondaCode = nook.Code{
		Value: "rhn01"}
)

var (
	rhondaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rhonda"}

	rhondaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Renée"}

	rhondaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rhonda"}

	rhondaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Renée"}

	rhondaGermanName = nook.Name{
		Language: language.German,
		Value:    "Regina"}

	rhondaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Roby"}

	rhondaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ユメコ"}

	rhondaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "론다"}

	rhondaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ronda"}

	rhondaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Ронда"}

	rhondaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "梦梦"}

	rhondaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ronda"}

	rhondaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "夢夢"}
)

var (
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
	rhondaCharacter = nook.Character{
		Animal:   animal.Rhinoceros,
		Birthday: rhondaBirthday,
		Code:     rhondaCode,
		Key:      character.Rhonda,
		Gender:   gender.Female,
		Name:     rhondaName}
)

var (
	rhondaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bigfoot"}

	rhondaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grand pied"}

	rhondaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "grootvoet"}

	rhondaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grand pied"}

	rhondaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "stampf"}

	rhondaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "piedone"}

	rhondaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ンフ"}

	rhondaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "훗"}

	rhondaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "jip-jip"}

	rhondaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "босс"}

	rhondaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嗯噗"}

	rhondaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "jip-jip"}

	rhondaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗯噗"}
)

var (
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
	Rhonda = nook.Villager{
		Character:   rhondaCharacter,
		Personality: personality.Normal,
		Phrase:      rhondaPhrase}
)
