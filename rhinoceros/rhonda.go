package rhinoceros

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	rhondaBirthday = nook.Birthday{
		Day:   24,
		Month: time.January}
)

var (
	rhondaCode = nook.Code{
		Value: "rhn01"}
)

var (
	rhondaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rhonda"}

	rhondaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Renéegrand pied"}

	rhondaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rhondagrootvoet"}

	rhondaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Renéegrand pied"}

	rhondaGermanName = nook.Name{
		Language: language.German,
		Value:    "Reginastampf"}

	rhondaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Robypiedone"}

	rhondaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ユメコンフ"}

	rhondaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "론다훗"}

	rhondaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rondajip-jip"}

	rhondaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рондабосс"}

	rhondaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "梦梦嗯噗"}

	rhondaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rondajip-jip"}

	rhondaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "夢夢嗯噗"}
)

var (
	rhondaName = nook.Languages{
		language.AmericanEnglish:      rhondaAmericanEnglishName,
		language.CanadianFrench:       rhondaCanadianFrenchName,
		language.Dutch:                rhondaDutchName,
		language.French:               rhondaFrenchName,
		language.German:               rhondaGermanName,
		language.Italian:              rhondaItalianName,
		language.Japanese:             rhondaJapaneseName,
		language.Korean:               rhondaKoreanName,
		language.LatinAmericanSpanish: rhondaLatinAmericanSpanishName,
		language.Russian:              rhondaRussianName,
		language.SimplifiedChinese:    rhondaSimplifiedChineseName,
		language.Spanish:              rhondaSpanishName,
		language.TraditionalChinese:   rhondaTraditionalChineseName}
)

var (
	rhondaCharacter = nook.Character{
		Animal:   Rhinoceros,
		Birthday: rhondaBirthday,
		Code:     rhondaCode,
		Gender:   nook.Female,
		Name:     rhondaName}
)

var (
	rhondaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bigfoot"}

	rhondaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grand pied"}

	rhondaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "grootvoet"}

	rhondaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grand pied"}

	rhondaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "stampf"}

	rhondaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "piedone"}

	rhondaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ンフ"}

	rhondaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "훗"}

	rhondaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "jip-jip"}

	rhondaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "босс"}

	rhondaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嗯噗"}

	rhondaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "jip-jip"}

	rhondaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗯噗"}
)

var (
	rhondaPhrase = nook.Languages{
		language.AmericanEnglish:      rhondaAmericanEnglishName,
		language.CanadianFrench:       rhondaCanadianFrenchName,
		language.Dutch:                rhondaDutchName,
		language.French:               rhondaFrenchName,
		language.German:               rhondaGermanName,
		language.Italian:              rhondaItalianName,
		language.Japanese:             rhondaJapaneseName,
		language.Korean:               rhondaKoreanName,
		language.LatinAmericanSpanish: rhondaLatinAmericanSpanishName,
		language.Russian:              rhondaRussianName,
		language.SimplifiedChinese:    rhondaSimplifiedChineseName,
		language.Spanish:              rhondaSpanishName,
		language.TraditionalChinese:   rhondaTraditionalChineseName}
)

var (
	Rhonda = nook.Villager{
		Character:   rhondaCharacter,
		Personality: nook.Normal,
		Phrase:      rhondaPhrase}
)
