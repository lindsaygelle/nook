package anteater

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// nosegayBirthday represents Nosegay's birthday.
var (
	// nosegayBirthday represents nosegay birthday.
	nosegayBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

// nosegayCode represents Nosegay's unique code.
var (
	// nosegayCode represents nosegay code.
	nosegayCode = nook.Code{
		Value: ""}
)

// Different names for Nosegay in various languages.
var (
	// nosegayAmericanEnglishName represents nosegay american english name.
	nosegayAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nosegay"}

	// nosegayCanadianFrenchName represents nosegay canadian french name.
	nosegayCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// nosegayDutchName represents nosegay dutch name.
	nosegayDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// nosegayFrenchName represents nosegay french name.
	nosegayFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tamara"}

	// nosegayGermanName represents nosegay german name.
	nosegayGermanName = nook.Name{
		Language: language.German,
		Value:    "Hortense"}

	// nosegayItalianName represents nosegay italian name.
	nosegayItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Narica"}

	// nosegayJapaneseName represents nosegay japanese name.
	nosegayJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アリンコ"}

	// nosegayKoreanName represents nosegay korean name.
	nosegayKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// nosegayLatinAmericanSpanishName represents nosegay latin american spanish name.
	nosegayLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// nosegayRussianName represents nosegay russian name.
	nosegayRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// nosegaySimplifiedChineseName represents nosegay simplified chinese name.
	nosegaySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小蚁"}

	// nosegaySpanishName represents nosegay spanish name.
	nosegaySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marisa"}

	// nosegayTraditionalChineseName represents nosegay traditional chinese name.
	nosegayTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

// nosegayName represents Nosegay's name in different languages.
var (
	// nosegayName represents nosegay name.
	nosegayName = nook.Languages{
		language.AmericanEnglish:      nosegayAmericanEnglishName,
		language.CanadianFrench:       nosegayCanadianFrenchName,
		language.Dutch:                nosegayDutchName,
		language.French:               nosegayFrenchName,
		language.German:               nosegayGermanName,
		language.Italian:              nosegayItalianName,
		language.Japanese:             nosegayJapaneseName,
		language.Korean:               nosegayKoreanName,
		language.LatinAmericanSpanish: nosegayLatinAmericanSpanishName,
		language.Russian:              nosegayRussianName,
		language.SimplifiedChinese:    nosegaySimplifiedChineseName,
		language.Spanish:              nosegaySpanishName,
		language.TraditionalChinese:   nosegayTraditionalChineseName}
)

// nosegayCharacter represents Nosegay's character information.
var (
	// nosegayCharacter represents nosegay character.
	nosegayCharacter = nook.Character{
		Animal:   animal.Anteater,
		Birthday: nosegayBirthday,
		Code:     nosegayCode,
		Key:      character.Nosegay,
		Gender:   gender.Female,
		Name:     nosegayName,
		Special:  false}
)

// Different phrases spoken by Nosegay in various languages.
var (
	// nosegayAmericanEnglishPhrase represents nosegay american english phrase.
	nosegayAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hoooonk"}

	// nosegayCanadianFrenchPhrase represents nosegay canadian french phrase.
	nosegayCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// nosegayDutchPhrase represents nosegay dutch phrase.
	nosegayDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// nosegayFrenchPhrase represents nosegay french phrase.
	nosegayFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "hoooonk"}

	// nosegayGermanPhrase represents nosegay german phrase.
	nosegayGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnauf"}

	// nosegayItalianPhrase represents nosegay italian phrase.
	nosegayItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hoooonk"}

	// nosegayJapanesePhrase represents nosegay japanese phrase.
	nosegayJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でアリ"}

	// nosegayKoreanPhrase represents nosegay korean phrase.
	nosegayKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// nosegayLatinAmericanSpanishPhrase represents nosegay latin american spanish phrase.
	nosegayLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// nosegayRussianPhrase represents nosegay russian phrase.
	nosegayRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// nosegaySimplifiedChinesePhrase represents nosegay simplified chinese phrase.
	nosegaySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "确实"}

	// nosegaySpanishPhrase represents nosegay spanish phrase.
	nosegaySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "achís"}

	// nosegayTraditionalChinesePhrase represents nosegay traditional chinese phrase.
	nosegayTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

// nosegayPhrase represents Nosegay's phrases in different languages.
var (
	// nosegayPhrase represents nosegay phrase.
	nosegayPhrase = nook.Languages{
		language.AmericanEnglish:      nosegayAmericanEnglishPhrase,
		language.CanadianFrench:       nosegayCanadianFrenchPhrase,
		language.Dutch:                nosegayDutchPhrase,
		language.French:               nosegayFrenchPhrase,
		language.German:               nosegayGermanPhrase,
		language.Italian:              nosegayItalianPhrase,
		language.Japanese:             nosegayJapanesePhrase,
		language.Korean:               nosegayKoreanPhrase,
		language.LatinAmericanSpanish: nosegayLatinAmericanSpanishPhrase,
		language.Russian:              nosegayRussianPhrase,
		language.SimplifiedChinese:    nosegaySimplifiedChinesePhrase,
		language.Spanish:              nosegaySpanishPhrase,
		language.TraditionalChinese:   nosegayTraditionalChinesePhrase}
)

// Nosegay represents the character Nosegay with her complete information.
var (
	// Nosegay represents nosegay.
	Nosegay = nook.Villager{
		Character:   nosegayCharacter,
		Personality: personality.Normal,
		Phrase:      nosegayPhrase}
)
