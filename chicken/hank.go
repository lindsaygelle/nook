package chicken

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Brucecocorico"}

	hankGermanName = nook.Name{
		Language: language.German,
		Value:    "Galiardobakbak"}

	hankItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Pettoreboh boh"}

	hankJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "タツたじゃん"}

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
		Value:    "丁丁啾啾"}

	hankSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Claudiococorico"}

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
		Animal:   Chicken,
		Birthday: hankBirthday,
		Code:     hankCode,
		Gender:   nook.Male,
		Name:     hankName}
)

var (
	hankAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	hankCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	hankDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

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
		Value:    "N/A"}

	hankLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	hankRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	hankSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "啾啾"}

	hankSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cocorico"}

	hankTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	hankPhrase = nook.Languages{
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
	Hank = nook.Villager{
		Character:   hankCharacter,
		Personality: nook.Jock,
		Phrase:      hankPhrase}
)
