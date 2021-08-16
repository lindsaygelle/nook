package penguin

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	hopperBirthday = nook.Birthday{
		Day:   6,
		Month: time.April}
)

var (
	hopperCode = nook.Code{
		Value: "pgn03"}
)

var (
	hopperAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hopper"}

	hopperCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Victor"}

	hopperDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Hopper"}

	hopperFrenchName = nook.Name{
		Language: language.French,
		Value:    "Victor"}

	hopperGermanName = nook.Name{
		Language: language.German,
		Value:    "Hauke"}

	hopperItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pinguaio"}

	hopperJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ダルマン"}

	hopperKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "달만이"}

	hopperLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Güiñón"}

	hopperRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Хоппер"}

	hopperSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "达满"}

	hopperSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Güiñón"}

	hopperTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "達滿"}
)

var (
	hopperName = nook.Languages{
		language.AmericanEnglish:      hopperAmericanEnglishName,
		language.CanadianFrench:       hopperCanadianFrenchName,
		language.Dutch:                hopperDutchName,
		language.French:               hopperFrenchName,
		language.German:               hopperGermanName,
		language.Italian:              hopperItalianName,
		language.Japanese:             hopperJapaneseName,
		language.Korean:               hopperKoreanName,
		language.LatinAmericanSpanish: hopperLatinAmericanSpanishName,
		language.Russian:              hopperRussianName,
		language.SimplifiedChinese:    hopperSimplifiedChineseName,
		language.Spanish:              hopperSpanishName,
		language.TraditionalChinese:   hopperTraditionalChineseName}
)

var (
	hopperCharacter = nook.Character{
		Animal:   animal.Penguin,
		Birthday: hopperBirthday,
		Code:     hopperCode,
		Gender:   gender.Male,
		Name:     hopperName}
)

var (
	hopperAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "slushie"}

	hopperCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "manette"}

	hopperDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ĳsje"}

	hopperFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "manette"}

	hopperGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "slaschi"}

	hopperItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sguish"}

	hopperJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ちぅねん"}

	hopperKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "아따"}

	hopperLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "esploch"}

	hopperRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ледок"}

	hopperSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "禅"}

	hopperSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "esploch"}

	hopperTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "禪"}
)

var (
	hopperPhrase = nook.Languages{
		language.AmericanEnglish:      hopperAmericanEnglishName,
		language.CanadianFrench:       hopperCanadianFrenchName,
		language.Dutch:                hopperDutchName,
		language.French:               hopperFrenchName,
		language.German:               hopperGermanName,
		language.Italian:              hopperItalianName,
		language.Japanese:             hopperJapaneseName,
		language.Korean:               hopperKoreanName,
		language.LatinAmericanSpanish: hopperLatinAmericanSpanishName,
		language.Russian:              hopperRussianName,
		language.SimplifiedChinese:    hopperSimplifiedChineseName,
		language.Spanish:              hopperSpanishName,
		language.TraditionalChinese:   hopperTraditionalChineseName}
)

var (
	Hopper = nook.Villager{
		Character:   hopperCharacter,
		Personality: personality.Cranky,
		Phrase:      hopperPhrase}
)
