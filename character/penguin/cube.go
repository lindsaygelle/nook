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
	// cubeBirthday represents cube birthday.
	cubeBirthday = nook.Birthday{
		Day:   29,
		Month: time.January}
)

var (
	// cubeCode represents cube code.
	cubeCode = nook.Code{
		Value: "pgn02"}
)

var (
	// cubeAmericanEnglishName represents cube american english name.
	cubeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cube"}

	// cubeCanadianFrenchName represents cube canadian french name.
	cubeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cube"}

	// cubeDutchName represents cube dutch name.
	cubeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cube"}

	// cubeFrenchName represents cube french name.
	cubeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cube"}

	// cubeGermanName represents cube german name.
	cubeGermanName = nook.Name{
		Language: language.German,
		Value:    "Cube"}

	// cubeItalianName represents cube italian name.
	cubeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cubetto"}

	// cubeJapaneseName represents cube japanese name.
	cubeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ビス"}

	// cubeKoreanName represents cube korean name.
	cubeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "빙수"}

	// cubeLatinAmericanSpanishName represents cube latin american spanish name.
	cubeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cube"}

	// cubeRussianName represents cube russian name.
	cubeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кьюб"}

	// cubeSimplifiedChineseName represents cube simplified chinese name.
	cubeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗斯"}

	// cubeSpanishName represents cube spanish name.
	cubeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cube"}

	// cubeTraditionalChineseName represents cube traditional chinese name.
	cubeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羅斯"}
)

var (
	// cubeName represents cube name.
	cubeName = nook.Languages{
		language.AmericanEnglish:      cubeAmericanEnglishName,
		language.CanadianFrench:       cubeCanadianFrenchName,
		language.Dutch:                cubeDutchName,
		language.French:               cubeFrenchName,
		language.German:               cubeGermanName,
		language.Italian:              cubeItalianName,
		language.Japanese:             cubeJapaneseName,
		language.Korean:               cubeKoreanName,
		language.LatinAmericanSpanish: cubeLatinAmericanSpanishName,
		language.Russian:              cubeRussianName,
		language.SimplifiedChinese:    cubeSimplifiedChineseName,
		language.Spanish:              cubeSpanishName,
		language.TraditionalChinese:   cubeTraditionalChineseName}
)

var (
	// cubeCharacter represents cube character.
	cubeCharacter = nook.Character{
		Animal:   animal.Penguin,
		Birthday: cubeBirthday,
		Code:     cubeCode,
		Key:      character.Cube,
		Gender:   gender.Male,
		Name:     cubeName,
		Special:  false}
)

var (
	// cubeAmericanEnglishPhrase represents cube american english phrase.
	cubeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "d-d-dude"}

	// cubeCanadianFrenchPhrase represents cube canadian french phrase.
	cubeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "cococopain"}

	// cubeDutchPhrase represents cube dutch phrase.
	cubeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "k-k-kerel"}

	// cubeFrenchPhrase represents cube french phrase.
	cubeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "cococopain"}

	// cubeGermanPhrase represents cube german phrase.
	cubeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "d-d-dude"}

	// cubeItalianPhrase represents cube italian phrase.
	cubeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "brivido"}

	// cubeJapanesePhrase represents cube japanese phrase.
	cubeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ペンペン"}

	// cubeKoreanPhrase represents cube korean phrase.
	cubeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "땡땡"}

	// cubeLatinAmericanSpanishPhrase represents cube latin american spanish phrase.
	cubeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "waap"}

	// cubeRussianPhrase represents cube russian phrase.
	cubeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ч-чубрик"}

	// cubeSimplifiedChinesePhrase represents cube simplified chinese phrase.
	cubeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "十字"}

	// cubeSpanishPhrase represents cube spanish phrase.
	cubeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "waap"}

	// cubeTraditionalChinesePhrase represents cube traditional chinese phrase.
	cubeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "十字"}
)

var (
	// cubePhrase represents cube phrase.
	cubePhrase = nook.Languages{
		language.AmericanEnglish:      cubeAmericanEnglishPhrase,
		language.CanadianFrench:       cubeCanadianFrenchPhrase,
		language.Dutch:                cubeDutchPhrase,
		language.French:               cubeFrenchPhrase,
		language.German:               cubeGermanPhrase,
		language.Italian:              cubeItalianPhrase,
		language.Japanese:             cubeJapanesePhrase,
		language.Korean:               cubeKoreanPhrase,
		language.LatinAmericanSpanish: cubeLatinAmericanSpanishPhrase,
		language.Russian:              cubeRussianPhrase,
		language.SimplifiedChinese:    cubeSimplifiedChinesePhrase,
		language.Spanish:              cubeSpanishPhrase,
		language.TraditionalChinese:   cubeTraditionalChinesePhrase}
)

var (
	// Cube represents cube.
	Cube = nook.Villager{
		Character:   cubeCharacter,
		Personality: personality.Lazy,
		Phrase:      cubePhrase}
)
