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
	// wadeBirthday represents wade birthday.
	wadeBirthday = nook.Birthday{
		Day:   30,
		Month: time.October}
)

var (
	// wadeCode represents wade code.
	wadeCode = nook.Code{
		Value: "pgn09"}
)

var (
	// wadeAmericanEnglishName represents wade american english name.
	wadeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Wade"}

	// wadeCanadianFrenchName represents wade canadian french name.
	wadeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Miglou"}

	// wadeDutchName represents wade dutch name.
	wadeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Wade"}

	// wadeFrenchName represents wade french name.
	wadeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Miglou"}

	// wadeGermanName represents wade german name.
	wadeGermanName = nook.Name{
		Language: language.German,
		Value:    "Staksi"}

	// wadeItalianName represents wade italian name.
	wadeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Plinio"}

	// wadeJapaneseName represents wade japanese name.
	wadeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カマボコ"}

	// wadeKoreanName represents wade korean name.
	wadeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "호떡"}

	// wadeLatinAmericanSpanishName represents wade latin american spanish name.
	wadeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Petín"}

	// wadeRussianName represents wade russian name.
	wadeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Уэйд"}

	// wadeSimplifiedChineseName represents wade simplified chinese name.
	wadeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "竹轮"}

	// wadeSpanishName represents wade spanish name.
	wadeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Petín"}

	// wadeTraditionalChineseName represents wade traditional chinese name.
	wadeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "竹輪"}
)

var (
	// wadeName represents wade name.
	wadeName = nook.Languages{
		language.AmericanEnglish:      wadeAmericanEnglishName,
		language.CanadianFrench:       wadeCanadianFrenchName,
		language.Dutch:                wadeDutchName,
		language.French:               wadeFrenchName,
		language.German:               wadeGermanName,
		language.Italian:              wadeItalianName,
		language.Japanese:             wadeJapaneseName,
		language.Korean:               wadeKoreanName,
		language.LatinAmericanSpanish: wadeLatinAmericanSpanishName,
		language.Russian:              wadeRussianName,
		language.SimplifiedChinese:    wadeSimplifiedChineseName,
		language.Spanish:              wadeSpanishName,
		language.TraditionalChinese:   wadeTraditionalChineseName}
)

var (
	// wadeCharacter represents wade character.
	wadeCharacter = nook.Character{
		Animal:   animal.Penguin,
		Birthday: wadeBirthday,
		Code:     wadeCode,
		Key:      character.Wade,
		Gender:   gender.Male,
		Name:     wadeName,
		Special:  false}
)

var (
	// wadeAmericanEnglishPhrase represents wade american english phrase.
	wadeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "so it goes"}

	// wadeCanadianFrenchPhrase represents wade canadian french phrase.
	wadeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "blanbec"}

	// wadeDutchPhrase represents wade dutch phrase.
	wadeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zo is 't"}

	// wadeFrenchPhrase represents wade french phrase.
	wadeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "blanbec"}

	// wadeGermanPhrase represents wade german phrase.
	wadeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "wackel"}

	// wadeItalianPhrase represents wade italian phrase.
	wadeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "uhiii uuuh"}

	// wadeJapanesePhrase represents wade japanese phrase.
	wadeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だからね"}

	// wadeKoreanPhrase represents wade korean phrase.
	wadeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "뒤집어"}

	// wadeLatinAmericanSpanishPhrase represents wade latin american spanish phrase.
	wadeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "espoink"}

	// wadeRussianPhrase represents wade russian phrase.
	wadeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "вот так вот"}

	// wadeSimplifiedChinesePhrase represents wade simplified chinese phrase.
	wadeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "所以啦"}

	// wadeSpanishPhrase represents wade spanish phrase.
	wadeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "personaje"}

	// wadeTraditionalChinesePhrase represents wade traditional chinese phrase.
	wadeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "所以啦"}
)

var (
	// wadePhrase represents wade phrase.
	wadePhrase = nook.Languages{
		language.AmericanEnglish:      wadeAmericanEnglishPhrase,
		language.CanadianFrench:       wadeCanadianFrenchPhrase,
		language.Dutch:                wadeDutchPhrase,
		language.French:               wadeFrenchPhrase,
		language.German:               wadeGermanPhrase,
		language.Italian:              wadeItalianPhrase,
		language.Japanese:             wadeJapanesePhrase,
		language.Korean:               wadeKoreanPhrase,
		language.LatinAmericanSpanish: wadeLatinAmericanSpanishPhrase,
		language.Russian:              wadeRussianPhrase,
		language.SimplifiedChinese:    wadeSimplifiedChinesePhrase,
		language.Spanish:              wadeSpanishPhrase,
		language.TraditionalChinese:   wadeTraditionalChinesePhrase}
)

var (
	// Wade represents wade.
	Wade = nook.Villager{
		Character:   wadeCharacter,
		Personality: personality.Lazy,
		Phrase:      wadePhrase}
)
