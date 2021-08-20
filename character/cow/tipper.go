package cow

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
	tipperBirthday = nook.Birthday{
		Day:   25,
		Month: time.August}
)

var (
	tipperCode = nook.Code{
		Value: "cow01"}
)

var (
	tipperAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tipper"}

	tipperCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Valé"}

	tipperDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tipper"}

	tipperFrenchName = nook.Name{
		Language: language.French,
		Value:    "Valé"}

	tipperGermanName = nook.Name{
		Language: language.German,
		Value:    "Angela"}

	tipperItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tipper"}

	tipperJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "まきば"}

	tipperKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마틸다"}

	tipperLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pinta"}

	tipperRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Типпер"}

	tipperSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "牧牧"}

	tipperSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pinta"}

	tipperTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "牧牧"}
)

var (
	tipperName = nook.Languages{
		language.AmericanEnglish:      tipperAmericanEnglishName,
		language.CanadianFrench:       tipperCanadianFrenchName,
		language.Dutch:                tipperDutchName,
		language.French:               tipperFrenchName,
		language.German:               tipperGermanName,
		language.Italian:              tipperItalianName,
		language.Japanese:             tipperJapaneseName,
		language.Korean:               tipperKoreanName,
		language.LatinAmericanSpanish: tipperLatinAmericanSpanishName,
		language.Russian:              tipperRussianName,
		language.SimplifiedChinese:    tipperSimplifiedChineseName,
		language.Spanish:              tipperSpanishName,
		language.TraditionalChinese:   tipperTraditionalChineseName}
)

var (
	tipperCharacter = nook.Character{
		Animal:   animal.Cow,
		Birthday: tipperBirthday,
		Code:     tipperCode,
		Key:      character.Tipper,
		Gender:   gender.Female,
		Name:     tipperName,
		Special:  false}
)

var (
	tipperAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pushy"}

	tipperCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "choupette"}

	tipperDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "boelala"}

	tipperFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "choupette"}

	tipperGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "dingdong"}

	tipperItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sì sì sì"}

	tipperJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ミルミル"}

	tipperKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우유우유"}

	tipperLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "muamuuu"}

	tipperRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "теленок"}

	tipperSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "牛奶"}

	tipperSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cuernitos"}

	tipperTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "牛奶"}
)

var (
	tipperPhrase = nook.Languages{
		language.AmericanEnglish:      tipperAmericanEnglishPhrase,
		language.CanadianFrench:       tipperCanadianFrenchPhrase,
		language.Dutch:                tipperDutchPhrase,
		language.French:               tipperFrenchPhrase,
		language.German:               tipperGermanPhrase,
		language.Italian:              tipperItalianPhrase,
		language.Japanese:             tipperJapanesePhrase,
		language.Korean:               tipperKoreanPhrase,
		language.LatinAmericanSpanish: tipperLatinAmericanSpanishPhrase,
		language.Russian:              tipperRussianPhrase,
		language.SimplifiedChinese:    tipperSimplifiedChinesePhrase,
		language.Spanish:              tipperSpanishPhrase,
		language.TraditionalChinese:   tipperTraditionalChinesePhrase}
)

var (
	Tipper = nook.Villager{
		Character:   tipperCharacter,
		Personality: personality.Snooty,
		Phrase:      tipperPhrase}
)
