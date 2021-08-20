package penguin

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
		Value:    "Laurie"}

	sprinkleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sprinkle"}

	sprinkleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Laurie"}

	sprinkleGermanName = nook.Name{
		Language: language.German,
		Value:    "Svenja"}

	sprinkleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Brina"}

	sprinkleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フラッペ"}

	sprinkleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "크리미"}

	sprinkleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rizolda"}

	sprinkleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Спринкл"}

	sprinkleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "冰莎"}

	sprinkleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rizolda"}

	sprinkleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "冰莎"}
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
		Animal:   animal.Penguin,
		Birthday: sprinkleBirthday,
		Code:     sprinkleCode,
		Key:      character.Sprinkle,
		Gender:   gender.Female,
		Name:     sprinkleName,
		Special:  false}
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
		language.AmericanEnglish:      sprinkleAmericanEnglishPhrase,
		language.CanadianFrench:       sprinkleCanadianFrenchPhrase,
		language.Dutch:                sprinkleDutchPhrase,
		language.French:               sprinkleFrenchPhrase,
		language.German:               sprinkleGermanPhrase,
		language.Italian:              sprinkleItalianPhrase,
		language.Japanese:             sprinkleJapanesePhrase,
		language.Korean:               sprinkleKoreanPhrase,
		language.LatinAmericanSpanish: sprinkleLatinAmericanSpanishPhrase,
		language.Russian:              sprinkleRussianPhrase,
		language.SimplifiedChinese:    sprinkleSimplifiedChinesePhrase,
		language.Spanish:              sprinkleSpanishPhrase,
		language.TraditionalChinese:   sprinkleTraditionalChinesePhrase}
)

var (
	Sprinkle = nook.Villager{
		Character:   sprinkleCharacter,
		Personality: personality.Peppy,
		Phrase:      sprinklePhrase}
)
