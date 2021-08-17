package penguin

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
		Value:    "Nora"}

	floDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Flo"}

	floFrenchName = nook.Name{
		Language: language.French,
		Value:    "Nora"}

	floGermanName = nook.Name{
		Language: language.German,
		Value:    "Susanne"}

	floItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Nives"}

	floJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "レイラ"}

	floKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "레이라"}

	floLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nieves"}

	floRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фло"}

	floSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蕾拉"}

	floSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Nieves"}

	floTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蕾拉"}
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
		Animal:   animal.Penguin,
		Birthday: floBirthday,
		Code:     floCode,
		Key:      character.Flo,
		Gender:   gender.Female,
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
		language.AmericanEnglish:      floAmericanEnglishPhrase,
		language.CanadianFrench:       floCanadianFrenchPhrase,
		language.Dutch:                floDutchPhrase,
		language.French:               floFrenchPhrase,
		language.German:               floGermanPhrase,
		language.Italian:              floItalianPhrase,
		language.Japanese:             floJapanesePhrase,
		language.Korean:               floKoreanPhrase,
		language.LatinAmericanSpanish: floLatinAmericanSpanishPhrase,
		language.Russian:              floRussianPhrase,
		language.SimplifiedChinese:    floSimplifiedChinesePhrase,
		language.Spanish:              floSpanishPhrase,
		language.TraditionalChinese:   floTraditionalChinesePhrase}
)

var (
	Flo = nook.Villager{
		Character:   floCharacter,
		Personality: personality.BigSister,
		Phrase:      floPhrase}
)
