package penguin

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
	puckBirthday = nook.Birthday{
		Day:   21,
		Month: time.February}
)

var (
	puckCode = nook.Code{
		Value: "pgn06"}
)

var (
	puckAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Puck"}

	puckCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Hervé"}

	puckDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Puck"}

	puckFrenchName = nook.Name{
		Language: language.French,
		Value:    "Hervé"}

	puckGermanName = nook.Name{
		Language: language.German,
		Value:    "Puck"}

	puckItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Glacido"}

	puckJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ホッケー"}

	puckKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "하키"}

	puckLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fredo"}

	puckRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пак"}

	puckSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哈奇"}

	puckSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Fredo"}

	puckTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哈奇"}
)

var (
	puckName = nook.Languages{
		language.AmericanEnglish:      puckAmericanEnglishName,
		language.CanadianFrench:       puckCanadianFrenchName,
		language.Dutch:                puckDutchName,
		language.French:               puckFrenchName,
		language.German:               puckGermanName,
		language.Italian:              puckItalianName,
		language.Japanese:             puckJapaneseName,
		language.Korean:               puckKoreanName,
		language.LatinAmericanSpanish: puckLatinAmericanSpanishName,
		language.Russian:              puckRussianName,
		language.SimplifiedChinese:    puckSimplifiedChineseName,
		language.Spanish:              puckSpanishName,
		language.TraditionalChinese:   puckTraditionalChineseName}
)

var (
	puckCharacter = nook.Character{
		Animal:   animal.Penguin,
		Birthday: puckBirthday,
		Code:     puckCode,
		Key:      character.Puck,
		Gender:   gender.Male,
		Name:     puckName,
		Special:  false}
)

var (
	puckAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "brrrrrrrrr"}

	puckCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "brrrrrrrrr"}

	puckDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "brrrrrrrrr"}

	puckFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "brrrrrrrrr"}

	puckGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brrrrrrrrr"}

	puckItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "brrrrrrrrr"}

	puckJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "さぶー"}

	puckKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "엣취"}

	puckLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "brrr-uah"}

	puckRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "брр"}

	puckSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "候补"}

	puckSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "sardinilla"}

	puckTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "候補"}
)

var (
	puckPhrase = nook.Languages{
		language.AmericanEnglish:      puckAmericanEnglishPhrase,
		language.CanadianFrench:       puckCanadianFrenchPhrase,
		language.Dutch:                puckDutchPhrase,
		language.French:               puckFrenchPhrase,
		language.German:               puckGermanPhrase,
		language.Italian:              puckItalianPhrase,
		language.Japanese:             puckJapanesePhrase,
		language.Korean:               puckKoreanPhrase,
		language.LatinAmericanSpanish: puckLatinAmericanSpanishPhrase,
		language.Russian:              puckRussianPhrase,
		language.SimplifiedChinese:    puckSimplifiedChinesePhrase,
		language.Spanish:              puckSpanishPhrase,
		language.TraditionalChinese:   puckTraditionalChinesePhrase}
)

var (
	Puck = nook.Villager{
		Character:   puckCharacter,
		Personality: personality.Lazy,
		Phrase:      puckPhrase}
)
