package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	henryBirthday = nook.Birthday{
		Day:   21,
		Month: time.September}
)

var (
	henryCode = nook.Code{
		Value: "flg19"}
)

var (
	henryAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Henry"}

	henryCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Henrikôwa"}

	henryDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Henrysnurkwaak"}

	henryFrenchName = nook.Name{
		Language: language.French,
		Value:    "Henripustule"}

	henryGermanName = nook.Name{
		Language: language.German,
		Value:    "Freddyquappi"}

	henryItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Renatocroac"}

	henryJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヘンリーむにゃ"}

	henryKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "헨리음냐음냐"}

	henryLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sapertocruacruá"}

	henryRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Генриква-храп"}

	henrySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "亨利滴咕"}

	henrySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sapertocruacruá"}

	henryTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "亨利滴咕"}
)

var (
	henryName = nook.Languages{
		language.AmericanEnglish:      henryAmericanEnglishName,
		language.CanadianFrench:       henryCanadianFrenchName,
		language.Dutch:                henryDutchName,
		language.French:               henryFrenchName,
		language.German:               henryGermanName,
		language.Italian:              henryItalianName,
		language.Japanese:             henryJapaneseName,
		language.Korean:               henryKoreanName,
		language.LatinAmericanSpanish: henryLatinAmericanSpanishName,
		language.Russian:              henryRussianName,
		language.SimplifiedChinese:    henrySimplifiedChineseName,
		language.Spanish:              henrySpanishName,
		language.TraditionalChinese:   henryTraditionalChineseName}
)

var (
	henryCharacter = nook.Character{
		Animal:   Frog,
		Birthday: henryBirthday,
		Code:     henryCode,
		Gender:   nook.Male,
		Name:     henryName}
)

var (
	henryAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snoozit"}

	henryCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "kôwa"}

	henryDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snurkwaak"}

	henryFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pustule"}

	henryGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quappi"}

	henryItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "croac"}

	henryJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "むにゃ"}

	henryKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "음냐음냐"}

	henryLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cruacruá"}

	henryRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ква-храп"}

	henrySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "滴咕"}

	henrySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cruacruá"}

	henryTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "滴咕"}
)

var (
	henryPhrase = nook.Languages{
		language.AmericanEnglish:      henryAmericanEnglishName,
		language.CanadianFrench:       henryCanadianFrenchName,
		language.Dutch:                henryDutchName,
		language.French:               henryFrenchName,
		language.German:               henryGermanName,
		language.Italian:              henryItalianName,
		language.Japanese:             henryJapaneseName,
		language.Korean:               henryKoreanName,
		language.LatinAmericanSpanish: henryLatinAmericanSpanishName,
		language.Russian:              henryRussianName,
		language.SimplifiedChinese:    henrySimplifiedChineseName,
		language.Spanish:              henrySpanishName,
		language.TraditionalChinese:   henryTraditionalChineseName}
)

var (
	Henry = nook.Villager{
		Character:   henryCharacter,
		Personality: nook.Smug,
		Phrase:      henryPhrase}
)
