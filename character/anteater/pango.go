package anteater

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	pangoBirthday = nook.Birthday{
		Day:   9,
		Month: time.November}
)

var (
	pangoCode = nook.Code{
		Value: "ant02"}
)

var (
	pangoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pango"}

	pangoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mathilda"}

	pangoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Pango"}

	pangoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mathilda"}

	pangoGermanName = nook.Name{
		Language: language.German,
		Value:    "Mathilda"}

	pangoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Carlotta"}

	pangoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パトラ"}

	pangoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "패트라"}

	pangoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Aspidora"}

	pangoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Панго"}

	pangoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "佩希"}

	pangoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Aspidora"}

	pangoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "佩希"}
)

var (
	pangoName = nook.Languages{
		language.AmericanEnglish:      pangoAmericanEnglishName,
		language.CanadianFrench:       pangoCanadianFrenchName,
		language.Dutch:                pangoDutchName,
		language.French:               pangoFrenchName,
		language.German:               pangoGermanName,
		language.Italian:              pangoItalianName,
		language.Japanese:             pangoJapaneseName,
		language.Korean:               pangoKoreanName,
		language.LatinAmericanSpanish: pangoLatinAmericanSpanishName,
		language.Russian:              pangoRussianName,
		language.SimplifiedChinese:    pangoSimplifiedChineseName,
		language.Spanish:              pangoSpanishName,
		language.TraditionalChinese:   pangoTraditionalChineseName}
)

var (
	pangoCharacter = nook.Character{
		Animal:   animal.Anteater,
		Birthday: pangoBirthday,
		Code:     pangoCode,
		Gender:   gender.Female,
		Name:     pangoName}
)

var (
	pangoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snooooof"}

	pangoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pouuuuuf"}

	pangoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snufffff"}

	pangoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pouuuuuf"}

	pangoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnieef"}

	pangoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "snuuf"}

	pangoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だっしー"}

	pangoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "라지요"}

	pangoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "snuf-snuf"}

	pangoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "снуф-снуф"}

	pangoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "希希"}

	pangoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "snuf-snuf"}

	pangoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "希希"}
)

var (
	pangoPhrase = nook.Languages{
		language.AmericanEnglish:      pangoAmericanEnglishName,
		language.CanadianFrench:       pangoCanadianFrenchName,
		language.Dutch:                pangoDutchName,
		language.French:               pangoFrenchName,
		language.German:               pangoGermanName,
		language.Italian:              pangoItalianName,
		language.Japanese:             pangoJapaneseName,
		language.Korean:               pangoKoreanName,
		language.LatinAmericanSpanish: pangoLatinAmericanSpanishName,
		language.Russian:              pangoRussianName,
		language.SimplifiedChinese:    pangoSimplifiedChineseName,
		language.Spanish:              pangoSpanishName,
		language.TraditionalChinese:   pangoTraditionalChineseName}
)

var (
	Pango = nook.Villager{
		Character:   pangoCharacter,
		Personality: personality.Peppy,
		Phrase:      pangoPhrase}
)