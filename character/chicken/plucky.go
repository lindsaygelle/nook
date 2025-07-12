package chicken

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
	// pluckyBirthday represents plucky birthday.
	pluckyBirthday = nook.Birthday{
		Day:   12,
		Month: time.October}
)

var (
	// pluckyCode represents plucky code.
	pluckyCode = nook.Code{
		Value: "chn10"}
)

var (
	// pluckyAmericanEnglishName represents plucky american english name.
	pluckyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Plucky"}

	// pluckyCanadianFrenchName represents plucky canadian french name.
	pluckyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Poulette"}

	// pluckyDutchName represents plucky dutch name.
	pluckyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Plucky"}

	// pluckyFrenchName represents plucky french name.
	pluckyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Poulette"}

	// pluckyGermanName represents plucky german name.
	pluckyGermanName = nook.Name{
		Language: language.German,
		Value:    "Mareile"}

	// pluckyItalianName represents plucky italian name.
	pluckyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cocca"}

	// pluckyJapaneseName represents plucky japanese name.
	pluckyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パタヤ"}

	// pluckyKoreanName represents plucky korean name.
	pluckyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "파타야"}

	// pluckyLatinAmericanSpanishName represents plucky latin american spanish name.
	pluckyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Herminia"}

	// pluckyRussianName represents plucky russian name.
	pluckyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Плаки"}

	// pluckySimplifiedChineseName represents plucky simplified chinese name.
	pluckySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "帕塔雅"}

	// pluckySpanishName represents plucky spanish name.
	pluckySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Herminia"}

	// pluckyTraditionalChineseName represents plucky traditional chinese name.
	pluckyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "帕塔雅"}
)

var (
	// pluckyName represents plucky name.
	pluckyName = nook.Languages{
		language.AmericanEnglish:      pluckyAmericanEnglishName,
		language.CanadianFrench:       pluckyCanadianFrenchName,
		language.Dutch:                pluckyDutchName,
		language.French:               pluckyFrenchName,
		language.German:               pluckyGermanName,
		language.Italian:              pluckyItalianName,
		language.Japanese:             pluckyJapaneseName,
		language.Korean:               pluckyKoreanName,
		language.LatinAmericanSpanish: pluckyLatinAmericanSpanishName,
		language.Russian:              pluckyRussianName,
		language.SimplifiedChinese:    pluckySimplifiedChineseName,
		language.Spanish:              pluckySpanishName,
		language.TraditionalChinese:   pluckyTraditionalChineseName}
)

var (
	// pluckyCharacter represents plucky character.
	pluckyCharacter = nook.Character{
		Animal:   animal.Chicken,
		Birthday: pluckyBirthday,
		Code:     pluckyCode,
		Key:      character.Plucky,
		Gender:   gender.Female,
		Name:     pluckyName,
		Special:  false}
)

var (
	// pluckyAmericanEnglishPhrase represents plucky american english phrase.
	pluckyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chicky-poo"}

	// pluckyCanadianFrenchPhrase represents plucky canadian french phrase.
	pluckyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "galine"}

	// pluckyDutchPhrase represents plucky dutch phrase.
	pluckyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kippetje"}

	// pluckyFrenchPhrase represents plucky french phrase.
	pluckyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "galine"}

	// pluckyGermanPhrase represents plucky german phrase.
	pluckyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hühnermist"}

	// pluckyItalianPhrase represents plucky italian phrase.
	pluckyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "zampoli"}

	// pluckyJapanesePhrase represents plucky japanese phrase.
	pluckyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "どうだい"}

	// pluckyKoreanPhrase represents plucky korean phrase.
	pluckyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그게어때"}

	// pluckyLatinAmericanSpanishPhrase represents plucky latin american spanish phrase.
	pluckyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "corocó"}

	// pluckyRussianPhrase represents plucky russian phrase.
	pluckyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "щип-щип"}

	// pluckySimplifiedChinesePhrase represents plucky simplified chinese phrase.
	pluckySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "泰泰"}

	// pluckySpanishPhrase represents plucky spanish phrase.
	pluckySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tesorito"}

	// pluckyTraditionalChinesePhrase represents plucky traditional chinese phrase.
	pluckyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "泰泰"}
)

var (
	// pluckyPhrase represents plucky phrase.
	pluckyPhrase = nook.Languages{
		language.AmericanEnglish:      pluckyAmericanEnglishPhrase,
		language.CanadianFrench:       pluckyCanadianFrenchPhrase,
		language.Dutch:                pluckyDutchPhrase,
		language.French:               pluckyFrenchPhrase,
		language.German:               pluckyGermanPhrase,
		language.Italian:              pluckyItalianPhrase,
		language.Japanese:             pluckyJapanesePhrase,
		language.Korean:               pluckyKoreanPhrase,
		language.LatinAmericanSpanish: pluckyLatinAmericanSpanishPhrase,
		language.Russian:              pluckyRussianPhrase,
		language.SimplifiedChinese:    pluckySimplifiedChinesePhrase,
		language.Spanish:              pluckySpanishPhrase,
		language.TraditionalChinese:   pluckyTraditionalChinesePhrase}
)

var (
	// Plucky represents plucky.
	Plucky = nook.Villager{
		Character:   pluckyCharacter,
		Personality: personality.BigSister,
		Phrase:      pluckyPhrase}
)
