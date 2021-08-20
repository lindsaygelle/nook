package koala

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
		Value:    "Kolala"}

	canberraDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Canberra"}

	canberraFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kolala"}

	canberraGermanName = nook.Name{
		Language: language.German,
		Value:    "Caroline"}

	canberraItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Canberra"}

	canberraJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キャンベラ"}

	canberraKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "캔버라"}

	canberraLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Canberra"}

	canberraRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Канберра"}

	canberraSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "简培菈"}

	canberraSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Canberra"}

	canberraTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "簡培菈"}
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
		Animal:   animal.Koala,
		Birthday: canberraBirthday,
		Code:     canberraCode,
		Key:      character.Canberra,
		Gender:   gender.Female,
		Name:     canberraName,
		Special:  false}
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
		language.AmericanEnglish:      canberraAmericanEnglishPhrase,
		language.CanadianFrench:       canberraCanadianFrenchPhrase,
		language.Dutch:                canberraDutchPhrase,
		language.French:               canberraFrenchPhrase,
		language.German:               canberraGermanPhrase,
		language.Italian:              canberraItalianPhrase,
		language.Japanese:             canberraJapanesePhrase,
		language.Korean:               canberraKoreanPhrase,
		language.LatinAmericanSpanish: canberraLatinAmericanSpanishPhrase,
		language.Russian:              canberraRussianPhrase,
		language.SimplifiedChinese:    canberraSimplifiedChinesePhrase,
		language.Spanish:              canberraSpanishPhrase,
		language.TraditionalChinese:   canberraTraditionalChinesePhrase}
)

var (
	Canberra = nook.Villager{
		Character:   canberraCharacter,
		Personality: personality.BigSister,
		Phrase:      canberraPhrase}
)
