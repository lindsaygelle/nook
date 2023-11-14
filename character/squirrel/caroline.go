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
	// carolineBirthday represents Caroline's birthday (July 15th).
	carolineBirthday = nook.Birthday{
		Day:   15,
		Month: time.July}
)

var (
	// carolineCode represents Caroline's unique code ("squ06").
	carolineCode = nook.Code{
		Value: "squ06"}
)

var (
	// carolineAmericanEnglishName represents Caroline's name in American English.
	carolineAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Caroline"}

	// carolineCanadianFrenchName represents Caroline's name in Canadian French.
	carolineCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Isabelle"}

	// carolineDutchName represents Caroline's name in Dutch.
	carolineDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Caroline"}

	// carolineFrenchName represents Caroline's name in French.
	carolineFrenchName = nook.Name{
		Language: language.French,
		Value:    "Isabelle"}

	// carolineGermanName represents Caroline's name in German.
	carolineGermanName = nook.Name{
		Language: language.German,
		Value:    "Ricarda"}

	// carolineItalianName represents Caroline's name in Italian.
	carolineItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Carolina"}

	// carolineJapaneseName represents Caroline's name in Japanese.
	carolineJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キャロライン"}

	// carolineKoreanName represents Caroline's name in Korean.
	carolineKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "캐롤라인"}

	// carolineLatinAmericanSpanishName represents Caroline's name in Latin American Spanish.
	carolineLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mariló"}

	// carolineRussianName represents Caroline's name in Russian.
	carolineRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Каролина"}

	// carolineSimplifiedChineseName represents Caroline's name in Simplified Chinese.
	carolineSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "贾萝琳"}

	// carolineSpanishName represents Caroline's name in Spanish.
	carolineSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mariló"}

	// carolineTraditionalChineseName represents Caroline's name in Traditional Chinese.
	carolineTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "賈蘿琳"}
)

var (
	// carolineName represents Caroline's name in different languages.
	carolineName = nook.Languages{
		language.AmericanEnglish:      carolineAmericanEnglishName,
		language.CanadianFrench:       carolineCanadianFrenchName,
		language.Dutch:                carolineDutchName,
		language.French:               carolineFrenchName,
		language.German:               carolineGermanName,
		language.Italian:              carolineItalianName,
		language.Japanese:             carolineJapaneseName,
		language.Korean:               carolineKoreanName,
		language.LatinAmericanSpanish: carolineLatinAmericanSpanishName,
		language.Russian:              carolineRussianName,
		language.SimplifiedChinese:    carolineSimplifiedChineseName,
		language.Spanish:              carolineSpanishName,
		language.TraditionalChinese:   carolineTraditionalChineseName}
)

var (
	// carolineCharacter represents Caroline's character information.
	carolineCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: carolineBirthday,
		Code:     carolineCode,
		Key:      character.Caroline,
		Gender:   gender.Female,
		Name:     carolineName,
		Special:  false}
)

var (
	// carolineAmericanEnglishPhrase represents Caroline's phrase in American English.
	carolineAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hulaaaa"}

	// carolineCanadianFrenchPhrase represents Caroline's phrase in Canadian French.
	carolineCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "houlaaaa"}

	// carolineDutchPhrase represents Caroline's phrase in Dutch.
	carolineDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hoelaaa"}

	// carolineFrenchPhrase represents Caroline's phrase in French.
	carolineFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "houlaaaa"}

	// carolineGermanPhrase represents Caroline's phrase in German.
	carolineGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hulaaaa"}

	// carolineItalianPhrase represents Caroline's phrase in Italian.
	carolineItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ehilaaà"}

	// carolineJapanesePhrase represents Caroline's phrase in Japanese.
	carolineJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ユー"}

	// carolineKoreanPhrase represents Caroline's phrase in Korean.
	carolineKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "휴"}

	// carolineLatinAmericanSpanishPhrase represents Caroline's phrase in Latin American Spanish.
	carolineLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ulaaaa"}

	// carolineRussianPhrase represents Caroline's phrase in Russian.
	carolineRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "уа-а-а"}

	// carolineSimplifiedChinesePhrase represents Caroline's phrase in Simplified Chinese.
	carolineSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "你"}

	// carolineSpanishPhrase represents Caroline's phrase in Spanish.
	carolineSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ulaaaa"}

	// carolineTraditionalChinesePhrase represents Caroline's phrase in Traditional Chinese.
	carolineTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "你"}
)

var (
	// carolinePhrase represents Caroline's phrases in different languages.
	carolinePhrase = nook.Languages{
		language.AmericanEnglish:      carolineAmericanEnglishPhrase,
		language.CanadianFrench:       carolineCanadianFrenchPhrase,
		language.Dutch:                carolineDutchPhrase,
		language.French:               carolineFrenchPhrase,
		language.German:               carolineGermanPhrase,
		language.Italian:              carolineItalianPhrase,
		language.Japanese:             carolineJapanesePhrase,
		language.Korean:               carolineKoreanPhrase,
		language.LatinAmericanSpanish: carolineLatinAmericanSpanishPhrase,
		language.Russian:              carolineRussianPhrase,
		language.SimplifiedChinese:    carolineSimplifiedChinesePhrase,
		language.Spanish:              carolineSpanishPhrase,
		language.TraditionalChinese:   carolineTraditionalChinesePhrase}
)

var (
	// Caroline represents the character Caroline with her complete information.
	Caroline = nook.Villager{
		Character:   carolineCharacter,
		Personality: personality.Normal,
		Phrase:      carolinePhrase}
)
