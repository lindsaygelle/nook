package pig

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
	// hamboBirthday represents hambo birthday.
	hamboBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// hamboCode represents hambo code.
	hamboCode = nook.Code{
		Value: ""}
)

var (
	// hamboAmericanEnglishName represents hambo american english name.
	hamboAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hambo"}

	// hamboCanadianFrenchName represents hambo canadian french name.
	hamboCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// hamboDutchName represents hambo dutch name.
	hamboDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// hamboFrenchName represents hambo french name.
	hamboFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jambo"}

	// hamboGermanName represents hambo german name.
	hamboGermanName = nook.Name{
		Language: language.German,
		Value:    "Silvio"}

	// hamboItalianName represents hambo italian name.
	hamboItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Zampone"}

	// hamboJapaneseName represents hambo japanese name.
	hamboJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "はちまき"}

	// hamboKoreanName represents hambo korean name.
	hamboKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// hamboLatinAmericanSpanishName represents hambo latin american spanish name.
	hamboLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// hamboRussianName represents hambo russian name.
	hamboRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// hamboSimplifiedChineseName represents hambo simplified chinese name.
	hamboSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "头巾"}

	// hamboSpanishName represents hambo spanish name.
	hamboSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Marcial"}

	// hamboTraditionalChineseName represents hambo traditional chinese name.
	hamboTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// hamboName represents hambo name.
	hamboName = nook.Languages{
		language.AmericanEnglish:      hamboAmericanEnglishName,
		language.CanadianFrench:       hamboCanadianFrenchName,
		language.Dutch:                hamboDutchName,
		language.French:               hamboFrenchName,
		language.German:               hamboGermanName,
		language.Italian:              hamboItalianName,
		language.Japanese:             hamboJapaneseName,
		language.Korean:               hamboKoreanName,
		language.LatinAmericanSpanish: hamboLatinAmericanSpanishName,
		language.Russian:              hamboRussianName,
		language.SimplifiedChinese:    hamboSimplifiedChineseName,
		language.Spanish:              hamboSpanishName,
		language.TraditionalChinese:   hamboTraditionalChineseName}
)

var (
	// hamboCharacter represents hambo character.
	hamboCharacter = nook.Character{
		Animal:   animal.Pig,
		Birthday: hamboBirthday,
		Code:     hamboCode,
		Key:      character.Hambo,
		Gender:   gender.Male,
		Name:     hamboName,
		Special:  false}
)

var (
	// hamboAmericanEnglishPhrase represents hambo american english phrase.
	hamboAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "yo"}

	// hamboCanadianFrenchPhrase represents hambo canadian french phrase.
	hamboCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// hamboDutchPhrase represents hambo dutch phrase.
	hamboDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// hamboFrenchPhrase represents hambo french phrase.
	hamboFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pti boudin"}

	// hamboGermanPhrase represents hambo german phrase.
	hamboGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hajaaa"}

	// hamboItalianPhrase represents hambo italian phrase.
	hamboItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "yo"}

	// hamboJapanesePhrase represents hambo japanese phrase.
	hamboJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だス"}

	// hamboKoreanPhrase represents hambo korean phrase.
	hamboKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// hamboLatinAmericanSpanishPhrase represents hambo latin american spanish phrase.
	hamboLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// hamboRussianPhrase represents hambo russian phrase.
	hamboRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// hamboSimplifiedChinesePhrase represents hambo simplified chinese phrase.
	hamboSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哈罗"}

	// hamboSpanishPhrase represents hambo spanish phrase.
	hamboSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "morcillita"}

	// hamboTraditionalChinesePhrase represents hambo traditional chinese phrase.
	hamboTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// hamboPhrase represents hambo phrase.
	hamboPhrase = nook.Languages{
		language.AmericanEnglish:      hamboAmericanEnglishPhrase,
		language.CanadianFrench:       hamboCanadianFrenchPhrase,
		language.Dutch:                hamboDutchPhrase,
		language.French:               hamboFrenchPhrase,
		language.German:               hamboGermanPhrase,
		language.Italian:              hamboItalianPhrase,
		language.Japanese:             hamboJapanesePhrase,
		language.Korean:               hamboKoreanPhrase,
		language.LatinAmericanSpanish: hamboLatinAmericanSpanishPhrase,
		language.Russian:              hamboRussianPhrase,
		language.SimplifiedChinese:    hamboSimplifiedChinesePhrase,
		language.Spanish:              hamboSpanishPhrase,
		language.TraditionalChinese:   hamboTraditionalChinesePhrase}
)

var (
	// Hambo represents hambo.
	Hambo = nook.Villager{
		Character:   hamboCharacter,
		Personality: personality.Jock,
		Phrase:      hamboPhrase}
)
