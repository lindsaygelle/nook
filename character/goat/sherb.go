package goat

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
	// sherbBirthday represents sherb birthday.
	sherbBirthday = nook.Birthday{
		Day:   18,
		Month: time.January}
)

var (
	// sherbCode represents sherb code.
	sherbCode = nook.Code{
		Value: "goa09"}
)

var (
	// sherbAmericanEnglishName represents sherb american english name.
	sherbAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sherb"}

	// sherbCanadianFrenchName represents sherb canadian french name.
	sherbCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Capri"}

	// sherbDutchName represents sherb dutch name.
	sherbDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sherb"}

	// sherbFrenchName represents sherb french name.
	sherbFrenchName = nook.Name{
		Language: language.French,
		Value:    "Capri"}

	// sherbGermanName represents sherb german name.
	sherbGermanName = nook.Name{
		Language: language.German,
		Value:    "Morpheus"}

	// sherbItalianName represents sherb italian name.
	sherbItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Capraldo"}

	// sherbJapaneseName represents sherb japanese name.
	sherbJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "レム"}

	// sherbKoreanName represents sherb korean name.
	sherbKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "래미"}

	// sherbLatinAmericanSpanishName represents sherb latin american spanish name.
	sherbLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Morfeo"}

	// sherbRussianName represents sherb russian name.
	sherbRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Шерб"}

	// sherbSimplifiedChineseName represents sherb simplified chinese name.
	sherbSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雷姆"}

	// sherbSpanishName represents sherb spanish name.
	sherbSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Morfeo"}

	// sherbTraditionalChineseName represents sherb traditional chinese name.
	sherbTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雷姆"}
)

var (
	// sherbName represents sherb name.
	sherbName = nook.Languages{
		language.AmericanEnglish:      sherbAmericanEnglishName,
		language.CanadianFrench:       sherbCanadianFrenchName,
		language.Dutch:                sherbDutchName,
		language.French:               sherbFrenchName,
		language.German:               sherbGermanName,
		language.Italian:              sherbItalianName,
		language.Japanese:             sherbJapaneseName,
		language.Korean:               sherbKoreanName,
		language.LatinAmericanSpanish: sherbLatinAmericanSpanishName,
		language.Russian:              sherbRussianName,
		language.SimplifiedChinese:    sherbSimplifiedChineseName,
		language.Spanish:              sherbSpanishName,
		language.TraditionalChinese:   sherbTraditionalChineseName}
)

var (
	// sherbCharacter represents sherb character.
	sherbCharacter = nook.Character{
		Animal:   animal.Goat,
		Birthday: sherbBirthday,
		Code:     sherbCode,
		Key:      character.Sherb,
		Gender:   gender.Male,
		Name:     sherbName,
		Special:  false}
)

var (
	// sherbAmericanEnglishPhrase represents sherb american english phrase.
	sherbAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bawwww"}

	// sherbCanadianFrenchPhrase represents sherb canadian french phrase.
	sherbCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bââââille"}

	// sherbDutchPhrase represents sherb dutch phrase.
	sherbDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "oefff"}

	// sherbFrenchPhrase represents sherb french phrase.
	sherbFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bââââille"}

	// sherbGermanPhrase represents sherb german phrase.
	sherbGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schlummer"}

	// sherbItalianPhrase represents sherb italian phrase.
	sherbItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "beeewn"}

	// sherbJapanesePhrase represents sherb japanese phrase.
	sherbJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ふわぁ"}

	// sherbKoreanPhrase represents sherb korean phrase.
	sherbKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "흐아앙"}

	// sherbLatinAmericanSpanishPhrase represents sherb latin american spanish phrase.
	sherbLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bostezzz"}

	// sherbRussianPhrase represents sherb russian phrase.
	sherbRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мемеша"}

	// sherbSimplifiedChinesePhrase represents sherb simplified chinese phrase.
	sherbSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "轻飘"}

	// sherbSpanishPhrase represents sherb spanish phrase.
	sherbSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bostezzz"}

	// sherbTraditionalChinesePhrase represents sherb traditional chinese phrase.
	sherbTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "輕飄"}
)

var (
	// sherbPhrase represents sherb phrase.
	sherbPhrase = nook.Languages{
		language.AmericanEnglish:      sherbAmericanEnglishPhrase,
		language.CanadianFrench:       sherbCanadianFrenchPhrase,
		language.Dutch:                sherbDutchPhrase,
		language.French:               sherbFrenchPhrase,
		language.German:               sherbGermanPhrase,
		language.Italian:              sherbItalianPhrase,
		language.Japanese:             sherbJapanesePhrase,
		language.Korean:               sherbKoreanPhrase,
		language.LatinAmericanSpanish: sherbLatinAmericanSpanishPhrase,
		language.Russian:              sherbRussianPhrase,
		language.SimplifiedChinese:    sherbSimplifiedChinesePhrase,
		language.Spanish:              sherbSpanishPhrase,
		language.TraditionalChinese:   sherbTraditionalChinesePhrase}
)

var (
	// Sherb represents sherb.
	Sherb = nook.Villager{
		Character:   sherbCharacter,
		Personality: personality.Lazy,
		Phrase:      sherbPhrase}
)
