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
	// callyBirthday represents Cally's birthday (September 4th).
	callyBirthday = nook.Birthday{
		Day:   4,
		Month: time.September}
)

var (
	// callyCode represents Cally's unique code ("squ11").
	callyCode = nook.Code{
		Value: "squ11"}
)

var (
	// callyAmericanEnglishName represents Cally's name in American English.
	callyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cally"}

	// callyCanadianFrenchName represents Cally's name in Canadian French.
	callyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Célia"}

	// callyDutchName represents Cally's name in Dutch.
	callyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cally"}

	// callyFrenchName represents Cally's name in French.
	callyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Célia"}

	// callyGermanName represents Cally's name in German.
	callyGermanName = nook.Name{
		Language: language.German,
		Value:    "Hörnchen"}

	// callyItalianName represents Cally's name in Italian.
	callyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rosa"}

	// callyJapaneseName represents Cally's name in Japanese.
	callyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パセリ"}

	// callyKoreanName represents Cally's name in Korean.
	callyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "파슬리"}

	// callyLatinAmericanSpanishName represents Cally's name in Latin American Spanish.
	callyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Almendra"}

	// callyRussianName represents Cally's name in Russian.
	callyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кэлли"}

	// callySimplifiedChineseName represents Cally's name in Simplified Chinese.
	callySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小芹"}

	// callySpanishName represents Cally's name in Spanish.
	callySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Almendra"}

	// callyTraditionalChineseName represents Cally's name in Traditional Chinese.
	callyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小芹"}
)

var (
	// callyName represents Cally's name in different languages.
	callyName = nook.Languages{
		language.AmericanEnglish:      callyAmericanEnglishName,
		language.CanadianFrench:       callyCanadianFrenchName,
		language.Dutch:                callyDutchName,
		language.French:               callyFrenchName,
		language.German:               callyGermanName,
		language.Italian:              callyItalianName,
		language.Japanese:             callyJapaneseName,
		language.Korean:               callyKoreanName,
		language.LatinAmericanSpanish: callyLatinAmericanSpanishName,
		language.Russian:              callyRussianName,
		language.SimplifiedChinese:    callySimplifiedChineseName,
		language.Spanish:              callySpanishName,
		language.TraditionalChinese:   callyTraditionalChineseName}
)

var (
	// callyCharacter represents Cally's character information.
	callyCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: callyBirthday,
		Code:     callyCode,
		Key:      character.Cally,
		Gender:   gender.Female,
		Name:     callyName,
		Special:  false}
)

var (
	// callyAmericanEnglishPhrase represents Cally's phrase in American English.
	callyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "WHEE"}

	// callyCanadianFrenchPhrase represents Cally's phrase in Canadian French.
	callyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "cacahuète"}

	// callyDutchPhrase represents Cally's phrase in Dutch.
	callyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "cashew"}

	// callyFrenchPhrase represents Cally's phrase in French.
	callyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "yeeeees"}

	// callyGermanPhrase represents Cally's phrase in German.
	callyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "huiii"}

	// callyItalianPhrase represents Cally's phrase in Italian.
	callyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "SCRUNCH"}

	// callyJapanesePhrase represents Cally's phrase in Japanese.
	callyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ララー"}

	// callyKoreanPhrase represents Cally's phrase in Korean.
	callyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "랄랄라"}

	// callyLatinAmericanSpanishPhrase represents Cally's phrase in Latin American Spanish.
	callyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ñip-ñip"}

	// callyRussianPhrase represents Cally's phrase in Russian.
	callyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ух ты"}

	// callySimplifiedChinesePhrase represents Cally's phrase in Simplified Chinese.
	callySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啦啦"}

	// callySpanishPhrase represents Cally's phrase in Spanish.
	callySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "macadamia"}

	// callyTraditionalChinesePhrase represents Cally's phrase in Traditional Chinese.
	callyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啦啦"}
)

var (
	// callyPhrase represents Cally's phrases in different languages.
	callyPhrase = nook.Languages{
		language.AmericanEnglish:      callyAmericanEnglishPhrase,
		language.CanadianFrench:       callyCanadianFrenchPhrase,
		language.Dutch:                callyDutchPhrase,
		language.French:               callyFrenchPhrase,
		language.German:               callyGermanPhrase,
		language.Italian:              callyItalianPhrase,
		language.Japanese:             callyJapanesePhrase,
		language.Korean:               callyKoreanPhrase,
		language.LatinAmericanSpanish: callyLatinAmericanSpanishPhrase,
		language.Russian:              callyRussianPhrase,
		language.SimplifiedChinese:    callySimplifiedChinesePhrase,
		language.Spanish:              callySpanishPhrase,
		language.TraditionalChinese:   callyTraditionalChinesePhrase}
)

var (
	// Cally represents the character Cally with her complete information.
	Cally = nook.Villager{
		Character:   callyCharacter,
		Personality: personality.Normal,
		Phrase:      callyPhrase}
)
