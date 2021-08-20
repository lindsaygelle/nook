package chicken

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
		Value:    "Jo"}

	broffinaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Broffina"}

	broffinaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jo"}

	broffinaGermanName = nook.Name{
		Language: language.German,
		Value:    "Elfriede"}

	broffinaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Concetta"}

	broffinaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カサンドラ"}

	broffinaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "히킨"}

	broffinaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Brunilda"}

	broffinaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Броффина"}

	broffinaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "凯西"}

	broffinaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Brunilda"}

	broffinaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "凱西"}
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
		Animal:   animal.Chicken,
		Birthday: broffinaBirthday,
		Code:     broffinaCode,
		Key:      character.Broffina,
		Gender:   gender.Female,
		Name:     broffinaName,
		Special:  false}
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
		language.AmericanEnglish:      broffinaAmericanEnglishPhrase,
		language.CanadianFrench:       broffinaCanadianFrenchPhrase,
		language.Dutch:                broffinaDutchPhrase,
		language.French:               broffinaFrenchPhrase,
		language.German:               broffinaGermanPhrase,
		language.Italian:              broffinaItalianPhrase,
		language.Japanese:             broffinaJapanesePhrase,
		language.Korean:               broffinaKoreanPhrase,
		language.LatinAmericanSpanish: broffinaLatinAmericanSpanishPhrase,
		language.Russian:              broffinaRussianPhrase,
		language.SimplifiedChinese:    broffinaSimplifiedChinesePhrase,
		language.Spanish:              broffinaSpanishPhrase,
		language.TraditionalChinese:   broffinaTraditionalChinesePhrase}
)

var (
	Broffina = nook.Villager{
		Character:   broffinaCharacter,
		Personality: personality.Snooty,
		Phrase:      broffinaPhrase}
)
