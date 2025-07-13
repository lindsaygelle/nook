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
	// rioBirthday represents rio birthday.
	rioBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// rioCode represents rio code.
	rioCode = nook.Code{
		Value: ""}
)

var (
	// rioAmericanEnglishName represents rio american english name.
	rioAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rio"}

	// rioCanadianFrenchName represents rio canadian french name.
	rioCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// rioDutchName represents rio dutch name.
	rioDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// rioFrenchName represents rio french name.
	rioFrenchName = nook.Name{
		Language: language.French,
		Value:    "Estrella"}

	// rioGermanName represents rio german name.
	rioGermanName = nook.Name{
		Language: language.German,
		Value:    "Verona"}

	// rioItalianName represents rio italian name.
	rioItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rita"}

	// rioJapaneseName represents rio japanese name.
	rioJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "デジャネイロ"}

	// rioKoreanName represents rio korean name.
	rioKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// rioLatinAmericanSpanishName represents rio latin american spanish name.
	rioLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// rioRussianName represents rio russian name.
	rioRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// rioSimplifiedChineseName represents rio simplified chinese name.
	rioSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "欢欢"}

	// rioSpanishName represents rio spanish name.
	rioSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rio"}

	// rioTraditionalChineseName represents rio traditional chinese name.
	rioTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// rioName represents rio name.
	rioName = nook.Languages{
		language.AmericanEnglish:      rioAmericanEnglishName,
		language.CanadianFrench:       rioCanadianFrenchName,
		language.Dutch:                rioDutchName,
		language.French:               rioFrenchName,
		language.German:               rioGermanName,
		language.Italian:              rioItalianName,
		language.Japanese:             rioJapaneseName,
		language.Korean:               rioKoreanName,
		language.LatinAmericanSpanish: rioLatinAmericanSpanishName,
		language.Russian:              rioRussianName,
		language.SimplifiedChinese:    rioSimplifiedChineseName,
		language.Spanish:              rioSpanishName,
		language.TraditionalChinese:   rioTraditionalChineseName}
)

var (
	// rioCharacter represents rio character.
	rioCharacter = nook.Character{
		Animal:   animal.Ostrich,
		Birthday: rioBirthday,
		Code:     rioCode,
		Key:      character.Rio,
		Gender:   gender.Female,
		Name:     rioName,
		Special:  false}
)

var (
	// rioAmericanEnglishPhrase represents rio american english phrase.
	rioAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "li'l chick"}

	// rioCanadianFrenchPhrase represents rio canadian french phrase.
	rioCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// rioDutchPhrase represents rio dutch phrase.
	rioDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// rioFrenchPhrase represents rio french phrase.
	rioFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bon bref"}

	// rioGermanPhrase represents rio german phrase.
	rioGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "küken"}

	// rioItalianPhrase represents rio italian phrase.
	rioItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bocconcino"}

	// rioJapanesePhrase represents rio japanese phrase.
	rioJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "リン"}

	// rioKoreanPhrase represents rio korean phrase.
	rioKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// rioLatinAmericanSpanishPhrase represents rio latin american spanish phrase.
	rioLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// rioRussianPhrase represents rio russian phrase.
	rioRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// rioSimplifiedChinesePhrase represents rio simplified chinese phrase.
	rioSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "铃"}

	// rioSpanishPhrase represents rio spanish phrase.
	rioSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "sambasamba"}

	// rioTraditionalChinesePhrase represents rio traditional chinese phrase.
	rioTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// rioPhrase represents rio phrase.
	rioPhrase = nook.Languages{
		language.AmericanEnglish:      rioAmericanEnglishPhrase,
		language.CanadianFrench:       rioCanadianFrenchPhrase,
		language.Dutch:                rioDutchPhrase,
		language.French:               rioFrenchPhrase,
		language.German:               rioGermanPhrase,
		language.Italian:              rioItalianPhrase,
		language.Japanese:             rioJapanesePhrase,
		language.Korean:               rioKoreanPhrase,
		language.LatinAmericanSpanish: rioLatinAmericanSpanishPhrase,
		language.Russian:              rioRussianPhrase,
		language.SimplifiedChinese:    rioSimplifiedChinesePhrase,
		language.Spanish:              rioSpanishPhrase,
		language.TraditionalChinese:   rioTraditionalChinesePhrase}
)

var (
	// Rio represents rio.
	Rio = nook.Villager{
		Character:   rioCharacter,
		Personality: personality.Peppy,
		Phrase:      rioPhrase}
)
