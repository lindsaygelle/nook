package mouse

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
	// pennyBirthday represents penny birthday.
	pennyBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// pennyCode represents penny code.
	pennyCode = nook.Code{
		Value: ""}
)

var (
	// pennyAmericanEnglishName represents penny american english name.
	pennyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Penny"}

	// pennyCanadianFrenchName represents penny canadian french name.
	pennyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Barbara"}

	// pennyDutchName represents penny dutch name.
	pennyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// pennyFrenchName represents penny french name.
	pennyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Barbara"}

	// pennyGermanName represents penny german name.
	pennyGermanName = nook.Name{
		Language: language.German,
		Value:    "Susi"}

	// pennyItalianName represents penny italian name.
	pennyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Moretta"}

	// pennyJapaneseName represents penny japanese name.
	pennyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ネズこ"}

	// pennyKoreanName represents penny korean name.
	pennyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// pennyLatinAmericanSpanishName represents penny latin american spanish name.
	pennyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Penélope"}

	// pennyRussianName represents penny russian name.
	pennyRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// pennySimplifiedChineseName represents penny simplified chinese name.
	pennySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小小"}

	// pennySpanishName represents penny spanish name.
	pennySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Penélope"}

	// pennyTraditionalChineseName represents penny traditional chinese name.
	pennyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// pennyName represents penny name.
	pennyName = nook.Languages{
		language.AmericanEnglish:      pennyAmericanEnglishName,
		language.CanadianFrench:       pennyCanadianFrenchName,
		language.Dutch:                pennyDutchName,
		language.French:               pennyFrenchName,
		language.German:               pennyGermanName,
		language.Italian:              pennyItalianName,
		language.Japanese:             pennyJapaneseName,
		language.Korean:               pennyKoreanName,
		language.LatinAmericanSpanish: pennyLatinAmericanSpanishName,
		language.Russian:              pennyRussianName,
		language.SimplifiedChinese:    pennySimplifiedChineseName,
		language.Spanish:              pennySpanishName,
		language.TraditionalChinese:   pennyTraditionalChineseName}
)

var (
	// pennyCharacter represents penny character.
	pennyCharacter = nook.Character{
		Animal:   animal.Mouse,
		Birthday: pennyBirthday,
		Code:     pennyCode,
		Key:      character.Penny,
		Gender:   gender.Female,
		Name:     pennyName,
		Special:  false}
)

var (
	// pennyAmericanEnglishPhrase represents penny american english phrase.
	pennyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ska-WEAK"}

	// pennyCanadianFrenchPhrase represents penny canadian french phrase.
	pennyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "oh pétard!"}

	// pennyDutchPhrase represents penny dutch phrase.
	pennyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// pennyFrenchPhrase represents penny french phrase.
	pennyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "oh pétard!"}

	// pennyGermanPhrase represents penny german phrase.
	pennyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "piiquie"}

	// pennyItalianPhrase represents penny italian phrase.
	pennyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "Unknown"}

	// pennyJapanesePhrase represents penny japanese phrase.
	pennyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのさ"}

	// pennyKoreanPhrase represents penny korean phrase.
	pennyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// pennyLatinAmericanSpanishPhrase represents penny latin american spanish phrase.
	pennyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "carambita"}

	// pennyRussianPhrase represents penny russian phrase.
	pennyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// pennySimplifiedChinesePhrase represents penny simplified chinese phrase.
	pennySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Unknown"}

	// pennySpanishPhrase represents penny spanish phrase.
	pennySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "carambita"}

	// pennyTraditionalChinesePhrase represents penny traditional chinese phrase.
	pennyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// pennyPhrase represents penny phrase.
	pennyPhrase = nook.Languages{
		language.AmericanEnglish:      pennyAmericanEnglishPhrase,
		language.CanadianFrench:       pennyCanadianFrenchPhrase,
		language.Dutch:                pennyDutchPhrase,
		language.French:               pennyFrenchPhrase,
		language.German:               pennyGermanPhrase,
		language.Italian:              pennyItalianPhrase,
		language.Japanese:             pennyJapanesePhrase,
		language.Korean:               pennyKoreanPhrase,
		language.LatinAmericanSpanish: pennyLatinAmericanSpanishPhrase,
		language.Russian:              pennyRussianPhrase,
		language.SimplifiedChinese:    pennySimplifiedChinesePhrase,
		language.Spanish:              pennySpanishPhrase,
		language.TraditionalChinese:   pennyTraditionalChinesePhrase}
)

var (
	// Penny represents penny.
	Penny = nook.Villager{
		Character:   pennyCharacter,
		Personality: personality.Peppy,
		Phrase:      pennyPhrase}
)
