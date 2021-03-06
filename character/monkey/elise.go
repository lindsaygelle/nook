package monkey

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
	eliseBirthday = nook.Birthday{
		Day:   21,
		Month: time.March}
)

var (
	eliseCode = nook.Code{
		Value: "mnk05"}
)

var (
	eliseAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Elise"}

	eliseCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Élise"}

	eliseDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Elise"}

	eliseFrenchName = nook.Name{
		Language: language.French,
		Value:    "Élise"}

	eliseGermanName = nook.Name{
		Language: language.German,
		Value:    "Steffi"}

	eliseItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Elisa"}

	eliseJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "モンこ"}

	eliseKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "몽자"}

	eliseLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Mayra"}

	eliseRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Элиза"}

	eliseSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "孟琪"}

	eliseSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mayra"}

	eliseTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "孟琪"}
)

var (
	eliseName = nook.Languages{
		language.AmericanEnglish:      eliseAmericanEnglishName,
		language.CanadianFrench:       eliseCanadianFrenchName,
		language.Dutch:                eliseDutchName,
		language.French:               eliseFrenchName,
		language.German:               eliseGermanName,
		language.Italian:              eliseItalianName,
		language.Japanese:             eliseJapaneseName,
		language.Korean:               eliseKoreanName,
		language.LatinAmericanSpanish: eliseLatinAmericanSpanishName,
		language.Russian:              eliseRussianName,
		language.SimplifiedChinese:    eliseSimplifiedChineseName,
		language.Spanish:              eliseSpanishName,
		language.TraditionalChinese:   eliseTraditionalChineseName}
)

var (
	eliseCharacter = nook.Character{
		Animal:   animal.Monkey,
		Birthday: eliseBirthday,
		Code:     eliseCode,
		Key:      character.Elise,
		Gender:   gender.Female,
		Name:     eliseName,
		Special:  false}
)

var (
	eliseAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "puh-lease"}

	eliseCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "steup"}

	eliseDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kom op zeg"}

	eliseFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "steup"}

	eliseGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "strebstreb"}

	eliseItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "per favore"}

	eliseJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だモン"}

	eliseKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "몽몽"}

	eliseLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "uh-lalá"}

	eliseRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "само собой"}

	eliseSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "孟孟"}

	eliseSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ricura"}

	eliseTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "孟孟"}
)

var (
	elisePhrase = nook.Languages{
		language.AmericanEnglish:      eliseAmericanEnglishPhrase,
		language.CanadianFrench:       eliseCanadianFrenchPhrase,
		language.Dutch:                eliseDutchPhrase,
		language.French:               eliseFrenchPhrase,
		language.German:               eliseGermanPhrase,
		language.Italian:              eliseItalianPhrase,
		language.Japanese:             eliseJapanesePhrase,
		language.Korean:               eliseKoreanPhrase,
		language.LatinAmericanSpanish: eliseLatinAmericanSpanishPhrase,
		language.Russian:              eliseRussianPhrase,
		language.SimplifiedChinese:    eliseSimplifiedChinesePhrase,
		language.Spanish:              eliseSpanishPhrase,
		language.TraditionalChinese:   eliseTraditionalChinesePhrase}
)

var (
	Elise = nook.Villager{
		Character:   eliseCharacter,
		Personality: personality.Snooty,
		Phrase:      elisePhrase}
)
