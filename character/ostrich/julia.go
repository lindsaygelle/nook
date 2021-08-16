package ostrich

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
	juliaBirthday = nook.Birthday{
		Day:   31,
		Month: time.July}
)

var (
	juliaCode = nook.Code{
		Value: "ost05"}
)

var (
	juliaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Julia"}

	juliaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Julie"}

	juliaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Julia"}

	juliaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Julie"}

	juliaGermanName = nook.Name{
		Language: language.German,
		Value:    "Julia"}

	juliaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Giulia"}

	juliaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジュリア"}

	juliaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "줄리아"}

	juliaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Julia"}

	juliaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Джулия"}

	juliaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "朱莉亚"}

	juliaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Julia"}

	juliaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "朱莉亞"}
)

var (
	juliaName = nook.Languages{
		language.AmericanEnglish:      juliaAmericanEnglishName,
		language.CanadianFrench:       juliaCanadianFrenchName,
		language.Dutch:                juliaDutchName,
		language.French:               juliaFrenchName,
		language.German:               juliaGermanName,
		language.Italian:              juliaItalianName,
		language.Japanese:             juliaJapaneseName,
		language.Korean:               juliaKoreanName,
		language.LatinAmericanSpanish: juliaLatinAmericanSpanishName,
		language.Russian:              juliaRussianName,
		language.SimplifiedChinese:    juliaSimplifiedChineseName,
		language.Spanish:              juliaSpanishName,
		language.TraditionalChinese:   juliaTraditionalChineseName}
)

var (
	juliaCharacter = nook.Character{
		Animal:   animal.Ostrich,
		Birthday: juliaBirthday,
		Code:     juliaCode,
		Key:      character.Julia,
		Gender:   gender.Female,
		Name:     juliaName}
)

var (
	juliaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "dahling"}

	juliaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "trutruche"}

	juliaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "schatje"}

	juliaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "trutruche"}

	juliaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "püh"}

	juliaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "blaaah"}

	juliaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "やだわね"}

	juliaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어머몰라"}

	juliaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "churri"}

	juliaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "дарлинг"}

	juliaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "吼唷"}

	juliaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "churri"}

	juliaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "吼唷"}
)

var (
	juliaPhrase = nook.Languages{
		language.AmericanEnglish:      juliaAmericanEnglishName,
		language.CanadianFrench:       juliaCanadianFrenchName,
		language.Dutch:                juliaDutchName,
		language.French:               juliaFrenchName,
		language.German:               juliaGermanName,
		language.Italian:              juliaItalianName,
		language.Japanese:             juliaJapaneseName,
		language.Korean:               juliaKoreanName,
		language.LatinAmericanSpanish: juliaLatinAmericanSpanishName,
		language.Russian:              juliaRussianName,
		language.SimplifiedChinese:    juliaSimplifiedChineseName,
		language.Spanish:              juliaSpanishName,
		language.TraditionalChinese:   juliaTraditionalChineseName}
)

var (
	Julia = nook.Villager{
		Character:   juliaCharacter,
		Personality: personality.Snooty,
		Phrase:      juliaPhrase}
)
