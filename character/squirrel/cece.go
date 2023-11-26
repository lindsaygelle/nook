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
	// ceceBirthday represents Cece's birthday.
	ceceBirthday = nook.Birthday{
		Day:   28,
		Month: time.May}
)

var (
	// ceceCode represents Cece's unique code.
	ceceCode = nook.Code{
		Value: "squ19"}
)

var (
	// ceceAmericanEnglishName represents Cece's name in American English.
	ceceAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cece"}

	// ceceCanadianFrenchName represents Cece's name in Canadian French.
	ceceCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kala"}

	// ceceDutchName represents Cece's name in Dutch.
	ceceDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// ceceFrenchName represents Cece's name in French.
	ceceFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kala"}

	// ceceGermanName represents Cece's name in German.
	ceceGermanName = nook.Name{
		Language: language.German,
		Value:    "A-Chan"}

	// ceceItalianName represents Cece's name in Italian.
	ceceItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cecilia"}

	// ceceJapaneseName represents Cece's name in Japanese.
	ceceJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "なぎさ"}

	// ceceKoreanName represents Cece's name in Korean.
	ceceKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "나기사"}

	// ceceLatinAmericanSpanishName represents Cece's name in Latin American Spanish.
	ceceLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Arduna"}

	// ceceRussianName represents Cece's name in Russian.
	ceceRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// ceceSimplifiedChineseName represents Cece's name in Simplified Chinese.
	ceceSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// ceceSpanishName represents Cece's name in Spanish.
	ceceSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Arduna"}

	// ceceTraditionalChineseName represents Cece's name in Traditional Chinese.
	ceceTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// ceceName represents Cece's name in different languages.
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
	// ceceCharacter represents Cece's character information.
	ceceCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: ceceBirthday,
		Code:     ceceCode,
		Key:      character.Cece,
		Gender:   gender.Female,
		Name:     ceceName,
		Special:  false}
)

var (
	// ceceAmericanEnglishPhrase represents Cece's phrase in American English.
	ceceAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "stay fresh"}

	// ceceCanadianFrenchPhrase represents Cece's phrase in Canadian French.
	ceceCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "maréôte"}

	// ceceDutchPhrase represents Cece's phrase in Dutch.
	ceceDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// ceceFrenchPhrase represents Cece's phrase in French.
	ceceFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "maréôte"}

	// ceceGermanPhrase represents Cece's phrase in German.
	ceceGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "aioli"}

	// ceceItalianPhrase represents Cece's phrase in Italian.
	ceceItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "calacala"}

	// ceceJapanesePhrase represents Cece's phrase in Japanese.
	ceceJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "よろしく"}

	// ceceKoreanPhrase represents Cece's phrase in Korean.
	ceceKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "잘지내자"}

	// ceceLatinAmericanSpanishPhrase represents Cece's phrase in Latin American Spanish.
	ceceLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cosmar"}

	// ceceRussianPhrase represents Cece's phrase in Russian.
	ceceRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// ceceSimplifiedChinesePhrase represents Cece's phrase in Simplified Chinese.
	ceceSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// ceceSpanishPhrase represents Cece's phrase in Spanish.
	ceceSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cosmar"}

	// ceceTraditionalChinesePhrase represents Cece's phrase in Traditional Chinese.
	ceceTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// cecePhrase represents Cece's phrases in different languages.
	cecePhrase = nook.Languages{
		language.AmericanEnglish:      ceceAmericanEnglishPhrase,
		language.CanadianFrench:       ceceCanadianFrenchPhrase,
		language.Dutch:                ceceDutchPhrase,
		language.French:               ceceFrenchPhrase,
		language.German:               ceceGermanPhrase,
		language.Italian:              ceceItalianPhrase,
		language.Japanese:             ceceJapanesePhrase,
		language.Korean:               ceceKoreanPhrase,
		language.LatinAmericanSpanish: ceceLatinAmericanSpanishPhrase,
		language.Russian:              ceceRussianPhrase,
		language.SimplifiedChinese:    ceceSimplifiedChinesePhrase,
		language.Spanish:              ceceSpanishPhrase,
		language.TraditionalChinese:   ceceTraditionalChinesePhrase}
)

var (
	// Cece represents the character Cece with her complete information.
	Cece = nook.Villager{
		Character:   ceceCharacter,
		Personality: personality.Peppy,
		Phrase:      cecePhrase}
)
