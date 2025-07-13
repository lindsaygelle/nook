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
	// luchaBirthday represents lucha birthday.
	luchaBirthday = nook.Birthday{
		Day:   12,
		Month: time.December}
)

var (
	// luchaCode represents lucha code.
	luchaCode = nook.Code{
		Value: "brd15"}
)

var (
	// luchaAmericanEnglishName represents lucha american english name.
	luchaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lucha"}

	// luchaCanadianFrenchName represents lucha canadian french name.
	luchaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Condor"}

	// luchaDutchName represents lucha dutch name.
	luchaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lucha"}

	// luchaFrenchName represents lucha french name.
	luchaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Condor"}

	// luchaGermanName represents lucha german name.
	luchaGermanName = nook.Name{
		Language: language.German,
		Value:    "Lukas"}

	// luchaItalianName represents lucha italian name.
	luchaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Loreto"}

	// luchaJapaneseName represents lucha japanese name.
	luchaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マスカラス"}

	// luchaKoreanName represents lucha korean name.
	luchaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마스카라스"}

	// luchaLatinAmericanSpanishName represents lucha latin american spanish name.
	luchaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Plumerio"}

	// luchaRussianName represents lucha russian name.
	luchaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Луча"}

	// luchaSimplifiedChineseName represents lucha simplified chinese name.
	luchaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "摔角鸦"}

	// luchaSpanishName represents lucha spanish name.
	luchaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Plumerio"}

	// luchaTraditionalChineseName represents lucha traditional chinese name.
	luchaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "摔角鴉"}
)

var (
	// luchaName represents lucha name.
	luchaName = nook.Languages{
		language.AmericanEnglish:      luchaAmericanEnglishName,
		language.CanadianFrench:       luchaCanadianFrenchName,
		language.Dutch:                luchaDutchName,
		language.French:               luchaFrenchName,
		language.German:               luchaGermanName,
		language.Italian:              luchaItalianName,
		language.Japanese:             luchaJapaneseName,
		language.Korean:               luchaKoreanName,
		language.LatinAmericanSpanish: luchaLatinAmericanSpanishName,
		language.Russian:              luchaRussianName,
		language.SimplifiedChinese:    luchaSimplifiedChineseName,
		language.Spanish:              luchaSpanishName,
		language.TraditionalChinese:   luchaTraditionalChineseName}
)

var (
	// luchaCharacter represents lucha character.
	luchaCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: luchaBirthday,
		Code:     luchaCode,
		Key:      character.Lucha,
		Gender:   gender.Male,
		Name:     luchaName,
		Special:  false}
)

var (
	// luchaAmericanEnglishPhrase represents lucha american english phrase.
	luchaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cacaw"}

	// luchaCanadianFrenchPhrase represents lucha canadian french phrase.
	luchaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "yaaaah"}

	// luchaDutchPhrase represents lucha dutch phrase.
	luchaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kakauw"}

	// luchaFrenchPhrase represents lucha french phrase.
	luchaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "yaaaah"}

	// luchaGermanPhrase represents lucha german phrase.
	luchaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mauser"}

	// luchaItalianPhrase represents lucha italian phrase.
	luchaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "lorito"}

	// luchaJapanesePhrase represents lucha japanese phrase.
	luchaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "カァー"}

	// luchaKoreanPhrase represents lucha korean phrase.
	luchaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "까아악"}

	// luchaLatinAmericanSpanishPhrase represents lucha latin american spanish phrase.
	luchaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "pío"}

	// luchaRussianPhrase represents lucha russian phrase.
	luchaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кар-ко"}

	// luchaSimplifiedChinesePhrase represents lucha simplified chinese phrase.
	luchaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘎"}

	// luchaSpanishPhrase represents lucha spanish phrase.
	luchaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "picopico"}

	// luchaTraditionalChinesePhrase represents lucha traditional chinese phrase.
	luchaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘎"}
)

var (
	// luchaPhrase represents lucha phrase.
	luchaPhrase = nook.Languages{
		language.AmericanEnglish:      luchaAmericanEnglishPhrase,
		language.CanadianFrench:       luchaCanadianFrenchPhrase,
		language.Dutch:                luchaDutchPhrase,
		language.French:               luchaFrenchPhrase,
		language.German:               luchaGermanPhrase,
		language.Italian:              luchaItalianPhrase,
		language.Japanese:             luchaJapanesePhrase,
		language.Korean:               luchaKoreanPhrase,
		language.LatinAmericanSpanish: luchaLatinAmericanSpanishPhrase,
		language.Russian:              luchaRussianPhrase,
		language.SimplifiedChinese:    luchaSimplifiedChinesePhrase,
		language.Spanish:              luchaSpanishPhrase,
		language.TraditionalChinese:   luchaTraditionalChinesePhrase}
)

var (
	// Lucha represents lucha.
	Lucha = nook.Villager{
		Character:   luchaCharacter,
		Personality: personality.Smug,
		Phrase:      luchaPhrase}
)
