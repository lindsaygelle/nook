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
	// jacobBirthday represents jacob birthday.
	jacobBirthday = nook.Birthday{
		Day:   24,
		Month: time.August}
)

var (
	// jacobCode represents jacob code.
	jacobCode = nook.Code{
		Value: "brd11"}
)

var (
	// jacobAmericanEnglishName represents jacob american english name.
	jacobAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Jacob Jakey"}

	// jacobCanadianFrenchName represents jacob canadian french name.
	jacobCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jacob"}

	// jacobDutchName represents jacob dutch name.
	jacobDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Jacob"}

	// jacobFrenchName represents jacob french name.
	jacobFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jacob"}

	// jacobGermanName represents jacob german name.
	jacobGermanName = nook.Name{
		Language: language.German,
		Value:    "Jesko"}

	// jacobItalianName represents jacob italian name.
	jacobItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Valerio"}

	// jacobJapaneseName represents jacob japanese name.
	jacobJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジャコテン"}

	// jacobKoreanName represents jacob korean name.
	jacobKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "야곱"}

	// jacobLatinAmericanSpanishName represents jacob latin american spanish name.
	jacobLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Repollo"}

	// jacobRussianName represents jacob russian name.
	jacobRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джейкоб"}

	// jacobSimplifiedChineseName represents jacob simplified chinese name.
	jacobSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "余板"}

	// jacobSpanishName represents jacob spanish name.
	jacobSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Repollo"}

	// jacobTraditionalChineseName represents jacob traditional chinese name.
	jacobTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "余板"}
)

var (
	// jacobName represents jacob name.
	jacobName = nook.Languages{
		language.AmericanEnglish:      jacobAmericanEnglishName,
		language.CanadianFrench:       jacobCanadianFrenchName,
		language.Dutch:                jacobDutchName,
		language.French:               jacobFrenchName,
		language.German:               jacobGermanName,
		language.Italian:              jacobItalianName,
		language.Japanese:             jacobJapaneseName,
		language.Korean:               jacobKoreanName,
		language.LatinAmericanSpanish: jacobLatinAmericanSpanishName,
		language.Russian:              jacobRussianName,
		language.SimplifiedChinese:    jacobSimplifiedChineseName,
		language.Spanish:              jacobSpanishName,
		language.TraditionalChinese:   jacobTraditionalChineseName}
)

var (
	// jacobCharacter represents jacob character.
	jacobCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: jacobBirthday,
		Code:     jacobCode,
		Key:      character.Jacob,
		Gender:   gender.Male,
		Name:     jacobName,
		Special:  false}
)

var (
	// jacobAmericanEnglishPhrase represents jacob american english phrase.
	jacobAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ya feel"}

	// jacobCanadianFrenchPhrase represents jacob canadian french phrase.
	jacobCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "asticô"}

	// jacobDutchPhrase represents jacob dutch phrase.
	jacobDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snappie"}

	// jacobFrenchPhrase represents jacob french phrase.
	jacobFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "asticô"}

	// jacobGermanPhrase represents jacob german phrase.
	jacobGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "tirilü"}

	// jacobItalianPhrase represents jacob italian phrase.
	jacobItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ciiipito"}

	// jacobJapanesePhrase represents jacob japanese phrase.
	jacobJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "っつーの"}

	// jacobKoreanPhrase represents jacob korean phrase.
	jacobKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그거야"}

	// jacobLatinAmericanSpanishPhrase represents jacob latin american spanish phrase.
	jacobLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chiurp"}

	// jacobRussianPhrase represents jacob russian phrase.
	jacobRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "чуррп"}

	// jacobSimplifiedChinesePhrase represents jacob simplified chinese phrase.
	jacobSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "话说"}

	// jacobSpanishPhrase represents jacob spanish phrase.
	jacobSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pelusilla"}

	// jacobTraditionalChinesePhrase represents jacob traditional chinese phrase.
	jacobTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "話說"}
)

var (
	// jacobPhrase represents jacob phrase.
	jacobPhrase = nook.Languages{
		language.AmericanEnglish:      jacobAmericanEnglishPhrase,
		language.CanadianFrench:       jacobCanadianFrenchPhrase,
		language.Dutch:                jacobDutchPhrase,
		language.French:               jacobFrenchPhrase,
		language.German:               jacobGermanPhrase,
		language.Italian:              jacobItalianPhrase,
		language.Japanese:             jacobJapanesePhrase,
		language.Korean:               jacobKoreanPhrase,
		language.LatinAmericanSpanish: jacobLatinAmericanSpanishPhrase,
		language.Russian:              jacobRussianPhrase,
		language.SimplifiedChinese:    jacobSimplifiedChinesePhrase,
		language.Spanish:              jacobSpanishPhrase,
		language.TraditionalChinese:   jacobTraditionalChinesePhrase}
)

var (
	// Jacob represents jacob.
	Jacob = nook.Villager{
		Character:   jacobCharacter,
		Personality: personality.Lazy,
		Phrase:      jacobPhrase}
)
