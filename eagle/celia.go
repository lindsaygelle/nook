package eagle

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	celiaBirthday = nook.Birthday{
		Day:   25,
		Month: time.March}
)

var (
	celiaCode = nook.Code{
		Value: "pbr09"}
)

var (
	celiaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Celia"}

	celiaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Garancebebec"}

	celiaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Celiapluimpje"}

	celiaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Garancebebec"}

	celiaGermanName = nook.Name{
		Language: language.German,
		Value:    "Lorakraaack"}

	celiaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Celiaah bon"}

	celiaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ティファニーそうね"}

	celiaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "티파니그래맞아"}

	celiaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Jazmínplumí"}

	celiaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Селияперышко"}

	celiaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "蒂芬妮是呢"}

	celiaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Jazmínplumis"}

	celiaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "蒂芬妮是呢"}
)

var (
	celiaName = nook.Languages{
		language.AmericanEnglish:      celiaAmericanEnglishName,
		language.CanadianFrench:       celiaCanadianFrenchName,
		language.Dutch:                celiaDutchName,
		language.French:               celiaFrenchName,
		language.German:               celiaGermanName,
		language.Italian:              celiaItalianName,
		language.Japanese:             celiaJapaneseName,
		language.Korean:               celiaKoreanName,
		language.LatinAmericanSpanish: celiaLatinAmericanSpanishName,
		language.Russian:              celiaRussianName,
		language.SimplifiedChinese:    celiaSimplifiedChineseName,
		language.Spanish:              celiaSpanishName,
		language.TraditionalChinese:   celiaTraditionalChineseName}
)

var (
	celiaCharacter = nook.Character{
		Animal:   Eagle,
		Birthday: celiaBirthday,
		Code:     celiaCode,
		Gender:   nook.Female,
		Name:     celiaName}
)

var (
	celiaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "feathers"}

	celiaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bebec"}

	celiaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pluimpje"}

	celiaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bebec"}

	celiaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kraaack"}

	celiaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ah bon"}

	celiaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "そうね"}

	celiaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그래맞아"}

	celiaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "plumí"}

	celiaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "перышко"}

	celiaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是呢"}

	celiaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "plumis"}

	celiaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是呢"}
)

var (
	celiaPhrase = nook.Languages{
		language.AmericanEnglish:      celiaAmericanEnglishName,
		language.CanadianFrench:       celiaCanadianFrenchName,
		language.Dutch:                celiaDutchName,
		language.French:               celiaFrenchName,
		language.German:               celiaGermanName,
		language.Italian:              celiaItalianName,
		language.Japanese:             celiaJapaneseName,
		language.Korean:               celiaKoreanName,
		language.LatinAmericanSpanish: celiaLatinAmericanSpanishName,
		language.Russian:              celiaRussianName,
		language.SimplifiedChinese:    celiaSimplifiedChineseName,
		language.Spanish:              celiaSpanishName,
		language.TraditionalChinese:   celiaTraditionalChineseName}
)

var (
	Celia = nook.Villager{
		Character:   celiaCharacter,
		Personality: nook.Normal,
		Phrase:      celiaPhrase}
)
