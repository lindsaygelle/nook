package gorilla

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	alBirthday = nook.Birthday{
		Day:   18,
		Month: time.October}
)

var (
	alCode = nook.Code{
		Value: "gor08"}
)

var (
	alAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Al"}

	alCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gustave"}

	alDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Al"}

	alFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gustave"}

	alGermanName = nook.Name{
		Language: language.German,
		Value:    "Kokong"}

	alItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gregorio"}

	alJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "たもつ"}

	alKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "우락"}

	alLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Álex"}

	alRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Эл"}

	alSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿保"}

	alSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Álex"}

	alTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿保"}
)

var (
	alName = nook.Languages{
		language.AmericanEnglish:      alAmericanEnglishName,
		language.CanadianFrench:       alCanadianFrenchName,
		language.Dutch:                alDutchName,
		language.French:               alFrenchName,
		language.German:               alGermanName,
		language.Italian:              alItalianName,
		language.Japanese:             alJapaneseName,
		language.Korean:               alKoreanName,
		language.LatinAmericanSpanish: alLatinAmericanSpanishName,
		language.Russian:              alRussianName,
		language.SimplifiedChinese:    alSimplifiedChineseName,
		language.Spanish:              alSpanishName,
		language.TraditionalChinese:   alTraditionalChineseName}
)

var (
	alCharacter = nook.Character{
		Animal:   Gorilla,
		Birthday: alBirthday,
		Code:     alCode,
		Gender:   nook.Male,
		Name:     alName}
)

var (
	alAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hoo hoo ha"}

	alCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ouba-oub"}

	alDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "aiaiai"}

	alFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ouba-oub"}

	alGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "uga-uga"}

	alItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bananao"}

	alJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウヒョッ"}

	alKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "부락"}

	alLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "juju-já"}

	alRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ух-ух-ух"}

	alSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呜呼"}

	alSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chocolate"}

	alTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗚呼"}
)

var (
	alPhrase = nook.Languages{
		language.AmericanEnglish:      alAmericanEnglishName,
		language.CanadianFrench:       alCanadianFrenchName,
		language.Dutch:                alDutchName,
		language.French:               alFrenchName,
		language.German:               alGermanName,
		language.Italian:              alItalianName,
		language.Japanese:             alJapaneseName,
		language.Korean:               alKoreanName,
		language.LatinAmericanSpanish: alLatinAmericanSpanishName,
		language.Russian:              alRussianName,
		language.SimplifiedChinese:    alSimplifiedChineseName,
		language.Spanish:              alSpanishName,
		language.TraditionalChinese:   alTraditionalChineseName}
)

var (
	Al = nook.Villager{
		Character:   alCharacter,
		Personality: nook.Lazy,
		Phrase:      alPhrase}
)
