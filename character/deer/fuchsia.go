package deer

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
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
		Value:    "Rosanne"}

	fuchsiaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Fuchsia"}

	fuchsiaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rosanne"}

	fuchsiaGermanName = nook.Name{
		Language: language.German,
		Value:    "Selina"}

	fuchsiaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fuxia"}

	fuchsiaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ジェシカ"}

	fuchsiaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "제시카"}

	fuchsiaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rosalina"}

	fuchsiaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Фуксия"}

	fuchsiaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "洁西卡"}

	fuchsiaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rosalina"}

	fuchsiaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "潔西卡"}
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
		Animal:   animal.Deer,
		Birthday: fuchsiaBirthday,
		Code:     fuchsiaCode,
		Gender:   gender.Female,
		Name:     fuchsiaName}
)

var (
	fuchsiaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "girlfriend"}

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
		Personality: personality.BigSister,
		Phrase:      fuchsiaPhrase}
)
