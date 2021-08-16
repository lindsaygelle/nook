package cow

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
		Value:    "Norma"}

	normaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Norma"}

	normaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Norma"}

	normaGermanName = nook.Name{
		Language: language.German,
		Value:    "Nelly"}

	normaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Norma"}

	normaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "いさこ"}

	normaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "미자"}

	normaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Norma"}

	normaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Норма"}

	normaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "晨曦"}

	normaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Norma"}

	normaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "晨曦"}
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
		Animal:   animal.Cow,
		Birthday: normaBirthday,
		Code:     normaCode,
		Key:      character.Norma,
		Gender:   gender.Female,
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
		Personality: personality.Normal,
		Phrase:      normaPhrase}
)
