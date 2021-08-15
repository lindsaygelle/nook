package tiger

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	biancaBirthday = nook.Birthday{
		Day:   13,
		Month: time.December}
)

var (
	biancaCode = nook.Code{
		Value: "tig06"}
)

var (
	biancaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bianca"}

	biancaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Noémiela vie"}

	biancaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Biancawitje"}

	biancaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Noémiecroky"}

	biancaGermanName = nook.Name{
		Language: language.German,
		Value:    "Bettinaprrr"}

	biancaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Grattinayeppa"}

	biancaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "コユキでヒョウ"}

	biancaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "백희그래호"}

	biancaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Biancaruarri"}

	biancaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бьянкахлоп-хлоп"}

	biancaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小雪雪豹"}

	biancaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Biancaguiñito"}

	biancaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小雪雪豹"}
)

var (
	biancaName = nook.Languages{
		language.AmericanEnglish:      biancaAmericanEnglishName,
		language.CanadianFrench:       biancaCanadianFrenchName,
		language.Dutch:                biancaDutchName,
		language.French:               biancaFrenchName,
		language.German:               biancaGermanName,
		language.Italian:              biancaItalianName,
		language.Japanese:             biancaJapaneseName,
		language.Korean:               biancaKoreanName,
		language.LatinAmericanSpanish: biancaLatinAmericanSpanishName,
		language.Russian:              biancaRussianName,
		language.SimplifiedChinese:    biancaSimplifiedChineseName,
		language.Spanish:              biancaSpanishName,
		language.TraditionalChinese:   biancaTraditionalChineseName}
)

var (
	biancaCharacter = nook.Character{
		Animal:   Tiger,
		Birthday: biancaBirthday,
		Code:     biancaCode,
		Gender:   nook.Female,
		Name:     biancaName}
)

var (
	biancaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "glimmer"}

	biancaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "la vie"}

	biancaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "witje"}

	biancaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "croky"}

	biancaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "prrr"}

	biancaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "yeppa"}

	biancaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でヒョウ"}

	biancaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그래호"}

	biancaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ruarri"}

	biancaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хлоп-хлоп"}

	biancaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雪豹"}

	biancaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "guiñito"}

	biancaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雪豹"}
)

var (
	biancaPhrase = nook.Languages{
		language.AmericanEnglish:      biancaAmericanEnglishName,
		language.CanadianFrench:       biancaCanadianFrenchName,
		language.Dutch:                biancaDutchName,
		language.French:               biancaFrenchName,
		language.German:               biancaGermanName,
		language.Italian:              biancaItalianName,
		language.Japanese:             biancaJapaneseName,
		language.Korean:               biancaKoreanName,
		language.LatinAmericanSpanish: biancaLatinAmericanSpanishName,
		language.Russian:              biancaRussianName,
		language.SimplifiedChinese:    biancaSimplifiedChineseName,
		language.Spanish:              biancaSpanishName,
		language.TraditionalChinese:   biancaTraditionalChineseName}
)

var (
	Bianca = nook.Villager{
		Character:   biancaCharacter,
		Personality: nook.Peppy,
		Phrase:      biancaPhrase}
)
