package bird

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
	flashBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	flashCode = nook.Code{
		Value: ""}
)

var (
	flashAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Flash"}

	flashCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	flashDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	flashFrenchName = nook.Name{
		Language: language.French,
		Value:    "Maurice"}

	flashGermanName = nook.Name{
		Language: language.German,
		Value:    "Flo"}

	flashItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Baleno"}

	flashJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "みちる"}

	flashKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	flashLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	flashRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	flashSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	flashSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Mauro"}

	flashTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	flashName = nook.Languages{
		language.AmericanEnglish:      flashAmericanEnglishName,
		language.CanadianFrench:       flashCanadianFrenchName,
		language.Dutch:                flashDutchName,
		language.French:               flashFrenchName,
		language.German:               flashGermanName,
		language.Italian:              flashItalianName,
		language.Japanese:             flashJapaneseName,
		language.Korean:               flashKoreanName,
		language.LatinAmericanSpanish: flashLatinAmericanSpanishName,
		language.Russian:              flashRussianName,
		language.SimplifiedChinese:    flashSimplifiedChineseName,
		language.Spanish:              flashSpanishName,
		language.TraditionalChinese:   flashTraditionalChineseName}
)

var (
	flashCharacter = nook.Character{
		Animal:   animal.Bird,
		Birthday: flashBirthday,
		Code:     flashCode,
		Key:      character.Flash,
		Gender:   gender.Male,
		Name:     flashName}
)

var (
	flashAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "babe"}

	flashCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	flashDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	flashFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "cuicui"}

	flashGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schätzchen"}

	flashItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "flash"}

	flashJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "チルチル"}

	flashKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	flashLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	flashRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	flashSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	flashSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cari"}

	flashTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	flashPhrase = nook.Languages{
		language.AmericanEnglish:      flashAmericanEnglishPhrase,
		language.CanadianFrench:       flashCanadianFrenchPhrase,
		language.Dutch:                flashDutchPhrase,
		language.French:               flashFrenchPhrase,
		language.German:               flashGermanPhrase,
		language.Italian:              flashItalianPhrase,
		language.Japanese:             flashJapanesePhrase,
		language.Korean:               flashKoreanPhrase,
		language.LatinAmericanSpanish: flashLatinAmericanSpanishPhrase,
		language.Russian:              flashRussianPhrase,
		language.SimplifiedChinese:    flashSimplifiedChinesePhrase,
		language.Spanish:              flashSpanishPhrase,
		language.TraditionalChinese:   flashTraditionalChinesePhrase}
)

var (
	Flash = nook.Villager{
		Character:   flashCharacter,
		Personality: personality.Cranky,
		Phrase:      flashPhrase}
)
