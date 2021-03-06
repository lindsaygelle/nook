package lion

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
		Value:    "Léo"}

	rexDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rex"}

	rexFrenchName = nook.Name{
		Language: language.French,
		Value:    "Léo"}

	rexGermanName = nook.Name{
		Language: language.German,
		Value:    "Rex"}

	rexItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rex"}

	rexJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サンデー"}

	rexKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "렉스"}

	rexLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Leoncio"}

	rexRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рекс"}

	rexSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "晴天"}

	rexSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Leoncio"}

	rexTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "晴天"}
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
		Animal:   animal.Lion,
		Birthday: rexBirthday,
		Code:     rexCode,
		Key:      character.Rex,
		Gender:   gender.Male,
		Name:     rexName,
		Special:  false}
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
		language.AmericanEnglish:      rexAmericanEnglishPhrase,
		language.CanadianFrench:       rexCanadianFrenchPhrase,
		language.Dutch:                rexDutchPhrase,
		language.French:               rexFrenchPhrase,
		language.German:               rexGermanPhrase,
		language.Italian:              rexItalianPhrase,
		language.Japanese:             rexJapanesePhrase,
		language.Korean:               rexKoreanPhrase,
		language.LatinAmericanSpanish: rexLatinAmericanSpanishPhrase,
		language.Russian:              rexRussianPhrase,
		language.SimplifiedChinese:    rexSimplifiedChinesePhrase,
		language.Spanish:              rexSpanishPhrase,
		language.TraditionalChinese:   rexTraditionalChinesePhrase}
)

var (
	Rex = nook.Villager{
		Character:   rexCharacter,
		Personality: personality.Lazy,
		Phrase:      rexPhrase}
)
