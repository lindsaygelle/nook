package cat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	kikiBirthday = nook.Birthday{
		Day:   8,
		Month: time.October}
)

var (
	kikiCode = nook.Code{
		Value: "cat04"}
)

var (
	kikiAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kiki"}

	kikiCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kikichti minou"}

	kikiDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kikipoesjemauw"}

	kikiFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kikichti minou"}

	kikiGermanName = nook.Name{
		Language: language.German,
		Value:    "Kikimiausi"}

	kikiItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Kikimieaoo"}

	kikiJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キャビアだニ"}

	kikiKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "캐비어뭐라지요"}

	kikiLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ágatamichifiú"}

	kikiRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кикикисуля"}

	kikiSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "余子酱酱酱"}

	kikiSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ágatagalletita"}

	kikiTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "余子醬醬醬"}
)

var (
	kikiName = nook.Languages{
		language.AmericanEnglish:      kikiAmericanEnglishName,
		language.CanadianFrench:       kikiCanadianFrenchName,
		language.Dutch:                kikiDutchName,
		language.French:               kikiFrenchName,
		language.German:               kikiGermanName,
		language.Italian:              kikiItalianName,
		language.Japanese:             kikiJapaneseName,
		language.Korean:               kikiKoreanName,
		language.LatinAmericanSpanish: kikiLatinAmericanSpanishName,
		language.Russian:              kikiRussianName,
		language.SimplifiedChinese:    kikiSimplifiedChineseName,
		language.Spanish:              kikiSpanishName,
		language.TraditionalChinese:   kikiTraditionalChineseName}
)

var (
	kikiCharacter = nook.Character{
		Animal:   Cat,
		Birthday: kikiBirthday,
		Code:     kikiCode,
		Gender:   nook.Female,
		Name:     kikiName}
)

var (
	kikiAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "kitty cat"}

	kikiCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chti minou"}

	kikiDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "poesjemauw"}

	kikiFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "chti minou"}

	kikiGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "miausi"}

	kikiItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "mieaoo"}

	kikiJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だニ"}

	kikiKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "뭐라지요"}

	kikiLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "michifiú"}

	kikiRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кисуля"}

	kikiSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "酱酱"}

	kikiSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "galletita"}

	kikiTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "醬醬"}
)

var (
	kikiPhrase = nook.Languages{
		language.AmericanEnglish:      kikiAmericanEnglishName,
		language.CanadianFrench:       kikiCanadianFrenchName,
		language.Dutch:                kikiDutchName,
		language.French:               kikiFrenchName,
		language.German:               kikiGermanName,
		language.Italian:              kikiItalianName,
		language.Japanese:             kikiJapaneseName,
		language.Korean:               kikiKoreanName,
		language.LatinAmericanSpanish: kikiLatinAmericanSpanishName,
		language.Russian:              kikiRussianName,
		language.SimplifiedChinese:    kikiSimplifiedChineseName,
		language.Spanish:              kikiSpanishName,
		language.TraditionalChinese:   kikiTraditionalChineseName}
)

var (
	Kiki = nook.Villager{
		Character:   kikiCharacter,
		Personality: nook.Normal,
		Phrase:      kikiPhrase}
)
