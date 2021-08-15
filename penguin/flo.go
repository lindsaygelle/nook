package penguin

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	floBirthday = nook.Birthday{
		Day:   2,
		Month: time.September}
)

var (
	floCode = nook.Code{
		Value: "pgn13"}
)

var (
	floAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Flo"}

	floCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Norafrigogo"}

	floDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Floduh"}

	floFrenchName = nook.Name{
		Language: language.French,
		Value:    "Noragolri"}

	floGermanName = nook.Name{
		Language: language.German,
		Value:    "Susanneflatsch"}

	floItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Nivesiglù"}

	floJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "レイラじゃんよ"}

	floKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "레이라맞잖아"}

	floLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nievesfresqui"}

	floRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Флочхи"}

	floSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蕾拉这样唷"}

	floSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Nievesfresqui"}

	floTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蕾拉這樣唷"}
)

var (
	floName = nook.Languages{
		language.AmericanEnglish:      floAmericanEnglishName,
		language.CanadianFrench:       floCanadianFrenchName,
		language.Dutch:                floDutchName,
		language.French:               floFrenchName,
		language.German:               floGermanName,
		language.Italian:              floItalianName,
		language.Japanese:             floJapaneseName,
		language.Korean:               floKoreanName,
		language.LatinAmericanSpanish: floLatinAmericanSpanishName,
		language.Russian:              floRussianName,
		language.SimplifiedChinese:    floSimplifiedChineseName,
		language.Spanish:              floSpanishName,
		language.TraditionalChinese:   floTraditionalChineseName}
)

var (
	floCharacter = nook.Character{
		Animal:   Penguin,
		Birthday: floBirthday,
		Code:     floCode,
		Gender:   nook.Female,
		Name:     floName}
)

var (
	floAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cha"}

	floCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "frigogo"}

	floDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "duh"}

	floFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "golri"}

	floGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "flatsch"}

	floItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "iglù"}

	floJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "じゃんよ"}

	floKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "맞잖아"}

	floLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fresqui"}

	floRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "чхи"}

	floSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "这样唷"}

	floSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "fresqui"}

	floTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "這樣唷"}
)

var (
	floPhrase = nook.Languages{
		language.AmericanEnglish:      floAmericanEnglishName,
		language.CanadianFrench:       floCanadianFrenchName,
		language.Dutch:                floDutchName,
		language.French:               floFrenchName,
		language.German:               floGermanName,
		language.Italian:              floItalianName,
		language.Japanese:             floJapaneseName,
		language.Korean:               floKoreanName,
		language.LatinAmericanSpanish: floLatinAmericanSpanishName,
		language.Russian:              floRussianName,
		language.SimplifiedChinese:    floSimplifiedChineseName,
		language.Spanish:              floSpanishName,
		language.TraditionalChinese:   floTraditionalChineseName}
)

var (
	Flo = nook.Villager{
		Character:   floCharacter,
		Personality: nook.BigSister,
		Phrase:      floPhrase}
)
