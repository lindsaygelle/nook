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
	// vicheBirthday represents Viche's birthday.
	vicheBirthday = nook.Birthday{
		Day:   7,
		Month: time.July}
)

var (
	// vicheCode represents Viche's unique code.
	vicheCode = nook.Code{
		Value: "squ20"}
)

var (
	// vicheAmericanEnglishName represents Viche's name in American English.
	vicheAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Viché"}

	// vicheCanadianFrenchName represents Viche's name in Canadian French.
	vicheCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mara"}

	// vicheDutchName represents Viche's name in Dutch.
	vicheDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// vicheFrenchName represents Viche's name in French.
	vicheFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mara"}

	// vicheGermanName represents Viche's name in German.
	vicheGermanName = nook.Name{
		Language: language.German,
		Value:    "L-Pyon"}

	// vicheItalianName represents Viche's name in Italian.
	vicheItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Vicky"}

	// vicheJapaneseName represents Viche's name in Japanese.
	vicheJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "みさき"}

	// vicheKoreanName represents Viche's name in Korean.
	vicheKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "미사키"}

	// vicheLatinAmericanSpanishName represents Viche's name in Latin American Spanish.
	vicheLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ardelta"}

	// vicheRussianName represents Viche's name in Russian.
	vicheRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// vicheSimplifiedChineseName represents Viche's name in Simplified Chinese.
	vicheSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// vicheSpanishName represents Viche's name in Spanish.
	vicheSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ardelta"}

	// vicheTraditionalChineseName represents Viche's name in Traditional Chinese.
	vicheTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// vicheName represents Viche's name in different languages.
	vicheName = nook.Languages{
		language.AmericanEnglish:      vicheAmericanEnglishName,
		language.CanadianFrench:       vicheCanadianFrenchName,
		language.Dutch:                vicheDutchName,
		language.French:               vicheFrenchName,
		language.German:               vicheGermanName,
		language.Italian:              vicheItalianName,
		language.Japanese:             vicheJapaneseName,
		language.Korean:               vicheKoreanName,
		language.LatinAmericanSpanish: vicheLatinAmericanSpanishName,
		language.Russian:              vicheRussianName,
		language.SimplifiedChinese:    vicheSimplifiedChineseName,
		language.Spanish:              vicheSpanishName,
		language.TraditionalChinese:   vicheTraditionalChineseName}
)

var (
	// vicheCharacter represents Viche's character information.
	vicheCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: vicheBirthday,
		Code:     vicheCode,
		Key:      character.Viche,
		Gender:   gender.Female,
		Name:     vicheName,
		Special:  false}
)

var (
	// vicheAmericanEnglishPhrase represents Viche's phrase in American English.
	vicheAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "splatastic"}

	// vicheCanadianFrenchPhrase represents Viche's phrase in Canadian French.
	vicheCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "marébasse"}

	// vicheDutchPhrase represents Viche's phrase in Dutch.
	vicheDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// vicheFrenchPhrase represents Viche's phrase in French.
	vicheFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "marébasse"}

	// vicheGermanPhrase represents Viche's phrase in German.
	vicheGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "limone"}

	// vicheItalianPhrase represents Viche's phrase in Italian.
	vicheItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "maromaro"}

	// vicheJapanesePhrase represents Viche's phrase in Japanese.
	vicheJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なんね"}

	// vicheKoreanPhrase represents Viche's phrase in Korean.
	vicheKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그거좋아"}

	// vicheLatinAmericanSpanishPhrase represents Viche's phrase in Latin American Spanish.
	vicheLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "costina"}

	// vicheRussianPhrase represents Viche's phrase in Russian.
	vicheRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// vicheSimplifiedChinesePhrase represents Viche's phrase in Simplified Chinese.
	vicheSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// vicheSpanishPhrase represents Viche's phrase in Spanish.
	vicheSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "costina"}

	// vicheTraditionalChinesePhrase represents Viche's phrase in Traditional Chinese.
	vicheTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// vichePhrase represents Viche's phrase in different languages.
	vichePhrase = nook.Languages{
		language.AmericanEnglish:      vicheAmericanEnglishPhrase,
		language.CanadianFrench:       vicheCanadianFrenchPhrase,
		language.Dutch:                vicheDutchPhrase,
		language.French:               vicheFrenchPhrase,
		language.German:               vicheGermanPhrase,
		language.Italian:              vicheItalianPhrase,
		language.Japanese:             vicheJapanesePhrase,
		language.Korean:               vicheKoreanPhrase,
		language.LatinAmericanSpanish: vicheLatinAmericanSpanishPhrase,
		language.Russian:              vicheRussianPhrase,
		language.SimplifiedChinese:    vicheSimplifiedChinesePhrase,
		language.Spanish:              vicheSpanishPhrase,
		language.TraditionalChinese:   vicheTraditionalChinesePhrase}
)

var (
	// Viche represents the character Viche with her complete information.
	Viche = nook.Villager{
		Character:   vicheCharacter,
		Personality: personality.Normal,
		Phrase:      vichePhrase}
)
