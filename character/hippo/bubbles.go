package hippo

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
	bubblesBirthday = nook.Birthday{
		Day:   18,
		Month: time.September}
)

var (
	bubblesCode = nook.Code{
		Value: "hip02"}
)

var (
	bubblesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bubbles"}

	bubblesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Hippy"}

	bubblesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bubbles"}

	bubblesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Hippy"}

	bubblesGermanName = nook.Name{
		Language: language.German,
		Value:    "Christin"}

	bubblesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ippolita"}

	bubblesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チャコ"}

	bubblesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "차코"}

	bubblesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Hipólita"}

	bubblesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Баблс"}

	bubblesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "曹珂"}

	bubblesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Hipólita"}

	bubblesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "曹珂"}
)

var (
	bubblesName = nook.Languages{
		language.AmericanEnglish:      bubblesAmericanEnglishName,
		language.CanadianFrench:       bubblesCanadianFrenchName,
		language.Dutch:                bubblesDutchName,
		language.French:               bubblesFrenchName,
		language.German:               bubblesGermanName,
		language.Italian:              bubblesItalianName,
		language.Japanese:             bubblesJapaneseName,
		language.Korean:               bubblesKoreanName,
		language.LatinAmericanSpanish: bubblesLatinAmericanSpanishName,
		language.Russian:              bubblesRussianName,
		language.SimplifiedChinese:    bubblesSimplifiedChineseName,
		language.Spanish:              bubblesSpanishName,
		language.TraditionalChinese:   bubblesTraditionalChineseName}
)

var (
	bubblesCharacter = nook.Character{
		Animal:   animal.Hippo,
		Birthday: bubblesBirthday,
		Code:     bubblesCode,
		Key:      character.Bubbles,
		Gender:   gender.Female,
		Name:     bubblesName}
)

var (
	bubblesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hipster"}

	bubblesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "hip hop là"}

	bubblesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hipster"}

	bubblesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "hip hop là"}

	bubblesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "blupper"}

	bubblesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hiippy"}

	bubblesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でガンス"}

	bubblesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "히포포"}

	bubblesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "glop-glop"}

	bubblesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хипстер"}

	bubblesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是是"}

	bubblesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "hippie"}

	bubblesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是是"}
)

var (
	bubblesPhrase = nook.Languages{
		language.AmericanEnglish:      bubblesAmericanEnglishName,
		language.CanadianFrench:       bubblesCanadianFrenchName,
		language.Dutch:                bubblesDutchName,
		language.French:               bubblesFrenchName,
		language.German:               bubblesGermanName,
		language.Italian:              bubblesItalianName,
		language.Japanese:             bubblesJapaneseName,
		language.Korean:               bubblesKoreanName,
		language.LatinAmericanSpanish: bubblesLatinAmericanSpanishName,
		language.Russian:              bubblesRussianName,
		language.SimplifiedChinese:    bubblesSimplifiedChineseName,
		language.Spanish:              bubblesSpanishName,
		language.TraditionalChinese:   bubblesTraditionalChineseName}
)

var (
	Bubbles = nook.Villager{
		Character:   bubblesCharacter,
		Personality: personality.Peppy,
		Phrase:      bubblesPhrase}
)
