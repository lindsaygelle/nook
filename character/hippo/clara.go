package hippo

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
	// claraBirthday represents clara birthday.
	claraBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// claraCode represents clara code.
	claraCode = nook.Code{
		Value: ""}
)

var (
	// claraAmericanEnglishName represents clara american english name.
	claraAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Clara"}

	// claraCanadianFrenchName represents clara canadian french name.
	claraCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// claraDutchName represents clara dutch name.
	claraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// claraFrenchName represents clara french name.
	claraFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// claraGermanName represents clara german name.
	claraGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// claraItalianName represents clara italian name.
	claraItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// claraJapaneseName represents clara japanese name.
	claraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クララ"}

	// claraKoreanName represents clara korean name.
	claraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// claraLatinAmericanSpanishName represents clara latin american spanish name.
	claraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// claraRussianName represents clara russian name.
	claraRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// claraSimplifiedChineseName represents clara simplified chinese name.
	claraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// claraSpanishName represents clara spanish name.
	claraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// claraTraditionalChineseName represents clara traditional chinese name.
	claraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// claraName represents clara name.
	claraName = nook.Languages{
		language.AmericanEnglish:      claraAmericanEnglishName,
		language.CanadianFrench:       claraCanadianFrenchName,
		language.Dutch:                claraDutchName,
		language.French:               claraFrenchName,
		language.German:               claraGermanName,
		language.Italian:              claraItalianName,
		language.Japanese:             claraJapaneseName,
		language.Korean:               claraKoreanName,
		language.LatinAmericanSpanish: claraLatinAmericanSpanishName,
		language.Russian:              claraRussianName,
		language.SimplifiedChinese:    claraSimplifiedChineseName,
		language.Spanish:              claraSpanishName,
		language.TraditionalChinese:   claraTraditionalChineseName}
)

var (
	// claraCharacter represents clara character.
	claraCharacter = nook.Character{
		Animal:   animal.Hippo,
		Birthday: claraBirthday,
		Code:     claraCode,
		Key:      character.Clara,
		Gender:   gender.Female,
		Name:     claraName,
		Special:  false}
)

var (
	// claraAmericanEnglishPhrase represents clara american english phrase.
	claraAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "うふふ"}

	// claraCanadianFrenchPhrase represents clara canadian french phrase.
	claraCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// claraDutchPhrase represents clara dutch phrase.
	claraDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// claraFrenchPhrase represents clara french phrase.
	claraFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// claraGermanPhrase represents clara german phrase.
	claraGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// claraItalianPhrase represents clara italian phrase.
	claraItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// claraJapanesePhrase represents clara japanese phrase.
	claraJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "うふふ"}

	// claraKoreanPhrase represents clara korean phrase.
	claraKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// claraLatinAmericanSpanishPhrase represents clara latin american spanish phrase.
	claraLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// claraRussianPhrase represents clara russian phrase.
	claraRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// claraSimplifiedChinesePhrase represents clara simplified chinese phrase.
	claraSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// claraSpanishPhrase represents clara spanish phrase.
	claraSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// claraTraditionalChinesePhrase represents clara traditional chinese phrase.
	claraTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// claraPhrase represents clara phrase.
	claraPhrase = nook.Languages{
		language.AmericanEnglish:      claraAmericanEnglishPhrase,
		language.CanadianFrench:       claraCanadianFrenchPhrase,
		language.Dutch:                claraDutchPhrase,
		language.French:               claraFrenchPhrase,
		language.German:               claraGermanPhrase,
		language.Italian:              claraItalianPhrase,
		language.Japanese:             claraJapanesePhrase,
		language.Korean:               claraKoreanPhrase,
		language.LatinAmericanSpanish: claraLatinAmericanSpanishPhrase,
		language.Russian:              claraRussianPhrase,
		language.SimplifiedChinese:    claraSimplifiedChinesePhrase,
		language.Spanish:              claraSpanishPhrase,
		language.TraditionalChinese:   claraTraditionalChinesePhrase}
)

var (
	// Clara represents clara.
	Clara = nook.Villager{
		Character:   claraCharacter,
		Personality: personality.Normal,
		Phrase:      claraPhrase}
)
