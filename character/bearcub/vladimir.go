package bearcub

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
		Value:    "Vladimir"}

	vladimirDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Vladimir"}

	vladimirFrenchName = nook.Name{
		Language: language.French,
		Value:    "Vladimir"}

	vladimirGermanName = nook.Name{
		Language: language.German,
		Value:    "Vladimir"}

	vladimirItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Vladimir"}

	vladimirJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ガビ"}

	vladimirKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "곰비"}

	vladimirLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Vladimir"}

	vladimirRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Владимир"}

	vladimirSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘉弼"}

	vladimirSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Vladimir"}

	vladimirTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘉弼"}
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
		Animal:   animal.Bearcub,
		Birthday: vladimirBirthday,
		Code:     vladimirCode,
		Key:      character.Vladimir,
		Gender:   gender.Male,
		Name:     vladimirName,
		Special:  false}
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
		language.AmericanEnglish:      vladimirAmericanEnglishPhrase,
		language.CanadianFrench:       vladimirCanadianFrenchPhrase,
		language.Dutch:                vladimirDutchPhrase,
		language.French:               vladimirFrenchPhrase,
		language.German:               vladimirGermanPhrase,
		language.Italian:              vladimirItalianPhrase,
		language.Japanese:             vladimirJapanesePhrase,
		language.Korean:               vladimirKoreanPhrase,
		language.LatinAmericanSpanish: vladimirLatinAmericanSpanishPhrase,
		language.Russian:              vladimirRussianPhrase,
		language.SimplifiedChinese:    vladimirSimplifiedChinesePhrase,
		language.Spanish:              vladimirSpanishPhrase,
		language.TraditionalChinese:   vladimirTraditionalChinesePhrase}
)

var (
	Vladimir = nook.Villager{
		Character:   vladimirCharacter,
		Personality: personality.Cranky,
		Phrase:      vladimirPhrase}
)
