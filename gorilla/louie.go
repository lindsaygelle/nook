package gorilla

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	louieBirthday = nook.Birthday{
		Day:   26,
		Month: time.March}
)

var (
	louieCode = nook.Code{
		Value: "gor04"}
)

var (
	louieAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Louie"}

	louieCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Louistonneau"}

	louieDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Louieoe-oe-ah"}

	louieFrenchName = nook.Name{
		Language: language.French,
		Value:    "Louistonneau"}

	louieGermanName = nook.Name{
		Language: language.German,
		Value:    "Ludwigugh-ugh"}

	louieItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Loubum bum"}

	louieJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "マッスルコング"}

	louieKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "머슬하압"}

	louieLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lougorilérico"}

	louieRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Луиух-ух-ах"}

	louieSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "壮壮钢钢"}

	louieSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lougorilérico"}

	louieTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "壯壯鋼鋼"}
)

var (
	louieName = nook.Languages{
		language.AmericanEnglish:      louieAmericanEnglishName,
		language.CanadianFrench:       louieCanadianFrenchName,
		language.Dutch:                louieDutchName,
		language.French:               louieFrenchName,
		language.German:               louieGermanName,
		language.Italian:              louieItalianName,
		language.Japanese:             louieJapaneseName,
		language.Korean:               louieKoreanName,
		language.LatinAmericanSpanish: louieLatinAmericanSpanishName,
		language.Russian:              louieRussianName,
		language.SimplifiedChinese:    louieSimplifiedChineseName,
		language.Spanish:              louieSpanishName,
		language.TraditionalChinese:   louieTraditionalChineseName}
)

var (
	louieCharacter = nook.Character{
		Animal:   Gorilla,
		Birthday: louieBirthday,
		Code:     louieCode,
		Gender:   nook.Male,
		Name:     louieName}
)

var (
	louieAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "tootsodelayhoo hoo ha"}

	louieCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "tonneau"}

	louieDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "oe-oe-ah"}

	louieFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "tonneau"}

	louieGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "ugh-ugh"}

	louieItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bum bum"}

	louieJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "コング"}

	louieKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "하압"}

	louieLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "gorilérico"}

	louieRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ух-ух-ах"}

	louieSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "钢钢"}

	louieSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "gorilérico"}

	louieTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鋼鋼"}
)

var (
	louiePhrase = nook.Languages{
		language.AmericanEnglish:      louieAmericanEnglishName,
		language.CanadianFrench:       louieCanadianFrenchName,
		language.Dutch:                louieDutchName,
		language.French:               louieFrenchName,
		language.German:               louieGermanName,
		language.Italian:              louieItalianName,
		language.Japanese:             louieJapaneseName,
		language.Korean:               louieKoreanName,
		language.LatinAmericanSpanish: louieLatinAmericanSpanishName,
		language.Russian:              louieRussianName,
		language.SimplifiedChinese:    louieSimplifiedChineseName,
		language.Spanish:              louieSpanishName,
		language.TraditionalChinese:   louieTraditionalChineseName}
)

var (
	Louie = nook.Villager{
		Character:   louieCharacter,
		Personality: nook.Jock,
		Phrase:      louiePhrase}
)
