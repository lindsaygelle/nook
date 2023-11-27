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
	// peanutBirthday represents Peanut's birthday.
	peanutBirthday = nook.Birthday{
		Day:   8,
		Month: time.June}
)

var (
	// peanutCode represents Peanut's unique code.
	peanutCode = nook.Code{
		Value: "squ00"}
)

var (
	// peanutAmericanEnglishName represents Peanut's name in American English.
	peanutAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Peanut"}

	// peanutCanadianFrenchName represents Peanut's name in Canadian French.
	peanutCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rachida"}

	// peanutDutchName represents Peanut's name in Dutch.
	peanutDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Peanut"}

	// peanutFrenchName represents Peanut's name in French.
	peanutFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rachida"}

	// peanutGermanName represents Peanut's name in German.
	peanutGermanName = nook.Name{
		Language: language.German,
		Value:    "Paulina"}

	// peanutItalianName represents Peanut's name in Italian.
	peanutItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rachele"}

	// peanutJapaneseName represents Peanut's name in Japanese.
	peanutJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ももこ"}

	// peanutKoreanName represents Peanut's name in Korean.
	peanutKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "핑키"}

	// peanutLatinAmericanSpanishName represents Peanut's name in Latin American Spanish.
	peanutLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Belinda"}

	// peanutRussianName represents Peanut's name in Russian.
	peanutRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пинат"}

	// peanutSimplifiedChineseName represents Peanut's name in Simplified Chinese.
	peanutSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小桃"}

	// peanutSpanishName represents Peanut's name in Spanish.
	peanutSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Belinda"}

	// peanutTraditionalChineseName represents Peanut's name in Traditional Chinese.
	peanutTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小桃"}
)

var (
	// peanutName represents Peanut's name in different languages.
	peanutName = nook.Languages{
		language.AmericanEnglish:      peanutAmericanEnglishName,
		language.CanadianFrench:       peanutCanadianFrenchName,
		language.Dutch:                peanutDutchName,
		language.French:               peanutFrenchName,
		language.German:               peanutGermanName,
		language.Italian:              peanutItalianName,
		language.Japanese:             peanutJapaneseName,
		language.Korean:               peanutKoreanName,
		language.LatinAmericanSpanish: peanutLatinAmericanSpanishName,
		language.Russian:              peanutRussianName,
		language.SimplifiedChinese:    peanutSimplifiedChineseName,
		language.Spanish:              peanutSpanishName,
		language.TraditionalChinese:   peanutTraditionalChineseName}
)

var (
	// peanutCharacter represents Peanut's character information.
	peanutCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: peanutBirthday,
		Code:     peanutCode,
		Key:      character.Peanut,
		Gender:   gender.Female,
		Name:     peanutName,
		Special:  false}
)

var (
	// peanutAmericanEnglishPhrase represents Peanut's phrase in American English.
	peanutAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "slacker"}

	// peanutCanadianFrenchPhrase represents Peanut's phrase in Canadian French.
	peanutCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "noisette"}

	// peanutDutchPhrase represents Peanut's phrase in Dutch.
	peanutDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "luilak"}

	// peanutFrenchPhrase represents Peanut's phrase in French.
	peanutFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "noisette"}

	// peanutGermanPhrase represents Peanut's phrase in German.
	peanutGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "faulenzer"}

	// peanutItalianPhrase represents Peanut's phrase in Italian.
	peanutItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sgronch"}

	// peanutJapanesePhrase represents Peanut's phrase in Japanese.
	peanutJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのよ"}

	// peanutKoreanPhrase represents Peanut's phrase in Korean.
	peanutKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "거얌"}

	// peanutLatinAmericanSpanishPhrase represents Peanut's phrase in Latin American Spanish.
	peanutLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "esñik"}

	// peanutRussianPhrase represents Peanut's phrase in Russian.
	peanutRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "лень"}

	// peanutSimplifiedChinesePhrase represents Peanut's phrase in Simplified Chinese.
	peanutSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "就是唷"}

	// peanutSpanishPhrase represents Peanut's phrase in Spanish.
	peanutSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "almendrita"}

	// peanutTraditionalChinesePhrase represents Peanut's phrase in Traditional Chinese.
	peanutTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "就是唷"}
)

var (
	// peanutPhrase represents Peanut's phrase in different languages.
	peanutPhrase = nook.Languages{
		language.AmericanEnglish:      peanutAmericanEnglishPhrase,
		language.CanadianFrench:       peanutCanadianFrenchPhrase,
		language.Dutch:                peanutDutchPhrase,
		language.French:               peanutFrenchPhrase,
		language.German:               peanutGermanPhrase,
		language.Italian:              peanutItalianPhrase,
		language.Japanese:             peanutJapanesePhrase,
		language.Korean:               peanutKoreanPhrase,
		language.LatinAmericanSpanish: peanutLatinAmericanSpanishPhrase,
		language.Russian:              peanutRussianPhrase,
		language.SimplifiedChinese:    peanutSimplifiedChinesePhrase,
		language.Spanish:              peanutSpanishPhrase,
		language.TraditionalChinese:   peanutTraditionalChinesePhrase}
)

var (
	// Peanut represents the character Peanut with her complete information.
	Peanut = nook.Villager{
		Character:   peanutCharacter,
		Personality: personality.Peppy,
		Phrase:      peanutPhrase}
)
