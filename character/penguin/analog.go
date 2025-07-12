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
	// analogBirthday represents analog birthday.
	analogBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// analogCode represents analog code.
	analogCode = nook.Code{
		Value: ""}
)

var (
	// analogAmericanEnglishName represents analog american english name.
	analogAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Analog"}

	// analogCanadianFrenchName represents analog canadian french name.
	analogCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// analogDutchName represents analog dutch name.
	analogDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// analogFrenchName represents analog french name.
	analogFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// analogGermanName represents analog german name.
	analogGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// analogItalianName represents analog italian name.
	analogItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// analogJapaneseName represents analog japanese name.
	analogJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アナログ"}

	// analogKoreanName represents analog korean name.
	analogKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// analogLatinAmericanSpanishName represents analog latin american spanish name.
	analogLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// analogRussianName represents analog russian name.
	analogRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// analogSimplifiedChineseName represents analog simplified chinese name.
	analogSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// analogSpanishName represents analog spanish name.
	analogSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// analogTraditionalChineseName represents analog traditional chinese name.
	analogTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// analogName represents analog name.
	analogName = nook.Languages{
		language.AmericanEnglish:      analogAmericanEnglishName,
		language.CanadianFrench:       analogCanadianFrenchName,
		language.Dutch:                analogDutchName,
		language.French:               analogFrenchName,
		language.German:               analogGermanName,
		language.Italian:              analogItalianName,
		language.Japanese:             analogJapaneseName,
		language.Korean:               analogKoreanName,
		language.LatinAmericanSpanish: analogLatinAmericanSpanishName,
		language.Russian:              analogRussianName,
		language.SimplifiedChinese:    analogSimplifiedChineseName,
		language.Spanish:              analogSpanishName,
		language.TraditionalChinese:   analogTraditionalChineseName}
)

var (
	// analogCharacter represents analog character.
	analogCharacter = nook.Character{
		Animal:   animal.Penguin,
		Birthday: analogBirthday,
		Code:     analogCode,
		Key:      character.Analog,
		Gender:   gender.Male,
		Name:     analogName,
		Special:  false}
)

var (
	// analogAmericanEnglishPhrase represents analog american english phrase.
	analogAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "だもんで"}

	// analogCanadianFrenchPhrase represents analog canadian french phrase.
	analogCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// analogDutchPhrase represents analog dutch phrase.
	analogDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// analogFrenchPhrase represents analog french phrase.
	analogFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// analogGermanPhrase represents analog german phrase.
	analogGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// analogItalianPhrase represents analog italian phrase.
	analogItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// analogJapanesePhrase represents analog japanese phrase.
	analogJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だもんで"}

	// analogKoreanPhrase represents analog korean phrase.
	analogKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// analogLatinAmericanSpanishPhrase represents analog latin american spanish phrase.
	analogLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// analogRussianPhrase represents analog russian phrase.
	analogRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// analogSimplifiedChinesePhrase represents analog simplified chinese phrase.
	analogSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// analogSpanishPhrase represents analog spanish phrase.
	analogSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// analogTraditionalChinesePhrase represents analog traditional chinese phrase.
	analogTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// analogPhrase represents analog phrase.
	analogPhrase = nook.Languages{
		language.AmericanEnglish:      analogAmericanEnglishPhrase,
		language.CanadianFrench:       analogCanadianFrenchPhrase,
		language.Dutch:                analogDutchPhrase,
		language.French:               analogFrenchPhrase,
		language.German:               analogGermanPhrase,
		language.Italian:              analogItalianPhrase,
		language.Japanese:             analogJapanesePhrase,
		language.Korean:               analogKoreanPhrase,
		language.LatinAmericanSpanish: analogLatinAmericanSpanishPhrase,
		language.Russian:              analogRussianPhrase,
		language.SimplifiedChinese:    analogSimplifiedChinesePhrase,
		language.Spanish:              analogSpanishPhrase,
		language.TraditionalChinese:   analogTraditionalChinesePhrase}
)

var (
	// Analog represents analog.
	Analog = nook.Villager{
		Character:   analogCharacter,
		Personality: personality.Jock,
		Phrase:      analogPhrase}
)
