package hippo

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Hippyhip hop là"}

	bubblesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bubbleshipster"}

	bubblesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Hippyhip hop là"}

	bubblesGermanName = nook.Name{
		Language: language.German,
		Value:    "Christinblupper"}

	bubblesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ippolitahiippy"}

	bubblesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チャコでガンス"}

	bubblesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "차코히포포"}

	bubblesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Hipólitaglop-glop"}

	bubblesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Баблсхипстер"}

	bubblesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "曹珂是是"}

	bubblesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Hipólitahippie"}

	bubblesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "曹珂是是"}
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
		Animal:   Hippo,
		Birthday: bubblesBirthday,
		Code:     bubblesCode,
		Gender:   nook.Female,
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
		Personality: nook.Peppy,
		Phrase:      bubblesPhrase}
)
