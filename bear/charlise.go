package bear

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	charliseBirthday = nook.Birthday{
		Day:   17,
		Month: time.April}
)

var (
	charliseCode = nook.Code{
		Value: "bea12"}
)

var (
	charliseAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Charlise"}

	charliseCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Zabouchacha"}

	charliseDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Charlisepff"}

	charliseFrenchName = nook.Name{
		Language: language.French,
		Value:    "Zabouverlue"}

	charliseGermanName = nook.Name{
		Language: language.German,
		Value:    "Tabeabärbär"}

	charliseItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Orsolasbuff"}

	charliseJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チャーミーよいしょ"}

	charliseKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "챠미으랏차차"}

	charliseLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Charourrrgh"}

	charliseRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Шарлизого-го"}

	charliseSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "恰咪唷咻"}

	charliseSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Charocomolooyes"}

	charliseTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "恰咪唷咻"}
)

var (
	charliseName = nook.Languages{
		language.AmericanEnglish:      charliseAmericanEnglishName,
		language.CanadianFrench:       charliseCanadianFrenchName,
		language.Dutch:                charliseDutchName,
		language.French:               charliseFrenchName,
		language.German:               charliseGermanName,
		language.Italian:              charliseItalianName,
		language.Japanese:             charliseJapaneseName,
		language.Korean:               charliseKoreanName,
		language.LatinAmericanSpanish: charliseLatinAmericanSpanishName,
		language.Russian:              charliseRussianName,
		language.SimplifiedChinese:    charliseSimplifiedChineseName,
		language.Spanish:              charliseSpanishName,
		language.TraditionalChinese:   charliseTraditionalChineseName}
)

var (
	charliseCharacter = nook.Character{
		Animal:   Bear,
		Birthday: charliseBirthday,
		Code:     charliseCode,
		Gender:   nook.Female,
		Name:     charliseName}
)

var (
	charliseAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "urgh"}

	charliseCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chacha"}

	charliseDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pff"}

	charliseFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "verlue"}

	charliseGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bärbär"}

	charliseItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sbuff"}

	charliseJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "よいしょ"}

	charliseKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "으랏차차"}

	charliseLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "urrrgh"}

	charliseRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ого-го"}

	charliseSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "唷咻"}

	charliseSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "comolooyes"}

	charliseTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "唷咻"}
)

var (
	charlisePhrase = nook.Languages{
		language.AmericanEnglish:      charliseAmericanEnglishName,
		language.CanadianFrench:       charliseCanadianFrenchName,
		language.Dutch:                charliseDutchName,
		language.French:               charliseFrenchName,
		language.German:               charliseGermanName,
		language.Italian:              charliseItalianName,
		language.Japanese:             charliseJapaneseName,
		language.Korean:               charliseKoreanName,
		language.LatinAmericanSpanish: charliseLatinAmericanSpanishName,
		language.Russian:              charliseRussianName,
		language.SimplifiedChinese:    charliseSimplifiedChineseName,
		language.Spanish:              charliseSpanishName,
		language.TraditionalChinese:   charliseTraditionalChineseName}
)

var (
	Charlise = nook.Villager{
		Character:   charliseCharacter,
		Personality: nook.BigSister,
		Phrase:      charlisePhrase}
)
