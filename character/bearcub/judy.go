package bearcub

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	judyBirthday = nook.Birthday{
		Day:   10,
		Month: time.March}
)

var (
	judyCode = nook.Code{
		Value: "cbr19"}
)

var (
	judyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Judy"}

	judyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Laura"}

	judyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Judy"}

	judyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Laura"}

	judyGermanName = nook.Name{
		Language: language.German,
		Value:    "Misuzu"}

	judyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Misuzu"}

	judyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "みすず"}

	judyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "미애"}

	judyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rosezna"}

	judyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джуди"}

	judySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "美玲"}

	judySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rosezna"}

	judyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "美玲"}
)

var (
	judyName = nook.Languages{
		language.AmericanEnglish:      judyAmericanEnglishName,
		language.CanadianFrench:       judyCanadianFrenchName,
		language.Dutch:                judyDutchName,
		language.French:               judyFrenchName,
		language.German:               judyGermanName,
		language.Italian:              judyItalianName,
		language.Japanese:             judyJapaneseName,
		language.Korean:               judyKoreanName,
		language.LatinAmericanSpanish: judyLatinAmericanSpanishName,
		language.Russian:              judyRussianName,
		language.SimplifiedChinese:    judySimplifiedChineseName,
		language.Spanish:              judySpanishName,
		language.TraditionalChinese:   judyTraditionalChineseName}
)

var (
	judyCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: judyBirthday,
		Code:     judyCode,
		Gender:   gender.Female,
		Name:     judyName}
)

var (
	judyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "myohmy"}

	judyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "oh là là"}

	judyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "nee maar"}

	judyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "oh là là"}

	judyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "träum"}

	judyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "misu misu"}

	judyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "あらら"}

	judyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어머머"}

	judyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "uyuyuy"}

	judyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "охохонюшки"}

	judySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哎呀"}

	judySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "qué cosas"}

	judyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哎呀"}
)

var (
	judyPhrase = nook.Languages{
		language.AmericanEnglish:      judyAmericanEnglishName,
		language.CanadianFrench:       judyCanadianFrenchName,
		language.Dutch:                judyDutchName,
		language.French:               judyFrenchName,
		language.German:               judyGermanName,
		language.Italian:              judyItalianName,
		language.Japanese:             judyJapaneseName,
		language.Korean:               judyKoreanName,
		language.LatinAmericanSpanish: judyLatinAmericanSpanishName,
		language.Russian:              judyRussianName,
		language.SimplifiedChinese:    judySimplifiedChineseName,
		language.Spanish:              judySpanishName,
		language.TraditionalChinese:   judyTraditionalChineseName}
)

var (
	Judy = nook.Villager{
		Character:   judyCharacter,
		Personality: personality.Snooty,
		Phrase:      judyPhrase}
)
