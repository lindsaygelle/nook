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
	rioBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	rioCode = nook.Code{
		Value: ""}
)

var (
	rioAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rio"}

	rioCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	rioDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	rioFrenchName = nook.Name{
		Language: language.French,
		Value:    "Estrella"}

	rioGermanName = nook.Name{
		Language: language.German,
		Value:    "Verona"}

	rioItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rita"}

	rioJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "デジャネイロ"}

	rioKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	rioLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	rioRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	rioSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "欢欢"}

	rioSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rio"}

	rioTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
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
	rioAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "li'l chick"}

	rioCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	rioDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	rioFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bon bref"}

	rioGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "küken"}

	rioItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bocconcino"}

	rioJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "リン"}

	rioKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	rioLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	rioRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	rioSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "铃"}

	rioSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "sambasamba"}

	rioTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
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
	Rio = nook.Villager{
		Character:   rioCharacter,
		Personality: personality.Peppy,
		Phrase:      rioPhrase}
)
