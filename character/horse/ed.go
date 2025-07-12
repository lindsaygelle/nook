package horse

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
	// edBirthday represents ed birthday.
	edBirthday = nook.Birthday{
		Day:   16,
		Month: time.September}
)

var (
	// edCode represents ed code.
	edCode = nook.Code{
		Value: "hrs06"}
)

var (
	// edAmericanEnglishName represents ed american english name.
	edAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ed"}

	// edCanadianFrenchName represents ed canadian french name.
	edCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Édouard"}

	// edDutchName represents ed dutch name.
	edDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ed"}

	// edFrenchName represents ed french name.
	edFrenchName = nook.Name{
		Language: language.French,
		Value:    "Édouard"}

	// edGermanName represents ed german name.
	edGermanName = nook.Name{
		Language: language.German,
		Value:    "Hermann"}

	// edItalianName represents ed italian name.
	edItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Codino"}

	// edJapaneseName represents ed japanese name.
	edJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キザノホマレ"}

	// edKoreanName represents ed korean name.
	edKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "꺼벙"}

	// edLatinAmericanSpanishName represents ed latin american spanish name.
	edLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Crinaldo"}

	// edRussianName represents ed russian name.
	edRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Эд"}

	// edSimplifiedChineseName represents ed simplified chinese name.
	edSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "马誉"}

	// edSpanishName represents ed spanish name.
	edSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Crinaldo"}

	// edTraditionalChineseName represents ed traditional chinese name.
	edTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "馬譽"}
)

var (
	// edName represents ed name.
	edName = nook.Languages{
		language.AmericanEnglish:      edAmericanEnglishName,
		language.CanadianFrench:       edCanadianFrenchName,
		language.Dutch:                edDutchName,
		language.French:               edFrenchName,
		language.German:               edGermanName,
		language.Italian:              edItalianName,
		language.Japanese:             edJapaneseName,
		language.Korean:               edKoreanName,
		language.LatinAmericanSpanish: edLatinAmericanSpanishName,
		language.Russian:              edRussianName,
		language.SimplifiedChinese:    edSimplifiedChineseName,
		language.Spanish:              edSpanishName,
		language.TraditionalChinese:   edTraditionalChineseName}
)

var (
	// edCharacter represents ed character.
	edCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: edBirthday,
		Code:     edCode,
		Key:      character.Ed,
		Gender:   gender.Male,
		Name:     edName,
		Special:  false}
)

var (
	// edAmericanEnglishPhrase represents ed american english phrase.
	edAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "greenhorn"}

	// edCanadianFrenchPhrase represents ed canadian french phrase.
	edCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tagada"}

	// edDutchPhrase represents ed dutch phrase.
	edDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ju"}

	// edFrenchPhrase represents ed french phrase.
	edFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tagada"}

	// edGermanPhrase represents ed german phrase.
	edGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "greenhorn"}

	// edItalianPhrase represents ed italian phrase.
	edItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pirulì"}

	// edJapanesePhrase represents ed japanese phrase.
	edJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "じゃない"}

	// edKoreanPhrase represents ed korean phrase.
	edKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "게슴츠레"}

	// edLatinAmericanSpanishPhrase represents ed latin american spanish phrase.
	edLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "jo-joy"}

	// edRussianPhrase represents ed russian phrase.
	edRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "новичок"}

	// edSimplifiedChinesePhrase represents ed simplified chinese phrase.
	edSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "不是啦"}

	// edSpanishPhrase represents ed spanish phrase.
	edSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "jo-joy"}

	// edTraditionalChinesePhrase represents ed traditional chinese phrase.
	edTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "不是啦"}
)

var (
	// edPhrase represents ed phrase.
	edPhrase = nook.Languages{
		language.AmericanEnglish:      edAmericanEnglishPhrase,
		language.CanadianFrench:       edCanadianFrenchPhrase,
		language.Dutch:                edDutchPhrase,
		language.French:               edFrenchPhrase,
		language.German:               edGermanPhrase,
		language.Italian:              edItalianPhrase,
		language.Japanese:             edJapanesePhrase,
		language.Korean:               edKoreanPhrase,
		language.LatinAmericanSpanish: edLatinAmericanSpanishPhrase,
		language.Russian:              edRussianPhrase,
		language.SimplifiedChinese:    edSimplifiedChinesePhrase,
		language.Spanish:              edSpanishPhrase,
		language.TraditionalChinese:   edTraditionalChinesePhrase}
)

var (
	// Ed represents ed.
	Ed = nook.Villager{
		Character:   edCharacter,
		Personality: personality.Smug,
		Phrase:      edPhrase}
)
