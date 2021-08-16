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
	wadeBirthday = nook.Birthday{
		Day:   30,
		Month: time.October}
)

var (
	wadeCode = nook.Code{
		Value: "pgn09"}
)

var (
	wadeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Wade"}

	wadeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Miglou"}

	wadeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Wade"}

	wadeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Miglou"}

	wadeGermanName = nook.Name{
		Language: language.German,
		Value:    "Staksi"}

	wadeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Plinio"}

	wadeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カマボコ"}

	wadeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "호떡"}

	wadeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Petín"}

	wadeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Уэйд"}

	wadeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "竹轮"}

	wadeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Petín"}

	wadeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "竹輪"}
)

var (
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
	wadeCharacter = nook.Character{
		Animal:   animal.Penguin,
		Birthday: wadeBirthday,
		Code:     wadeCode,
		Key:      character.Wade,
		Gender:   gender.Male,
		Name:     wadeName}
)

var (
	wadeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "so it goes"}

	wadeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "blanbec"}

	wadeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zo is 't"}

	wadeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "blanbec"}

	wadeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "wackel"}

	wadeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "uhiii uuuh"}

	wadeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だからね"}

	wadeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "뒤집어"}

	wadeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "espoink"}

	wadeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "вот так вот"}

	wadeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "所以啦"}

	wadeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "personaje"}

	wadeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "所以啦"}
)

var (
	wadePhrase = nook.Languages{
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
	Wade = nook.Villager{
		Character:   wadeCharacter,
		Personality: personality.Lazy,
		Phrase:      wadePhrase}
)
