package cat

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
		Value:    "Kitty"}

	kittyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kitty"}

	kittyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kitty"}

	kittyGermanName = nook.Name{
		Language: language.German,
		Value:    "Karen"}

	kittyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ester"}

	kittyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ショコラ"}

	kittyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "쇼콜라"}

	kittyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Kasandra"}

	kittyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Китти"}

	kittySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "朱古莉"}

	kittySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Kasandra"}

	kittyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "朱古莉"}
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
		Animal:   animal.Cat,
		Birthday: kittyBirthday,
		Code:     kittyCode,
		Key:      character.Kitty,
		Gender:   gender.Female,
		Name:     kittyName,
		Special:  false}
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
		language.AmericanEnglish:      kittyAmericanEnglishPhrase,
		language.CanadianFrench:       kittyCanadianFrenchPhrase,
		language.Dutch:                kittyDutchPhrase,
		language.French:               kittyFrenchPhrase,
		language.German:               kittyGermanPhrase,
		language.Italian:              kittyItalianPhrase,
		language.Japanese:             kittyJapanesePhrase,
		language.Korean:               kittyKoreanPhrase,
		language.LatinAmericanSpanish: kittyLatinAmericanSpanishPhrase,
		language.Russian:              kittyRussianPhrase,
		language.SimplifiedChinese:    kittySimplifiedChinesePhrase,
		language.Spanish:              kittySpanishPhrase,
		language.TraditionalChinese:   kittyTraditionalChinesePhrase}
)

var (
	Kitty = nook.Villager{
		Character:   kittyCharacter,
		Personality: personality.Snooty,
		Phrase:      kittyPhrase}
)
