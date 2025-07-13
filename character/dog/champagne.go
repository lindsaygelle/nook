package dog

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
	// champagneBirthday represents champagne birthday.
	champagneBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// champagneCode represents champagne code.
	champagneCode = nook.Code{
		Value: ""}
)

var (
	// champagneAmericanEnglishName represents champagne american english name.
	champagneAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Champagne"}

	// champagneCanadianFrenchName represents champagne canadian french name.
	champagneCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// champagneDutchName represents champagne dutch name.
	champagneDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// champagneFrenchName represents champagne french name.
	champagneFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	// champagneGermanName represents champagne german name.
	champagneGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	// champagneItalianName represents champagne italian name.
	champagneItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// champagneJapaneseName represents champagne japanese name.
	champagneJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シャンパン"}

	// champagneKoreanName represents champagne korean name.
	champagneKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// champagneLatinAmericanSpanishName represents champagne latin american spanish name.
	champagneLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// champagneRussianName represents champagne russian name.
	champagneRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// champagneSimplifiedChineseName represents champagne simplified chinese name.
	champagneSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// champagneSpanishName represents champagne spanish name.
	champagneSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// champagneTraditionalChineseName represents champagne traditional chinese name.
	champagneTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// champagneName represents champagne name.
	champagneName = nook.Languages{
		language.AmericanEnglish:      champagneAmericanEnglishName,
		language.CanadianFrench:       champagneCanadianFrenchName,
		language.Dutch:                champagneDutchName,
		language.French:               champagneFrenchName,
		language.German:               champagneGermanName,
		language.Italian:              champagneItalianName,
		language.Japanese:             champagneJapaneseName,
		language.Korean:               champagneKoreanName,
		language.LatinAmericanSpanish: champagneLatinAmericanSpanishName,
		language.Russian:              champagneRussianName,
		language.SimplifiedChinese:    champagneSimplifiedChineseName,
		language.Spanish:              champagneSpanishName,
		language.TraditionalChinese:   champagneTraditionalChineseName}
)

var (
	// champagneCharacter represents champagne character.
	champagneCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: champagneBirthday,
		Code:     champagneCode,
		Key:      character.Champagne,
		Gender:   gender.Male,
		Name:     champagneName,
		Special:  false}
)

var (
	// champagneAmericanEnglishPhrase represents champagne american english phrase.
	champagneAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ったく"}

	// champagneCanadianFrenchPhrase represents champagne canadian french phrase.
	champagneCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// champagneDutchPhrase represents champagne dutch phrase.
	champagneDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// champagneFrenchPhrase represents champagne french phrase.
	champagneFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	// champagneGermanPhrase represents champagne german phrase.
	champagneGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	// champagneItalianPhrase represents champagne italian phrase.
	champagneItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	// champagneJapanesePhrase represents champagne japanese phrase.
	champagneJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ったく"}

	// champagneKoreanPhrase represents champagne korean phrase.
	champagneKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// champagneLatinAmericanSpanishPhrase represents champagne latin american spanish phrase.
	champagneLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// champagneRussianPhrase represents champagne russian phrase.
	champagneRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// champagneSimplifiedChinesePhrase represents champagne simplified chinese phrase.
	champagneSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// champagneSpanishPhrase represents champagne spanish phrase.
	champagneSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	// champagneTraditionalChinesePhrase represents champagne traditional chinese phrase.
	champagneTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// champagnePhrase represents champagne phrase.
	champagnePhrase = nook.Languages{
		language.AmericanEnglish:      champagneAmericanEnglishPhrase,
		language.CanadianFrench:       champagneCanadianFrenchPhrase,
		language.Dutch:                champagneDutchPhrase,
		language.French:               champagneFrenchPhrase,
		language.German:               champagneGermanPhrase,
		language.Italian:              champagneItalianPhrase,
		language.Japanese:             champagneJapanesePhrase,
		language.Korean:               champagneKoreanPhrase,
		language.LatinAmericanSpanish: champagneLatinAmericanSpanishPhrase,
		language.Russian:              champagneRussianPhrase,
		language.SimplifiedChinese:    champagneSimplifiedChinesePhrase,
		language.Spanish:              champagneSpanishPhrase,
		language.TraditionalChinese:   champagneTraditionalChinesePhrase}
)

var (
	// Champagne represents champagne.
	Champagne = nook.Villager{
		Character:   champagneCharacter,
		Personality: personality.Cranky,
		Phrase:      champagnePhrase}
)
