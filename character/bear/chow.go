package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// chowBirthday represents Chow's birthday.
var (
	// chowBirthday represents chow birthday.
	chowBirthday = nook.Birthday{
		Day:   22,
		Month: time.July}
)

// chowCode represents Chow's unique code.
var (
	// chowCode represents chow code.
	chowCode = nook.Code{
		Value: "bea03"}
)

// Different names for Chow in various languages.
var (
	// chowAmericanEnglishName represents chow american english name.
	chowAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chow"}

	// chowCanadianFrenchName represents chow canadian french name.
	chowCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Chulin"}

	// chowDutchName represents chow dutch name.
	chowDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Chow"}

	// chowFrenchName represents chow french name.
	chowFrenchName = nook.Name{
		Language: language.French,
		Value:    "Chulin"}

	// chowGermanName represents chow german name.
	chowGermanName = nook.Name{
		Language: language.German,
		Value:    "Chang"}

	// chowItalianName represents chow italian name.
	chowItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Chowchow"}

	// chowJapaneseName represents chow japanese name.
	chowJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チャウヤン"}

	// chowKoreanName represents chow korean name.
	chowKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "츄양"}

	// chowLatinAmericanSpanishName represents chow latin american spanish name.
	chowLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pando"}

	// chowRussianName represents chow russian name.
	chowRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Чау"}

	// chowSimplifiedChineseName represents chow simplified chinese name.
	chowSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "朝阳"}

	// chowSpanishName represents chow spanish name.
	chowSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pando"}

	// chowTraditionalChineseName represents chow traditional chinese name.
	chowTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "朝陽"}
)

// chowName represents Chow's name in different languages.
var (
	// chowName represents chow name.
	chowName = nook.Languages{
		language.AmericanEnglish:      chowAmericanEnglishName,
		language.CanadianFrench:       chowCanadianFrenchName,
		language.Dutch:                chowDutchName,
		language.French:               chowFrenchName,
		language.German:               chowGermanName,
		language.Italian:              chowItalianName,
		language.Japanese:             chowJapaneseName,
		language.Korean:               chowKoreanName,
		language.LatinAmericanSpanish: chowLatinAmericanSpanishName,
		language.Russian:              chowRussianName,
		language.SimplifiedChinese:    chowSimplifiedChineseName,
		language.Spanish:              chowSpanishName,
		language.TraditionalChinese:   chowTraditionalChineseName}
)

// chowCharacter represents Chow's character information.
var (
	// chowCharacter represents chow character.
	chowCharacter = nook.Character{
		Animal:   animal.Bear,
		Birthday: chowBirthday,
		Code:     chowCode,
		Key:      character.Chow,
		Gender:   gender.Male,
		Name:     chowName,
		Special:  false}
)

// Different phrases spoken by Chow in various languages.
var (
	// chowAmericanEnglishPhrase represents chow american english phrase.
	chowAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "aiya"}

	// chowCanadianFrenchPhrase represents chow canadian french phrase.
	chowCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "lala"}

	// chowDutchPhrase represents chow dutch phrase.
	chowDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kiai"}

	// chowFrenchPhrase represents chow french phrase.
	chowFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "lala"}

	// chowGermanPhrase represents chow german phrase.
	chowGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hiijaa"}

	// chowItalianPhrase represents chow italian phrase.
	chowItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ehilà"}

	// chowJapanesePhrase represents chow japanese phrase.
	chowJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アルヨ"}

	// chowKoreanPhrase represents chow korean phrase.
	chowKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그럼"}

	// chowLatinAmericanSpanishPhrase represents chow latin american spanish phrase.
	chowLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grinchu"}

	// chowRussianPhrase represents chow russian phrase.
	chowRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ай-яй"}

	// chowSimplifiedChinesePhrase represents chow simplified chinese phrase.
	chowSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "有喔"}

	// chowSpanishPhrase represents chow spanish phrase.
	chowSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ándale-oso"}

	// chowTraditionalChinesePhrase represents chow traditional chinese phrase.
	chowTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "有喔"}
)

// chowPhrase represents Chow's phrases in different languages.
var (
	// chowPhrase represents chow phrase.
	chowPhrase = nook.Languages{
		language.AmericanEnglish:      chowAmericanEnglishPhrase,
		language.CanadianFrench:       chowCanadianFrenchPhrase,
		language.Dutch:                chowDutchPhrase,
		language.French:               chowFrenchPhrase,
		language.German:               chowGermanPhrase,
		language.Italian:              chowItalianPhrase,
		language.Japanese:             chowJapanesePhrase,
		language.Korean:               chowKoreanPhrase,
		language.LatinAmericanSpanish: chowLatinAmericanSpanishPhrase,
		language.Russian:              chowRussianPhrase,
		language.SimplifiedChinese:    chowSimplifiedChinesePhrase,
		language.Spanish:              chowSpanishPhrase,
		language.TraditionalChinese:   chowTraditionalChinesePhrase}
)

// Chow represents the character Chow with his complete information.
var (
	// Chow represents chow.
	Chow = nook.Villager{
		Character:   chowCharacter,
		Personality: personality.Cranky,
		Phrase:      chowPhrase}
)
