package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	purrlBirthday = nook.Birthday{
		Day:   29,
		Month: time.May}
)

var (
	purrlCode = nook.Code{
		Value: "cat07"}
)

var (
	purrlAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Purrl"}

	purrlCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Perlechaton"}

	purrlDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Purrlkitten"}

	purrlFrenchName = nook.Name{
		Language: language.French,
		Value:    "Perlechaton"}

	purrlGermanName = nook.Name{
		Language: language.German,
		Value:    "Frankakätzchen"}

	purrlItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Felidiamicetto"}

	purrlJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "たまふんっ"}

	purrlKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "타마냠"}

	purrlLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Wandamichimiu"}

	purrlRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Перлкис-кис"}

	purrlSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小玉嗯嗯"}

	purrlSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Wandatontaina"}

	purrlTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小玉嗯嗯"}
)

var (
	purrlName = nook.Languages{
		language.AmericanEnglish:      purrlAmericanEnglishName,
		language.CanadianFrench:       purrlCanadianFrenchName,
		language.Dutch:                purrlDutchName,
		language.French:               purrlFrenchName,
		language.German:               purrlGermanName,
		language.Italian:              purrlItalianName,
		language.Japanese:             purrlJapaneseName,
		language.Korean:               purrlKoreanName,
		language.LatinAmericanSpanish: purrlLatinAmericanSpanishName,
		language.Russian:              purrlRussianName,
		language.SimplifiedChinese:    purrlSimplifiedChineseName,
		language.Spanish:              purrlSpanishName,
		language.TraditionalChinese:   purrlTraditionalChineseName}
)

var (
	purrlCharacter = nook.Character{
		Animal:   Cat,
		Birthday: purrlBirthday,
		Code:     purrlCode,
		Gender:   nook.Female,
		Name:     purrlName}
)

var (
	purrlAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "kitten"}

	purrlCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chaton"}

	purrlDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kitten"}

	purrlFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "chaton"}

	purrlGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kätzchen"}

	purrlItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "micetto"}

	purrlJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ふんっ"}

	purrlKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "냠"}

	purrlLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "michimiu"}

	purrlRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кис-кис"}

	purrlSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嗯嗯"}

	purrlSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tontaina"}

	purrlTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗯嗯"}
)

var (
	purrlPhrase = nook.Languages{
		language.AmericanEnglish:      purrlAmericanEnglishName,
		language.CanadianFrench:       purrlCanadianFrenchName,
		language.Dutch:                purrlDutchName,
		language.French:               purrlFrenchName,
		language.German:               purrlGermanName,
		language.Italian:              purrlItalianName,
		language.Japanese:             purrlJapaneseName,
		language.Korean:               purrlKoreanName,
		language.LatinAmericanSpanish: purrlLatinAmericanSpanishName,
		language.Russian:              purrlRussianName,
		language.SimplifiedChinese:    purrlSimplifiedChineseName,
		language.Spanish:              purrlSpanishName,
		language.TraditionalChinese:   purrlTraditionalChineseName}
)

var (
	Purrl = nook.Villager{
		Character:   purrlCharacter,
		Personality: nook.Snooty,
		Phrase:      purrlPhrase}
)
