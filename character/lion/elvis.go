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
		Value:    "Elvis"}

	elvisDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Elvis"}

	elvisFrenchName = nook.Name{
		Language: language.French,
		Value:    "Elvis"}

	elvisGermanName = nook.Name{
		Language: language.German,
		Value:    "Leonardo"}

	elvisItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Elvis"}

	elvisJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キング"}

	elvisKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "킹"}

	elvisLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Elvis"}

	elvisRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Элвис"}

	elvisSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "皇狮"}

	elvisSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Elvis"}

	elvisTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "皇獅"}
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
		Animal:   animal.Lion,
		Birthday: elvisBirthday,
		Code:     elvisCode,
		Key:      character.Elvis,
		Gender:   gender.Male,
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
		language.AmericanEnglish:      elvisAmericanEnglishPhrase,
		language.CanadianFrench:       elvisCanadianFrenchPhrase,
		language.Dutch:                elvisDutchPhrase,
		language.French:               elvisFrenchPhrase,
		language.German:               elvisGermanPhrase,
		language.Italian:              elvisItalianPhrase,
		language.Japanese:             elvisJapanesePhrase,
		language.Korean:               elvisKoreanPhrase,
		language.LatinAmericanSpanish: elvisLatinAmericanSpanishPhrase,
		language.Russian:              elvisRussianPhrase,
		language.SimplifiedChinese:    elvisSimplifiedChinesePhrase,
		language.Spanish:              elvisSpanishPhrase,
		language.TraditionalChinese:   elvisTraditionalChinesePhrase}
)

var (
	Elvis = nook.Villager{
		Character:   elvisCharacter,
		Personality: personality.Cranky,
		Phrase:      elvisPhrase}
)
