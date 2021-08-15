package tiger

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	bangleBirthday = nook.Birthday{
		Day:   27,
		Month: time.August}
)

var (
	bangleCode = nook.Code{
		Value: "tig03"}
)

var (
	bangleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bangle"}

	bangleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Bengalema souris"}

	bangleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Banglegrommf"}

	bangleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Bengalema souris"}

	bangleGermanName = nook.Name{
		Language: language.German,
		Value:    "Tamaraschnurf"}

	bangleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tigriziagrouf"}

	bangleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ルーズなのぉー"}

	bangleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "루주쪼옥쪽"}

	bangleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Felinagri-gruá"}

	bangleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бэнглр-р-рф"}

	bangleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "卢姿是喔"}

	bangleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Felinami vida"}

	bangleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "盧姿是喔"}
)

var (
	bangleName = nook.Languages{
		language.AmericanEnglish:      bangleAmericanEnglishName,
		language.CanadianFrench:       bangleCanadianFrenchName,
		language.Dutch:                bangleDutchName,
		language.French:               bangleFrenchName,
		language.German:               bangleGermanName,
		language.Italian:              bangleItalianName,
		language.Japanese:             bangleJapaneseName,
		language.Korean:               bangleKoreanName,
		language.LatinAmericanSpanish: bangleLatinAmericanSpanishName,
		language.Russian:              bangleRussianName,
		language.SimplifiedChinese:    bangleSimplifiedChineseName,
		language.Spanish:              bangleSpanishName,
		language.TraditionalChinese:   bangleTraditionalChineseName}
)

var (
	bangleCharacter = nook.Character{
		Animal:   Tiger,
		Birthday: bangleBirthday,
		Code:     bangleCode,
		Gender:   nook.Female,
		Name:     bangleName}
)

var (
	bangleAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "growf"}

	bangleCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ma souris"}

	bangleDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "grommf"}

	bangleFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ma souris"}

	bangleGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnurf"}

	bangleItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "grouf"}

	bangleJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なのぉー"}

	bangleKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "쪼옥쪽"}

	bangleLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gri-gruá"}

	bangleRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "р-р-рф"}

	bangleSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "是喔"}

	bangleSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mi vida"}

	bangleTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "是喔"}
)

var (
	banglePhrase = nook.Languages{
		language.AmericanEnglish:      bangleAmericanEnglishName,
		language.CanadianFrench:       bangleCanadianFrenchName,
		language.Dutch:                bangleDutchName,
		language.French:               bangleFrenchName,
		language.German:               bangleGermanName,
		language.Italian:              bangleItalianName,
		language.Japanese:             bangleJapaneseName,
		language.Korean:               bangleKoreanName,
		language.LatinAmericanSpanish: bangleLatinAmericanSpanishName,
		language.Russian:              bangleRussianName,
		language.SimplifiedChinese:    bangleSimplifiedChineseName,
		language.Spanish:              bangleSpanishName,
		language.TraditionalChinese:   bangleTraditionalChineseName}
)

var (
	Bangle = nook.Villager{
		Character:   bangleCharacter,
		Personality: nook.Peppy,
		Phrase:      banglePhrase}
)
