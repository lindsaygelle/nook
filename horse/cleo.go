package horse

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Cléabeauté"}

	cleoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cleosnoepje"}

	cleoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cléabeauté"}

	cleoGermanName = nook.Name{
		Language: language.German,
		Value:    "Birgitplätzchen"}

	cleoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Susannadolcetto"}

	cleoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アイソトープあそばせ"}

	cleoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아이소토프훗훗"}

	cleoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Clotildehop-puá"}

	cleoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Клеорафинад"}

	cleoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小铀游游"}

	cleoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Clotildebomboncito"}

	cleoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小鈾遊遊"}
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
		Animal:   Horse,
		Birthday: cleoBirthday,
		Code:     cleoCode,
		Gender:   nook.Female,
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
	Cleo = nook.Villager{
		Character:   cleoCharacter,
		Personality: nook.Snooty,
		Phrase:      cleoPhrase}
)
