package lion

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	elvisBirthday = nook.Birthday{
		Day:   23,
		Month: time.July}
)

var (
	elvisCode = nook.Code{
		Value: "lon01"}
)

var (
	elvisAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Elvis"}

	elvisCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Elvisbébé"}

	elvisDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Elvisaloha"}

	elvisFrenchName = nook.Name{
		Language: language.French,
		Value:    "Elvisbébé"}

	elvisGermanName = nook.Name{
		Language: language.German,
		Value:    "Leonardogrolll"}

	elvisItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Elvisunh-hunh"}

	elvisJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キングダロガ"}

	elvisKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "킹안그냐"}

	elvisLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Elvisgroar"}

	elvisRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Элвисбуги-вуги"}

	elvisSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "皇狮听懂吧"}

	elvisSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Elvisgroar"}

	elvisTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "皇獅聽懂吧"}
)

var (
	elvisName = nook.Languages{
		language.AmericanEnglish:      elvisAmericanEnglishName,
		language.CanadianFrench:       elvisCanadianFrenchName,
		language.Dutch:                elvisDutchName,
		language.French:               elvisFrenchName,
		language.German:               elvisGermanName,
		language.Italian:              elvisItalianName,
		language.Japanese:             elvisJapaneseName,
		language.Korean:               elvisKoreanName,
		language.LatinAmericanSpanish: elvisLatinAmericanSpanishName,
		language.Russian:              elvisRussianName,
		language.SimplifiedChinese:    elvisSimplifiedChineseName,
		language.Spanish:              elvisSpanishName,
		language.TraditionalChinese:   elvisTraditionalChineseName}
)

var (
	elvisCharacter = nook.Character{
		Animal:   Lion,
		Birthday: elvisBirthday,
		Code:     elvisCode,
		Gender:   nook.Male,
		Name:     elvisName}
)

var (
	elvisAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "unh-hunh"}

	elvisCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bébé"}

	elvisDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "aloha"}

	elvisFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bébé"}

	elvisGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "grolll"}

	elvisItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "unh-hunh"}

	elvisJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ダロガ"}

	elvisKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "안그냐"}

	elvisLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "groar"}

	elvisRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "буги-вуги"}

	elvisSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "听懂吧"}

	elvisSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "groar"}

	elvisTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "聽懂吧"}
)

var (
	elvisPhrase = nook.Languages{
		language.AmericanEnglish:      elvisAmericanEnglishName,
		language.CanadianFrench:       elvisCanadianFrenchName,
		language.Dutch:                elvisDutchName,
		language.French:               elvisFrenchName,
		language.German:               elvisGermanName,
		language.Italian:              elvisItalianName,
		language.Japanese:             elvisJapaneseName,
		language.Korean:               elvisKoreanName,
		language.LatinAmericanSpanish: elvisLatinAmericanSpanishName,
		language.Russian:              elvisRussianName,
		language.SimplifiedChinese:    elvisSimplifiedChineseName,
		language.Spanish:              elvisSpanishName,
		language.TraditionalChinese:   elvisTraditionalChineseName}
)

var (
	Elvis = nook.Villager{
		Character:   elvisCharacter,
		Personality: nook.Cranky,
		Phrase:      elvisPhrase}
)
