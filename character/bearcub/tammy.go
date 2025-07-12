package bearcub

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
	// tammyBirthday represents tammy birthday.
	tammyBirthday = nook.Birthday{
		Day:   23,
		Month: time.June}
)

var (
	// tammyCode represents tammy code.
	tammyCode = nook.Code{
		Value: "cbr17"}
)

var (
	// tammyAmericanEnglishName represents tammy american english name.
	tammyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tammy"}

	// tammyCanadianFrenchName represents tammy canadian french name.
	tammyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Vanessa"}

	// tammyDutchName represents tammy dutch name.
	tammyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tammy"}

	// tammyFrenchName represents tammy french name.
	tammyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Vanessa"}

	// tammyGermanName represents tammy german name.
	tammyGermanName = nook.Name{
		Language: language.German,
		Value:    "Tatjana"}

	// tammyItalianName represents tammy italian name.
	tammyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tamara"}

	// tammyJapaneseName represents tammy japanese name.
	tammyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アネッサ"}

	// tammyKoreanName represents tammy korean name.
	tammyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아네사"}

	// tammyLatinAmericanSpanishName represents tammy latin american spanish name.
	tammyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aída"}

	// tammyRussianName represents tammy russian name.
	tammyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тамми"}

	// tammySimplifiedChineseName represents tammy simplified chinese name.
	tammySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "然姐"}

	// tammySpanishName represents tammy spanish name.
	tammySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aída"}

	// tammyTraditionalChineseName represents tammy traditional chinese name.
	tammyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "然姐"}
)

var (
	// tammyName represents tammy name.
	tammyName = nook.Languages{
		language.AmericanEnglish:      tammyAmericanEnglishName,
		language.CanadianFrench:       tammyCanadianFrenchName,
		language.Dutch:                tammyDutchName,
		language.French:               tammyFrenchName,
		language.German:               tammyGermanName,
		language.Italian:              tammyItalianName,
		language.Japanese:             tammyJapaneseName,
		language.Korean:               tammyKoreanName,
		language.LatinAmericanSpanish: tammyLatinAmericanSpanishName,
		language.Russian:              tammyRussianName,
		language.SimplifiedChinese:    tammySimplifiedChineseName,
		language.Spanish:              tammySpanishName,
		language.TraditionalChinese:   tammyTraditionalChineseName}
)

var (
	// tammyCharacter represents tammy character.
	tammyCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: tammyBirthday,
		Code:     tammyCode,
		Key:      character.Tammy,
		Gender:   gender.Female,
		Name:     tammyName,
		Special:  false}
)

var (
	// tammyAmericanEnglishPhrase represents tammy american english phrase.
	tammyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ya heard"}

	// tammyCanadianFrenchPhrase represents tammy canadian french phrase.
	tammyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grifouille"}

	// tammyDutchPhrase represents tammy dutch phrase.
	tammyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "weet je"}

	// tammyFrenchPhrase represents tammy french phrase.
	tammyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grifouille"}

	// tammyGermanPhrase represents tammy german phrase.
	tammyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bossi"}

	// tammyItalianPhrase represents tammy italian phrase.
	tammyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ohè"}

	// tammyJapanesePhrase represents tammy japanese phrase.
	tammyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だってサ"}

	// tammyKoreanPhrase represents tammy korean phrase.
	tammyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "차라리"}

	// tammyLatinAmericanSpanishPhrase represents tammy latin american spanish phrase.
	tammyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tararí"}

	// tammyRussianPhrase represents tammy russian phrase.
	tammyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "понимаешь"}

	// tammySimplifiedChinesePhrase represents tammy simplified chinese phrase.
	tammySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "竟然说"}

	// tammySpanishPhrase represents tammy spanish phrase.
	tammySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tararí"}

	// tammyTraditionalChinesePhrase represents tammy traditional chinese phrase.
	tammyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "竟然說"}
)

var (
	// tammyPhrase represents tammy phrase.
	tammyPhrase = nook.Languages{
		language.AmericanEnglish:      tammyAmericanEnglishPhrase,
		language.CanadianFrench:       tammyCanadianFrenchPhrase,
		language.Dutch:                tammyDutchPhrase,
		language.French:               tammyFrenchPhrase,
		language.German:               tammyGermanPhrase,
		language.Italian:              tammyItalianPhrase,
		language.Japanese:             tammyJapanesePhrase,
		language.Korean:               tammyKoreanPhrase,
		language.LatinAmericanSpanish: tammyLatinAmericanSpanishPhrase,
		language.Russian:              tammyRussianPhrase,
		language.SimplifiedChinese:    tammySimplifiedChinesePhrase,
		language.Spanish:              tammySpanishPhrase,
		language.TraditionalChinese:   tammyTraditionalChinesePhrase}
)

var (
	// Tammy represents tammy.
	Tammy = nook.Villager{
		Character:   tammyCharacter,
		Personality: personality.BigSister,
		Phrase:      tammyPhrase}
)
