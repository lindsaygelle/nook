package koala

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	canberraBirthday = nook.Birthday{
		Day:   14,
		Month: time.May}
)

var (
	canberraCode = nook.Code{
		Value: "kal08"}
)

var (
	canberraAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Canberra"}

	canberraCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kolalazyva"}

	canberraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Canberraechnie"}

	canberraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kolalazyva"}

	canberraGermanName = nook.Name{
		Language: language.German,
		Value:    "Carolinekolala"}

	canberraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Canberraeurekalì"}

	canberraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キャンベラえ～"}

	canberraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "캔버라어어"}

	canberraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Canberrakoalalá"}

	canberraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Канберраох-хо"}

	canberraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "简培菈咦"}

	canberraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Canberrakolalá"}

	canberraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "簡培菈咦～"}
)

var (
	canberraName = nook.Languages{
		language.AmericanEnglish:      canberraAmericanEnglishName,
		language.CanadianFrench:       canberraCanadianFrenchName,
		language.Dutch:                canberraDutchName,
		language.French:               canberraFrenchName,
		language.German:               canberraGermanName,
		language.Italian:              canberraItalianName,
		language.Japanese:             canberraJapaneseName,
		language.Korean:               canberraKoreanName,
		language.LatinAmericanSpanish: canberraLatinAmericanSpanishName,
		language.Russian:              canberraRussianName,
		language.SimplifiedChinese:    canberraSimplifiedChineseName,
		language.Spanish:              canberraSpanishName,
		language.TraditionalChinese:   canberraTraditionalChineseName}
)

var (
	canberraCharacter = nook.Character{
		Animal:   Koala,
		Birthday: canberraBirthday,
		Code:     canberraCode,
		Gender:   nook.Female,
		Name:     canberraName}
)

var (
	canberraAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nuh uh"}

	canberraCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "zyva"}

	canberraDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "echnie"}

	canberraFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "zyva"}

	canberraGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kolala"}

	canberraItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "eurekalì"}

	canberraJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "え～"}

	canberraKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어어"}

	canberraLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "koalalá"}

	canberraRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ох-хо"}

	canberraSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咦"}

	canberraSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "kolalá"}

	canberraTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "咦～"}
)

var (
	canberraPhrase = nook.Languages{
		language.AmericanEnglish:      canberraAmericanEnglishName,
		language.CanadianFrench:       canberraCanadianFrenchName,
		language.Dutch:                canberraDutchName,
		language.French:               canberraFrenchName,
		language.German:               canberraGermanName,
		language.Italian:              canberraItalianName,
		language.Japanese:             canberraJapaneseName,
		language.Korean:               canberraKoreanName,
		language.LatinAmericanSpanish: canberraLatinAmericanSpanishName,
		language.Russian:              canberraRussianName,
		language.SimplifiedChinese:    canberraSimplifiedChineseName,
		language.Spanish:              canberraSpanishName,
		language.TraditionalChinese:   canberraTraditionalChineseName}
)

var (
	Canberra = nook.Villager{
		Character:   canberraCharacter,
		Personality: nook.BigSister,
		Phrase:      canberraPhrase}
)
