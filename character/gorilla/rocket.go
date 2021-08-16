package gorilla

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
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
		Value:    "Gertrude"}

	rocketDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rocket"}

	rocketFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gertrude"}

	rocketGermanName = nook.Name{
		Language: language.German,
		Value:    "Katrin"}

	rocketItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Kinga"}

	rocketJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "４ごう"}

	rocketKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "4호"}

	rocketLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Gloria"}

	rocketRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рокет"}

	rocketSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿四"}

	rocketSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Gloria"}

	rocketTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿四"}
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
		Animal:   animal.Gorilla,
		Birthday: rocketBirthday,
		Code:     rocketCode,
		Gender:   gender.Female,
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
		Personality: personality.BigSister,
		Phrase:      rocketPhrase}
)
