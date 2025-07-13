package duck

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
	// frecklesBirthday represents freckles birthday.
	frecklesBirthday = nook.Birthday{
		Day:   19,
		Month: time.February}
)

var (
	// frecklesCode represents freckles code.
	frecklesCode = nook.Code{
		Value: "duk07"}
)

var (
	// frecklesAmericanEnglishName represents freckles american english name.
	frecklesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Freckles"}

	// frecklesCanadianFrenchName represents freckles canadian french name.
	frecklesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Caro"}

	// frecklesDutchName represents freckles dutch name.
	frecklesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Freckles"}

	// frecklesFrenchName represents freckles french name.
	frecklesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Caro"}

	// frecklesGermanName represents freckles german name.
	frecklesGermanName = nook.Name{
		Language: language.German,
		Value:    "Quacks"}

	// frecklesItalianName represents freckles italian name.
	frecklesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Semola"}

	// frecklesJapaneseName represents freckles japanese name.
	frecklesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マグロ"}

	// frecklesKoreanName represents freckles korean name.
	frecklesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "다랑어"}

	// frecklesLatinAmericanSpanishName represents freckles latin american spanish name.
	frecklesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rebeca"}

	// frecklesRussianName represents freckles russian name.
	frecklesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фреклс"}

	// frecklesSimplifiedChineseName represents freckles simplified chinese name.
	frecklesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小鲔"}

	// frecklesSpanishName represents freckles spanish name.
	frecklesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rebeca"}

	// frecklesTraditionalChineseName represents freckles traditional chinese name.
	frecklesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小鮪"}
)

var (
	// frecklesName represents freckles name.
	frecklesName = nook.Languages{
		language.AmericanEnglish:      frecklesAmericanEnglishName,
		language.CanadianFrench:       frecklesCanadianFrenchName,
		language.Dutch:                frecklesDutchName,
		language.French:               frecklesFrenchName,
		language.German:               frecklesGermanName,
		language.Italian:              frecklesItalianName,
		language.Japanese:             frecklesJapaneseName,
		language.Korean:               frecklesKoreanName,
		language.LatinAmericanSpanish: frecklesLatinAmericanSpanishName,
		language.Russian:              frecklesRussianName,
		language.SimplifiedChinese:    frecklesSimplifiedChineseName,
		language.Spanish:              frecklesSpanishName,
		language.TraditionalChinese:   frecklesTraditionalChineseName}
)

var (
	// frecklesCharacter represents freckles character.
	frecklesCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: frecklesBirthday,
		Code:     frecklesCode,
		Key:      character.Freckles,
		Gender:   gender.Female,
		Name:     frecklesName,
		Special:  false}
)

var (
	// frecklesAmericanEnglishPhrase represents freckles american english phrase.
	frecklesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ducky"}

	// frecklesCanadianFrenchPhrase represents freckles canadian french phrase.
	frecklesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "poussin"}

	// frecklesDutchPhrase represents freckles dutch phrase.
	frecklesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "eendje"}

	// frecklesFrenchPhrase represents freckles french phrase.
	frecklesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "poussin"}

	// frecklesGermanPhrase represents freckles german phrase.
	frecklesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnatter"}

	// frecklesItalianPhrase represents freckles italian phrase.
	frecklesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "qua qua"}

	// frecklesJapanesePhrase represents freckles japanese phrase.
	frecklesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ぎょぎょ"}

	// frecklesKoreanPhrase represents freckles korean phrase.
	frecklesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "물물"}

	// frecklesLatinAmericanSpanishPhrase represents freckles latin american spanish phrase.
	frecklesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cui-cuá"}

	// frecklesRussianPhrase represents freckles russian phrase.
	frecklesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "утенок"}

	// frecklesSimplifiedChinesePhrase represents freckles simplified chinese phrase.
	frecklesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鱼鱼"}

	// frecklesSpanishPhrase represents freckles spanish phrase.
	frecklesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cui-cuá"}

	// frecklesTraditionalChinesePhrase represents freckles traditional chinese phrase.
	frecklesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "魚魚"}
)

var (
	// frecklesPhrase represents freckles phrase.
	frecklesPhrase = nook.Languages{
		language.AmericanEnglish:      frecklesAmericanEnglishPhrase,
		language.CanadianFrench:       frecklesCanadianFrenchPhrase,
		language.Dutch:                frecklesDutchPhrase,
		language.French:               frecklesFrenchPhrase,
		language.German:               frecklesGermanPhrase,
		language.Italian:              frecklesItalianPhrase,
		language.Japanese:             frecklesJapanesePhrase,
		language.Korean:               frecklesKoreanPhrase,
		language.LatinAmericanSpanish: frecklesLatinAmericanSpanishPhrase,
		language.Russian:              frecklesRussianPhrase,
		language.SimplifiedChinese:    frecklesSimplifiedChinesePhrase,
		language.Spanish:              frecklesSpanishPhrase,
		language.TraditionalChinese:   frecklesTraditionalChinesePhrase}
)

var (
	// Freckles represents freckles.
	Freckles = nook.Villager{
		Character:   frecklesCharacter,
		Personality: personality.Peppy,
		Phrase:      frecklesPhrase}
)
