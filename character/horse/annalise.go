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
	// annaliseBirthday represents annalise birthday.
	annaliseBirthday = nook.Birthday{
		Day:   2,
		Month: time.December}
)

var (
	// annaliseCode represents annalise code.
	annaliseCode = nook.Code{
		Value: "hrs09"}
)

var (
	// annaliseAmericanEnglishName represents annalise american english name.
	annaliseAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Annalise"}

	// annaliseCanadianFrenchName represents annalise canadian french name.
	annaliseCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Âne-lise"}

	// annaliseDutchName represents annalise dutch name.
	annaliseDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Annalise"}

	// annaliseFrenchName represents annalise french name.
	annaliseFrenchName = nook.Name{
		Language: language.French,
		Value:    "Âne-lise"}

	// annaliseGermanName represents annalise german name.
	annaliseGermanName = nook.Name{
		Language: language.German,
		Value:    "Annerose"}

	// annaliseItalianName represents annalise italian name.
	annaliseItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lisa"}

	// annaliseJapaneseName represents annalise japanese name.
	annaliseJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シルブプレ"}

	// annaliseKoreanName represents annalise korean name.
	annaliseKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "실부플레"}

	// annaliseLatinAmericanSpanishName represents annalise latin american spanish name.
	annaliseLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Isadora"}

	// annaliseRussianName represents annalise russian name.
	annaliseRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лиза"}

	// annaliseSimplifiedChineseName represents annalise simplified chinese name.
	annaliseSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "舒芙蕾"}

	// annaliseSpanishName represents annalise spanish name.
	annaliseSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Isadora"}

	// annaliseTraditionalChineseName represents annalise traditional chinese name.
	annaliseTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "舒芙蕾"}
)

var (
	// annaliseName represents annalise name.
	annaliseName = nook.Languages{
		language.AmericanEnglish:      annaliseAmericanEnglishName,
		language.CanadianFrench:       annaliseCanadianFrenchName,
		language.Dutch:                annaliseDutchName,
		language.French:               annaliseFrenchName,
		language.German:               annaliseGermanName,
		language.Italian:              annaliseItalianName,
		language.Japanese:             annaliseJapaneseName,
		language.Korean:               annaliseKoreanName,
		language.LatinAmericanSpanish: annaliseLatinAmericanSpanishName,
		language.Russian:              annaliseRussianName,
		language.SimplifiedChinese:    annaliseSimplifiedChineseName,
		language.Spanish:              annaliseSpanishName,
		language.TraditionalChinese:   annaliseTraditionalChineseName}
)

var (
	// annaliseCharacter represents annalise character.
	annaliseCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: annaliseBirthday,
		Code:     annaliseCode,
		Key:      character.Annalise,
		Gender:   gender.Female,
		Name:     annaliseName,
		Special:  false}
)

var (
	// annaliseAmericanEnglishPhrase represents annalise american english phrase.
	annaliseAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nipper"}

	// annaliseCanadianFrenchPhrase represents annalise canadian french phrase.
	annaliseCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "hippique"}

	// annaliseDutchPhrase represents annalise dutch phrase.
	annaliseDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "veulen"}

	// annaliseFrenchPhrase represents annalise french phrase.
	annaliseFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "galapia"}

	// annaliseGermanPhrase represents annalise german phrase.
	annaliseGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mausezahn"}

	// annaliseItalianPhrase represents annalise italian phrase.
	annaliseItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "palio"}

	// annaliseJapanesePhrase represents annalise japanese phrase.
	annaliseJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "サヴァ"}

	// annaliseKoreanPhrase represents annalise korean phrase.
	annaliseKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "봉쥬르"}

	// annaliseLatinAmericanSpanishPhrase represents annalise latin american spanish phrase.
	annaliseLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "nih-nah"}

	// annaliseRussianPhrase represents annalise russian phrase.
	annaliseRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кусь-кусь"}

	// annaliseSimplifiedChinesePhrase represents annalise simplified chinese phrase.
	annaliseSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "舒服"}

	// annaliseSpanishPhrase represents annalise spanish phrase.
	annaliseSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "criatura"}

	// annaliseTraditionalChinesePhrase represents annalise traditional chinese phrase.
	annaliseTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "舒服"}
)

var (
	// annalisePhrase represents annalise phrase.
	annalisePhrase = nook.Languages{
		language.AmericanEnglish:      annaliseAmericanEnglishPhrase,
		language.CanadianFrench:       annaliseCanadianFrenchPhrase,
		language.Dutch:                annaliseDutchPhrase,
		language.French:               annaliseFrenchPhrase,
		language.German:               annaliseGermanPhrase,
		language.Italian:              annaliseItalianPhrase,
		language.Japanese:             annaliseJapanesePhrase,
		language.Korean:               annaliseKoreanPhrase,
		language.LatinAmericanSpanish: annaliseLatinAmericanSpanishPhrase,
		language.Russian:              annaliseRussianPhrase,
		language.SimplifiedChinese:    annaliseSimplifiedChinesePhrase,
		language.Spanish:              annaliseSpanishPhrase,
		language.TraditionalChinese:   annaliseTraditionalChinesePhrase}
)

var (
	// Annalise represents annalise.
	Annalise = nook.Villager{
		Character:   annaliseCharacter,
		Personality: personality.Snooty,
		Phrase:      annalisePhrase}
)
