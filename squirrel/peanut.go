package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Rachidanoisette"}

	peanutDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Peanutluilak"}

	peanutFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rachidanoisette"}

	peanutGermanName = nook.Name{
		Language: language.German,
		Value:    "Paulinafaulenzer"}

	peanutItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rachelesgronch"}

	peanutJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ももこなのよ"}

	peanutKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "핑키거얌"}

	peanutLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Belindaesñik"}

	peanutRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пинатлень"}

	peanutSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小桃就是唷"}

	peanutSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Belindaalmendrita"}

	peanutTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小桃就是唷"}
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
		Animal:   Squirrel,
		Birthday: peanutBirthday,
		Code:     peanutCode,
		Gender:   nook.Female,
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
	Peanut = nook.Villager{
		Character:   peanutCharacter,
		Personality: nook.Peppy,
		Phrase:      peanutPhrase}
)
