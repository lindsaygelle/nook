package rabbit

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	genjiBirthday = nook.Birthday{
		Day:   21,
		Month: time.January}
)

var (
	genjiCode = nook.Code{
		Value: "rbt08"}
)

var (
	genjiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Genji"}

	genjiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kalimochi"}

	genjiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Genjimochi"}

	genjiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kalimochi"}

	genjiGermanName = nook.Name{
		Language: language.German,
		Value:    "Akigoseimas"}

	genjiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Genjimochi"}

	genjiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゲンジまろ"}

	genjiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "토시땡글"}

	genjiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sumomochi"}

	genjiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гэндзизайпончик"}

	genjiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "儒林臣"}

	genjiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sumomochi"}

	genjiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "儒林臣"}
)

var (
	genjiName = nook.Languages{
		language.AmericanEnglish:      genjiAmericanEnglishName,
		language.CanadianFrench:       genjiCanadianFrenchName,
		language.Dutch:                genjiDutchName,
		language.French:               genjiFrenchName,
		language.German:               genjiGermanName,
		language.Italian:              genjiItalianName,
		language.Japanese:             genjiJapaneseName,
		language.Korean:               genjiKoreanName,
		language.LatinAmericanSpanish: genjiLatinAmericanSpanishName,
		language.Russian:              genjiRussianName,
		language.SimplifiedChinese:    genjiSimplifiedChineseName,
		language.Spanish:              genjiSpanishName,
		language.TraditionalChinese:   genjiTraditionalChineseName}
)

var (
	genjiCharacter = nook.Character{
		Animal:   Rabbit,
		Birthday: genjiBirthday,
		Code:     genjiCode,
		Gender:   nook.Male,
		Name:     genjiName}
)

var (
	genjiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "otakumochi"}

	genjiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mochi"}

	genjiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mochi"}

	genjiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mochi"}

	genjiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "goseimas"}

	genjiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mochi"}

	genjiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "まろ"}

	genjiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "땡글"}

	genjiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "mochi"}

	genjiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "зайпончик"}

	genjiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "臣"}

	genjiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mochi"}

	genjiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "臣"}
)

var (
	genjiPhrase = nook.Languages{
		language.AmericanEnglish:      genjiAmericanEnglishName,
		language.CanadianFrench:       genjiCanadianFrenchName,
		language.Dutch:                genjiDutchName,
		language.French:               genjiFrenchName,
		language.German:               genjiGermanName,
		language.Italian:              genjiItalianName,
		language.Japanese:             genjiJapaneseName,
		language.Korean:               genjiKoreanName,
		language.LatinAmericanSpanish: genjiLatinAmericanSpanishName,
		language.Russian:              genjiRussianName,
		language.SimplifiedChinese:    genjiSimplifiedChineseName,
		language.Spanish:              genjiSpanishName,
		language.TraditionalChinese:   genjiTraditionalChineseName}
)

var (
	Genji = nook.Villager{
		Character:   genjiCharacter,
		Personality: nook.Jock,
		Phrase:      genjiPhrase}
)
