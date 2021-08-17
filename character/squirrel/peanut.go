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
	peanutBirthday = nook.Birthday{
		Day:   8,
		Month: time.June}
)

var (
	peanutCode = nook.Code{
		Value: "squ00"}
)

var (
	peanutAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Peanut"}

	peanutCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rachida"}

	peanutDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Peanut"}

	peanutFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rachida"}

	peanutGermanName = nook.Name{
		Language: language.German,
		Value:    "Paulina"}

	peanutItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rachele"}

	peanutJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ももこ"}

	peanutKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "핑키"}

	peanutLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Belinda"}

	peanutRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пинат"}

	peanutSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小桃"}

	peanutSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Belinda"}

	peanutTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小桃"}
)

var (
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
	peanutCharacter = nook.Character{
		Animal:   animal.Squirrel,
		Birthday: peanutBirthday,
		Code:     peanutCode,
		Key:      character.Peanut,
		Gender:   gender.Female,
		Name:     peanutName}
)

var (
	peanutAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "slacker"}

	peanutCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "noisette"}

	peanutDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "luilak"}

	peanutFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "noisette"}

	peanutGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "faulenzer"}

	peanutItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sgronch"}

	peanutJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのよ"}

	peanutKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "거얌"}

	peanutLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "esñik"}

	peanutRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "лень"}

	peanutSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "就是唷"}

	peanutSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "almendrita"}

	peanutTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "就是唷"}
)

var (
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
	Peanut = nook.Villager{
		Character:   peanutCharacter,
		Personality: personality.Peppy,
		Phrase:      peanutPhrase}
)
