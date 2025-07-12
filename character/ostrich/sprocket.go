package ostrich

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
	// sprocketBirthday represents sprocket birthday.
	sprocketBirthday = nook.Birthday{
		Day:   1,
		Month: time.December}
)

var (
	// sprocketCode represents sprocket code.
	sprocketCode = nook.Code{
		Value: "ost03"}
)

var (
	// sprocketAmericanEnglishName represents sprocket american english name.
	sprocketAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sprocket"}

	// sprocketCanadianFrenchName represents sprocket canadian french name.
	sprocketCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Laflèche"}

	// sprocketDutchName represents sprocket dutch name.
	sprocketDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sprocket"}

	// sprocketFrenchName represents sprocket french name.
	sprocketFrenchName = nook.Name{
		Language: language.French,
		Value:    "Laflèche"}

	// sprocketGermanName represents sprocket german name.
	sprocketGermanName = nook.Name{
		Language: language.German,
		Value:    "Lutz"}

	// sprocketItalianName represents sprocket italian name.
	sprocketItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Frankie"}

	// sprocketJapaneseName represents sprocket japanese name.
	sprocketJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ヘルツ"}

	// sprocketKoreanName represents sprocket korean name.
	sprocketKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "헤르츠"}

	// sprocketLatinAmericanSpanishName represents sprocket latin american spanish name.
	sprocketLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ráfaga"}

	// sprocketRussianName represents sprocket russian name.
	sprocketRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Спрокет"}

	// sprocketSimplifiedChineseName represents sprocket simplified chinese name.
	sprocketSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鹤兹"}

	// sprocketSpanishName represents sprocket spanish name.
	sprocketSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ráfaga"}

	// sprocketTraditionalChineseName represents sprocket traditional chinese name.
	sprocketTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鶴茲"}
)

var (
	// sprocketName represents sprocket name.
	sprocketName = nook.Languages{
		language.AmericanEnglish:      sprocketAmericanEnglishName,
		language.CanadianFrench:       sprocketCanadianFrenchName,
		language.Dutch:                sprocketDutchName,
		language.French:               sprocketFrenchName,
		language.German:               sprocketGermanName,
		language.Italian:              sprocketItalianName,
		language.Japanese:             sprocketJapaneseName,
		language.Korean:               sprocketKoreanName,
		language.LatinAmericanSpanish: sprocketLatinAmericanSpanishName,
		language.Russian:              sprocketRussianName,
		language.SimplifiedChinese:    sprocketSimplifiedChineseName,
		language.Spanish:              sprocketSpanishName,
		language.TraditionalChinese:   sprocketTraditionalChineseName}
)

var (
	// sprocketCharacter represents sprocket character.
	sprocketCharacter = nook.Character{
		Animal:   animal.Ostrich,
		Birthday: sprocketBirthday,
		Code:     sprocketCode,
		Key:      character.Sprocket,
		Gender:   gender.Male,
		Name:     sprocketName,
		Special:  false}
)

var (
	// sprocketAmericanEnglishPhrase represents sprocket american english phrase.
	sprocketAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "zort"}

	// sprocketCanadianFrenchPhrase represents sprocket canadian french phrase.
	sprocketCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pik pik"}

	// sprocketDutchPhrase represents sprocket dutch phrase.
	sprocketDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "klik-kluk"}

	// sprocketFrenchPhrase represents sprocket french phrase.
	sprocketFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pik pik"}

	// sprocketGermanPhrase represents sprocket german phrase.
	sprocketGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "affenzahn"}

	// sprocketItalianPhrase represents sprocket italian phrase.
	sprocketItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "niiun"}

	// sprocketJapanesePhrase represents sprocket japanese phrase.
	sprocketJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だメカ"}

	// sprocketKoreanPhrase represents sprocket korean phrase.
	sprocketKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "치지직"}

	// sprocketLatinAmericanSpanishPhrase represents sprocket latin american spanish phrase.
	sprocketLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "chiuuun"}

	// sprocketRussianPhrase represents sprocket russian phrase.
	sprocketRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "дзынь"}

	// sprocketSimplifiedChinesePhrase represents sprocket simplified chinese phrase.
	sprocketSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "机械"}

	// sprocketSpanishPhrase represents sprocket spanish phrase.
	sprocketSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "chiuuun"}

	// sprocketTraditionalChinesePhrase represents sprocket traditional chinese phrase.
	sprocketTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "機械"}
)

var (
	// sprocketPhrase represents sprocket phrase.
	sprocketPhrase = nook.Languages{
		language.AmericanEnglish:      sprocketAmericanEnglishPhrase,
		language.CanadianFrench:       sprocketCanadianFrenchPhrase,
		language.Dutch:                sprocketDutchPhrase,
		language.French:               sprocketFrenchPhrase,
		language.German:               sprocketGermanPhrase,
		language.Italian:              sprocketItalianPhrase,
		language.Japanese:             sprocketJapanesePhrase,
		language.Korean:               sprocketKoreanPhrase,
		language.LatinAmericanSpanish: sprocketLatinAmericanSpanishPhrase,
		language.Russian:              sprocketRussianPhrase,
		language.SimplifiedChinese:    sprocketSimplifiedChinesePhrase,
		language.Spanish:              sprocketSpanishPhrase,
		language.TraditionalChinese:   sprocketTraditionalChinesePhrase}
)

var (
	// Sprocket represents sprocket.
	Sprocket = nook.Villager{
		Character:   sprocketCharacter,
		Personality: personality.Jock,
		Phrase:      sprocketPhrase}
)
