package pig

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	sporkBirthday = nook.Birthday{
		Day:   3,
		Month: time.September}
)

var (
	sporkCode = nook.Code{
		Value: "pig05"}
)

var (
	sporkAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Spork Crackle"}

	sporkCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Justin"}

	sporkDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Crackle"}

	sporkFrenchName = nook.Name{
		Language: language.French,
		Value:    "Justin"}

	sporkGermanName = nook.Name{
		Language: language.German,
		Value:    "Schwarte"}

	sporkItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cicciolo"}

	sporkJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ポーク"}

	sporkKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "포크"}

	sporkLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Espork"}

	sporkRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Крэкл"}

	sporkSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "肉肉"}

	sporkSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Espork"}

	sporkTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "肉肉"}
)

var (
	sporkName = nook.Languages{
		language.AmericanEnglish:      sporkAmericanEnglishName,
		language.CanadianFrench:       sporkCanadianFrenchName,
		language.Dutch:                sporkDutchName,
		language.French:               sporkFrenchName,
		language.German:               sporkGermanName,
		language.Italian:              sporkItalianName,
		language.Japanese:             sporkJapaneseName,
		language.Korean:               sporkKoreanName,
		language.LatinAmericanSpanish: sporkLatinAmericanSpanishName,
		language.Russian:              sporkRussianName,
		language.SimplifiedChinese:    sporkSimplifiedChineseName,
		language.Spanish:              sporkSpanishName,
		language.TraditionalChinese:   sporkTraditionalChineseName}
)

var (
	sporkCharacter = nook.Character{
		Animal:   Pig,
		Birthday: sporkBirthday,
		Code:     sporkCode,
		Gender:   nook.Male,
		Name:     sporkName}
)

var (
	sporkAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snork"}

	sporkCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "diantre"}

	sporkDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snurk"}

	sporkFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "diantre"}

	sporkGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grunz"}

	sporkItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "snork"}

	sporkJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だブー"}

	sporkKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꼬르륵"}

	sporkLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "pork"}

	sporkRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрю"}

	sporkSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "噗"}

	sporkSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pork"}

	sporkTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "噗"}
)

var (
	sporkPhrase = nook.Languages{
		language.AmericanEnglish:      sporkAmericanEnglishName,
		language.CanadianFrench:       sporkCanadianFrenchName,
		language.Dutch:                sporkDutchName,
		language.French:               sporkFrenchName,
		language.German:               sporkGermanName,
		language.Italian:              sporkItalianName,
		language.Japanese:             sporkJapaneseName,
		language.Korean:               sporkKoreanName,
		language.LatinAmericanSpanish: sporkLatinAmericanSpanishName,
		language.Russian:              sporkRussianName,
		language.SimplifiedChinese:    sporkSimplifiedChineseName,
		language.Spanish:              sporkSpanishName,
		language.TraditionalChinese:   sporkTraditionalChineseName}
)

var (
	Spork = nook.Villager{
		Character:   sporkCharacter,
		Personality: nook.Lazy,
		Phrase:      sporkPhrase}
)
