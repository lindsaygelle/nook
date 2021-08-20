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
		Value:    "Perle"}

	purrlDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Purrl"}

	purrlFrenchName = nook.Name{
		Language: language.French,
		Value:    "Perle"}

	purrlGermanName = nook.Name{
		Language: language.German,
		Value:    "Franka"}

	purrlItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Felidia"}

	purrlJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "たま"}

	purrlKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "타마"}

	purrlLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Wanda"}

	purrlRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Перл"}

	purrlSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小玉"}

	purrlSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Wanda"}

	purrlTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小玉"}
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
		Animal:   animal.Cat,
		Birthday: purrlBirthday,
		Code:     purrlCode,
		Key:      character.Purrl,
		Gender:   gender.Female,
		Name:     purrlName,
		Special:  false}
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
		language.AmericanEnglish:      purrlAmericanEnglishPhrase,
		language.CanadianFrench:       purrlCanadianFrenchPhrase,
		language.Dutch:                purrlDutchPhrase,
		language.French:               purrlFrenchPhrase,
		language.German:               purrlGermanPhrase,
		language.Italian:              purrlItalianPhrase,
		language.Japanese:             purrlJapanesePhrase,
		language.Korean:               purrlKoreanPhrase,
		language.LatinAmericanSpanish: purrlLatinAmericanSpanishPhrase,
		language.Russian:              purrlRussianPhrase,
		language.SimplifiedChinese:    purrlSimplifiedChinesePhrase,
		language.Spanish:              purrlSpanishPhrase,
		language.TraditionalChinese:   purrlTraditionalChinesePhrase}
)

var (
	Purrl = nook.Villager{
		Character:   purrlCharacter,
		Personality: personality.Snooty,
		Phrase:      purrlPhrase}
)
