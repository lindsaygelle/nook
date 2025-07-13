package bird

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
	// twiggyBirthday represents twiggy birthday.
	twiggyBirthday = nook.Birthday{
		Day:   13,
		Month: time.July}
)

var (
	// twiggyCode represents twiggy code.
	twiggyCode = nook.Code{
		Value: "brd03"}
)

var (
	// twiggyAmericanEnglishName represents twiggy american english name.
	twiggyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Twiggy"}

	// twiggyCanadianFrenchName represents twiggy canadian french name.
	twiggyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Titi"}

	// twiggyDutchName represents twiggy dutch name.
	twiggyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Twiggy"}

	// twiggyFrenchName represents twiggy french name.
	twiggyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Titi"}

	// twiggyGermanName represents twiggy german name.
	twiggyGermanName = nook.Name{
		Language: language.German,
		Value:    "Twiggy"}

	// twiggyItalianName represents twiggy italian name.
	twiggyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Titti"}

	// twiggyJapaneseName represents twiggy japanese name.
	twiggyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ピーチク"}

	// twiggyKoreanName represents twiggy korean name.
	twiggyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "핀틱"}

	// twiggyLatinAmericanSpanishName represents twiggy latin american spanish name.
	twiggyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tití"}

	// twiggyRussianName represents twiggy russian name.
	twiggyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Твигги"}

	// twiggySimplifiedChineseName represents twiggy simplified chinese name.
	twiggySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "叽叽"}

	// twiggySpanishName represents twiggy spanish name.
	twiggySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tití"}

	// twiggyTraditionalChineseName represents twiggy traditional chinese name.
	twiggyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘰嘰"}
)

var (
	// twiggyName represents twiggy name.
	twiggyName = nook.Languages{
		language.AmericanEnglish:      twiggyAmericanEnglishName,
		language.CanadianFrench:       twiggyCanadianFrenchName,
		language.Dutch:                twiggyDutchName,
		language.French:               twiggyFrenchName,
		language.German:               twiggyGermanName,
		language.Italian:              twiggyItalianName,
		language.Japanese:             twiggyJapaneseName,
		language.Korean:               twiggyKoreanName,
		language.LatinAmericanSpanish: twiggyLatinAmericanSpanishName,
		language.Russian:              twiggyRussianName,
		language.SimplifiedChinese:    twiggySimplifiedChineseName,
		language.Spanish:              twiggySpanishName,
		language.TraditionalChinese:   twiggyTraditionalChineseName}
)

var (
	// twiggyCharacter represents twiggy character.
	twiggyCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: twiggyBirthday,
		Code:     twiggyCode,
		Key:      character.Twiggy,
		Gender:   gender.Female,
		Name:     twiggyName,
		Special:  false}
)

var (
	// twiggyAmericanEnglishPhrase represents twiggy american english phrase.
	twiggyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cheepers"}

	// twiggyCanadianFrenchPhrase represents twiggy canadian french phrase.
	twiggyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "coolos"}

	// twiggyDutchPhrase represents twiggy dutch phrase.
	twiggyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "fwieeet"}

	// twiggyFrenchPhrase represents twiggy french phrase.
	twiggyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "coolos"}

	// twiggyGermanPhrase represents twiggy german phrase.
	twiggyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zwirtschi"}

	// twiggyItalianPhrase represents twiggy italian phrase.
	twiggyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ciiip"}

	// twiggyJapanesePhrase represents twiggy japanese phrase.
	twiggyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ッピ"}

	// twiggyKoreanPhrase represents twiggy korean phrase.
	twiggyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "크"}

	// twiggyLatinAmericanSpanishPhrase represents twiggy latin american spanish phrase.
	twiggyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tirití"}

	// twiggyRussianPhrase represents twiggy russian phrase.
	twiggyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тю-вить"}

	// twiggySimplifiedChinesePhrase represents twiggy simplified chinese phrase.
	twiggySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "叽"}

	// twiggySpanishPhrase represents twiggy spanish phrase.
	twiggySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tirití"}

	// twiggyTraditionalChinesePhrase represents twiggy traditional chinese phrase.
	twiggyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘰"}
)

var (
	// twiggyPhrase represents twiggy phrase.
	twiggyPhrase = nook.Languages{
		language.AmericanEnglish:      twiggyAmericanEnglishPhrase,
		language.CanadianFrench:       twiggyCanadianFrenchPhrase,
		language.Dutch:                twiggyDutchPhrase,
		language.French:               twiggyFrenchPhrase,
		language.German:               twiggyGermanPhrase,
		language.Italian:              twiggyItalianPhrase,
		language.Japanese:             twiggyJapanesePhrase,
		language.Korean:               twiggyKoreanPhrase,
		language.LatinAmericanSpanish: twiggyLatinAmericanSpanishPhrase,
		language.Russian:              twiggyRussianPhrase,
		language.SimplifiedChinese:    twiggySimplifiedChinesePhrase,
		language.Spanish:              twiggySpanishPhrase,
		language.TraditionalChinese:   twiggyTraditionalChinesePhrase}
)

var (
	// Twiggy represents twiggy.
	Twiggy = nook.Villager{
		Character:   twiggyCharacter,
		Personality: personality.Peppy,
		Phrase:      twiggyPhrase}
)
