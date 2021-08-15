package cow

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Valéchoupette"}

	tipperDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tipperboelala"}

	tipperFrenchName = nook.Name{
		Language: language.French,
		Value:    "Valéchoupette"}

	tipperGermanName = nook.Name{
		Language: language.German,
		Value:    "Angeladingdong"}

	tipperItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tippersì sì sì"}

	tipperJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "まきばミルミル"}

	tipperKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마틸다우유우유"}

	tipperLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pintamuamuuu"}

	tipperRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Типпертеленок"}

	tipperSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "牧牧牛奶"}

	tipperSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pintacuernitos"}

	tipperTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "牧牧牛奶"}
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
		Animal:   Cow,
		Birthday: tipperBirthday,
		Code:     tipperCode,
		Gender:   nook.Female,
		Name:     tipperName}
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
	Tipper = nook.Villager{
		Character:   tipperCharacter,
		Personality: nook.Snooty,
		Phrase:      tipperPhrase}
)
