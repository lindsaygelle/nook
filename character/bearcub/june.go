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
	// juneBirthday represents june birthday.
	juneBirthday = nook.Birthday{
		Day:   21,
		Month: time.May}
)

var (
	// juneCode represents june code.
	juneCode = nook.Code{
		Value: "cbr13"}
)

var (
	// juneAmericanEnglishName represents june american english name.
	juneAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "June"}

	// juneCanadianFrenchName represents june canadian french name.
	juneCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Agnès"}

	// juneDutchName represents june dutch name.
	juneDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "June"}

	// juneFrenchName represents june french name.
	juneFrenchName = nook.Name{
		Language: language.French,
		Value:    "Agnès"}

	// juneGermanName represents june german name.
	juneGermanName = nook.Name{
		Language: language.German,
		Value:    "Juna"}

	// juneItalianName represents june italian name.
	juneItalianName = nook.Name{
		Language: language.Italian,
		Value:    "June"}

	// juneJapaneseName represents june japanese name.
	juneJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "メイ"}

	// juneKoreanName represents june korean name.
	juneKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "메이"}

	// juneLatinAmericanSpanishName represents june latin american spanish name.
	juneLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Hanalulú"}

	// juneRussianName represents june russian name.
	juneRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джун"}

	// juneSimplifiedChineseName represents june simplified chinese name.
	juneSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿妹"}

	// juneSpanishName represents june spanish name.
	juneSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Hanalulú"}

	// juneTraditionalChineseName represents june traditional chinese name.
	juneTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿妹"}
)

var (
	// juneName represents june name.
	juneName = nook.Languages{
		language.AmericanEnglish:      juneAmericanEnglishName,
		language.CanadianFrench:       juneCanadianFrenchName,
		language.Dutch:                juneDutchName,
		language.French:               juneFrenchName,
		language.German:               juneGermanName,
		language.Italian:              juneItalianName,
		language.Japanese:             juneJapaneseName,
		language.Korean:               juneKoreanName,
		language.LatinAmericanSpanish: juneLatinAmericanSpanishName,
		language.Russian:              juneRussianName,
		language.SimplifiedChinese:    juneSimplifiedChineseName,
		language.Spanish:              juneSpanishName,
		language.TraditionalChinese:   juneTraditionalChineseName}
)

var (
	// juneCharacter represents june character.
	juneCharacter = nook.Character{
		Animal:   animal.Bearcub,
		Birthday: juneBirthday,
		Code:     juneCode,
		Key:      character.June,
		Gender:   gender.Female,
		Name:     juneName,
		Special:  false}
)

var (
	// juneAmericanEnglishPhrase represents june american english phrase.
	juneAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "rainbow"}

	// juneCanadianFrenchPhrase represents june canadian french phrase.
	juneCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "cacou"}

	// juneDutchPhrase represents june dutch phrase.
	juneDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "wolkje"}

	// juneFrenchPhrase represents june french phrase.
	juneFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "cacou"}

	// juneGermanPhrase represents june german phrase.
	juneGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schönheit"}

	// juneItalianPhrase represents june italian phrase.
	juneItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grrande"}

	// juneJapanesePhrase represents june japanese phrase.
	juneJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "アルネ"}

	// juneKoreanPhrase represents june korean phrase.
	juneKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "있죠"}

	// juneLatinAmericanSpanishPhrase represents june latin american spanish phrase.
	juneLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "lurilá"}

	// juneRussianPhrase represents june russian phrase.
	juneRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "радуга-дуга"}

	// juneSimplifiedChinesePhrase represents june simplified chinese phrase.
	juneSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "有耶"}

	// juneSpanishPhrase represents june spanish phrase.
	juneSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cielo"}

	// juneTraditionalChinesePhrase represents june traditional chinese phrase.
	juneTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "有耶"}
)

var (
	// junePhrase represents june phrase.
	junePhrase = nook.Languages{
		language.AmericanEnglish:      juneAmericanEnglishPhrase,
		language.CanadianFrench:       juneCanadianFrenchPhrase,
		language.Dutch:                juneDutchPhrase,
		language.French:               juneFrenchPhrase,
		language.German:               juneGermanPhrase,
		language.Italian:              juneItalianPhrase,
		language.Japanese:             juneJapanesePhrase,
		language.Korean:               juneKoreanPhrase,
		language.LatinAmericanSpanish: juneLatinAmericanSpanishPhrase,
		language.Russian:              juneRussianPhrase,
		language.SimplifiedChinese:    juneSimplifiedChinesePhrase,
		language.Spanish:              juneSpanishPhrase,
		language.TraditionalChinese:   juneTraditionalChinesePhrase}
)

var (
	// June represents june.
	June = nook.Villager{
		Character:   juneCharacter,
		Personality: personality.Normal,
		Phrase:      junePhrase}
)
