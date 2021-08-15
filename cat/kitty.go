package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	kittyBirthday = nook.Birthday{
		Day:   15,
		Month: time.February}
)

var (
	kittyCode = nook.Code{
		Value: "cat14"}
)

var (
	kittyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kitty"}

	kittyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kittymiaaaouh"}

	kittyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kittygiechel"}

	kittyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kittymiaaaouh"}

	kittyGermanName = nook.Name{
		Language: language.German,
		Value:    "Karenmiezmiez"}

	kittyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Esterfusolo"}

	kittyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ショコラフフ"}

	kittyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "쇼콜라자갸"}

	kittyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Kasandramiaul"}

	kittyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Киттимр-мр-мр"}

	kittySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "朱古莉古古"}

	kittySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Kasandracascabel"}

	kittyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "朱古莉古古"}
)

var (
	kittyName = nook.Languages{
		language.AmericanEnglish:      kittyAmericanEnglishName,
		language.CanadianFrench:       kittyCanadianFrenchName,
		language.Dutch:                kittyDutchName,
		language.French:               kittyFrenchName,
		language.German:               kittyGermanName,
		language.Italian:              kittyItalianName,
		language.Japanese:             kittyJapaneseName,
		language.Korean:               kittyKoreanName,
		language.LatinAmericanSpanish: kittyLatinAmericanSpanishName,
		language.Russian:              kittyRussianName,
		language.SimplifiedChinese:    kittySimplifiedChineseName,
		language.Spanish:              kittySpanishName,
		language.TraditionalChinese:   kittyTraditionalChineseName}
)

var (
	kittyCharacter = nook.Character{
		Animal:   Cat,
		Birthday: kittyBirthday,
		Code:     kittyCode,
		Gender:   nook.Female,
		Name:     kittyName}
)

var (
	kittyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "mrowrr"}

	kittyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "miaaaouh"}

	kittyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "giechel"}

	kittyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "miaaaouh"}

	kittyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "miezmiez"}

	kittyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "fusolo"}

	kittyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "フフ"}

	kittyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "자갸"}

	kittyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "miaul"}

	kittyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мр-мр-мр"}

	kittySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "古古"}

	kittySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cascabel"}

	kittyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "古古"}
)

var (
	kittyPhrase = nook.Languages{
		language.AmericanEnglish:      kittyAmericanEnglishName,
		language.CanadianFrench:       kittyCanadianFrenchName,
		language.Dutch:                kittyDutchName,
		language.French:               kittyFrenchName,
		language.German:               kittyGermanName,
		language.Italian:              kittyItalianName,
		language.Japanese:             kittyJapaneseName,
		language.Korean:               kittyKoreanName,
		language.LatinAmericanSpanish: kittyLatinAmericanSpanishName,
		language.Russian:              kittyRussianName,
		language.SimplifiedChinese:    kittySimplifiedChineseName,
		language.Spanish:              kittySpanishName,
		language.TraditionalChinese:   kittyTraditionalChineseName}
)

var (
	Kitty = nook.Villager{
		Character:   kittyCharacter,
		Personality: nook.Snooty,
		Phrase:      kittyPhrase}
)
