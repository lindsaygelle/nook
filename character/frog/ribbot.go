package frog

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
	// ribbotBirthday represents ribbot birthday.
	ribbotBirthday = nook.Birthday{
		Day:   13,
		Month: time.February}
)

var (
	// ribbotCode represents ribbot code.
	ribbotCode = nook.Code{
		Value: "flg01"}
)

var (
	// ribbotAmericanEnglishName represents ribbot american english name.
	ribbotAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ribbot"}

	// ribbotCanadianFrenchName represents ribbot canadian french name.
	ribbotCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Crabot"}

	// ribbotDutchName represents ribbot dutch name.
	ribbotDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ribbot"}

	// ribbotFrenchName represents ribbot french name.
	ribbotFrenchName = nook.Name{
		Language: language.French,
		Value:    "Crabot"}

	// ribbotGermanName represents ribbot german name.
	ribbotGermanName = nook.Name{
		Language: language.German,
		Value:    "Robbi"}

	// ribbotItalianName represents ribbot italian name.
	ribbotItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ranabot"}

	// ribbotJapaneseName represents ribbot japanese name.
	ribbotJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ガチャ"}

	// ribbotKoreanName represents ribbot korean name.
	ribbotKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "철컥"}

	// ribbotLatinAmericanSpanishName represents ribbot latin american spanish name.
	ribbotLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ranobot"}

	// ribbotRussianName represents ribbot russian name.
	ribbotRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Риббот"}

	// ribbotSimplifiedChineseName represents ribbot simplified chinese name.
	ribbotSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "锵锵"}

	// ribbotSpanishName represents ribbot spanish name.
	ribbotSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ranobot"}

	// ribbotTraditionalChineseName represents ribbot traditional chinese name.
	ribbotTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鏘鏘"}
)

var (
	// ribbotName represents ribbot name.
	ribbotName = nook.Languages{
		language.AmericanEnglish:      ribbotAmericanEnglishName,
		language.CanadianFrench:       ribbotCanadianFrenchName,
		language.Dutch:                ribbotDutchName,
		language.French:               ribbotFrenchName,
		language.German:               ribbotGermanName,
		language.Italian:              ribbotItalianName,
		language.Japanese:             ribbotJapaneseName,
		language.Korean:               ribbotKoreanName,
		language.LatinAmericanSpanish: ribbotLatinAmericanSpanishName,
		language.Russian:              ribbotRussianName,
		language.SimplifiedChinese:    ribbotSimplifiedChineseName,
		language.Spanish:              ribbotSpanishName,
		language.TraditionalChinese:   ribbotTraditionalChineseName}
)

var (
	// ribbotCharacter represents ribbot character.
	ribbotCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: ribbotBirthday,
		Code:     ribbotCode,
		Key:      character.Ribbot,
		Gender:   gender.Male,
		Name:     ribbotName,
		Special:  false}
)

var (
	// ribbotAmericanEnglishPhrase represents ribbot american english phrase.
	ribbotAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "toady"}

	// ribbotCanadianFrenchPhrase represents ribbot canadian french phrase.
	ribbotCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "croac"}

	// ribbotDutchPhrase represents ribbot dutch phrase.
	ribbotDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kwwwrrrk"}

	// ribbotFrenchPhrase represents ribbot french phrase.
	ribbotFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "croac"}

	// ribbotGermanPhrase represents ribbot german phrase.
	ribbotGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quarrrk"}

	// ribbotItalianPhrase represents ribbot italian phrase.
	ribbotItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cracrà"}

	// ribbotJapanesePhrase represents ribbot japanese phrase.
	ribbotJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だロボ"}

	// ribbotKoreanPhrase represents ribbot korean phrase.
	ribbotKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오버"}

	// ribbotLatinAmericanSpanishPhrase represents ribbot latin american spanish phrase.
	ribbotLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "crobit"}

	// ribbotRussianPhrase represents ribbot russian phrase.
	ribbotRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "квак-дзынь"}

	// ribbotSimplifiedChinesePhrase represents ribbot simplified chinese phrase.
	ribbotSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "机器"}

	// ribbotSpanishPhrase represents ribbot spanish phrase.
	ribbotSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "crobit"}

	// ribbotTraditionalChinesePhrase represents ribbot traditional chinese phrase.
	ribbotTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "機器"}
)

var (
	// ribbotPhrase represents ribbot phrase.
	ribbotPhrase = nook.Languages{
		language.AmericanEnglish:      ribbotAmericanEnglishPhrase,
		language.CanadianFrench:       ribbotCanadianFrenchPhrase,
		language.Dutch:                ribbotDutchPhrase,
		language.French:               ribbotFrenchPhrase,
		language.German:               ribbotGermanPhrase,
		language.Italian:              ribbotItalianPhrase,
		language.Japanese:             ribbotJapanesePhrase,
		language.Korean:               ribbotKoreanPhrase,
		language.LatinAmericanSpanish: ribbotLatinAmericanSpanishPhrase,
		language.Russian:              ribbotRussianPhrase,
		language.SimplifiedChinese:    ribbotSimplifiedChinesePhrase,
		language.Spanish:              ribbotSpanishPhrase,
		language.TraditionalChinese:   ribbotTraditionalChinesePhrase}
)

var (
	// Ribbot represents ribbot.
	Ribbot = nook.Villager{
		Character:   ribbotCharacter,
		Personality: personality.Jock,
		Phrase:      ribbotPhrase}
)
