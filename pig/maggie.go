package pig

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	maggieBirthday = nook.Birthday{
		Day:   3,
		Month: time.September}
)

var (
	maggieCode = nook.Code{
		Value: "pig10"}
)

var (
	maggieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Maggie"}

	maggieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Marjoriegrouik"}

	maggieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Maggieknor-knor"}

	maggieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Marjoriegrouik"}

	maggieGermanName = nook.Name{
		Language: language.German,
		Value:    "Magdakringel"}

	maggieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Marcellasissì"}

	maggieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マーガレットうん"}

	maggieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마가렛유후"}

	maggieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Margaríoinki"}

	maggieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Мэггихрюля"}

	maggieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "玛格嗯"}

	maggieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Margarígirasol"}

	maggieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "瑪格嗯"}
)

var (
	maggieName = nook.Languages{
		language.AmericanEnglish:      maggieAmericanEnglishName,
		language.CanadianFrench:       maggieCanadianFrenchName,
		language.Dutch:                maggieDutchName,
		language.French:               maggieFrenchName,
		language.German:               maggieGermanName,
		language.Italian:              maggieItalianName,
		language.Japanese:             maggieJapaneseName,
		language.Korean:               maggieKoreanName,
		language.LatinAmericanSpanish: maggieLatinAmericanSpanishName,
		language.Russian:              maggieRussianName,
		language.SimplifiedChinese:    maggieSimplifiedChineseName,
		language.Spanish:              maggieSpanishName,
		language.TraditionalChinese:   maggieTraditionalChineseName}
)

var (
	maggieCharacter = nook.Character{
		Animal:   Pig,
		Birthday: maggieBirthday,
		Code:     maggieCode,
		Gender:   nook.Female,
		Name:     maggieName}
)

var (
	maggieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "schep"}

	maggieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "grouik"}

	maggieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "knor-knor"}

	maggieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "grouik"}

	maggieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kringel"}

	maggieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "sissì"}

	maggieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "うん"}

	maggieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "유후"}

	maggieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "oinki"}

	maggieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "хрюля"}

	maggieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嗯"}

	maggieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "girasol"}

	maggieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嗯"}
)

var (
	maggiePhrase = nook.Languages{
		language.AmericanEnglish:      maggieAmericanEnglishName,
		language.CanadianFrench:       maggieCanadianFrenchName,
		language.Dutch:                maggieDutchName,
		language.French:               maggieFrenchName,
		language.German:               maggieGermanName,
		language.Italian:              maggieItalianName,
		language.Japanese:             maggieJapaneseName,
		language.Korean:               maggieKoreanName,
		language.LatinAmericanSpanish: maggieLatinAmericanSpanishName,
		language.Russian:              maggieRussianName,
		language.SimplifiedChinese:    maggieSimplifiedChineseName,
		language.Spanish:              maggieSpanishName,
		language.TraditionalChinese:   maggieTraditionalChineseName}
)

var (
	Maggie = nook.Villager{
		Character:   maggieCharacter,
		Personality: nook.Normal,
		Phrase:      maggiePhrase}
)
