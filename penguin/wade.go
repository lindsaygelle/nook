package penguin

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Migloublanbec"}

	wadeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Wadezo is 't"}

	wadeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Migloublanbec"}

	wadeGermanName = nook.Name{
		Language: language.German,
		Value:    "Staksiwackel"}

	wadeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pliniouhiii uuuh"}

	wadeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カマボコだからね"}

	wadeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "호떡뒤집어"}

	wadeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Petínespoink"}

	wadeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Уэйдвот так вот"}

	wadeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "竹轮所以啦"}

	wadeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Petínpersonaje"}

	wadeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "竹輪所以啦"}
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
		Animal:   Penguin,
		Birthday: wadeBirthday,
		Code:     wadeCode,
		Gender:   nook.Male,
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
		Personality: nook.Lazy,
		Phrase:      wadePhrase}
)
