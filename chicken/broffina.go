package chicken

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	broffinaBirthday = nook.Birthday{
		Day:   24,
		Month: time.October}
)

var (
	broffinaCode = nook.Code{
		Value: "chn12"}
)

var (
	broffinaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Broffina"}

	broffinaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jomon œuf"}

	broffinaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Broffinakakel"}

	broffinaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jomon œuf"}

	broffinaGermanName = nook.Name{
		Language: language.German,
		Value:    "Elfriedeeieiei"}

	broffinaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Concettaovodè"}

	broffinaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カサンドラケッコー"}

	broffinaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "히킨오싹오싹"}

	broffinaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Brunildakiriquicó"}

	broffinaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Броффинаки-ки-ки"}

	broffinaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "凯西咯咯"}

	broffinaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Brunildapitas"}

	broffinaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "凱西咯咯"}
)

var (
	broffinaName = nook.Languages{
		language.AmericanEnglish:      broffinaAmericanEnglishName,
		language.CanadianFrench:       broffinaCanadianFrenchName,
		language.Dutch:                broffinaDutchName,
		language.French:               broffinaFrenchName,
		language.German:               broffinaGermanName,
		language.Italian:              broffinaItalianName,
		language.Japanese:             broffinaJapaneseName,
		language.Korean:               broffinaKoreanName,
		language.LatinAmericanSpanish: broffinaLatinAmericanSpanishName,
		language.Russian:              broffinaRussianName,
		language.SimplifiedChinese:    broffinaSimplifiedChineseName,
		language.Spanish:              broffinaSpanishName,
		language.TraditionalChinese:   broffinaTraditionalChineseName}
)

var (
	broffinaCharacter = nook.Character{
		Animal:   Chicken,
		Birthday: broffinaBirthday,
		Code:     broffinaCode,
		Gender:   nook.Female,
		Name:     broffinaName}
)

var (
	broffinaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cluckadoo"}

	broffinaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon œuf"}

	broffinaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kakel"}

	broffinaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon œuf"}

	broffinaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "eieiei"}

	broffinaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "ovodè"}

	broffinaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ケッコー"}

	broffinaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오싹오싹"}

	broffinaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "kiriquicó"}

	broffinaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ки-ки-ки"}

	broffinaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咯咯"}

	broffinaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "pitas"}

	broffinaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咯咯"}
)

var (
	broffinaPhrase = nook.Languages{
		language.AmericanEnglish:      broffinaAmericanEnglishName,
		language.CanadianFrench:       broffinaCanadianFrenchName,
		language.Dutch:                broffinaDutchName,
		language.French:               broffinaFrenchName,
		language.German:               broffinaGermanName,
		language.Italian:              broffinaItalianName,
		language.Japanese:             broffinaJapaneseName,
		language.Korean:               broffinaKoreanName,
		language.LatinAmericanSpanish: broffinaLatinAmericanSpanishName,
		language.Russian:              broffinaRussianName,
		language.SimplifiedChinese:    broffinaSimplifiedChineseName,
		language.Spanish:              broffinaSpanishName,
		language.TraditionalChinese:   broffinaTraditionalChineseName}
)

var (
	Broffina = nook.Villager{
		Character:   broffinaCharacter,
		Personality: nook.Snooty,
		Phrase:      broffinaPhrase}
)
