package cow

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
	// carrotBirthday represents carrot birthday.
	carrotBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// carrotCode represents carrot code.
	carrotCode = nook.Code{
		Value: ""}
)

var (
	// carrotAmericanEnglishName represents carrot american english name.
	carrotAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Carrot"}

	// carrotCanadianFrenchName represents carrot canadian french name.
	carrotCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// carrotDutchName represents carrot dutch name.
	carrotDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// carrotFrenchName represents carrot french name.
	carrotFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// carrotGermanName represents carrot german name.
	carrotGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// carrotItalianName represents carrot italian name.
	carrotItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// carrotJapaneseName represents carrot japanese name.
	carrotJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キャロット"}

	// carrotKoreanName represents carrot korean name.
	carrotKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// carrotLatinAmericanSpanishName represents carrot latin american spanish name.
	carrotLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// carrotRussianName represents carrot russian name.
	carrotRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// carrotSimplifiedChineseName represents carrot simplified chinese name.
	carrotSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// carrotSpanishName represents carrot spanish name.
	carrotSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// carrotTraditionalChineseName represents carrot traditional chinese name.
	carrotTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// carrotName represents carrot name.
	carrotName = nook.Languages{
		language.AmericanEnglish:      carrotAmericanEnglishName,
		language.CanadianFrench:       carrotCanadianFrenchName,
		language.Dutch:                carrotDutchName,
		language.French:               carrotFrenchName,
		language.German:               carrotGermanName,
		language.Italian:              carrotItalianName,
		language.Japanese:             carrotJapaneseName,
		language.Korean:               carrotKoreanName,
		language.LatinAmericanSpanish: carrotLatinAmericanSpanishName,
		language.Russian:              carrotRussianName,
		language.SimplifiedChinese:    carrotSimplifiedChineseName,
		language.Spanish:              carrotSpanishName,
		language.TraditionalChinese:   carrotTraditionalChineseName}
)

var (
	// carrotCharacter represents carrot character.
	carrotCharacter = nook.Character{
		Animal:   animal.Cow,
		Birthday: carrotBirthday,
		Code:     carrotCode,
		Key:      character.Carrot,
		Gender:   gender.Female,
		Name:     carrotName,
		Special:  false}
)

var (
	// carrotAmericanEnglishPhrase represents carrot american english phrase.
	carrotAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "きゃはっ"}

	// carrotCanadianFrenchPhrase represents carrot canadian french phrase.
	carrotCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// carrotDutchPhrase represents carrot dutch phrase.
	carrotDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// carrotFrenchPhrase represents carrot french phrase.
	carrotFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// carrotGermanPhrase represents carrot german phrase.
	carrotGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// carrotItalianPhrase represents carrot italian phrase.
	carrotItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// carrotJapanesePhrase represents carrot japanese phrase.
	carrotJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "きゃはっ"}

	// carrotKoreanPhrase represents carrot korean phrase.
	carrotKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// carrotLatinAmericanSpanishPhrase represents carrot latin american spanish phrase.
	carrotLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// carrotRussianPhrase represents carrot russian phrase.
	carrotRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// carrotSimplifiedChinesePhrase represents carrot simplified chinese phrase.
	carrotSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// carrotSpanishPhrase represents carrot spanish phrase.
	carrotSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// carrotTraditionalChinesePhrase represents carrot traditional chinese phrase.
	carrotTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// carrotPhrase represents carrot phrase.
	carrotPhrase = nook.Languages{
		language.AmericanEnglish:      carrotAmericanEnglishPhrase,
		language.CanadianFrench:       carrotCanadianFrenchPhrase,
		language.Dutch:                carrotDutchPhrase,
		language.French:               carrotFrenchPhrase,
		language.German:               carrotGermanPhrase,
		language.Italian:              carrotItalianPhrase,
		language.Japanese:             carrotJapanesePhrase,
		language.Korean:               carrotKoreanPhrase,
		language.LatinAmericanSpanish: carrotLatinAmericanSpanishPhrase,
		language.Russian:              carrotRussianPhrase,
		language.SimplifiedChinese:    carrotSimplifiedChinesePhrase,
		language.Spanish:              carrotSpanishPhrase,
		language.TraditionalChinese:   carrotTraditionalChinesePhrase}
)

var (
	// Carrot represents carrot.
	Carrot = nook.Villager{
		Character:   carrotCharacter,
		Personality: personality.Normal,
		Phrase:      carrotPhrase}
)
