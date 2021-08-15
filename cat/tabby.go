package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	tabbyBirthday = nook.Birthday{
		Day:   13,
		Month: time.August}
)

var (
	tabbyCode = nook.Code{
		Value: "cat12"}
)

var (
	tabbyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tabby"}

	tabbyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tigrimi-AOUUH"}

	tabbyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tabbymi-WAUW"}

	tabbyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tigrimi-AOUUH"}

	tabbyGermanName = nook.Name{
		Language: language.German,
		Value:    "Zitamrriau"}

	tabbyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "LiscameOW"}

	tabbyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トラこにゃは"}

	tabbyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "호냥이삐뽀"}

	tabbyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Lianaanimiaul"}

	tabbyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тэббимя-ого"}

	tabbySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小虎喵吼"}

	tabbySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Lianaanimiaul"}

	tabbyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小虎喵吼"}
)

var (
	tabbyName = nook.Languages{
		language.AmericanEnglish:      tabbyAmericanEnglishName,
		language.CanadianFrench:       tabbyCanadianFrenchName,
		language.Dutch:                tabbyDutchName,
		language.French:               tabbyFrenchName,
		language.German:               tabbyGermanName,
		language.Italian:              tabbyItalianName,
		language.Japanese:             tabbyJapaneseName,
		language.Korean:               tabbyKoreanName,
		language.LatinAmericanSpanish: tabbyLatinAmericanSpanishName,
		language.Russian:              tabbyRussianName,
		language.SimplifiedChinese:    tabbySimplifiedChineseName,
		language.Spanish:              tabbySpanishName,
		language.TraditionalChinese:   tabbyTraditionalChineseName}
)

var (
	tabbyCharacter = nook.Character{
		Animal:   Cat,
		Birthday: tabbyBirthday,
		Code:     tabbyCode,
		Gender:   nook.Female,
		Name:     tabbyName}
)

var (
	tabbyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "me-WOW"}

	tabbyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mi-AOUUH"}

	tabbyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mi-WAUW"}

	tabbyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mi-AOUUH"}

	tabbyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mrriau"}

	tabbyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "meOW"}

	tabbyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "にゃは"}

	tabbyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "삐뽀"}

	tabbyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "animiaul"}

	tabbyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мя-ого"}

	tabbySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "喵吼"}

	tabbySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "animiaul"}

	tabbyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "喵吼"}
)

var (
	tabbyPhrase = nook.Languages{
		language.AmericanEnglish:      tabbyAmericanEnglishName,
		language.CanadianFrench:       tabbyCanadianFrenchName,
		language.Dutch:                tabbyDutchName,
		language.French:               tabbyFrenchName,
		language.German:               tabbyGermanName,
		language.Italian:              tabbyItalianName,
		language.Japanese:             tabbyJapaneseName,
		language.Korean:               tabbyKoreanName,
		language.LatinAmericanSpanish: tabbyLatinAmericanSpanishName,
		language.Russian:              tabbyRussianName,
		language.SimplifiedChinese:    tabbySimplifiedChineseName,
		language.Spanish:              tabbySpanishName,
		language.TraditionalChinese:   tabbyTraditionalChineseName}
)

var (
	Tabby = nook.Villager{
		Character:   tabbyCharacter,
		Personality: nook.Peppy,
		Phrase:      tabbyPhrase}
)
