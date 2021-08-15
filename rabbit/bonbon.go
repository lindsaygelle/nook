package rabbit

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	bonbonBirthday = nook.Birthday{
		Day:   3,
		Month: time.March}
)

var (
	bonbonCode = nook.Code{
		Value: "rbt17"}
)

var (
	bonbonAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bonbon"}

	bonbonCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Sylvettenananère"}

	bonbonDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Bonbonheer-lijk"}

	bonbonFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sylvettenananère"}

	bonbonGermanName = nook.Name{
		Language: language.German,
		Value:    "Henrikehüppi"}

	bonbonItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Lollyhurrah"}

	bonbonJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ミミィヤバッ"}

	bonbonKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "미미어뜨케"}

	bonbonLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Chocolatoh lalá"}

	bonbonRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Бонбонвкусняшка"}

	bonbonSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "妙妙不妙"}

	bonbonSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Chocolatoh lalá"}

	bonbonTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "妙妙不妙"}
)

var (
	bonbonName = nook.Languages{
		language.AmericanEnglish:      bonbonAmericanEnglishName,
		language.CanadianFrench:       bonbonCanadianFrenchName,
		language.Dutch:                bonbonDutchName,
		language.French:               bonbonFrenchName,
		language.German:               bonbonGermanName,
		language.Italian:              bonbonItalianName,
		language.Japanese:             bonbonJapaneseName,
		language.Korean:               bonbonKoreanName,
		language.LatinAmericanSpanish: bonbonLatinAmericanSpanishName,
		language.Russian:              bonbonRussianName,
		language.SimplifiedChinese:    bonbonSimplifiedChineseName,
		language.Spanish:              bonbonSpanishName,
		language.TraditionalChinese:   bonbonTraditionalChineseName}
)

var (
	bonbonCharacter = nook.Character{
		Animal:   Rabbit,
		Birthday: bonbonBirthday,
		Code:     bonbonCode,
		Gender:   nook.Female,
		Name:     bonbonName}
)

var (
	bonbonAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "deelish"}

	bonbonCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "nananère"}

	bonbonDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "heer-lijk"}

	bonbonFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "nananère"}

	bonbonGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hüppi"}

	bonbonItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "hurrah"}

	bonbonJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ヤバッ"}

	bonbonKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "어뜨케"}

	bonbonLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "oh lalá"}

	bonbonRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "вкусняшка"}

	bonbonSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "不妙"}

	bonbonSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "oh lalá"}

	bonbonTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "不妙"}
)

var (
	bonbonPhrase = nook.Languages{
		language.AmericanEnglish:      bonbonAmericanEnglishName,
		language.CanadianFrench:       bonbonCanadianFrenchName,
		language.Dutch:                bonbonDutchName,
		language.French:               bonbonFrenchName,
		language.German:               bonbonGermanName,
		language.Italian:              bonbonItalianName,
		language.Japanese:             bonbonJapaneseName,
		language.Korean:               bonbonKoreanName,
		language.LatinAmericanSpanish: bonbonLatinAmericanSpanishName,
		language.Russian:              bonbonRussianName,
		language.SimplifiedChinese:    bonbonSimplifiedChineseName,
		language.Spanish:              bonbonSpanishName,
		language.TraditionalChinese:   bonbonTraditionalChineseName}
)

var (
	Bonbon = nook.Villager{
		Character:   bonbonCharacter,
		Personality: nook.Peppy,
		Phrase:      bonbonPhrase}
)
