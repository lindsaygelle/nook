package chicken

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
	hankBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	hankCode = nook.Code{
		Value: ""}
)

var (
	hankAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hank"}

	hankCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	hankDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	hankFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bruce"}

	hankGermanName = nook.Name{
		Language: language.German,
		Value:    "Galiardo"}

	hankItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pettore"}

	hankJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タツた"}

	hankKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	hankLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	hankRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	hankSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "丁丁"}

	hankSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Claudio"}

	hankTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	hankName = nook.Languages{
		language.AmericanEnglish:      hankAmericanEnglishName,
		language.CanadianFrench:       hankCanadianFrenchName,
		language.Dutch:                hankDutchName,
		language.French:               hankFrenchName,
		language.German:               hankGermanName,
		language.Italian:              hankItalianName,
		language.Japanese:             hankJapaneseName,
		language.Korean:               hankKoreanName,
		language.LatinAmericanSpanish: hankLatinAmericanSpanishName,
		language.Russian:              hankRussianName,
		language.SimplifiedChinese:    hankSimplifiedChineseName,
		language.Spanish:              hankSpanishName,
		language.TraditionalChinese:   hankTraditionalChineseName}
)

var (
	hankCharacter = nook.Character{
		Animal:   animal.Chicken,
		Birthday: hankBirthday,
		Code:     hankCode,
		Key:      character.Hank,
		Gender:   gender.Male,
		Name:     hankName,
		Special:  false}
)

var (
	hankAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "buhk buhk"}

	hankCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	hankDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	hankFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "cocorico"}

	hankGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bakbak"}

	hankItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "boh boh"}

	hankJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "じゃん"}

	hankKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	hankLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	hankRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	hankSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啾啾"}

	hankSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cocorico"}

	hankTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	hankPhrase = nook.Languages{
		language.AmericanEnglish:      hankAmericanEnglishPhrase,
		language.CanadianFrench:       hankCanadianFrenchPhrase,
		language.Dutch:                hankDutchPhrase,
		language.French:               hankFrenchPhrase,
		language.German:               hankGermanPhrase,
		language.Italian:              hankItalianPhrase,
		language.Japanese:             hankJapanesePhrase,
		language.Korean:               hankKoreanPhrase,
		language.LatinAmericanSpanish: hankLatinAmericanSpanishPhrase,
		language.Russian:              hankRussianPhrase,
		language.SimplifiedChinese:    hankSimplifiedChinesePhrase,
		language.Spanish:              hankSpanishPhrase,
		language.TraditionalChinese:   hankTraditionalChinesePhrase}
)

var (
	Hank = nook.Villager{
		Character:   hankCharacter,
		Personality: personality.Jock,
		Phrase:      hankPhrase}
)
