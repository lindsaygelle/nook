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
		Value:    "Henri"}

	henryDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Henry"}

	henryFrenchName = nook.Name{
		Language: language.French,
		Value:    "Henri"}

	henryGermanName = nook.Name{
		Language: language.German,
		Value:    "Freddy"}

	henryItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Renato"}

	henryJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヘンリー"}

	henryKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "헨리"}

	henryLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Saperto"}

	henryRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Генри"}

	henrySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "亨利"}

	henrySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Saperto"}

	henryTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "亨利"}
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
		Animal:   animal.Frog,
		Birthday: henryBirthday,
		Code:     henryCode,
		Key:      character.Henry,
		Gender:   gender.Male,
		Name:     henryName,
		Special:  false}
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
		language.AmericanEnglish:      henryAmericanEnglishPhrase,
		language.CanadianFrench:       henryCanadianFrenchPhrase,
		language.Dutch:                henryDutchPhrase,
		language.French:               henryFrenchPhrase,
		language.German:               henryGermanPhrase,
		language.Italian:              henryItalianPhrase,
		language.Japanese:             henryJapanesePhrase,
		language.Korean:               henryKoreanPhrase,
		language.LatinAmericanSpanish: henryLatinAmericanSpanishPhrase,
		language.Russian:              henryRussianPhrase,
		language.SimplifiedChinese:    henrySimplifiedChinesePhrase,
		language.Spanish:              henrySpanishPhrase,
		language.TraditionalChinese:   henryTraditionalChinesePhrase}
)

var (
	Henry = nook.Villager{
		Character:   henryCharacter,
		Personality: personality.Smug,
		Phrase:      henryPhrase}
)
