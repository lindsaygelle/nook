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
	// knoxBirthday represents knox birthday.
	knoxBirthday = nook.Birthday{
		Day:   23,
		Month: time.November}
)

var (
	// knoxCode represents knox code.
	knoxCode = nook.Code{
		Value: "chn11"}
)

var (
	// knoxAmericanEnglishName represents knox american english name.
	knoxAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Knox"}

	// knoxCanadianFrenchName represents knox canadian french name.
	knoxCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Wolfram"}

	// knoxDutchName represents knox dutch name.
	knoxDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Knox"}

	// knoxFrenchName represents knox french name.
	knoxFrenchName = nook.Name{
		Language: language.French,
		Value:    "Wolfram"}

	// knoxGermanName represents knox german name.
	knoxGermanName = nook.Name{
		Language: language.German,
		Value:    "Tschiwi"}

	// knoxItalianName represents knox italian name.
	knoxItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Kalin"}

	// knoxJapaneseName represents knox japanese name.
	knoxJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キンカク"}

	// knoxKoreanName represents knox korean name.
	knoxKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "금끼오"}

	// knoxLatinAmericanSpanishName represents knox latin american spanish name.
	knoxLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tono"}

	// knoxRussianName represents knox russian name.
	knoxRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Нокс"}

	// knoxSimplifiedChineseName represents knox simplified chinese name.
	knoxSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "金阁"}

	// knoxSpanishName represents knox spanish name.
	knoxSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tono"}

	// knoxTraditionalChineseName represents knox traditional chinese name.
	knoxTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "金閣"}
)

var (
	// knoxName represents knox name.
	knoxName = nook.Languages{
		language.AmericanEnglish:      knoxAmericanEnglishName,
		language.CanadianFrench:       knoxCanadianFrenchName,
		language.Dutch:                knoxDutchName,
		language.French:               knoxFrenchName,
		language.German:               knoxGermanName,
		language.Italian:              knoxItalianName,
		language.Japanese:             knoxJapaneseName,
		language.Korean:               knoxKoreanName,
		language.LatinAmericanSpanish: knoxLatinAmericanSpanishName,
		language.Russian:              knoxRussianName,
		language.SimplifiedChinese:    knoxSimplifiedChineseName,
		language.Spanish:              knoxSpanishName,
		language.TraditionalChinese:   knoxTraditionalChineseName}
)

var (
	// knoxCharacter represents knox character.
	knoxCharacter = nook.Character{
		Animal:   animal.Chicken,
		Birthday: knoxBirthday,
		Code:     knoxCode,
		Key:      character.Knox,
		Gender:   gender.Male,
		Name:     knoxName,
		Special:  false}
)

var (
	// knoxAmericanEnglishPhrase represents knox american english phrase.
	knoxAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cluckling"}

	// knoxCanadianFrenchPhrase represents knox canadian french phrase.
	knoxCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ra-ta-tak"}

	// knoxDutchPhrase represents knox dutch phrase.
	knoxDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "nestei"}

	// knoxFrenchPhrase represents knox french phrase.
	knoxFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ra-ta-tak"}

	// knoxGermanPhrase represents knox german phrase.
	knoxGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knusprig"}

	// knoxItalianPhrase represents knox italian phrase.
	knoxItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "uovole"}

	// knoxJapanesePhrase represents knox japanese phrase.
	knoxJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "せいッ"}

	// knoxKoreanPhrase represents knox korean phrase.
	knoxKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "하앗"}

	// knoxLatinAmericanSpanishPhrase represents knox latin american spanish phrase.
	knoxLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cocorocó"}

	// knoxRussianPhrase represents knox russian phrase.
	knoxRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "короко"}

	// knoxSimplifiedChinesePhrase represents knox simplified chinese phrase.
	knoxSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "等待"}

	// knoxSpanishPhrase represents knox spanish phrase.
	knoxSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mazorca"}

	// knoxTraditionalChinesePhrase represents knox traditional chinese phrase.
	knoxTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "等待"}
)

var (
	// knoxPhrase represents knox phrase.
	knoxPhrase = nook.Languages{
		language.AmericanEnglish:      knoxAmericanEnglishPhrase,
		language.CanadianFrench:       knoxCanadianFrenchPhrase,
		language.Dutch:                knoxDutchPhrase,
		language.French:               knoxFrenchPhrase,
		language.German:               knoxGermanPhrase,
		language.Italian:              knoxItalianPhrase,
		language.Japanese:             knoxJapanesePhrase,
		language.Korean:               knoxKoreanPhrase,
		language.LatinAmericanSpanish: knoxLatinAmericanSpanishPhrase,
		language.Russian:              knoxRussianPhrase,
		language.SimplifiedChinese:    knoxSimplifiedChinesePhrase,
		language.Spanish:              knoxSpanishPhrase,
		language.TraditionalChinese:   knoxTraditionalChinesePhrase}
)

var (
	// Knox represents knox.
	Knox = nook.Villager{
		Character:   knoxCharacter,
		Personality: personality.Cranky,
		Phrase:      knoxPhrase}
)
