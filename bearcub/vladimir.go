package bearcub

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	vladimirBirthday = nook.Birthday{
		Day:   2,
		Month: time.August}
)

var (
	vladimirCode = nook.Code{
		Value: "cbr06"}
)

var (
	vladimirAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Vladimir"}

	vladimirCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Vladimirniet"}

	vladimirDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Vladimirnjet"}

	vladimirFrenchName = nook.Name{
		Language: language.French,
		Value:    "Vladimirniet"}

	vladimirGermanName = nook.Name{
		Language: language.German,
		Value:    "Vladimirnastrovje"}

	vladimirItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Vladimirnyet"}

	vladimirJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ガビやんけ"}

	vladimirKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "곰비참말로"}

	vladimirLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Vladimirnyet"}

	vladimirRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Владимирну не-е-ет"}

	vladimirSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘉弼唉唷"}

	vladimirSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Vladimirnyet"}

	vladimirTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘉弼唉唷"}
)

var (
	vladimirName = nook.Languages{
		language.AmericanEnglish:      vladimirAmericanEnglishName,
		language.CanadianFrench:       vladimirCanadianFrenchName,
		language.Dutch:                vladimirDutchName,
		language.French:               vladimirFrenchName,
		language.German:               vladimirGermanName,
		language.Italian:              vladimirItalianName,
		language.Japanese:             vladimirJapaneseName,
		language.Korean:               vladimirKoreanName,
		language.LatinAmericanSpanish: vladimirLatinAmericanSpanishName,
		language.Russian:              vladimirRussianName,
		language.SimplifiedChinese:    vladimirSimplifiedChineseName,
		language.Spanish:              vladimirSpanishName,
		language.TraditionalChinese:   vladimirTraditionalChineseName}
)

var (
	vladimirCharacter = nook.Character{
		Animal:   Bearcub,
		Birthday: vladimirBirthday,
		Code:     vladimirCode,
		Gender:   nook.Male,
		Name:     vladimirName}
)

var (
	vladimirAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nyet"}

	vladimirCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "niet"}

	vladimirDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "njet"}

	vladimirFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "niet"}

	vladimirGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "nastrovje"}

	vladimirItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "nyet"}

	vladimirJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やんけ"}

	vladimirKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "참말로"}

	vladimirLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "nyet"}

	vladimirRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ну не-е-ет"}

	vladimirSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唉唷"}

	vladimirSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "nyet"}

	vladimirTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唉唷"}
)

var (
	vladimirPhrase = nook.Languages{
		language.AmericanEnglish:      vladimirAmericanEnglishName,
		language.CanadianFrench:       vladimirCanadianFrenchName,
		language.Dutch:                vladimirDutchName,
		language.French:               vladimirFrenchName,
		language.German:               vladimirGermanName,
		language.Italian:              vladimirItalianName,
		language.Japanese:             vladimirJapaneseName,
		language.Korean:               vladimirKoreanName,
		language.LatinAmericanSpanish: vladimirLatinAmericanSpanishName,
		language.Russian:              vladimirRussianName,
		language.SimplifiedChinese:    vladimirSimplifiedChineseName,
		language.Spanish:              vladimirSpanishName,
		language.TraditionalChinese:   vladimirTraditionalChineseName}
)

var (
	Vladimir = nook.Villager{
		Character:   vladimirCharacter,
		Personality: nook.Cranky,
		Phrase:      vladimirPhrase}
)
