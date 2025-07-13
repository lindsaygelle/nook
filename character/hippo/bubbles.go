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
	// bubblesBirthday represents bubbles birthday.
	bubblesBirthday = nook.Birthday{
		Day:   18,
		Month: time.September}
)

var (
	// bubblesCode represents bubbles code.
	bubblesCode = nook.Code{
		Value: "hip02"}
)

var (
	// bubblesAmericanEnglishName represents bubbles american english name.
	bubblesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bubbles"}

	// bubblesCanadianFrenchName represents bubbles canadian french name.
	bubblesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Hippy"}

	// bubblesDutchName represents bubbles dutch name.
	bubblesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bubbles"}

	// bubblesFrenchName represents bubbles french name.
	bubblesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Hippy"}

	// bubblesGermanName represents bubbles german name.
	bubblesGermanName = nook.Name{
		Language: language.German,
		Value:    "Christin"}

	// bubblesItalianName represents bubbles italian name.
	bubblesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ippolita"}

	// bubblesJapaneseName represents bubbles japanese name.
	bubblesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チャコ"}

	// bubblesKoreanName represents bubbles korean name.
	bubblesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "차코"}

	// bubblesLatinAmericanSpanishName represents bubbles latin american spanish name.
	bubblesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Hipólita"}

	// bubblesRussianName represents bubbles russian name.
	bubblesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Баблс"}

	// bubblesSimplifiedChineseName represents bubbles simplified chinese name.
	bubblesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "曹珂"}

	// bubblesSpanishName represents bubbles spanish name.
	bubblesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Hipólita"}

	// bubblesTraditionalChineseName represents bubbles traditional chinese name.
	bubblesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "曹珂"}
)

var (
	// bubblesName represents bubbles name.
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
	// bubblesCharacter represents bubbles character.
	bubblesCharacter = nook.Character{
		Animal:   animal.Hippo,
		Birthday: bubblesBirthday,
		Code:     bubblesCode,
		Key:      character.Bubbles,
		Gender:   gender.Female,
		Name:     bubblesName,
		Special:  false}
)

var (
	// bubblesAmericanEnglishPhrase represents bubbles american english phrase.
	bubblesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hipster"}

	// bubblesCanadianFrenchPhrase represents bubbles canadian french phrase.
	bubblesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "hip hop là"}

	// bubblesDutchPhrase represents bubbles dutch phrase.
	bubblesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "hipster"}

	// bubblesFrenchPhrase represents bubbles french phrase.
	bubblesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "hip hop là"}

	// bubblesGermanPhrase represents bubbles german phrase.
	bubblesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "blupper"}

	// bubblesItalianPhrase represents bubbles italian phrase.
	bubblesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hiippy"}

	// bubblesJapanesePhrase represents bubbles japanese phrase.
	bubblesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でガンス"}

	// bubblesKoreanPhrase represents bubbles korean phrase.
	bubblesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "히포포"}

	// bubblesLatinAmericanSpanishPhrase represents bubbles latin american spanish phrase.
	bubblesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "glop-glop"}

	// bubblesRussianPhrase represents bubbles russian phrase.
	bubblesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хипстер"}

	// bubblesSimplifiedChinesePhrase represents bubbles simplified chinese phrase.
	bubblesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是是"}

	// bubblesSpanishPhrase represents bubbles spanish phrase.
	bubblesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "hippie"}

	// bubblesTraditionalChinesePhrase represents bubbles traditional chinese phrase.
	bubblesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是是"}
)

var (
	// bubblesPhrase represents bubbles phrase.
	bubblesPhrase = nook.Languages{
		language.AmericanEnglish:      bubblesAmericanEnglishPhrase,
		language.CanadianFrench:       bubblesCanadianFrenchPhrase,
		language.Dutch:                bubblesDutchPhrase,
		language.French:               bubblesFrenchPhrase,
		language.German:               bubblesGermanPhrase,
		language.Italian:              bubblesItalianPhrase,
		language.Japanese:             bubblesJapanesePhrase,
		language.Korean:               bubblesKoreanPhrase,
		language.LatinAmericanSpanish: bubblesLatinAmericanSpanishPhrase,
		language.Russian:              bubblesRussianPhrase,
		language.SimplifiedChinese:    bubblesSimplifiedChinesePhrase,
		language.Spanish:              bubblesSpanishPhrase,
		language.TraditionalChinese:   bubblesTraditionalChinesePhrase}
)

var (
	// Bubbles represents bubbles.
	Bubbles = nook.Villager{
		Character:   bubblesCharacter,
		Personality: personality.Peppy,
		Phrase:      bubblesPhrase}
)
