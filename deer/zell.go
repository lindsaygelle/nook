package deer

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	zellBirthday = nook.Birthday{
		Day:   7,
		Month: time.June}
)

var (
	zellCode = nook.Code{
		Value: "der02"}
)

var (
	zellAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Zell"}

	zellCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Régiswapiti"}

	zellDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Zellgazellig"}

	zellFrenchName = nook.Name{
		Language: language.French,
		Value:    "Régiswapiti"}

	zellGermanName = nook.Name{
		Language: language.German,
		Value:    "Waltersproing"}

	zellItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Antilioarriba"}

	zellJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ネルソンたしかに"}

	zellKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "넬슨슴사사"}

	zellLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Corvilobambín"}

	zellRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Зеллпрыг-прыг"}

	zellSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "森森鹿营"}

	zellSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Corvilobambín"}

	zellTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "森森鹿營"}
)

var (
	zellName = nook.Languages{
		language.AmericanEnglish:      zellAmericanEnglishName,
		language.CanadianFrench:       zellCanadianFrenchName,
		language.Dutch:                zellDutchName,
		language.French:               zellFrenchName,
		language.German:               zellGermanName,
		language.Italian:              zellItalianName,
		language.Japanese:             zellJapaneseName,
		language.Korean:               zellKoreanName,
		language.LatinAmericanSpanish: zellLatinAmericanSpanishName,
		language.Russian:              zellRussianName,
		language.SimplifiedChinese:    zellSimplifiedChineseName,
		language.Spanish:              zellSpanishName,
		language.TraditionalChinese:   zellTraditionalChineseName}
)

var (
	zellCharacter = nook.Character{
		Animal:   Deer,
		Birthday: zellBirthday,
		Code:     zellCode,
		Gender:   nook.Male,
		Name:     zellName}
)

var (
	zellAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pronk"}

	zellCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "wapiti"}

	zellDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "gazellig"}

	zellFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "wapiti"}

	zellGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "sproing"}

	zellItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "arriba"}

	zellJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "たしかに"}

	zellKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "슴사사"}

	zellLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bambín"}

	zellRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "прыг-прыг"}

	zellSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "鹿营"}

	zellSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bambín"}

	zellTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鹿營"}
)

var (
	zellPhrase = nook.Languages{
		language.AmericanEnglish:      zellAmericanEnglishName,
		language.CanadianFrench:       zellCanadianFrenchName,
		language.Dutch:                zellDutchName,
		language.French:               zellFrenchName,
		language.German:               zellGermanName,
		language.Italian:              zellItalianName,
		language.Japanese:             zellJapaneseName,
		language.Korean:               zellKoreanName,
		language.LatinAmericanSpanish: zellLatinAmericanSpanishName,
		language.Russian:              zellRussianName,
		language.SimplifiedChinese:    zellSimplifiedChineseName,
		language.Spanish:              zellSpanishName,
		language.TraditionalChinese:   zellTraditionalChineseName}
)

var (
	Zell = nook.Villager{
		Character:   zellCharacter,
		Personality: nook.Smug,
		Phrase:      zellPhrase}
)
