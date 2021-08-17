package cat

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
		Value:    "Tigri"}

	tabbyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tabby"}

	tabbyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tigri"}

	tabbyGermanName = nook.Name{
		Language: language.German,
		Value:    "Zita"}

	tabbyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lisca"}

	tabbyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "トラこ"}

	tabbyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "호냥이"}

	tabbyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Liana"}

	tabbyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Тэбби"}

	tabbySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "小虎"}

	tabbySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Liana"}

	tabbyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "小虎"}
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
		Animal:   animal.Cat,
		Birthday: tabbyBirthday,
		Code:     tabbyCode,
		Key:      character.Tabby,
		Gender:   gender.Female,
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
		language.AmericanEnglish:      tabbyAmericanEnglishPhrase,
		language.CanadianFrench:       tabbyCanadianFrenchPhrase,
		language.Dutch:                tabbyDutchPhrase,
		language.French:               tabbyFrenchPhrase,
		language.German:               tabbyGermanPhrase,
		language.Italian:              tabbyItalianPhrase,
		language.Japanese:             tabbyJapanesePhrase,
		language.Korean:               tabbyKoreanPhrase,
		language.LatinAmericanSpanish: tabbyLatinAmericanSpanishPhrase,
		language.Russian:              tabbyRussianPhrase,
		language.SimplifiedChinese:    tabbySimplifiedChinesePhrase,
		language.Spanish:              tabbySpanishPhrase,
		language.TraditionalChinese:   tabbyTraditionalChinesePhrase}
)

var (
	Tabby = nook.Villager{
		Character:   tabbyCharacter,
		Personality: personality.Peppy,
		Phrase:      tabbyPhrase}
)
