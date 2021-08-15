package cow

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	normaBirthday = nook.Birthday{
		Day:   20,
		Month: time.September}
)

var (
	normaCode = nook.Code{
		Value: "cow06"}
)

var (
	normaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Norma"}

	normaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Normameuh nan"}

	normaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Normaboehoe"}

	normaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Normameuh nan"}

	normaGermanName = nook.Name{
		Language: language.German,
		Value:    "Nellymuuuhi"}

	normaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Normamu muuu"}

	normaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "いさこうふ"}

	normaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "미자에헤"}

	normaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Normamu-mu"}

	normaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Нормадо-ре-му"}

	normaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "晨曦微笑"}

	normaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Normamu-mu"}

	normaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "晨曦微笑"}
)

var (
	normaName = nook.Languages{
		language.AmericanEnglish:      normaAmericanEnglishName,
		language.CanadianFrench:       normaCanadianFrenchName,
		language.Dutch:                normaDutchName,
		language.French:               normaFrenchName,
		language.German:               normaGermanName,
		language.Italian:              normaItalianName,
		language.Japanese:             normaJapaneseName,
		language.Korean:               normaKoreanName,
		language.LatinAmericanSpanish: normaLatinAmericanSpanishName,
		language.Russian:              normaRussianName,
		language.SimplifiedChinese:    normaSimplifiedChineseName,
		language.Spanish:              normaSpanishName,
		language.TraditionalChinese:   normaTraditionalChineseName}
)

var (
	normaCharacter = nook.Character{
		Animal:   Cow,
		Birthday: normaBirthday,
		Code:     normaCode,
		Gender:   nook.Female,
		Name:     normaName}
)

var (
	normaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "hoof hoo"}

	normaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "meuh nan"}

	normaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "boehoe"}

	normaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "meuh nan"}

	normaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "muuuhi"}

	normaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mu muuu"}

	normaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "うふ"}

	normaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "에헤"}

	normaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "mu-mu"}

	normaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "до-ре-му"}

	normaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "微笑"}

	normaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mu-mu"}

	normaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "微笑"}
)

var (
	normaPhrase = nook.Languages{
		language.AmericanEnglish:      normaAmericanEnglishName,
		language.CanadianFrench:       normaCanadianFrenchName,
		language.Dutch:                normaDutchName,
		language.French:               normaFrenchName,
		language.German:               normaGermanName,
		language.Italian:              normaItalianName,
		language.Japanese:             normaJapaneseName,
		language.Korean:               normaKoreanName,
		language.LatinAmericanSpanish: normaLatinAmericanSpanishName,
		language.Russian:              normaRussianName,
		language.SimplifiedChinese:    normaSimplifiedChineseName,
		language.Spanish:              normaSpanishName,
		language.TraditionalChinese:   normaTraditionalChineseName}
)

var (
	Norma = nook.Villager{
		Character:   normaCharacter,
		Personality: nook.Normal,
		Phrase:      normaPhrase}
)
