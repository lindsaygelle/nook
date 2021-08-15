package ostrich

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Estrellabon bref"}

	rioGermanName = nook.Name{
		Language: language.German,
		Value:    "Veronaküken"}

	rioItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ritabocconcino"}

	rioJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "デジャネイロリン"}

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
		Value:    "欢欢铃"}

	rioSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Riosambasamba"}

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
		Animal:   Ostrich,
		Birthday: rioBirthday,
		Code:     rioCode,
		Gender:   nook.Female,
		Name:     rioName}
)

var (
	rioAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	rioCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	rioDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

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
		Value:    "N/A"}

	rioLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	rioRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	rioSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "铃"}

	rioSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "sambasamba"}

	rioTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	rioPhrase = nook.Languages{
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
	Rio = nook.Villager{
		Character:   rioCharacter,
		Personality: nook.Peppy,
		Phrase:      rioPhrase}
)
