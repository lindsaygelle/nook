package gorilla

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
	// alBirthday represents al birthday.
	alBirthday = nook.Birthday{
		Day:   18,
		Month: time.October}
)

var (
	// alCode represents al code.
	alCode = nook.Code{
		Value: "gor08"}
)

var (
	// alAmericanEnglishName represents al american english name.
	alAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Al"}

	// alCanadianFrenchName represents al canadian french name.
	alCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gustave"}

	// alDutchName represents al dutch name.
	alDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Al"}

	// alFrenchName represents al french name.
	alFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gustave"}

	// alGermanName represents al german name.
	alGermanName = nook.Name{
		Language: language.German,
		Value:    "Kokong"}

	// alItalianName represents al italian name.
	alItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gregorio"}

	// alJapaneseName represents al japanese name.
	alJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "たもつ"}

	// alKoreanName represents al korean name.
	alKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "우락"}

	// alLatinAmericanSpanishName represents al latin american spanish name.
	alLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Álex"}

	// alRussianName represents al russian name.
	alRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Эл"}

	// alSimplifiedChineseName represents al simplified chinese name.
	alSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿保"}

	// alSpanishName represents al spanish name.
	alSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Álex"}

	// alTraditionalChineseName represents al traditional chinese name.
	alTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿保"}
)

var (
	// alName represents al name.
	alName = nook.Languages{
		language.AmericanEnglish:      alAmericanEnglishName,
		language.CanadianFrench:       alCanadianFrenchName,
		language.Dutch:                alDutchName,
		language.French:               alFrenchName,
		language.German:               alGermanName,
		language.Italian:              alItalianName,
		language.Japanese:             alJapaneseName,
		language.Korean:               alKoreanName,
		language.LatinAmericanSpanish: alLatinAmericanSpanishName,
		language.Russian:              alRussianName,
		language.SimplifiedChinese:    alSimplifiedChineseName,
		language.Spanish:              alSpanishName,
		language.TraditionalChinese:   alTraditionalChineseName}
)

var (
	// alCharacter represents al character.
	alCharacter = nook.Character{
		Animal:   animal.Gorilla,
		Birthday: alBirthday,
		Code:     alCode,
		Key:      character.Al,
		Gender:   gender.Male,
		Name:     alName,
		Special:  false}
)

var (
	// alAmericanEnglishPhrase represents al american english phrase.
	alAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hoo hoo ha"}

	// alCanadianFrenchPhrase represents al canadian french phrase.
	alCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ouba-oub"}

	// alDutchPhrase represents al dutch phrase.
	alDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "aiaiai"}

	// alFrenchPhrase represents al french phrase.
	alFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ouba-oub"}

	// alGermanPhrase represents al german phrase.
	alGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "uga-uga"}

	// alItalianPhrase represents al italian phrase.
	alItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bananao"}

	// alJapanesePhrase represents al japanese phrase.
	alJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ウヒョッ"}

	// alKoreanPhrase represents al korean phrase.
	alKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "부락"}

	// alLatinAmericanSpanishPhrase represents al latin american spanish phrase.
	alLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "juju-já"}

	// alRussianPhrase represents al russian phrase.
	alRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ух-ух-ух"}

	// alSimplifiedChinesePhrase represents al simplified chinese phrase.
	alSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呜呼"}

	// alSpanishPhrase represents al spanish phrase.
	alSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chocolate"}

	// alTraditionalChinesePhrase represents al traditional chinese phrase.
	alTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗚呼"}
)

var (
	// alPhrase represents al phrase.
	alPhrase = nook.Languages{
		language.AmericanEnglish:      alAmericanEnglishPhrase,
		language.CanadianFrench:       alCanadianFrenchPhrase,
		language.Dutch:                alDutchPhrase,
		language.French:               alFrenchPhrase,
		language.German:               alGermanPhrase,
		language.Italian:              alItalianPhrase,
		language.Japanese:             alJapanesePhrase,
		language.Korean:               alKoreanPhrase,
		language.LatinAmericanSpanish: alLatinAmericanSpanishPhrase,
		language.Russian:              alRussianPhrase,
		language.SimplifiedChinese:    alSimplifiedChinesePhrase,
		language.Spanish:              alSpanishPhrase,
		language.TraditionalChinese:   alTraditionalChinesePhrase}
)

var (
	// Al represents al.
	Al = nook.Villager{
		Character:   alCharacter,
		Personality: personality.Lazy,
		Phrase:      alPhrase}
)
