package gorilla

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	rocketBirthday = nook.Birthday{
		Day:   14,
		Month: time.April}
)

var (
	rocketCode = nook.Code{
		Value: "gor09"}
)

var (
	rocketAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rocket"}

	rocketCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gertrudepatouche"}

	rocketDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rocketvroem"}

	rocketFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gertrudepatouche"}

	rocketGermanName = nook.Name{
		Language: language.German,
		Value:    "Katrinbrahaha"}

	rocketItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Kingauhatà"}

	rocketJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "４ごうそりゃっ"}

	rocketKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "4호끼얍"}

	rocketLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gloriauga-uga"}

	rocketRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рокетна взлет"}

	rocketSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿四煞啊"}

	rocketSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gloriauga-uga"}

	rocketTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿四煞啊"}
)

var (
	rocketName = nook.Languages{
		language.AmericanEnglish:      rocketAmericanEnglishName,
		language.CanadianFrench:       rocketCanadianFrenchName,
		language.Dutch:                rocketDutchName,
		language.French:               rocketFrenchName,
		language.German:               rocketGermanName,
		language.Italian:              rocketItalianName,
		language.Japanese:             rocketJapaneseName,
		language.Korean:               rocketKoreanName,
		language.LatinAmericanSpanish: rocketLatinAmericanSpanishName,
		language.Russian:              rocketRussianName,
		language.SimplifiedChinese:    rocketSimplifiedChineseName,
		language.Spanish:              rocketSpanishName,
		language.TraditionalChinese:   rocketTraditionalChineseName}
)

var (
	rocketCharacter = nook.Character{
		Animal:   Gorilla,
		Birthday: rocketBirthday,
		Code:     rocketCode,
		Gender:   nook.Female,
		Name:     rocketName}
)

var (
	rocketAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "vroom"}

	rocketCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "patouche"}

	rocketDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "vroem"}

	rocketFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "patouche"}

	rocketGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "brahaha"}

	rocketItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "uhatà"}

	rocketJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "そりゃっ"}

	rocketKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "끼얍"}

	rocketLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "uga-uga"}

	rocketRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "на взлет"}

	rocketSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "煞啊"}

	rocketSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "uga-uga"}

	rocketTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "煞啊"}
)

var (
	rocketPhrase = nook.Languages{
		language.AmericanEnglish:      rocketAmericanEnglishName,
		language.CanadianFrench:       rocketCanadianFrenchName,
		language.Dutch:                rocketDutchName,
		language.French:               rocketFrenchName,
		language.German:               rocketGermanName,
		language.Italian:              rocketItalianName,
		language.Japanese:             rocketJapaneseName,
		language.Korean:               rocketKoreanName,
		language.LatinAmericanSpanish: rocketLatinAmericanSpanishName,
		language.Russian:              rocketRussianName,
		language.SimplifiedChinese:    rocketSimplifiedChineseName,
		language.Spanish:              rocketSpanishName,
		language.TraditionalChinese:   rocketTraditionalChineseName}
)

var (
	Rocket = nook.Villager{
		Character:   rocketCharacter,
		Personality: nook.BigSister,
		Phrase:      rocketPhrase}
)
