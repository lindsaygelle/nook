package pig

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	agnesBirthday = nook.Birthday{
		Day:   21,
		Month: time.April}
)

var (
	agnesCode = nook.Code{
		Value: "pig17"}
)

var (
	agnesAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Agnes"}

	agnesCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Pansysauciflard"}

	agnesDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Agnessnuffelaar"}

	agnesFrenchName = nook.Name{
		Language: language.French,
		Value:    "Pansysauciflard"}

	agnesGermanName = nook.Name{
		Language: language.German,
		Value:    "Noraschnorrz"}

	agnesItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Maiasplosh"}

	agnesJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "アグネスブフフ"}

	agnesKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "아그네스꿀꾸루"}

	agnesLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Nereadindoink"}

	agnesRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Агнесхрю-хрю"}

	agnesSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "古乃欣噗呼呼"}

	agnesSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Negreadindoink"}

	agnesTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "古乃欣噗呼呼"}
)

var (
	agnesName = nook.Languages{
		language.AmericanEnglish:      agnesAmericanEnglishName,
		language.CanadianFrench:       agnesCanadianFrenchName,
		language.Dutch:                agnesDutchName,
		language.French:               agnesFrenchName,
		language.German:               agnesGermanName,
		language.Italian:              agnesItalianName,
		language.Japanese:             agnesJapaneseName,
		language.Korean:               agnesKoreanName,
		language.LatinAmericanSpanish: agnesLatinAmericanSpanishName,
		language.Russian:              agnesRussianName,
		language.SimplifiedChinese:    agnesSimplifiedChineseName,
		language.Spanish:              agnesSpanishName,
		language.TraditionalChinese:   agnesTraditionalChineseName}
)

var (
	agnesCharacter = nook.Character{
		Animal:   Pig,
		Birthday: agnesBirthday,
		Code:     agnesCode,
		Gender:   nook.Female,
		Name:     agnesName}
)

var (
	agnesAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "snuffle"}

	agnesCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "sauciflard"}

	agnesDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "snuffelaar"}

	agnesFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "sauciflard"}

	agnesGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schnorrz"}

	agnesItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "splosh"}

	agnesJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ブフフ"}

	agnesKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "꿀꾸루"}

	agnesLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "dindoink"}

	agnesRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрю-хрю"}

	agnesSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "噗呼呼"}

	agnesSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "dindoink"}

	agnesTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "噗呼呼"}
)

var (
	agnesPhrase = nook.Languages{
		language.AmericanEnglish:      agnesAmericanEnglishName,
		language.CanadianFrench:       agnesCanadianFrenchName,
		language.Dutch:                agnesDutchName,
		language.French:               agnesFrenchName,
		language.German:               agnesGermanName,
		language.Italian:              agnesItalianName,
		language.Japanese:             agnesJapaneseName,
		language.Korean:               agnesKoreanName,
		language.LatinAmericanSpanish: agnesLatinAmericanSpanishName,
		language.Russian:              agnesRussianName,
		language.SimplifiedChinese:    agnesSimplifiedChineseName,
		language.Spanish:              agnesSpanishName,
		language.TraditionalChinese:   agnesTraditionalChineseName}
)

var (
	Agnes = nook.Villager{
		Character:   agnesCharacter,
		Personality: nook.BigSister,
		Phrase:      agnesPhrase}
)
