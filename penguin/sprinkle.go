package penguin

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	sprinkleBirthday = nook.Birthday{
		Day:   20,
		Month: time.February}
)

var (
	sprinkleCode = nook.Code{
		Value: "pgn14"}
)

var (
	sprinkleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sprinkle"}

	sprinkleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lauriemon glaçon"}

	sprinkleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sprinklefrappé"}

	sprinkleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lauriemon glaçon"}

	sprinkleGermanName = nook.Name{
		Language: language.German,
		Value:    "Svenjabibib"}

	sprinkleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Brinaunz unz"}

	sprinkleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フラッペヒヤリ"}

	sprinkleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "크리미꽁꽁"}

	sprinkleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rizoldafrapé"}

	sprinkleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Спринклпенка"}

	sprinkleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "冰莎凉爽"}

	sprinkleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rizoldaesmoquin"}

	sprinkleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "冰莎涼爽"}
)

var (
	sprinkleName = nook.Languages{
		language.AmericanEnglish:      sprinkleAmericanEnglishName,
		language.CanadianFrench:       sprinkleCanadianFrenchName,
		language.Dutch:                sprinkleDutchName,
		language.French:               sprinkleFrenchName,
		language.German:               sprinkleGermanName,
		language.Italian:              sprinkleItalianName,
		language.Japanese:             sprinkleJapaneseName,
		language.Korean:               sprinkleKoreanName,
		language.LatinAmericanSpanish: sprinkleLatinAmericanSpanishName,
		language.Russian:              sprinkleRussianName,
		language.SimplifiedChinese:    sprinkleSimplifiedChineseName,
		language.Spanish:              sprinkleSpanishName,
		language.TraditionalChinese:   sprinkleTraditionalChineseName}
)

var (
	sprinkleCharacter = nook.Character{
		Animal:   Penguin,
		Birthday: sprinkleBirthday,
		Code:     sprinkleCode,
		Gender:   nook.Female,
		Name:     sprinkleName}
)

var (
	sprinkleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "frappe"}

	sprinkleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon glaçon"}

	sprinkleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "frappé"}

	sprinkleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon glaçon"}

	sprinkleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bibib"}

	sprinkleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "unz unz"}

	sprinkleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヒヤリ"}

	sprinkleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꽁꽁"}

	sprinkleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "frapé"}

	sprinkleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "пенка"}

	sprinkleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "凉爽"}

	sprinkleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "esmoquin"}

	sprinkleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "涼爽"}
)

var (
	sprinklePhrase = nook.Languages{
		language.AmericanEnglish:      sprinkleAmericanEnglishName,
		language.CanadianFrench:       sprinkleCanadianFrenchName,
		language.Dutch:                sprinkleDutchName,
		language.French:               sprinkleFrenchName,
		language.German:               sprinkleGermanName,
		language.Italian:              sprinkleItalianName,
		language.Japanese:             sprinkleJapaneseName,
		language.Korean:               sprinkleKoreanName,
		language.LatinAmericanSpanish: sprinkleLatinAmericanSpanishName,
		language.Russian:              sprinkleRussianName,
		language.SimplifiedChinese:    sprinkleSimplifiedChineseName,
		language.Spanish:              sprinkleSpanishName,
		language.TraditionalChinese:   sprinkleTraditionalChineseName}
)

var (
	Sprinkle = nook.Villager{
		Character:   sprinkleCharacter,
		Personality: nook.Peppy,
		Phrase:      sprinklePhrase}
)
