package chicken

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	beckyBirthday = nook.Birthday{
		Day:   9,
		Month: time.December}
)

var (
	beckyCode = nook.Code{
		Value: "chn09"}
)

var (
	beckyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Becky"}

	beckyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Sonyacôôôôt"}

	beckyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Beckykuiken"}

	beckyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sonyacôôôôt"}

	beckyGermanName = nook.Name{
		Language: language.German,
		Value:    "Ingabokbok"}

	beckyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Annettacoccoricò"}

	beckyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アリアハレルヤ"}

	beckyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아리아도레미"}

	beckyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Raminacucurrús"}

	beckyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Беккимодненько"}

	beckySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咏旋哈雷路亚"}

	beckySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Raminacucurrús"}

	beckyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "詠旋哈雷路亞"}
)

var (
	beckyName = nook.Languages{
		language.AmericanEnglish:      beckyAmericanEnglishName,
		language.CanadianFrench:       beckyCanadianFrenchName,
		language.Dutch:                beckyDutchName,
		language.French:               beckyFrenchName,
		language.German:               beckyGermanName,
		language.Italian:              beckyItalianName,
		language.Japanese:             beckyJapaneseName,
		language.Korean:               beckyKoreanName,
		language.LatinAmericanSpanish: beckyLatinAmericanSpanishName,
		language.Russian:              beckyRussianName,
		language.SimplifiedChinese:    beckySimplifiedChineseName,
		language.Spanish:              beckySpanishName,
		language.TraditionalChinese:   beckyTraditionalChineseName}
)

var (
	beckyCharacter = nook.Character{
		Animal:   Chicken,
		Birthday: beckyBirthday,
		Code:     beckyCode,
		Gender:   nook.Female,
		Name:     beckyName}
)

var (
	beckyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "chicklet"}

	beckyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "côôôôt"}

	beckyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kuiken"}

	beckyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "côôôôt"}

	beckyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "bokbok"}

	beckyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "coccoricò"}

	beckyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ハレルヤ"}

	beckyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "도레미"}

	beckyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cucurrús"}

	beckyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "модненько"}

	beckySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哈雷路亚"}

	beckySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cucurrús"}

	beckyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "哈雷路亞"}
)

var (
	beckyPhrase = nook.Languages{
		language.AmericanEnglish:      beckyAmericanEnglishName,
		language.CanadianFrench:       beckyCanadianFrenchName,
		language.Dutch:                beckyDutchName,
		language.French:               beckyFrenchName,
		language.German:               beckyGermanName,
		language.Italian:              beckyItalianName,
		language.Japanese:             beckyJapaneseName,
		language.Korean:               beckyKoreanName,
		language.LatinAmericanSpanish: beckyLatinAmericanSpanishName,
		language.Russian:              beckyRussianName,
		language.SimplifiedChinese:    beckySimplifiedChineseName,
		language.Spanish:              beckySpanishName,
		language.TraditionalChinese:   beckyTraditionalChineseName}
)

var (
	Becky = nook.Villager{
		Character:   beckyCharacter,
		Personality: nook.Snooty,
		Phrase:      beckyPhrase}
)
