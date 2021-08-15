package lion

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	rexBirthday = nook.Birthday{
		Day:   24,
		Month: time.July}
)

var (
	rexCode = nook.Code{
		Value: "lon02"}
)

var (
	rexAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rex"}

	rexCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Léolionceau"}

	rexDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rexcoole kat"}

	rexFrenchName = nook.Name{
		Language: language.French,
		Value:    "Léolionceau"}

	rexGermanName = nook.Name{
		Language: language.German,
		Value:    "Rexcoolcat"}

	rexItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rexgrrringo"}

	rexJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サンデーだオン"}

	rexKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "렉스티라이온"}

	rexLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Leonciogrrreh"}

	rexRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рексспокойно"}

	rexSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "晴天狮狮"}

	rexSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Leonciomelenas"}

	rexTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "晴天獅獅"}
)

var (
	rexName = nook.Languages{
		language.AmericanEnglish:      rexAmericanEnglishName,
		language.CanadianFrench:       rexCanadianFrenchName,
		language.Dutch:                rexDutchName,
		language.French:               rexFrenchName,
		language.German:               rexGermanName,
		language.Italian:              rexItalianName,
		language.Japanese:             rexJapaneseName,
		language.Korean:               rexKoreanName,
		language.LatinAmericanSpanish: rexLatinAmericanSpanishName,
		language.Russian:              rexRussianName,
		language.SimplifiedChinese:    rexSimplifiedChineseName,
		language.Spanish:              rexSpanishName,
		language.TraditionalChinese:   rexTraditionalChineseName}
)

var (
	rexCharacter = nook.Character{
		Animal:   Lion,
		Birthday: rexBirthday,
		Code:     rexCode,
		Gender:   nook.Male,
		Name:     rexName}
)

var (
	rexAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cool cat"}

	rexCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "lionceau"}

	rexDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "coole kat"}

	rexFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "lionceau"}

	rexGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "coolcat"}

	rexItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grrringo"}

	rexJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だオン"}

	rexKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "티라이온"}

	rexLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grrreh"}

	rexRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "спокойно"}

	rexSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "狮狮"}

	rexSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "melenas"}

	rexTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "獅獅"}
)

var (
	rexPhrase = nook.Languages{
		language.AmericanEnglish:      rexAmericanEnglishName,
		language.CanadianFrench:       rexCanadianFrenchName,
		language.Dutch:                rexDutchName,
		language.French:               rexFrenchName,
		language.German:               rexGermanName,
		language.Italian:              rexItalianName,
		language.Japanese:             rexJapaneseName,
		language.Korean:               rexKoreanName,
		language.LatinAmericanSpanish: rexLatinAmericanSpanishName,
		language.Russian:              rexRussianName,
		language.SimplifiedChinese:    rexSimplifiedChineseName,
		language.Spanish:              rexSpanishName,
		language.TraditionalChinese:   rexTraditionalChineseName}
)

var (
	Rex = nook.Villager{
		Character:   rexCharacter,
		Personality: nook.Lazy,
		Phrase:      rexPhrase}
)
