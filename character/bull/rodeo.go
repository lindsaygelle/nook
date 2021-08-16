package bull

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	rodeoBirthday = nook.Birthday{
		Day:   29,
		Month: time.October}
)

var (
	rodeoCode = nook.Code{
		Value: "bul01"}
)

var (
	rodeoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rodeo"}

	rodeoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Flèche"}

	rodeoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rodeo"}

	rodeoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Flèche"}

	rodeoGermanName = nook.Name{
		Language: language.German,
		Value:    "Toro"}

	rodeoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rodeo"}

	rodeoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ロデオ"}

	rodeoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "로데오"}

	rodeoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rodeo"}

	rodeoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Родео"}

	rodeoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗迪欧"}

	rodeoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rodeo"}

	rodeoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羅迪歐"}
)

var (
	rodeoName = nook.Languages{
		language.AmericanEnglish:      rodeoAmericanEnglishName,
		language.CanadianFrench:       rodeoCanadianFrenchName,
		language.Dutch:                rodeoDutchName,
		language.French:               rodeoFrenchName,
		language.German:               rodeoGermanName,
		language.Italian:              rodeoItalianName,
		language.Japanese:             rodeoJapaneseName,
		language.Korean:               rodeoKoreanName,
		language.LatinAmericanSpanish: rodeoLatinAmericanSpanishName,
		language.Russian:              rodeoRussianName,
		language.SimplifiedChinese:    rodeoSimplifiedChineseName,
		language.Spanish:              rodeoSpanishName,
		language.TraditionalChinese:   rodeoTraditionalChineseName}
)

var (
	rodeoCharacter = nook.Character{
		Animal:   animal.Bull,
		Birthday: rodeoBirthday,
		Code:     rodeoCode,
		Gender:   gender.Male,
		Name:     rodeoName}
)

var (
	rodeoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chaps"}

	rodeoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "olé"}

	rodeoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "cowboy"}

	rodeoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "crac"}

	rodeoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "oléolé"}

	rodeoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "oléolé"}

	rodeoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "さすがに"}

	rodeoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "디용"}

	rodeoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "olé"}

	rodeoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ковбой"}

	rodeoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "了不起"}

	rodeoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "olé"}

	rodeoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "了不起"}
)

var (
	rodeoPhrase = nook.Languages{
		language.AmericanEnglish:      rodeoAmericanEnglishName,
		language.CanadianFrench:       rodeoCanadianFrenchName,
		language.Dutch:                rodeoDutchName,
		language.French:               rodeoFrenchName,
		language.German:               rodeoGermanName,
		language.Italian:              rodeoItalianName,
		language.Japanese:             rodeoJapaneseName,
		language.Korean:               rodeoKoreanName,
		language.LatinAmericanSpanish: rodeoLatinAmericanSpanishName,
		language.Russian:              rodeoRussianName,
		language.SimplifiedChinese:    rodeoSimplifiedChineseName,
		language.Spanish:              rodeoSpanishName,
		language.TraditionalChinese:   rodeoTraditionalChineseName}
)

var (
	Rodeo = nook.Villager{
		Character:   rodeoCharacter,
		Personality: personality.Lazy,
		Phrase:      rodeoPhrase}
)
