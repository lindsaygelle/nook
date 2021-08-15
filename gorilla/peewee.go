package gorilla

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	peeweeBirthday = nook.Birthday{
		Day:   11,
		Month: time.September}
)

var (
	peeweeCode = nook.Code{
		Value: "gor01"}
)

var (
	peeweeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Peewee"}

	peeweeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gogogadget"}

	peeweeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Peeweedwerg"}

	peeweeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gogogadget"}

	peeweeGermanName = nook.Name{
		Language: language.German,
		Value:    "Manfredzwerg"}

	peeweeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Kongbaaanane"}

	peeweeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ダンベルガオ"}

	peeweeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "덤벨겉멋"}

	peeweeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lotarcani"}

	peeweeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пивимелкота"}

	peeweeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哑铃高高"}

	peeweeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lotarcani"}

	peeweeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啞鈴高高"}
)

var (
	peeweeName = nook.Languages{
		language.AmericanEnglish:      peeweeAmericanEnglishName,
		language.CanadianFrench:       peeweeCanadianFrenchName,
		language.Dutch:                peeweeDutchName,
		language.French:               peeweeFrenchName,
		language.German:               peeweeGermanName,
		language.Italian:              peeweeItalianName,
		language.Japanese:             peeweeJapaneseName,
		language.Korean:               peeweeKoreanName,
		language.LatinAmericanSpanish: peeweeLatinAmericanSpanishName,
		language.Russian:              peeweeRussianName,
		language.SimplifiedChinese:    peeweeSimplifiedChineseName,
		language.Spanish:              peeweeSpanishName,
		language.TraditionalChinese:   peeweeTraditionalChineseName}
)

var (
	peeweeCharacter = nook.Character{
		Animal:   Gorilla,
		Birthday: peeweeBirthday,
		Code:     peeweeCode,
		Gender:   nook.Male,
		Name:     peeweeName}
)

var (
	peeweeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "li'l dudeli'l bitty baby"}

	peeweeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "gadget"}

	peeweeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "dwerg"}

	peeweeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "gadget"}

	peeweeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "zwerg"}

	peeweeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "baaanane"}

	peeweeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ガオ"}

	peeweeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "겉멋"}

	peeweeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cani"}

	peeweeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мелкота"}

	peeweeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "高高"}

	peeweeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cani"}

	peeweeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "高高"}
)

var (
	peeweePhrase = nook.Languages{
		language.AmericanEnglish:      peeweeAmericanEnglishName,
		language.CanadianFrench:       peeweeCanadianFrenchName,
		language.Dutch:                peeweeDutchName,
		language.French:               peeweeFrenchName,
		language.German:               peeweeGermanName,
		language.Italian:              peeweeItalianName,
		language.Japanese:             peeweeJapaneseName,
		language.Korean:               peeweeKoreanName,
		language.LatinAmericanSpanish: peeweeLatinAmericanSpanishName,
		language.Russian:              peeweeRussianName,
		language.SimplifiedChinese:    peeweeSimplifiedChineseName,
		language.Spanish:              peeweeSpanishName,
		language.TraditionalChinese:   peeweeTraditionalChineseName}
)

var (
	Peewee = nook.Villager{
		Character:   peeweeCharacter,
		Personality: nook.Cranky,
		Phrase:      peeweePhrase}
)
