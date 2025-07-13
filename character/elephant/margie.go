package elephant

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
	// margieBirthday represents margie birthday.
	margieBirthday = nook.Birthday{
		Day:   28,
		Month: time.January}
)

var (
	// margieCode represents margie code.
	margieCode = nook.Code{
		Value: "elp04"}
)

var (
	// margieAmericanEnglishName represents margie american english name.
	margieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Margie"}

	// margieCanadianFrenchName represents margie canadian french name.
	margieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Maguy"}

	// margieDutchName represents margie dutch name.
	margieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Margie"}

	// margieFrenchName represents margie french name.
	margieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maguy"}

	// margieGermanName represents margie german name.
	margieGermanName = nook.Name{
		Language: language.German,
		Value:    "Elisa"}

	// margieItalianName represents margie italian name.
	margieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Marianna"}

	// margieJapaneseName represents margie japanese name.
	margieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サリー"}

	// margieKoreanName represents margie korean name.
	margieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "샐리"}

	// margieLatinAmericanSpanishName represents margie latin american spanish name.
	margieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rita"}

	// margieRussianName represents margie russian name.
	margieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Марджи"}

	// margieSimplifiedChineseName represents margie simplified chinese name.
	margieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莎莉"}

	// margieSpanishName represents margie spanish name.
	margieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rita"}

	// margieTraditionalChineseName represents margie traditional chinese name.
	margieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莎莉"}
)

var (
	// margieName represents margie name.
	margieName = nook.Languages{
		language.AmericanEnglish:      margieAmericanEnglishName,
		language.CanadianFrench:       margieCanadianFrenchName,
		language.Dutch:                margieDutchName,
		language.French:               margieFrenchName,
		language.German:               margieGermanName,
		language.Italian:              margieItalianName,
		language.Japanese:             margieJapaneseName,
		language.Korean:               margieKoreanName,
		language.LatinAmericanSpanish: margieLatinAmericanSpanishName,
		language.Russian:              margieRussianName,
		language.SimplifiedChinese:    margieSimplifiedChineseName,
		language.Spanish:              margieSpanishName,
		language.TraditionalChinese:   margieTraditionalChineseName}
)

var (
	// margieCharacter represents margie character.
	margieCharacter = nook.Character{
		Animal:   animal.Elephant,
		Birthday: margieBirthday,
		Code:     margieCode,
		Key:      character.Margie,
		Gender:   gender.Female,
		Name:     margieName,
		Special:  false}
)

var (
	// margieAmericanEnglishPhrase represents margie american english phrase.
	margieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "tootie"}

	// margieCanadianFrenchPhrase represents margie canadian french phrase.
	margieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "toutie"}

	// margieDutchPhrase represents margie dutch phrase.
	margieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pwèèpie"}

	// margieFrenchPhrase represents margie french phrase.
	margieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "toutie"}

	// margieGermanPhrase represents margie german phrase.
	margieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "trampl"}

	// margieItalianPhrase represents margie italian phrase.
	margieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "tootie"}

	// margieJapanesePhrase represents margie japanese phrase.
	margieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "シャララ"}

	// margieKoreanPhrase represents margie korean phrase.
	margieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "샬랄라"}

	// margieLatinAmericanSpanishPhrase represents margie latin american spanish phrase.
	margieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "fiuf-fiuf"}

	// margieRussianPhrase represents margie russian phrase.
	margieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тру-у-ти"}

	// margieSimplifiedChinesePhrase represents margie simplified chinese phrase.
	margieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "莎啦啦"}

	// margieSpanishPhrase represents margie spanish phrase.
	margieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "flaqui"}

	// margieTraditionalChinesePhrase represents margie traditional chinese phrase.
	margieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "莎啦啦"}
)

var (
	// margiePhrase represents margie phrase.
	margiePhrase = nook.Languages{
		language.AmericanEnglish:      margieAmericanEnglishPhrase,
		language.CanadianFrench:       margieCanadianFrenchPhrase,
		language.Dutch:                margieDutchPhrase,
		language.French:               margieFrenchPhrase,
		language.German:               margieGermanPhrase,
		language.Italian:              margieItalianPhrase,
		language.Japanese:             margieJapanesePhrase,
		language.Korean:               margieKoreanPhrase,
		language.LatinAmericanSpanish: margieLatinAmericanSpanishPhrase,
		language.Russian:              margieRussianPhrase,
		language.SimplifiedChinese:    margieSimplifiedChinesePhrase,
		language.Spanish:              margieSpanishPhrase,
		language.TraditionalChinese:   margieTraditionalChinesePhrase}
)

var (
	// Margie represents margie.
	Margie = nook.Villager{
		Character:   margieCharacter,
		Personality: personality.Normal,
		Phrase:      margiePhrase}
)
