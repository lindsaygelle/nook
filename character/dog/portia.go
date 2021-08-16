package dog

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
	portiaBirthday = nook.Birthday{
		Day:   25,
		Month: time.October}
)

var (
	portiaCode = nook.Code{
		Value: "dog05"}
)

var (
	portiaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Portia"}

	portiaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Dalma"}

	portiaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Portia"}

	portiaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Dalma"}

	portiaGermanName = nook.Name{
		Language: language.German,
		Value:    "Isolde"}

	portiaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Porzia"}

	portiaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ブレンダ"}

	portiaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "블랜더"}

	portiaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Amanda"}

	portiaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Порция"}

	portiaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "傅兰"}

	portiaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Amanda"}

	portiaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "傅蘭"}
)

var (
	portiaName = nook.Languages{
		language.AmericanEnglish:      portiaAmericanEnglishName,
		language.CanadianFrench:       portiaCanadianFrenchName,
		language.Dutch:                portiaDutchName,
		language.French:               portiaFrenchName,
		language.German:               portiaGermanName,
		language.Italian:              portiaItalianName,
		language.Japanese:             portiaJapaneseName,
		language.Korean:               portiaKoreanName,
		language.LatinAmericanSpanish: portiaLatinAmericanSpanishName,
		language.Russian:              portiaRussianName,
		language.SimplifiedChinese:    portiaSimplifiedChineseName,
		language.Spanish:              portiaSpanishName,
		language.TraditionalChinese:   portiaTraditionalChineseName}
)

var (
	portiaCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: portiaBirthday,
		Code:     portiaCode,
		Key:      character.Portia,
		Gender:   gender.Female,
		Name:     portiaName}
)

var (
	portiaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "ruffian"}

	portiaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "brute"}

	portiaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "keffer"}

	portiaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "brute"}

	portiaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "frechi"}

	portiaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ronfietto"}

	portiaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "フンッ"}

	portiaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "흥"}

	portiaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "ruf-ruf"}

	portiaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "тяв-ряв"}

	portiaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "屋屋"}

	portiaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "ruf-ruf"}

	portiaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "屋屋"}
)

var (
	portiaPhrase = nook.Languages{
		language.AmericanEnglish:      portiaAmericanEnglishName,
		language.CanadianFrench:       portiaCanadianFrenchName,
		language.Dutch:                portiaDutchName,
		language.French:               portiaFrenchName,
		language.German:               portiaGermanName,
		language.Italian:              portiaItalianName,
		language.Japanese:             portiaJapaneseName,
		language.Korean:               portiaKoreanName,
		language.LatinAmericanSpanish: portiaLatinAmericanSpanishName,
		language.Russian:              portiaRussianName,
		language.SimplifiedChinese:    portiaSimplifiedChineseName,
		language.Spanish:              portiaSpanishName,
		language.TraditionalChinese:   portiaTraditionalChineseName}
)

var (
	Portia = nook.Villager{
		Character:   portiaCharacter,
		Personality: personality.Snooty,
		Phrase:      portiaPhrase}
)
