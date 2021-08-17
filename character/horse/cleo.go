package horse

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
	cleoBirthday = nook.Birthday{
		Day:   9,
		Month: time.February}
)

var (
	cleoCode = nook.Code{
		Value: "hrs07"}
)

var (
	cleoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cleo"}

	cleoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cléa"}

	cleoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cleo"}

	cleoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cléa"}

	cleoGermanName = nook.Name{
		Language: language.German,
		Value:    "Birgit"}

	cleoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Susanna"}

	cleoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アイソトープ"}

	cleoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아이소토프"}

	cleoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Clotilde"}

	cleoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Клео"}

	cleoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小铀"}

	cleoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Clotilde"}

	cleoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小鈾"}
)

var (
	cleoName = nook.Languages{
		language.AmericanEnglish:      cleoAmericanEnglishName,
		language.CanadianFrench:       cleoCanadianFrenchName,
		language.Dutch:                cleoDutchName,
		language.French:               cleoFrenchName,
		language.German:               cleoGermanName,
		language.Italian:              cleoItalianName,
		language.Japanese:             cleoJapaneseName,
		language.Korean:               cleoKoreanName,
		language.LatinAmericanSpanish: cleoLatinAmericanSpanishName,
		language.Russian:              cleoRussianName,
		language.SimplifiedChinese:    cleoSimplifiedChineseName,
		language.Spanish:              cleoSpanishName,
		language.TraditionalChinese:   cleoTraditionalChineseName}
)

var (
	cleoCharacter = nook.Character{
		Animal:   animal.Horse,
		Birthday: cleoBirthday,
		Code:     cleoCode,
		Key:      character.Cleo,
		Gender:   gender.Female,
		Name:     cleoName}
)

var (
	cleoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sugar"}

	cleoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "beauté"}

	cleoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snoepje"}

	cleoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "beauté"}

	cleoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "plätzchen"}

	cleoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "dolcetto"}

	cleoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "あそばせ"}

	cleoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "훗훗"}

	cleoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "hop-puá"}

	cleoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "рафинад"}

	cleoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "游游"}

	cleoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bomboncito"}

	cleoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "遊遊"}
)

var (
	cleoPhrase = nook.Languages{
		language.AmericanEnglish:      cleoAmericanEnglishPhrase,
		language.CanadianFrench:       cleoCanadianFrenchPhrase,
		language.Dutch:                cleoDutchPhrase,
		language.French:               cleoFrenchPhrase,
		language.German:               cleoGermanPhrase,
		language.Italian:              cleoItalianPhrase,
		language.Japanese:             cleoJapanesePhrase,
		language.Korean:               cleoKoreanPhrase,
		language.LatinAmericanSpanish: cleoLatinAmericanSpanishPhrase,
		language.Russian:              cleoRussianPhrase,
		language.SimplifiedChinese:    cleoSimplifiedChinesePhrase,
		language.Spanish:              cleoSpanishPhrase,
		language.TraditionalChinese:   cleoTraditionalChinesePhrase}
)

var (
	Cleo = nook.Villager{
		Character:   cleoCharacter,
		Personality: personality.Snooty,
		Phrase:      cleoPhrase}
)
