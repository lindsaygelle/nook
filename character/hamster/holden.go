package hamster

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
	// holdenBirthday represents holden birthday.
	holdenBirthday = nook.Birthday{
		Day:   11,
		Month: time.June}
)

var (
	// holdenCode represents holden code.
	holdenCode = nook.Code{
		Value: "ham08"}
)

var (
	// holdenAmericanEnglishName represents holden american english name.
	holdenAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Holden"}

	// holdenCanadianFrenchName represents holden canadian french name.
	holdenCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Glutin"}

	// holdenDutchName represents holden dutch name.
	holdenDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// holdenFrenchName represents holden french name.
	holdenFrenchName = nook.Name{
		Language: language.French,
		Value:    "Glutin"}

	// holdenGermanName represents holden german name.
	holdenGermanName = nook.Name{
		Language: language.German,
		Value:    "Holden"}

	// holdenItalianName represents holden italian name.
	holdenItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Collotto"}

	// holdenJapaneseName represents holden japanese name.
	holdenJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "のりぼう"}

	// holdenKoreanName represents holden korean name.
	holdenKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "노리보"}

	// holdenLatinAmericanSpanishName represents holden latin american spanish name.
	holdenLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Holden"}

	// holdenRussianName represents holden russian name.
	holdenRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// holdenSimplifiedChineseName represents holden simplified chinese name.
	holdenSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// holdenSpanishName represents holden spanish name.
	holdenSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Holden"}

	// holdenTraditionalChineseName represents holden traditional chinese name.
	holdenTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// holdenName represents holden name.
	holdenName = nook.Languages{
		language.AmericanEnglish:      holdenAmericanEnglishName,
		language.CanadianFrench:       holdenCanadianFrenchName,
		language.Dutch:                holdenDutchName,
		language.French:               holdenFrenchName,
		language.German:               holdenGermanName,
		language.Italian:              holdenItalianName,
		language.Japanese:             holdenJapaneseName,
		language.Korean:               holdenKoreanName,
		language.LatinAmericanSpanish: holdenLatinAmericanSpanishName,
		language.Russian:              holdenRussianName,
		language.SimplifiedChinese:    holdenSimplifiedChineseName,
		language.Spanish:              holdenSpanishName,
		language.TraditionalChinese:   holdenTraditionalChineseName}
)

var (
	// holdenCharacter represents holden character.
	holdenCharacter = nook.Character{
		Animal:   animal.Hamster,
		Birthday: holdenBirthday,
		Code:     holdenCode,
		Key:      character.Holden,
		Gender:   gender.Male,
		Name:     holdenName,
		Special:  false}
)

var (
	// holdenAmericanEnglishPhrase represents holden american english phrase.
	holdenAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "glue stick"}

	// holdenCanadianFrenchPhrase represents holden canadian french phrase.
	holdenCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "gluglu"}

	// holdenDutchPhrase represents holden dutch phrase.
	holdenDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// holdenFrenchPhrase represents holden french phrase.
	holdenFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "gluglu"}

	// holdenGermanPhrase represents holden german phrase.
	holdenGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "klebstift"}

	// holdenItalianPhrase represents holden italian phrase.
	holdenItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "incollo"}

	// holdenJapanesePhrase represents holden japanese phrase.
	holdenJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ノリノリ"}

	// holdenKoreanPhrase represents holden korean phrase.
	holdenKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "둠칫둠칫"}

	// holdenLatinAmericanSpanishPhrase represents holden latin american spanish phrase.
	holdenLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "tudiquesí"}

	// holdenRussianPhrase represents holden russian phrase.
	holdenRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// holdenSimplifiedChinesePhrase represents holden simplified chinese phrase.
	holdenSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	// holdenSpanishPhrase represents holden spanish phrase.
	holdenSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tudiquesí"}

	// holdenTraditionalChinesePhrase represents holden traditional chinese phrase.
	holdenTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// holdenPhrase represents holden phrase.
	holdenPhrase = nook.Languages{
		language.AmericanEnglish:      holdenAmericanEnglishPhrase,
		language.CanadianFrench:       holdenCanadianFrenchPhrase,
		language.Dutch:                holdenDutchPhrase,
		language.French:               holdenFrenchPhrase,
		language.German:               holdenGermanPhrase,
		language.Italian:              holdenItalianPhrase,
		language.Japanese:             holdenJapanesePhrase,
		language.Korean:               holdenKoreanPhrase,
		language.LatinAmericanSpanish: holdenLatinAmericanSpanishPhrase,
		language.Russian:              holdenRussianPhrase,
		language.SimplifiedChinese:    holdenSimplifiedChinesePhrase,
		language.Spanish:              holdenSpanishPhrase,
		language.TraditionalChinese:   holdenTraditionalChinesePhrase}
)

var (
	// Holden represents holden.
	Holden = nook.Villager{
		Character:   holdenCharacter,
		Personality: personality.Jock,
		Phrase:      holdenPhrase}
)
