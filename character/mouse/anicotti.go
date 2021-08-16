package mouse

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
	anicottiBirthday = nook.Birthday{
		Day:   24,
		Month: time.February}
)

var (
	anicottiCode = nook.Code{
		Value: "mus10"}
)

var (
	anicottiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Anicotti"}

	anicottiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Annie"}

	anicottiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Anicotti"}

	anicottiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Annie"}

	anicottiGermanName = nook.Name{
		Language: language.German,
		Value:    "Eva"}

	anicottiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Squitta"}

	anicottiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ラザニア"}

	anicottiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "라자냐"}

	anicottiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Clorinda"}

	anicottiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Аникотти"}

	anicottiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗萱儿"}

	anicottiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Clorinda"}

	anicottiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羅萱兒"}
)

var (
	anicottiName = nook.Languages{
		language.AmericanEnglish:      anicottiAmericanEnglishName,
		language.CanadianFrench:       anicottiCanadianFrenchName,
		language.Dutch:                anicottiDutchName,
		language.French:               anicottiFrenchName,
		language.German:               anicottiGermanName,
		language.Italian:              anicottiItalianName,
		language.Japanese:             anicottiJapaneseName,
		language.Korean:               anicottiKoreanName,
		language.LatinAmericanSpanish: anicottiLatinAmericanSpanishName,
		language.Russian:              anicottiRussianName,
		language.SimplifiedChinese:    anicottiSimplifiedChineseName,
		language.Spanish:              anicottiSpanishName,
		language.TraditionalChinese:   anicottiTraditionalChineseName}
)

var (
	anicottiCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: anicottiBirthday,
		Code:     anicottiCode,
		Key:      character.Anicotti,
		Gender:   gender.Female,
		Name:     anicottiName}
)

var (
	anicottiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cannoli"}

	anicottiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "dis donc"}

	anicottiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "cannoli"}

	anicottiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "dis donc"}

	anicottiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "piepser"}

	anicottiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cippoli"}

	anicottiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ルンルン"}

	anicottiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "룰룰"}

	anicottiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cloricló"}

	anicottiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "канноли"}

	anicottiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "开心"}

	anicottiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cloricló"}

	anicottiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "開心"}
)

var (
	anicottiPhrase = nook.Languages{
		language.AmericanEnglish:      anicottiAmericanEnglishName,
		language.CanadianFrench:       anicottiCanadianFrenchName,
		language.Dutch:                anicottiDutchName,
		language.French:               anicottiFrenchName,
		language.German:               anicottiGermanName,
		language.Italian:              anicottiItalianName,
		language.Japanese:             anicottiJapaneseName,
		language.Korean:               anicottiKoreanName,
		language.LatinAmericanSpanish: anicottiLatinAmericanSpanishName,
		language.Russian:              anicottiRussianName,
		language.SimplifiedChinese:    anicottiSimplifiedChineseName,
		language.Spanish:              anicottiSpanishName,
		language.TraditionalChinese:   anicottiTraditionalChineseName}
)

var (
	Anicotti = nook.Villager{
		Character:   anicottiCharacter,
		Personality: personality.Peppy,
		Phrase:      anicottiPhrase}
)
