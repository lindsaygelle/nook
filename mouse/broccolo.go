package mouse

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	broccoloBirthday = nook.Birthday{
		Day:   30,
		Month: time.June}
)

var (
	broccoloCode = nook.Code{
		Value: "mus12"}
)

var (
	broccoloAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Broccolo"}

	broccoloCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Stevenflim-flam"}

	broccoloDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Broccolosmakelijk"}

	broccoloFrenchName = nook.Name{
		Language: language.French,
		Value:    "Stevenflim-flam"}

	broccoloGermanName = nook.Name{
		Language: language.German,
		Value:    "Fritziknabba"}

	broccoloItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Francosquikko"}

	broccoloJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ブロッコリーピコ"}

	broccoloKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "브로콜리앗차"}

	broccoloLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Brócoliflinflán"}

	broccoloRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Брокколоням-ням"}

	broccoloSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "茜蓝微微"}

	broccoloSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Brócolidientes"}

	broccoloTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "茜藍微微"}
)

var (
	broccoloName = nook.Languages{
		language.AmericanEnglish:      broccoloAmericanEnglishName,
		language.CanadianFrench:       broccoloCanadianFrenchName,
		language.Dutch:                broccoloDutchName,
		language.French:               broccoloFrenchName,
		language.German:               broccoloGermanName,
		language.Italian:              broccoloItalianName,
		language.Japanese:             broccoloJapaneseName,
		language.Korean:               broccoloKoreanName,
		language.LatinAmericanSpanish: broccoloLatinAmericanSpanishName,
		language.Russian:              broccoloRussianName,
		language.SimplifiedChinese:    broccoloSimplifiedChineseName,
		language.Spanish:              broccoloSpanishName,
		language.TraditionalChinese:   broccoloTraditionalChineseName}
)

var (
	broccoloCharacter = nook.Character{
		Animal:   Mouse,
		Birthday: broccoloBirthday,
		Code:     broccoloCode,
		Gender:   nook.Male,
		Name:     broccoloName}
)

var (
	broccoloAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "eat it"}

	broccoloCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "flim-flam"}

	broccoloDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "smakelijk"}

	broccoloFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "flim-flam"}

	broccoloGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knabba"}

	broccoloItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "squikko"}

	broccoloJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ピコ"}

	broccoloKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "앗차"}

	broccoloLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "flinflán"}

	broccoloRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ням-ням"}

	broccoloSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "微微"}

	broccoloSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "dientes"}

	broccoloTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "微微"}
)

var (
	broccoloPhrase = nook.Languages{
		language.AmericanEnglish:      broccoloAmericanEnglishName,
		language.CanadianFrench:       broccoloCanadianFrenchName,
		language.Dutch:                broccoloDutchName,
		language.French:               broccoloFrenchName,
		language.German:               broccoloGermanName,
		language.Italian:              broccoloItalianName,
		language.Japanese:             broccoloJapaneseName,
		language.Korean:               broccoloKoreanName,
		language.LatinAmericanSpanish: broccoloLatinAmericanSpanishName,
		language.Russian:              broccoloRussianName,
		language.SimplifiedChinese:    broccoloSimplifiedChineseName,
		language.Spanish:              broccoloSpanishName,
		language.TraditionalChinese:   broccoloTraditionalChineseName}
)

var (
	Broccolo = nook.Villager{
		Character:   broccoloCharacter,
		Personality: nook.Lazy,
		Phrase:      broccoloPhrase}
)
