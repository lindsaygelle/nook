package frog

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
	// jeremiahBirthday represents jeremiah birthday.
	jeremiahBirthday = nook.Birthday{
		Day:   8,
		Month: time.July}
)

var (
	// jeremiahCode represents jeremiah code.
	jeremiahCode = nook.Code{
		Value: "flg07"}
)

var (
	// jeremiahAmericanEnglishName represents jeremiah american english name.
	jeremiahAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Jeremiah"}

	// jeremiahCanadianFrenchName represents jeremiah canadian french name.
	jeremiahCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jérémie"}

	// jeremiahDutchName represents jeremiah dutch name.
	jeremiahDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Jeremiah"}

	// jeremiahFrenchName represents jeremiah french name.
	jeremiahFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jérémie"}

	// jeremiahGermanName represents jeremiah german name.
	jeremiahGermanName = nook.Name{
		Language: language.German,
		Value:    "Jörg"}

	// jeremiahItalianName represents jeremiah italian name.
	jeremiahItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Geremia"}

	// jeremiahJapaneseName represents jeremiah japanese name.
	jeremiahJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クワトロ"}

	// jeremiahKoreanName represents jeremiah korean name.
	jeremiahKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "드리미"}

	// jeremiahLatinAmericanSpanishName represents jeremiah latin american spanish name.
	jeremiahLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jeremías"}

	// jeremiahRussianName represents jeremiah russian name.
	jeremiahRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джеремия"}

	// jeremiahSimplifiedChineseName represents jeremiah simplified chinese name.
	jeremiahSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿刻"}

	// jeremiahSpanishName represents jeremiah spanish name.
	jeremiahSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jeremías"}

	// jeremiahTraditionalChineseName represents jeremiah traditional chinese name.
	jeremiahTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿刻"}
)

var (
	// jeremiahName represents jeremiah name.
	jeremiahName = nook.Languages{
		language.AmericanEnglish:      jeremiahAmericanEnglishName,
		language.CanadianFrench:       jeremiahCanadianFrenchName,
		language.Dutch:                jeremiahDutchName,
		language.French:               jeremiahFrenchName,
		language.German:               jeremiahGermanName,
		language.Italian:              jeremiahItalianName,
		language.Japanese:             jeremiahJapaneseName,
		language.Korean:               jeremiahKoreanName,
		language.LatinAmericanSpanish: jeremiahLatinAmericanSpanishName,
		language.Russian:              jeremiahRussianName,
		language.SimplifiedChinese:    jeremiahSimplifiedChineseName,
		language.Spanish:              jeremiahSpanishName,
		language.TraditionalChinese:   jeremiahTraditionalChineseName}
)

var (
	// jeremiahCharacter represents jeremiah character.
	jeremiahCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: jeremiahBirthday,
		Code:     jeremiahCode,
		Key:      character.Jeremiah,
		Gender:   gender.Male,
		Name:     jeremiahName,
		Special:  false}
)

var (
	// jeremiahAmericanEnglishPhrase represents jeremiah american english phrase.
	jeremiahAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nee-deep"}

	// jeremiahCanadianFrenchPhrase represents jeremiah canadian french phrase.
	jeremiahCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "brooap"}

	// jeremiahDutchPhrase represents jeremiah dutch phrase.
	jeremiahDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "rikkekik"}

	// jeremiahFrenchPhrase represents jeremiah french phrase.
	jeremiahFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "brooap"}

	// jeremiahGermanPhrase represents jeremiah german phrase.
	jeremiahGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "nidip"}

	// jeremiahItalianPhrase represents jeremiah italian phrase.
	jeremiahItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "fri fri"}

	// jeremiahJapanesePhrase represents jeremiah japanese phrase.
	jeremiahJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "にゃむ"}

	// jeremiahKoreanPhrase represents jeremiah korean phrase.
	jeremiahKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "딩동댕"}

	// jeremiahLatinAmericanSpanishPhrase represents jeremiah latin american spanish phrase.
	jeremiahLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "croc-croc"}

	// jeremiahRussianPhrase represents jeremiah russian phrase.
	jeremiahRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "буль-буль"}

	// jeremiahSimplifiedChinesePhrase represents jeremiah simplified chinese phrase.
	jeremiahSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嚼嚼"}

	// jeremiahSpanishPhrase represents jeremiah spanish phrase.
	jeremiahSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "croc-croc"}

	// jeremiahTraditionalChinesePhrase represents jeremiah traditional chinese phrase.
	jeremiahTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嚼嚼"}
)

var (
	// jeremiahPhrase represents jeremiah phrase.
	jeremiahPhrase = nook.Languages{
		language.AmericanEnglish:      jeremiahAmericanEnglishPhrase,
		language.CanadianFrench:       jeremiahCanadianFrenchPhrase,
		language.Dutch:                jeremiahDutchPhrase,
		language.French:               jeremiahFrenchPhrase,
		language.German:               jeremiahGermanPhrase,
		language.Italian:              jeremiahItalianPhrase,
		language.Japanese:             jeremiahJapanesePhrase,
		language.Korean:               jeremiahKoreanPhrase,
		language.LatinAmericanSpanish: jeremiahLatinAmericanSpanishPhrase,
		language.Russian:              jeremiahRussianPhrase,
		language.SimplifiedChinese:    jeremiahSimplifiedChinesePhrase,
		language.Spanish:              jeremiahSpanishPhrase,
		language.TraditionalChinese:   jeremiahTraditionalChinesePhrase}
)

var (
	// Jeremiah represents jeremiah.
	Jeremiah = nook.Villager{
		Character:   jeremiahCharacter,
		Personality: personality.Lazy,
		Phrase:      jeremiahPhrase}
)
