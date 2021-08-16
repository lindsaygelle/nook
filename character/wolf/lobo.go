package wolf

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	loboBirthday = nook.Birthday{
		Day:   5,
		Month: time.November}
)

var (
	loboCode = nook.Code{
		Value: "wol01"}
)

var (
	loboAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lobo"}

	loboCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lobo"}

	loboDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lobo"}

	loboFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lobo"}

	loboGermanName = nook.Name{
		Language: language.German,
		Value:    "Lupo"}

	loboItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lupen"}

	loboJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ブンジロウ"}

	loboKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "늑태"}

	loboLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lupo"}

	loboRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лобо"}

	loboSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "潘奇隆"}

	loboSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lupo"}

	loboTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "潘奇隆"}
)

var (
	loboName = nook.Languages{
		language.AmericanEnglish:      loboAmericanEnglishName,
		language.CanadianFrench:       loboCanadianFrenchName,
		language.Dutch:                loboDutchName,
		language.French:               loboFrenchName,
		language.German:               loboGermanName,
		language.Italian:              loboItalianName,
		language.Japanese:             loboJapaneseName,
		language.Korean:               loboKoreanName,
		language.LatinAmericanSpanish: loboLatinAmericanSpanishName,
		language.Russian:              loboRussianName,
		language.SimplifiedChinese:    loboSimplifiedChineseName,
		language.Spanish:              loboSpanishName,
		language.TraditionalChinese:   loboTraditionalChineseName}
)

var (
	loboCharacter = nook.Character{
		Animal:   animal.Wolf,
		Birthday: loboBirthday,
		Code:     loboCode,
		Gender:   gender.Male,
		Name:     loboName}
)

var (
	loboAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ah-rooooo"}

	loboCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ah rooooo"}

	loboDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "aooooeee"}

	loboFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ah rooooo"}

	loboGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "auuuuuu"}

	loboItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "etchuuuu"}

	loboJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だぜよ"}

	loboKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "봐봐"}

	loboLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "auuuh"}

	loboRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ау-у-у"}

	loboSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "不然呢"}

	loboSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "auuuh"}

	loboTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "不然呢"}
)

var (
	loboPhrase = nook.Languages{
		language.AmericanEnglish:      loboAmericanEnglishName,
		language.CanadianFrench:       loboCanadianFrenchName,
		language.Dutch:                loboDutchName,
		language.French:               loboFrenchName,
		language.German:               loboGermanName,
		language.Italian:              loboItalianName,
		language.Japanese:             loboJapaneseName,
		language.Korean:               loboKoreanName,
		language.LatinAmericanSpanish: loboLatinAmericanSpanishName,
		language.Russian:              loboRussianName,
		language.SimplifiedChinese:    loboSimplifiedChineseName,
		language.Spanish:              loboSpanishName,
		language.TraditionalChinese:   loboTraditionalChineseName}
)

var (
	Lobo = nook.Villager{
		Character:   loboCharacter,
		Personality: personality.Cranky,
		Phrase:      loboPhrase}
)
