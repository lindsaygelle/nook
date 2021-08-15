package bull

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Flècheolé"}

	rodeoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rodeocowboy"}

	rodeoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Flèchecrac"}

	rodeoGermanName = nook.Name{
		Language: language.German,
		Value:    "Torooléolé"}

	rodeoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Rodeooléolé"}

	rodeoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ロデオさすがに"}

	rodeoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "로데오디용"}

	rodeoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rodeoolé"}

	rodeoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Родеоковбой"}

	rodeoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗迪欧了不起"}

	rodeoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rodeoolé"}

	rodeoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羅迪歐了不起"}
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
		Animal:   Bull,
		Birthday: rodeoBirthday,
		Code:     rodeoCode,
		Gender:   nook.Male,
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
		Personality: nook.Lazy,
		Phrase:      rodeoPhrase}
)
