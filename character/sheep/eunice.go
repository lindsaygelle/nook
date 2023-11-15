package sheep

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
	// euniceBirthday represents Eunice's birthday (April 3rd).
	euniceBirthday = nook.Birthday{
		Day:   3,
		Month: time.April}
)

var (
	// euniceCode represents Eunice's unique code ("shp02").
	euniceCode = nook.Code{
		Value: "shp02"}
)

var (
	// euniceAmericanEnglishName represents Eunice's name in American English.
	euniceAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Eunice"}

	// euniceCanadianFrenchName represents Eunice's name in Canadian French.
	euniceCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bérénice"}

	// euniceDutchName represents Eunice's name in Dutch.
	euniceDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Eunice"}

	// euniceFrenchName represents Eunice's name in French.
	euniceFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bérénice"}

	// euniceGermanName represents Eunice's name in German.
	euniceGermanName = nook.Name{
		Language: language.German,
		Value:    "Edith"}

	// euniceItalianName represents Eunice's name in Italian.
	euniceItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tenerina"}

	// euniceJapaneseName represents Eunice's name in Japanese.
	euniceJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "モヘア"}

	// euniceKoreanName represents Eunice's name in Korean.
	euniceKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "곱슬이"}

	// euniceLatinAmericanSpanishName represents Eunice's name in Latin American Spanish.
	euniceLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lanolina"}

	// euniceRussianName represents Eunice's name in Russian.
	euniceRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Юнис"}

	// euniceSimplifiedChineseName represents Eunice's name in Simplified Chinese.
	euniceSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "毛海儿"}

	// euniceSpanishName represents Eunice's name in Spanish.
	euniceSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lanolina"}

	// euniceTraditionalChineseName represents Eunice's name in Traditional Chinese.
	euniceTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "毛海兒"}
)

var (
	// euniceName represents Eunice's name in different languages.
	euniceName = nook.Languages{
		language.AmericanEnglish:      euniceAmericanEnglishName,
		language.CanadianFrench:       euniceCanadianFrenchName,
		language.Dutch:                euniceDutchName,
		language.French:               euniceFrenchName,
		language.German:               euniceGermanName,
		language.Italian:              euniceItalianName,
		language.Japanese:             euniceJapaneseName,
		language.Korean:               euniceKoreanName,
		language.LatinAmericanSpanish: euniceLatinAmericanSpanishName,
		language.Russian:              euniceRussianName,
		language.SimplifiedChinese:    euniceSimplifiedChineseName,
		language.Spanish:              euniceSpanishName,
		language.TraditionalChinese:   euniceTraditionalChineseName}
)

var (
	// euniceCharacter represents Eunice's character information.
	euniceCharacter = nook.Character{
		Animal:   animal.Sheep,
		Birthday: euniceBirthday,
		Code:     euniceCode,
		Key:      character.Eunice,
		Gender:   gender.Female,
		Name:     euniceName,
		Special:  false}
)

var (
	// euniceAmericanEnglishPhrase represents Eunice's phrase in American English.
	euniceAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "lambchop"}

	// euniceCanadianFrenchPhrase represents Eunice's phrase in Canadian French.
	euniceCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bêbêêêêh"}

	// euniceDutchPhrase represents Eunice's phrase in Dutch.
	euniceDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "krulletje"}

	// euniceFrenchPhrase represents Eunice's phrase in French.
	euniceFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bêbêêêêh"}

	// euniceGermanPhrase represents Eunice's phrase in German.
	euniceGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "la-a-amm"}

	// euniceItalianPhrase represents Eunice's phrase in Italian.
	euniceItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "laaagna"}

	// euniceJapanesePhrase represents Eunice's phrase in Japanese.
	euniceJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "メェー"}

	// euniceKoreanPhrase represents Eunice's phrase in Korean.
	euniceKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "메에에"}

	// euniceLatinAmericanSpanishPhrase represents Eunice's phrase in Latin American Spanish.
	euniceLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "yupitiii"}

	// euniceRussianPhrase represents Eunice's phrase in Russian.
	euniceRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ребрышко"}

	// euniceSimplifiedChinesePhrase represents Eunice's phrase in Simplified Chinese.
	euniceSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咩"}

	// euniceSpanishPhrase represents Eunice's phrase in Spanish.
	euniceSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "veeenga"}

	// euniceTraditionalChinesePhrase represents Eunice's phrase in Traditional Chinese.
	euniceTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咩"}
)

var (
	// eunicePhrase represents Eunice's phrases in different languages.
	eunicePhrase = nook.Languages{
		language.AmericanEnglish:      euniceAmericanEnglishPhrase,
		language.CanadianFrench:       euniceCanadianFrenchPhrase,
		language.Dutch:                euniceDutchPhrase,
		language.French:               euniceFrenchPhrase,
		language.German:               euniceGermanPhrase,
		language.Italian:              euniceItalianPhrase,
		language.Japanese:             euniceJapanesePhrase,
		language.Korean:               euniceKoreanPhrase,
		language.LatinAmericanSpanish: euniceLatinAmericanSpanishPhrase,
		language.Russian:              euniceRussianPhrase,
		language.SimplifiedChinese:    euniceSimplifiedChinesePhrase,
		language.Spanish:              euniceSpanishPhrase,
		language.TraditionalChinese:   euniceTraditionalChinesePhrase}
)

var (
	// Eunice represents the character Eunice with her complete information.
	Eunice = nook.Villager{
		Character:   euniceCharacter,
		Personality: personality.Normal,
		Phrase:      eunicePhrase}
)
