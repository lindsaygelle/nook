package gorilla

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
		Value:    "Gogo"}

	peeweeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Peewee"}

	peeweeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gogo"}

	peeweeGermanName = nook.Name{
		Language: language.German,
		Value:    "Manfred"}

	peeweeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Kong"}

	peeweeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ダンベル"}

	peeweeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "덤벨"}

	peeweeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lotar"}

	peeweeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Пиви"}

	peeweeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "哑铃"}

	peeweeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lotar"}

	peeweeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "啞鈴"}
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
		Animal:   animal.Gorilla,
		Birthday: peeweeBirthday,
		Code:     peeweeCode,
		Key:      character.Peewee,
		Gender:   gender.Male,
		Name:     peeweeName,
		Special:  false}
)

var (
	peeweeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "li'l dude"}

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
		language.AmericanEnglish:      peeweeAmericanEnglishPhrase,
		language.CanadianFrench:       peeweeCanadianFrenchPhrase,
		language.Dutch:                peeweeDutchPhrase,
		language.French:               peeweeFrenchPhrase,
		language.German:               peeweeGermanPhrase,
		language.Italian:              peeweeItalianPhrase,
		language.Japanese:             peeweeJapanesePhrase,
		language.Korean:               peeweeKoreanPhrase,
		language.LatinAmericanSpanish: peeweeLatinAmericanSpanishPhrase,
		language.Russian:              peeweeRussianPhrase,
		language.SimplifiedChinese:    peeweeSimplifiedChinesePhrase,
		language.Spanish:              peeweeSpanishPhrase,
		language.TraditionalChinese:   peeweeTraditionalChinesePhrase}
)

var (
	Peewee = nook.Villager{
		Character:   peeweeCharacter,
		Personality: personality.Cranky,
		Phrase:      peeweePhrase}
)
