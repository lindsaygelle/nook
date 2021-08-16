package squirrel

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
	ceceBirthday = nook.Birthday{
		Day:   28,
		Month: time.May}
)

var (
	ceceCode = nook.Code{
		Value: "squ19"}
)

var (
	ceceAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cece"}

	ceceCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kala"}

	ceceDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	ceceFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kala"}

	ceceGermanName = nook.Name{
		Language: language.German,
		Value:    "A-Chan"}

	ceceItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cecilia"}

	ceceJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "なぎさ"}

	ceceKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "나기사"}

	ceceLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Arduna"}

	ceceRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	ceceSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	ceceSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Arduna"}

	ceceTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	ceceName = nook.Languages{
		language.AmericanEnglish:      ceceAmericanEnglishName,
		language.CanadianFrench:       ceceCanadianFrenchName,
		language.Dutch:                ceceDutchName,
		language.French:               ceceFrenchName,
		language.German:               ceceGermanName,
		language.Italian:              ceceItalianName,
		language.Japanese:             ceceJapaneseName,
		language.Korean:               ceceKoreanName,
		language.LatinAmericanSpanish: ceceLatinAmericanSpanishName,
		language.Russian:              ceceRussianName,
		language.SimplifiedChinese:    ceceSimplifiedChineseName,
		language.Spanish:              ceceSpanishName,
		language.TraditionalChinese:   ceceTraditionalChineseName}
)

var (
	ceceCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: ceceBirthday,
		Code:     ceceCode,
		Key:      character.Cece,
		Gender:   gender.Female,
		Name:     ceceName}
)

var (
	ceceAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "stay fresh"}

	ceceCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "maréôte"}

	ceceDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	ceceFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "maréôte"}

	ceceGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "aioli"}

	ceceItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "calacala"}

	ceceJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "よろしく"}

	ceceKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "잘지내자"}

	ceceLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cosmar"}

	ceceRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	ceceSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	ceceSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cosmar"}

	ceceTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	cecePhrase = nook.Languages{
		language.AmericanEnglish:      ceceAmericanEnglishName,
		language.CanadianFrench:       ceceCanadianFrenchName,
		language.Dutch:                ceceDutchName,
		language.French:               ceceFrenchName,
		language.German:               ceceGermanName,
		language.Italian:              ceceItalianName,
		language.Japanese:             ceceJapaneseName,
		language.Korean:               ceceKoreanName,
		language.LatinAmericanSpanish: ceceLatinAmericanSpanishName,
		language.Russian:              ceceRussianName,
		language.SimplifiedChinese:    ceceSimplifiedChineseName,
		language.Spanish:              ceceSpanishName,
		language.TraditionalChinese:   ceceTraditionalChineseName}
)

var (
	Cece = nook.Villager{
		Character:   ceceCharacter,
		Personality: personality.Peppy,
		Phrase:      cecePhrase}
)
