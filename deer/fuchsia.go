package deer

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	fuchsiaBirthday = nook.Birthday{
		Day:   19,
		Month: time.September}
)

var (
	fuchsiaCode = nook.Code{
		Value: "der06"}
)

var (
	fuchsiaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Fuchsia"}

	fuchsiaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rosanneframbou"}

	fuchsiaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Fuchsiameid"}

	fuchsiaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rosannerosace"}

	fuchsiaGermanName = nook.Name{
		Language: language.German,
		Value:    "Selinazick"}

	fuchsiaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fuxiacervilà"}

	fuchsiaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジェシカなんしか"}

	fuchsiaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "제시카핑크크"}

	fuchsiaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rosalinabraaaam"}

	fuchsiaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фуксиязазноба"}

	fuchsiaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "洁西卡西卡"}

	fuchsiaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rosalinabraaaam"}

	fuchsiaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "潔西卡西卡"}
)

var (
	fuchsiaName = nook.Languages{
		language.AmericanEnglish:      fuchsiaAmericanEnglishName,
		language.CanadianFrench:       fuchsiaCanadianFrenchName,
		language.Dutch:                fuchsiaDutchName,
		language.French:               fuchsiaFrenchName,
		language.German:               fuchsiaGermanName,
		language.Italian:              fuchsiaItalianName,
		language.Japanese:             fuchsiaJapaneseName,
		language.Korean:               fuchsiaKoreanName,
		language.LatinAmericanSpanish: fuchsiaLatinAmericanSpanishName,
		language.Russian:              fuchsiaRussianName,
		language.SimplifiedChinese:    fuchsiaSimplifiedChineseName,
		language.Spanish:              fuchsiaSpanishName,
		language.TraditionalChinese:   fuchsiaTraditionalChineseName}
)

var (
	fuchsiaCharacter = nook.Character{
		Animal:   Deer,
		Birthday: fuchsiaBirthday,
		Code:     fuchsiaCode,
		Gender:   nook.Female,
		Name:     fuchsiaName}
)

var (
	fuchsiaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "girlfriendprecious"}

	fuchsiaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "frambou"}

	fuchsiaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "meid"}

	fuchsiaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "rosace"}

	fuchsiaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zick"}

	fuchsiaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cervilà"}

	fuchsiaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "なんしか"}

	fuchsiaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "핑크크"}

	fuchsiaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "braaaam"}

	fuchsiaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "зазноба"}

	fuchsiaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "西卡"}

	fuchsiaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "braaaam"}

	fuchsiaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "西卡"}
)

var (
	fuchsiaPhrase = nook.Languages{
		language.AmericanEnglish:      fuchsiaAmericanEnglishName,
		language.CanadianFrench:       fuchsiaCanadianFrenchName,
		language.Dutch:                fuchsiaDutchName,
		language.French:               fuchsiaFrenchName,
		language.German:               fuchsiaGermanName,
		language.Italian:              fuchsiaItalianName,
		language.Japanese:             fuchsiaJapaneseName,
		language.Korean:               fuchsiaKoreanName,
		language.LatinAmericanSpanish: fuchsiaLatinAmericanSpanishName,
		language.Russian:              fuchsiaRussianName,
		language.SimplifiedChinese:    fuchsiaSimplifiedChineseName,
		language.Spanish:              fuchsiaSpanishName,
		language.TraditionalChinese:   fuchsiaTraditionalChineseName}
)

var (
	Fuchsia = nook.Villager{
		Character:   fuchsiaCharacter,
		Personality: nook.BigSister,
		Phrase:      fuchsiaPhrase}
)
