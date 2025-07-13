package duck

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
	// joeyBirthday represents joey birthday.
	joeyBirthday = nook.Birthday{
		Day:   3,
		Month: time.January}
)

var (
	// joeyCode represents joey code.
	joeyCode = nook.Code{
		Value: "duk01"}
)

var (
	// joeyAmericanEnglishName represents joey american english name.
	joeyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Joey"}

	// joeyCanadianFrenchName represents joey canadian french name.
	joeyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Joseph"}

	// joeyDutchName represents joey dutch name.
	joeyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Joey"}

	// joeyFrenchName represents joey french name.
	joeyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Joseph"}

	// joeyGermanName represents joey german name.
	joeyGermanName = nook.Name{
		Language: language.German,
		Value:    "Kalle"}

	// joeyItalianName represents joey italian name.
	joeyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pino"}

	// joeyJapaneseName represents joey japanese name.
	joeyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "リチャード"}

	// joeyKoreanName represents joey korean name.
	joeyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "리처드"}

	// joeyLatinAmericanSpanishName represents joey latin american spanish name.
	joeyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pascual"}

	// joeyRussianName represents joey russian name.
	joeyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джоуи"}

	// joeySimplifiedChineseName represents joey simplified chinese name.
	joeySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "李察"}

	// joeySpanishName represents joey spanish name.
	joeySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pascual"}

	// joeyTraditionalChineseName represents joey traditional chinese name.
	joeyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "李察"}
)

var (
	// joeyName represents joey name.
	joeyName = nook.Languages{
		language.AmericanEnglish:      joeyAmericanEnglishName,
		language.CanadianFrench:       joeyCanadianFrenchName,
		language.Dutch:                joeyDutchName,
		language.French:               joeyFrenchName,
		language.German:               joeyGermanName,
		language.Italian:              joeyItalianName,
		language.Japanese:             joeyJapaneseName,
		language.Korean:               joeyKoreanName,
		language.LatinAmericanSpanish: joeyLatinAmericanSpanishName,
		language.Russian:              joeyRussianName,
		language.SimplifiedChinese:    joeySimplifiedChineseName,
		language.Spanish:              joeySpanishName,
		language.TraditionalChinese:   joeyTraditionalChineseName}
)

var (
	// joeyCharacter represents joey character.
	joeyCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: joeyBirthday,
		Code:     joeyCode,
		Key:      character.Joey,
		Gender:   gender.Male,
		Name:     joeyName,
		Special:  false}
)

var (
	// joeyAmericanEnglishPhrase represents joey american english phrase.
	joeyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bleeeeeck"}

	// joeyCanadianFrenchPhrase represents joey canadian french phrase.
	joeyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "beeeuuuurk"}

	// joeyDutchPhrase represents joey dutch phrase.
	joeyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kèèèk"}

	// joeyFrenchPhrase represents joey french phrase.
	joeyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "beeeuuuurk"}

	// joeyGermanPhrase represents joey german phrase.
	joeyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tschiiirp"}

	// joeyItalianPhrase represents joey italian phrase.
	joeyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "squack"}

	// joeyJapanesePhrase represents joey japanese phrase.
	joeyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でヤンス"}

	// joeyKoreanPhrase represents joey korean phrase.
	joeyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그래유"}

	// joeyLatinAmericanSpanishPhrase represents joey latin american spanish phrase.
	joeyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "uac-uac"}

	// joeyRussianPhrase represents joey russian phrase.
	joeyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кряк"}

	// joeySimplifiedChinesePhrase represents joey simplified chinese phrase.
	joeySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鸭鸭"}

	// joeySpanishPhrase represents joey spanish phrase.
	joeySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "uac-uac"}

	// joeyTraditionalChinesePhrase represents joey traditional chinese phrase.
	joeyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鴨鴨"}
)

var (
	// joeyPhrase represents joey phrase.
	joeyPhrase = nook.Languages{
		language.AmericanEnglish:      joeyAmericanEnglishPhrase,
		language.CanadianFrench:       joeyCanadianFrenchPhrase,
		language.Dutch:                joeyDutchPhrase,
		language.French:               joeyFrenchPhrase,
		language.German:               joeyGermanPhrase,
		language.Italian:              joeyItalianPhrase,
		language.Japanese:             joeyJapanesePhrase,
		language.Korean:               joeyKoreanPhrase,
		language.LatinAmericanSpanish: joeyLatinAmericanSpanishPhrase,
		language.Russian:              joeyRussianPhrase,
		language.SimplifiedChinese:    joeySimplifiedChinesePhrase,
		language.Spanish:              joeySpanishPhrase,
		language.TraditionalChinese:   joeyTraditionalChinesePhrase}
)

var (
	// Joey represents joey.
	Joey = nook.Villager{
		Character:   joeyCharacter,
		Personality: personality.Lazy,
		Phrase:      joeyPhrase}
)
