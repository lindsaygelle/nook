package horse

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
	// buckBirthday represents buck birthday.
	buckBirthday = nook.Birthday{
		Day:   4,
		Month: time.April}
)

var (
	// buckCode represents buck code.
	buckCode = nook.Code{
		Value: "hrs00"}
)

var (
	// buckAmericanEnglishName represents buck american english name.
	buckAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Buck"}

	// buckCanadianFrenchName represents buck canadian french name.
	buckCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Daniel"}

	// buckDutchName represents buck dutch name.
	buckDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Buck"}

	// buckFrenchName represents buck french name.
	buckFrenchName = nook.Name{
		Language: language.French,
		Value:    "Daniel"}

	// buckGermanName represents buck german name.
	buckGermanName = nook.Name{
		Language: language.German,
		Value:    "Rudi"}

	// buckItalianName represents buck italian name.
	buckItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Trottolo"}

	// buckJapaneseName represents buck japanese name.
	buckJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヴァヤシコフ"}

	// buckKoreanName represents buck korean name.
	buckKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "바야시코프"}

	// buckLatinAmericanSpanishName represents buck latin american spanish name.
	buckLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Trotón"}

	// buckRussianName represents buck russian name.
	buckRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бак"}

	// buckSimplifiedChineseName represents buck simplified chinese name.
	buckSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "威亚"}

	// buckSpanishName represents buck spanish name.
	buckSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Trotón"}

	// buckTraditionalChineseName represents buck traditional chinese name.
	buckTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "威亞"}
)

var (
	// buckName represents buck name.
	buckName = nook.Languages{
		language.AmericanEnglish:      buckAmericanEnglishName,
		language.CanadianFrench:       buckCanadianFrenchName,
		language.Dutch:                buckDutchName,
		language.French:               buckFrenchName,
		language.German:               buckGermanName,
		language.Italian:              buckItalianName,
		language.Japanese:             buckJapaneseName,
		language.Korean:               buckKoreanName,
		language.LatinAmericanSpanish: buckLatinAmericanSpanishName,
		language.Russian:              buckRussianName,
		language.SimplifiedChinese:    buckSimplifiedChineseName,
		language.Spanish:              buckSpanishName,
		language.TraditionalChinese:   buckTraditionalChineseName}
)

var (
	// buckCharacter represents buck character.
	buckCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: buckBirthday,
		Code:     buckCode,
		Key:      character.Buck,
		Gender:   gender.Male,
		Name:     buckName,
		Special:  false}
)

var (
	// buckAmericanEnglishPhrase represents buck american english phrase.
	buckAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pardner"}

	// buckCanadianFrenchPhrase represents buck canadian french phrase.
	buckCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "howdy"}

	// buckDutchPhrase represents buck dutch phrase.
	buckDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "draver"}

	// buckFrenchPhrase represents buck french phrase.
	buckFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "taty"}

	// buckGermanPhrase represents buck german phrase.
	buckGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "wiehie"}

	// buckItalianPhrase represents buck italian phrase.
	buckItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cowboy"}

	// buckJapanesePhrase represents buck japanese phrase.
	buckJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なんちて"}

	// buckKoreanPhrase represents buck korean phrase.
	buckKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "농담"}

	// buckLatinAmericanSpanishPhrase represents buck latin american spanish phrase.
	buckLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "jopop"}

	// buckRussianPhrase represents buck russian phrase.
	buckRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "партнер"}

	// buckSimplifiedChinesePhrase represents buck simplified chinese phrase.
	buckSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "搞笑"}

	// buckSpanishPhrase represents buck spanish phrase.
	buckSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "zanahoria"}

	// buckTraditionalChinesePhrase represents buck traditional chinese phrase.
	buckTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "搞笑"}
)

var (
	// buckPhrase represents buck phrase.
	buckPhrase = nook.Languages{
		language.AmericanEnglish:      buckAmericanEnglishPhrase,
		language.CanadianFrench:       buckCanadianFrenchPhrase,
		language.Dutch:                buckDutchPhrase,
		language.French:               buckFrenchPhrase,
		language.German:               buckGermanPhrase,
		language.Italian:              buckItalianPhrase,
		language.Japanese:             buckJapanesePhrase,
		language.Korean:               buckKoreanPhrase,
		language.LatinAmericanSpanish: buckLatinAmericanSpanishPhrase,
		language.Russian:              buckRussianPhrase,
		language.SimplifiedChinese:    buckSimplifiedChinesePhrase,
		language.Spanish:              buckSpanishPhrase,
		language.TraditionalChinese:   buckTraditionalChinesePhrase}
)

var (
	// Buck represents buck.
	Buck = nook.Villager{
		Character:   buckCharacter,
		Personality: personality.Jock,
		Phrase:      buckPhrase}
)
