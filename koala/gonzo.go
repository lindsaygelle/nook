package koala

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	gonzoBirthday = nook.Birthday{
		Day:   13,
		Month: time.October}
)

var (
	gonzoCode = nook.Code{
		Value: "kal04"}
)

var (
	gonzoAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gonzo"}

	gonzoCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Gonzo"}

	gonzoDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Gonzo"}

	gonzoFrenchName = nook.Name{
		Language: language.French,
		Value:    "Gonzo"}

	gonzoGermanName = nook.Name{
		Language: language.German,
		Value:    "Heribert"}

	gonzoItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Gonzo"}

	gonzoJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ゴンゾー"}

	gonzoKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "근성"}

	gonzoLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Bronco"}

	gonzoRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Гонзо"}

	gonzoSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "钢锁"}

	gonzoSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Bronco"}

	gonzoTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鋼鎖"}
)

var (
	gonzoName = nook.Languages{
		language.AmericanEnglish:      gonzoAmericanEnglishName,
		language.CanadianFrench:       gonzoCanadianFrenchName,
		language.Dutch:                gonzoDutchName,
		language.French:               gonzoFrenchName,
		language.German:               gonzoGermanName,
		language.Italian:              gonzoItalianName,
		language.Japanese:             gonzoJapaneseName,
		language.Korean:               gonzoKoreanName,
		language.LatinAmericanSpanish: gonzoLatinAmericanSpanishName,
		language.Russian:              gonzoRussianName,
		language.SimplifiedChinese:    gonzoSimplifiedChineseName,
		language.Spanish:              gonzoSpanishName,
		language.TraditionalChinese:   gonzoTraditionalChineseName}
)

var (
	gonzoCharacter = nook.Character{
		Animal:   Koala,
		Birthday: gonzoBirthday,
		Code:     gonzoCode,
		Gender:   nook.Male,
		Name:     gonzoName}
)

var (
	gonzoAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "mate"}

	gonzoCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "calyptus"}

	gonzoDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "partner"}

	gonzoFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "galopin"}

	gonzoGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "partner"}

	gonzoItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "dinamite"}

	gonzoJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だがや"}

	gonzoKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "버텨"}

	gonzoLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "jorf-jorf"}

	gonzoRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "напарник"}

	gonzoSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "钱钱"}

	gonzoSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "jorf-jorf"}

	gonzoTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "錢錢"}
)

var (
	gonzoPhrase = nook.Languages{
		language.AmericanEnglish:      gonzoAmericanEnglishName,
		language.CanadianFrench:       gonzoCanadianFrenchName,
		language.Dutch:                gonzoDutchName,
		language.French:               gonzoFrenchName,
		language.German:               gonzoGermanName,
		language.Italian:              gonzoItalianName,
		language.Japanese:             gonzoJapaneseName,
		language.Korean:               gonzoKoreanName,
		language.LatinAmericanSpanish: gonzoLatinAmericanSpanishName,
		language.Russian:              gonzoRussianName,
		language.SimplifiedChinese:    gonzoSimplifiedChineseName,
		language.Spanish:              gonzoSpanishName,
		language.TraditionalChinese:   gonzoTraditionalChineseName}
)

var (
	Gonzo = nook.Villager{
		Character:   gonzoCharacter,
		Personality: nook.Cranky,
		Phrase:      gonzoPhrase}
)
